FROM golang:1.12.5

MAINTAINER Pedro Pacheco <pedro.pacheco@soickm.com>

COPY src/main.go src/
COPY src/services src/services
COPY src/go-sensor src/go-sensor
COPY src/trace src/trace
COPY src/form bin/form
COPY src/util src/util
COPY src/types src/types
RUN go get -d github.com/opentracing/opentracing-go
RUN go get -d github.com/opentracing/opentracing-go/ext
RUN go get -d github.com/go-sql-driver/mysql
RUN go get -d github.com/felixge/httpsnoop
RUN go get -d github.com/looplab/fsm
RUN go get -d github.com/opentracing/basictracer-go
ENV GOBIN="/go/bin"
RUN CGO_ENABLED=0 GOOS=linux go install src/main.go
WORKDIR /go/bin/
ENV GOPORT=8080
CMD ["./main"]

