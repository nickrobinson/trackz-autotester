# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.6.0

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/eclipse/paho.mqtt.golang
RUN go get github.com/nickrobinson/trackz-autotester
RUN go install github.com/nickrobinson/trackz-autotester

CMD ["/go/bin/trackz-puller", "-server", "mqtt.isengard.io", "-port", "1883"]