FROM golang
WORKDIR /src/github.com/schmidtp0740/MOEI_Generator
COPY . .
RUN go build -o app
CMD ["./app"]
