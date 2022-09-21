FROM golang

WORKDIR /usr/app/test

COPY . .

EXPOSE 3000

CMD ["go", "run", "main.go"]



FROM rabbitmq:3.8-management
RUN rabbitmq-plugins enable --offline rabbitmq_mqtt rabbitmq_federation_management rabbitmq_stomp