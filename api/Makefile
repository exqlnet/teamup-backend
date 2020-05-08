VERSION=latest

swagger:
	dir=$(mktemp -d)
	git clone "https://github.com/go-swagger/go-swagger" "$$dir"
	cd "$dir"
	go install ./cmd/swagger

doc: swagger
	swagger generate spec -o ./swagger.json --scan-models
	swagger serve swagger.json -p 8887
build:
	static -src=swagger/dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./api-server . && chmod +x ./api-server
	docker build . -t teamup-api:$(VERSION)
run: build
	docker run -d teamup-api:$(VERSION)
test:
	echo "Api server test ok"