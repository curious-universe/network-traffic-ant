FROM golang:1.16-rc-buster

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

RUN echo \
        deb http://mirrors.aliyun.com/debian/ buster main non-free contrib \
        deb-src http://mirrors.aliyun.com/debian/ buster main non-free contrib \
        deb http://mirrors.aliyun.com/debian-security buster/updates main \
        deb-src http://mirrors.aliyun.com/debian-security buster/updates main \
        deb http://mirrors.aliyun.com/debian/ buster-updates main non-free contrib \
        deb-src http://mirrors.aliyun.com/debian/ buster-updates main non-free contrib \
        deb http://mirrors.aliyun.com/debian/ buster-backports main non-free contrib \
        deb-src http://mirrors.aliyun.com/debian/ buster-backports main non-free contrib \
    > /etc/apt/sources.list && \
    apt-get update -y \
    apt-get install -y net-tools

WORKDIR /www