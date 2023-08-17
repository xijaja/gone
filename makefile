# 编译本机可执行文件
build:
	make build-front
	make build-back
build-back:
	go build -o app -tags embed main.go
build-front:
	cd frontend && pnpm build && cd ..

# 交叉编译，嵌入前端静态资源
b-lin:
	GOOS=linux GOARCH=amd64 go build -o app -tags embed main.go
b-mac:
	GOOS=darwin GOARCH=arm64 go build -o app -tags embed main.go
b-mac-amd:
	GOOS=darwin GOARCH=amd64 go build -o app -tags embed main.go
b-win:
	GOOS=windows GOARCH=amd64 go build -o app.exe -tags embed main.go

# 交叉编译，嵌入 sqlite 数据库
b-lin-sqlite:
	GOOS=linux GOARCH=amd64 go build -o app -tags 'sqlite' main.go
b-mac-sqlite:
	GOOS=darwin GOARCH=arm64 go build -o app -tags 'sqlite' main.go
b-mac-amd-sqlite:
	GOOS=darwin GOARCH=amd64 go build -o app -tags 'sqlite' main.go
b-win-sqlite:
	GOOS=windows GOARCH=amd64 go build -o app.exe -tags 'sqlite' main.go

# 使用 docker 启动数据库
pg:
	docker run -itd \
	--name pg \
	-p 5432:5432 \
	-v postgres:/var/lib/postgresql/data \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_PASSWORD=postgrespassword \
	-e POSTGRES_DB=postgres \
	-e TZ=PRC \
	postgres

rs:
	docker run -itd \
	--name rs \
	-p 6379:6379 \
	-v redis:/data \
	--requirepass gG4lD1oL0gB6gA1a \
	redis

