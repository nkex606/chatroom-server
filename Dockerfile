FROM golang as builder
WORKDIR /go/src/github.com/nkex606/chatroom-server
COPY . .
RUN go get -d -v -t ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
ENV port="8080"
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/nkex606/chatroom-server/app .
CMD [ "./app" ] 