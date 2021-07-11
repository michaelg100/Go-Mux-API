FROM golang:1.16

WORKDIR /app
COPY . /app

RUN go get github.com/jdkato/prose/v2

CMD ["/bin/bash"]

