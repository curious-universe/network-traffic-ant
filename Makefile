include Makefile.common

.PHONY: debug-http server

default: server buildsucc

buildsucc:
	@echo Build Network-Traffic-Ant successfully!

debug-http:
	go build -o http/server http/main.go && ./http/server

debug-http2:
	go build -o http2/server http2/main.go && ./http2/server

server:
ifeq ($(TARGET), "")
	CGO_ENABLED=1 $(GOBUILD) $(RACE_FLAG) -ldflags '$(LDFLAGS)' -o bin/$(BIN_NAME) main.go
else
	CGO_ENABLED=1 $(GOBUILD) $(RACE_FLAG) -ldflags '$(LDFLAGS)' -o '$(TARGET)' main.go
endif

fmt:
	gofmt -w -s ./..