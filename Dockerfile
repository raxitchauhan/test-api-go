# FROM golang:1.19.3-alpine as builder
# COPY go.mod go.sum ./
# #/app/
# WORKDIR /go/src/
# COPY . .
# RUN go mod download
# # RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/bucketeer gitlab.com/idoko/bucketeer

# RUN go build -o /docker-gs-ping

# EXPOSE 8083

# CMD [ "/docker-gs-ping" ]
# # FROM alpine
# # RUN apk add --no-cache ca-certificates && update-ca-certificates
# # COPY --from=builder /app/build/bucketeer /usr/bin/bucketeer
# # EXPOSE 8080 8080
# # ENTRYPOINT ["/usr/bin/bucketeer"]

FROM golang:1.19

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd

CMD ["app"]