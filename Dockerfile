FROM golang:latest AS devbuilder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

RUN mkdir /api
ADD api /api/
WORKDIR /api

COPY api $GOPATH/src/mypackage/api/
WORKDIR $GOPATH/src/mypackage/api/

#RUN dep ensure --vendor-only
RUN dep ensure
RUN go build -o /go/bin/main
EXPOSE 8080
CMD ["/go/bin/main"]
