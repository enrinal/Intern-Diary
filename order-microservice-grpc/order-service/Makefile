build:
	protoc -I. --go_out=plugins=micro:. \
		proto/order/order.proto
	GOOS=linux GOARCH=amd64 go build -o hermes-service-order && \
    docker build -t hermes-service-order .

run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		hermes-service-order