build:
	go clean
	go fmt
	go vet
	go build
	echo "Binary created inside project directory."

install:
	go install

start_dev:
	ENV=dev go run main.go

start_dev_bin:
	ENV=dev ./p2e-background

test:
	go run test.go