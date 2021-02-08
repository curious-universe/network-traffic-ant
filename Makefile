
include Makefile.common

.PHONY: debug-http server

default: server buildsucc

buildsucc:
	@echo Build Network-Traffic-Ant successfully!

debug-http:
	go build -o http/server http/main.go && ./http/server

server:
ifeq ($(TARGET), "")
	CGO_ENABLED=1 $(GOBUILD) $(RACE_FLAG) -ldflags '$(LDFLAGS)' -o bin/nta main.go
else
	CGO_ENABLED=1 $(GOBUILD) $(RACE_FLAG) -ldflags '$(LDFLAGS)' -o '$(TARGET)' main.go
endif