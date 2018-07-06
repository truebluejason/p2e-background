build:
	go build

start_dev:
	ENV=dev go run main.go

test:
	go run test.go