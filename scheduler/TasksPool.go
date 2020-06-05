package scheduler

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/**
 *@author  wxn
 *@project ConcurrencyCron
 *@package ConcurrencyCron
 *@date    19-8-2 上午11:23
 */
type TasksPool interface {
	At(tm string) *task                                       //Run at some times
	AtTime(tm time.Time) *task                                //Run at some times
	Seconds() *task                                           //Run every few seconds
	Minutes() *task                                           //Run every few minutes
	Hours() *task                                             //Run every few hours
	Days() *task                                              //Run every few days
	Weekdays() *task                                          //Run every few weeks
	Monday() *task                                            //Run every few weeks on Monday
	Tuesday() *task                                           //Run every few weeks on Tuesday
	Wednesday() *task                                         //Run every few weeks on Wednesday
	Thursday() *task                                          //Run every few weeks on Thursday
	Friday() *task                                            //Run every few weeks on Friday
	Saturday() *task                                          //Run every few weeks on Saturday
	Sunday() *task                                            //Run every few weeks on Sunday
	JudgeRun(tm time.Time) bool                               //Determine if it is going to run
	Run(ticket TicketsPool, tm time.Time)                     //Run and judge the next run time
	GetNext() time.Time                                       //Get nest run time
	Do(taskFunc interface{}, params ...interface{}) *task     //Add a run function
	Remove(taskFunc interface{}, params ...interface{}) *task //Add a run function
	GetUuid() string                                          //get uuid
	Done() bool                                               //get once job done
	Once() bool
	GetFunInfo() string
}

type task struct {
	uuid            string        //uuid for each task
	interval        uint64        //interval foreach task
	atTime          time.Duration //The time of the task starts running e.g: 8:05
	dayRun          time.Time     //The last run time of the task
	latest          time.Time     //The last run time of the task
	next            time.Time     //The next run time of the task
	startDay        time.Weekday  //The weeks of the task starts running e.g: Monday
	funcName        string        //The name of the function that needs to be run
	funcVal         interface{}   //the function that needs to be run
	funcParam       []interface{} //Function parameters
	unit            string        //The unit of running interval
	once            bool          //run once
	done            bool          //once job done
	writer          io.Writer
	RunRemove       bool
	funcRemove      string        //The name of the function that needs to be run
	funcValRemove   interface{}   //the function that needs to be run
	funcParamRemove []interface{} //Function parameters
}

//create a task
func NewTask(interval uint64, writer io.Writer) TasksPool {
	jobId := uuid.NewV4()

	id := jobId.String()
	return &task{
		uuid:     id,
		interval: interval,
		latest:   time.Unix(0, 0),
		next:     time.Unix(0, 0),
		startDay: time.Sunday,
		writer:   writer,
	}
}

func NewOnceTask(writer io.Writer) TasksPool {
	jobId := uuid.NewV4()

	id := jobId.String()
	return &task{
		uuid:     id,
		interval: 1,
		latest:   time.Unix(0, 0),
		next:     time.Unix(0, 0),
		startDay: time.Sunday,
		unit:     "once",
		once:     true,
		writer:   writer,
	}
}

func (j *task) setUnit(unit string) *task {
	j.unit = unit
	return j
}

func (j *task) weekday(startDay time.Weekday) *task {
	j.startDay = startDay
	return j.Weekdays()
}

func (j *task) periodDuration() time.Duration {
	interval := time.Duration(j.interval)
	switch j.unit {
	case "seconds":
		return interval * time.Second
	case "minutes":
		return interval * time.Minute
	case "hours":
		return interval * time.Hour
	case "days":
		return interval * time.Hour * 24
	case "weeks":
		return interval * time.Hour * 24 * 7
	case "once":
		return interval * time.Hour * 24
	case "from":
		return 0
	}

	panic("No unit")
}

func (j *task) getNextRun() {

	now := time.Now()
	if j.latest == time.Unix(0, 0) {
		j.latest = now
	}
	switch j.unit {
	case "seconds":
		j.next = j.latest.Add(time.Duration(j.interval) * time.Second)
	case "minutes":
		j.next = j.latest.Add(time.Duration(j.interval) * time.Minute)
	case "hours":
		j.next = j.latest.Add(time.Duration(j.interval) * time.Hour)
	case "days":
		j.next = time.Date(j.latest.Year(), j.latest.Month(), j.latest.Day(), 0, 0, 0, 0, time.Local)
		j.next = j.next.Add(j.atTime)
	case "weeks":
		j.next = time.Date(j.latest.Year(), j.latest.Month(), j.latest.Day(), 0, 0, 0, 0, time.Local)
		dayDiff := int(j.startDay)
		dayDiff -= int(j.next.Weekday())
		if dayDiff != 0 {
			j.next = j.next.Add(time.Duration(dayDiff) * 24 * time.Hour)
		}
		j.next = j.next.Add(j.atTime)
	case "once":
		if !j.done {
			j.next = time.Date(j.latest.Year(), j.latest.Month(), j.latest.Day(), 0, 0, 0, 0, time.Local)
			j.next = j.next.Add(j.atTime)
		}
	case "from":
		if !j.done {
			j.next = j.dayRun
			j.next = j.next.Add(j.atTime)
		}
	}

	if !j.once {
		for j.next.Before(now) || j.next.Before(j.latest) {
			j.next = j.next.Add(j.periodDuration())
		}
	}

	fmt.Printf("Return next!")

}

func (j *task) GetNext() time.Time {
	return j.next
}

func (j *task) JudgeRun(tm time.Time) bool {
	if (j.once && !j.done) || !j.once {
		return tm.Unix() >= j.next.Unix()
	}
	return false
}

func (j *task) Run(ticket TicketsPool, tm time.Time) {
	defer func() {
		if p := recover(); p != nil {
			err, ok := interface{}(p).(error)
			var errMsg string
			if ok {
				errMsg = fmt.Sprintf("Async Call Panic! (error: %s)", err)
			} else {
				errMsg = fmt.Sprintf("Async Call Panic! (clue: %#v)", p)
			}
			fmt.Println(errMsg)
		}
		ticket.Return()
	}()
	taskFunc := reflect.ValueOf(j.funcVal)
	if len(j.funcParam) != taskFunc.Type().NumIn() {
		fmt.Println("param number error")
	}
	params := make([]reflect.Value, len(j.funcParam))
	for i, param := range j.funcParam {
		params[i] = reflect.ValueOf(param)
	}
	j.latest = tm
	if j.once {
		j.done = true

		if j.RunRemove {
			taskFunc := reflect.ValueOf(j.funcValRemove)
			if len(j.funcParamRemove) != taskFunc.Type().NumIn() {
				fmt.Println(j.writer, "param number error:need %d params given %d params", taskFunc.Type().NumIn(), len(j.funcParamRemove))
			}
			params := make([]reflect.Value, len(j.funcParamRemove))
			for i, param := range j.funcParamRemove {
				params[i] = reflect.ValueOf(param)
			}
			taskFunc.Call(params)
		}
	}
	j.getNextRun()
	taskFunc.Call(params)

}

func (j *task) Do(taskFunc interface{}, params ...interface{}) *task {
	tp := reflect.TypeOf(taskFunc)
	if tp.Kind() != reflect.Func {
		panic("only function can be schedule into the job queue.")
	}
	j.funcName = runtime.FuncForPC(reflect.ValueOf(taskFunc).Pointer()).Name()
	j.funcVal = taskFunc
	j.funcParam = params
	j.getNextRun()
	fmt.Println(j.writer, time.Now(), ": Remove cron,func: ", j.GetFunInfo())
	return j
}

func (j *task) Remove(taskFunc interface{}, params ...interface{}) *task {
	tp := reflect.TypeOf(taskFunc)
	if tp.Kind() != reflect.Func {
		panic("only function can be schedule into the job queue.")
	}
	j.funcRemove = runtime.FuncForPC(reflect.ValueOf(taskFunc).Pointer()).Name()
	j.funcValRemove = taskFunc
	j.funcParamRemove = params
	j.RunRemove = true
	j.getNextRun()
	fmt.Println(j.writer, time.Now(), ":created cron,func:", j.GetFunInfo())
	return j
}

func (j *task) Seconds() *task {
	return j.setUnit("seconds")
}

func (j *task) Minutes() *task {
	return j.setUnit("minutes")
}

func (j *task) Hours() *task {
	return j.setUnit("hours")
}

func (j *task) Days() *task {
	return j.setUnit("days")
}

func (j *task) Weekdays() *task {
	return j.setUnit("weeks")
}

func (j *task) Monday() (job *task) {
	return j.weekday(time.Monday)
}

func (j *task) Tuesday() *task {
	return j.weekday(time.Tuesday)
}

func (j *task) Wednesday() *task {
	return j.weekday(time.Wednesday)
}

func (j *task) Thursday() *task {
	return j.weekday(time.Thursday)
}

func (j *task) Friday() *task {
	return j.weekday(time.Friday)
}

func (j *task) Saturday() *task {
	return j.weekday(time.Saturday)
}

func (j *task) Sunday() *task {
	return j.weekday(time.Sunday)
}

func (j *task) GetUuid() string {
	return j.uuid
}

func formatTime(t string) (hour, min int, err error) {
	var er = errors.New("time format error")
	ts := strings.Split(t, ":")
	if len(ts) != 2 {
		err = er
		return
	}
	hourString := strings.Trim(ts[0], " ")
	minString := strings.Trim(ts[1], " ")

	if hour, err = strconv.Atoi(hourString); err != nil {
		return
	}
	if min, err = strconv.Atoi(minString); err != nil {
		return
	}

	if hour < 0 || hour > 23 || min < 0 || min > 59 {
		err = er
		return
	}
	return hour, min, nil
}

func (j *task) At(tm string) *task {
	hour, min, err := formatTime(tm)
	if err != nil {
		panic(err)
	}
	// save atTime start as duration from midnight
	j.atTime = time.Duration(hour)*time.Hour + time.Duration(min)*time.Minute
	return j
}

func (j *task) AtTime(tm time.Time) *task {
	j.dayRun = tm
	j.atTime = time.Duration(tm.Hour())*time.Hour + time.Duration(tm.Minute())*time.Minute
	return j
}

func (j *task) Done() bool {

	return j.done
}

func (j *task) Once() bool {
	return j.once
}

func (j *task) GetFunInfo() string {
	return fmt.Sprintf("%s(%v)", j.funcName, j.funcParam)
}
