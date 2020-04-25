build:
	export GOPROXY="https://goproxy.cn"
	go mod download
	cd user && make build
	cd api && make build

run: build
	cd user && make run
	cd api && make run
