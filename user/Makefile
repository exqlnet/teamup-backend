VERSION=latest

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./user-service . && chmod +x ./user-service
	docker build . -t registry.ncuos.com/591210216_qq.com/teamup-user:$(VERSION)
run: build
	docker run -d -e app-secret=${APP_SECRET} -e app-id=${APP_ID} teamup-user:$(VERSION)

test:
	echo "User service test ok"
