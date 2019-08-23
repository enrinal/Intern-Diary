FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD order.json /app/order.json
ADD hermes-cli-order /app/hermes-cli-order

RUN ls

CMD ["./hermes-cli-order"]