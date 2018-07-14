VERSION := ${VERSION}

build:
	[ -n "${VERSION}" ] # "Example Use: VERSION=1.0.0 make build"
	go clean
	go fmt
	go vet
	docker build . -t truebluejason/p2e-background:${VERSION}
	docker system prune --force

deploy:
	[ -n "${VERSION}" ] # "Example Use: VERSION=1.0.0 make deploy"
	docker tag truebluejason/p2e-background:${VERSION} truebluejason/p2e-background:latest
	docker push truebluejason/p2e-background:${VERSION}
	docker push truebluejason/p2e-background:latest
	@echo "Image 'truebluejason/p2e-background' with versions '${VERSION}' and 'latest' pushed."

install:
	go install

start_dev:
	ENV=dev go run main.go

test:
	echo ${HELLO}