FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum .

RUN go mod tidy

COPY . . 

RUN go build -o buildFile

FROM alpine

COPY --from=builder /app/buildFile ./buildFile

CMD ["./buildFile"]