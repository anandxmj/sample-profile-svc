FROM 	golang:1.18 AS builder
WORKDIR	/go/src/github.com/anandxmj/profile-service
COPY	main.go ./
COPY    go.mod ./
COPY    go.sum ./
RUN 	go mod download
RUN	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM	alpine:latest
WORKDIR	/root/
COPY 	--from=builder /go/src/github.com/anandxmj/profile-service/main ./
CMD 	["./main"]


