VERSION=1.0

build:
	protoc --proto_path=$$GOPATH/src:./proto --micro_out=./proto --go_out=./proto ./proto/user.proto
	go build -o ./user-service . && chmod +x ./user-service
run: build
	docker build . -t teamup-user:$(VERSION)
	docker run -d -e app-secret=${APP_SECRET} -e app-id=${APP_ID} teamup-user:$(VERSION)
