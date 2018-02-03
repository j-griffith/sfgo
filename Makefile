.PHONY: all types provider clean

all: types provider client

types:
	if [ ! -d ./vendor ]; then dep ensure; fi
	go build ./pkg/types

provider:
	if [ ! -d ./vendor ]; then dep ensure; fi
	go build ./pkg/provider

client:
	if [ ! -d ./vendor ]; then dep ensure; fi
	go build ./pkg/client

install:
	go install ./pkg/types
	go install ./pkg/provider
	go install ./pkg/client
