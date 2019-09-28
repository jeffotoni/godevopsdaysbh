#################################################
# Dockerfile distroless
#################################################
FROM golang:1.13.1 as builder
WORKDIR /go/src/
COPY . .
#RUN go install -v ./...
RUN go build -o api2 main2.go
RUN cp api2 /go/bin/api2

############################
# STEP 2 build a small image
############################
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/api2 /
CMD ["/api2"]
