.PHONY: all types provider clean

all: types provider

types:
	if [ ! -d ./vendor ]; then dep ensure; fi
	go build ./pkg/types

provider:
	if [ ! -d ./vendor ]; then dep ensure; fi
	go build ./pkg/provider

install:
	go install ./pkg/types
	go install ./pkg/provider
