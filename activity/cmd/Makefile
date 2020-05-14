VERSION="latest"

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./activity-service . && chmod +x ./activity-service
	docker build . -t teamup-activity:$(VERSION)