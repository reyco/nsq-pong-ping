FROM golang:1.4

RUN go get github.com/bitly/go-nsq 

# Copy the local package files to the container's workspace.
ADD . /ping


ENTRYPOINT ["go", "run", "/ping/ping.go"]


