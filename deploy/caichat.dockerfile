FROM golang:1.21

WORKDIR /opt/caichat

COPY go.mod /opt/caichat/
COPY go.sum /opt/caichat/
RUN go mod download

COPY cmd/ /opt/caichat/cmd/
COPY pkg/ /opt/caichat/pkg/
RUN go build -o /opt/caichat/bin/caichat cmd/caichat/main.go

CMD [ "/opt/caichat/bin/caichat" ]
