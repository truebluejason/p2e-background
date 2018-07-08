build:
	go clean
	go fmt
	go vet
	env GOOS=linux GOARCH=amd64 GOARM=7 go build
	echo "Linux specific binary created inside project directory."

install:
	go install

start_dev:
	ENV=dev go run main.go

start_dev_bin:
	ENV=dev ./p2e-background

test:
	go run test.go