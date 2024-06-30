FROM golang:1.22.4
WORKDIR /builds
COPY . .
RUN go get -d -v
RUN go build -o server .
CMD [ "./server" ]
EXPOSE 9800