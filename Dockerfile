FROM golang:1.16-alpine
ENV CGO_ENABLED=0
ENV GO111MODULE=on
# Alpine images do not have git so we install it.
RUN apk add git
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD ./ /app
## We specify that we now wish to execute
## any further commands inside our /app
## directory
WORKDIR /app
## Add this go mod download command to pull in any dependencies
RUN go mod download
## We run go build to compile the binary
## executable of our Go program
RUN go build -trimpath -o build/figment-sensor-log-processor ./cmd/figment-sensor-log-processor/*.go
## Our start command which kicks off
## our newly created binary executable
CMD ["sh", "-c", "./build/figment-sensor-log-processor"]
