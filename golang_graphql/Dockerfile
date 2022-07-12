FROM golang:1.18.3 as builder
ENV GOPATH /golang
WORKDIR /
COPY ./ /
RUN go build ./server.go

FROM debian:stretch 
COPY --from=0 /server .
COPY --from=0 ./db /db
COPY --from=0 ./app.env /
EXPOSE 8081
ENTRYPOINT ["./server"]