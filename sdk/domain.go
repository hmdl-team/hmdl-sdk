package sdk

import (
	auth "github.com/korylprince/go-ad-auth/v3"
)

type Config struct {
	Server   string ``
	Port     int
	BaseDN   string
	Security auth.SecurityType
}

type UserDomain struct {
	UserName   string
	Name       string
	Email      string
	Title      string
	Company    string
	Department string
}

func (u *Config) LoginDomain(userName, password string) (*UserDomain, error) {

	config := &auth.Config{
		Server:   u.Server,
		Port:     u.Port,
		BaseDN:   u.BaseDN,
		Security: u.Security,
	}

	status, entry, _, err := auth.AuthenticateExtended(config, userName, password, nil, nil)

	if err != nil {
		//handle err
		return nil, err
	}

	if !status {
		//handle failed authentication
		return nil, nil
	}

	name := entry.GetAttributeValue("name")
	mail := entry.GetAttributeValue("mail")
	title := entry.GetAttributeValue("title")
	department := entry.GetAttributeValue("department")
	company := entry.GetAttributeValue("company")

	return &UserDomain{
		UserName:   userName,
		Email:      mail,
		Title:      title,
		Department: department,
		Company:    company,
		Name:       name,
	}, nil

}
