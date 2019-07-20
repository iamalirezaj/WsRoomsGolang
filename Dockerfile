FROM golang:1.12

# Update and install curl
RUN apt-get update

# Creating work directory
RUN mkdir $GOPATH/src/websocket

# Adding project to work directory
ADD . $GOPATH/src/websocket

# Choosing work directory
WORKDIR $GOPATH/src/websocket

# Install project dependencies
RUN go get

# build project
RUN go build -o server .

# Expose websocket server port
EXPOSE 3000