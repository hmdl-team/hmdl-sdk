build:
	docker build -t demonst/hmdl-user-service:v1 .
deloy:
	docker login && docker build -t demonst/hmdl-user-service:v1 . && docker push demonst/hmdl-user-service:v1