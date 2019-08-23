build:
	GOOS=linux GOARCH=amd64 go build -o hermes-cli-order
	docker build -t hermes-cli-order .

run:
	docker run hermes-cli-order