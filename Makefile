build:
	go clean
	go fmt
	go vet
	env GOOS=linux GOARCH=amd64 GOARM=7 go build -o ${GOPATH}/bin/p2e-background
	@echo "Linux specific binary created inside '${GOPATH}/bin' directory."

install:
	go install

start_dev:
	ENV=dev go run main.go

start_dev_bin:
	ENV=dev ./p2e-background

test:
	echo ${HELLO}