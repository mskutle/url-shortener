FROM golang:1.21.1

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/api
RUN go build -o /url-shortener

EXPOSE 3000
CMD [ "/url-shortener" ]