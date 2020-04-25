build:
	export GOPROXY="https://goproxy.cn"
	go mod download
	go get github.com/golang/protobuf/{proto,protoc-gen-go}
	go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master
	cd user && make build
	cd api && make build

run: build
	cd user && make run
	cd api && make run
