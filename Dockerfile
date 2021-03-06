FROM golang:1.6

RUN go get -u github.com/kardianos/govendor

ADD . /go/src/github.com/kotfalya/erebus

WORKDIR /go/src/github.com/kotfalya/erebus

RUN govendor sync

RUN go install github.com/kotfalya/erebus/cmd/erebusd

ENV PATH /go/bin:$PATH

ENTRYPOINT ["erebusd", "-serviceName", "agent", "-groupName", "erebus"]

EXPOSE 14141 14142
