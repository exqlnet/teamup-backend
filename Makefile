build:
	export GOPROXY="https://goproxy.cn"
	go mod download
	cd user && make build
	cd api && make build

run: build
	docker-compose up

test:
	cd user && make test
	cd api && make test

deploy:
	docker-compose down
	docker-compose up -d
