# go-vue-example

example app using Vue.js and Go

## Usage

look make comannd


```
init/deps:
	go get -u github.com/golang/dep/cmd/dep

server/deps:
	dep ensure

server/run:
	go run main.go

front/deps:
	cd vue 
	npm install 
	cd ../

front/build:
	cd vue 
	npm run build
	cd ../

db/migrate:
	go run script/migrate.go

db/init:
	sh script/reset_db.sh
```