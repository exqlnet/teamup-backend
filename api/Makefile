VERSION=1.0

swagger:
	dir=$(mktemp -d)
	git clone "https://github.com/go-swagger/go-swagger" "$$dir"
	cd "$dir"
	go install ./cmd/swagger

doc: swagger
	swagger generate spec -o ./swagger.json --scan-models
	swagger serve swagger.json -p 8887

build:
	go build -o ./api-server . && chmod +x ./api-server

run: build
	docker build . -t teamup-api:$(VERSION)
	docker run -d teamup-api:$(VERSION)
