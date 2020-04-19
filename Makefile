build:
	docker build -t demonst/hmdl-user-service:v1 .
deloy:
	docker login && docker build -t demonst/hmdl-user-service:v1 . && docker push demonst/hmdl-user-service:v1
clean:
	rm pb/*.go
gitup:
	git add * && git commit -m "update" && git push
gen:
                                        	protoc -I${GOPATH}/src \
                                        		--go_out=plugins=grpc:pb \
                                        		--proto_path=proto proto/*.proto