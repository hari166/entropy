FROM golang:1.20-alpine3.18

WORKDIR /entropy

COPY . .

RUN go build -o entropy

CMD [ "tail","-f","/dev/null" ]