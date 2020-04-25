build:
	export GOPROXY="https://goproxy.cn"
	go mod download
	cd user && make build
	cd api && make build

run: build
	cd user && make run
	cd api && make run

test:
	cd user && make test
	cd api && make test

deploy:
	docker-compose down
	docker rmi registry.ncuos.com/591210216_qq.com/teamup-user:latest
	docker rmi registry.ncuos.com/591210216_qq.com/teamup-api:latest
	docker-compose up -d
