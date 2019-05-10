FROM golang:1.7.4
LABEL author Tim Urista

ENV SOURCES /go/src/github.com/timurista/go-auth-api/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ["main.go"]