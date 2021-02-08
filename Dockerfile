FROM golang:1.16-rc-buster

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /www