FROM golang

WORKDIR /opt/facade
COPY . .
RUN go build -o /opt/facade/facade cmd/facade/main.go

CMD [ "/opt/facade/facade" ]
