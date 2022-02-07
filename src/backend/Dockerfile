FROM golang

ENV SRC_DIR=/go/src/github.com/rabbice/movieapi/src/backend
ENV GOBIN=/go/bin

WORKDIR $GOBIN

ADD . $SRC_DIR

RUN cd /go/src/;
RUN go get github.com/go-sql-driver/mysql;
RUN go get github.com/gorilla/mux;

RUN go install github.com/rabbice/movieapi/src/backend@latest;
ENTRYPOINT [ "./backend" ]

EXPOSE 8000