.PHONY: helloserver

GOBIN = $(shell pwd)/bin

helloserver:
	go install ./cmd/helloserver
	@echo "Done building."
	@echo "Run \"$(GOBIN)/helloserver\" to launch helloserver."
