.PHONY: dev dev_http

dev:
	go build -o nta main.go && sudo ./nta

dev_http:
	go build -o http/server http/main.go && ./http/server