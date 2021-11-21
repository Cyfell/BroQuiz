FROM golang:alpine

# Get requirements
RUN go install github.com/freshautomations/stoml@latest
WORKDIR /go/src/github.com/Cyfell/BroQuiz

# Get requirements
COPY go.mod go.sum ./
RUN go mod download -x

# Install daemon
COPY . .
RUN go install -v ./cmd/...

FROM alpine

VOLUME /etc/broquiz  

EXPOSE 8080

COPY --from=0 /go/bin /usr/local/bin

COPY ./build/package /docker

CMD [ "sh", "/docker/entrypoint.sh" ]