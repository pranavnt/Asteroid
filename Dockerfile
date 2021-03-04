FROM golang:1.16

LABEL maintainer="pranavnt@outlook.com"
LABEL description="Intuitive NoSQL database."

COPY . /usr/src/app
WORKDIR /usr/src/app

RUN go get -v -t -d ./...
RUN go build -o asteroid .

EXPOSE 5555
CMD [ "./asteroid" ]
