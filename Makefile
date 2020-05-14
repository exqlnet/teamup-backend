build:
	export GOPROXY="https://goproxy.cn"
	go mod download
	cd config && make build
	cd user/cmd && make build
	cd activity/cmd && make build
	cd api/cmd && make build

run: build
	docker-compose up

test:
	cd user/cmd && make test
	cd api/cmd && make test

deploy:
	docker-compose down
	docker-compose up -d
