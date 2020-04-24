build:
	export GOPROXY="https://goproxy.cn"
	go mod download
	cd user && make build
