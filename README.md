# network-traffic-ant
Network traffic ant, base on pcap

# Docker

```shell
// build the image
docker build -t nta .

// run a container
docker run -it --name ntac --privileged -w /www -v$(PWD):/www nta bash

// next start container
docker start -i ntac
```


# Develop

## start a http server 

```shell
make debug-http
```

## build 

```shell
make
```

## demo

```shell
➜  network-traffic-ant git:(main) ✗ ./bin/nta version    
Version is v1.0.0
BuildTS is 2021-02-08 16:02:04
GitHash is 7df5bf307a23755bb2ed18aa64fc91a8cee72aff
GitBranch is main
```