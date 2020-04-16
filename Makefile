build:
	docker build -t demonst/hmdl-user-service:v1 .
deloy:
	docker login && docker build -t demonst/hmdl-user-service:v1 . && docker push demonst/hmdl-user-service:v1
gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
clean:
	rm pb/*.go
gitup:
	git add * && git commit -m "update" && git push
gen2:
	protoc -I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/lyft/protoc-gen-validate \
		--go_out=plugins=grpc:pb \
		--proto_path=proto proto/*.proto
	protoc -I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/lyft/protoc-gen-validate \
		--grpc-gateway_out=logtostderr=true:pb \
		--go_out=plugins=grpc:pb \
		--proto_path=proto proto/*.proto
