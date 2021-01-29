.PHONY: default build
APPS        := broker dbcreate
BLDDIR      ?= bin
BUILD_DATE  := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

.EXPORT_ALL_VARIABLES:
GO111MODULE  = on

default: build

build: clean $(APPS)

$(BLDDIR)/%:
	go build -o $@ ./cmd/$*

$(APPS): %: $(BLDDIR)/%

deps:
	@echo 'Installing go modules...'
	@go mod download

clean:
	@mkdir -p $(BLDDIR)
	@for app in $(APPS) ; do \
		rm -f $(BLDDIR)/$$app ; \
	done

format:
	@echo 'Formatting the code...'
	@gofmt -w .
	@goimports -local "github.com/sergeykoshlatuu/test_go_mqtt_sqlite" -w .
