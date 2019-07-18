BINARY=engine
test: 
	go test -v -cover -covermode=atomic ./...

engine:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} main.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t simple-order-system .

run:
	docker-compose up -d

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint