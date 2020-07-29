FROM golang
ADD . /go/src/github.com/lunhu/studentrecord
WORKDIR /go/src/github.com/lunhu/studentrecord
ENV GO111MODULE=on
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
RUN go install studentrecord
ENTRYPOINT /go/bin/studentrecord
EXPOSE 3000