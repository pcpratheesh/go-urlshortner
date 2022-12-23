FROM golang:1.19-alpine

RUN mkdir /build
ADD . /build/
WORKDIR /build

ENV URLSRT_BASE_URL=""

RUN go build -o url.shortner .

EXPOSE 8080

CMD [ "./url.shortner" ]
