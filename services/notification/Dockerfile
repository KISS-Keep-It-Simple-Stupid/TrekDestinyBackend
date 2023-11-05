FROM golang:1.20.4-alpine3.17

RUN mkdir app
WORKDIR /app
COPY ./*.mod /app/
COPY ./*.sum .
RUN go mod tidy
COPY . .
EXPOSE 8040
CMD [ "go" , "run" , "." ]
