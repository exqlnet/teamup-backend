swagger:
	dir=$(mktemp -d)
	git clone "https://github.com/go-swagger/go-swagger" "$$dir"
	cd "$dir"
	go install ./cmd/swagger

doc: swagger
	swagger generate spec -o ./swagger.json --scan-models
	swagger serve swagger.json -p 8887
