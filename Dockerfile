FROM golang

WORKDIR /usr/app

COPY . .

EXPOSE 3000

CMD ["go", "run", "main.go"]
