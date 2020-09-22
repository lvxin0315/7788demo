FROM golang

ENV TZ Asia/Shanghai

ENV GO111MODULE on

ENV GOPROXY https://goproxy.io

RUN echo 'Asia/Shanghai' >/etc/timezone

COPY . .

# 编译，然后运行
RUN go build -o main main.go

CMD ./main