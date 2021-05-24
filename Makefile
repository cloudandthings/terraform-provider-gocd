TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=cloudandthing.io
NAMESPACE=cicd
NAME=gocd
BINARY=terraform-provider-${NAME}
VERSION=0.2
OS_ARCH=darwin_amd64

GOFMT_FILES?=./internal/provider ./internal/hashcode

# For local testing, run `make testacc`
SERVER ?=http://127.0.0.1:8153/go/
TESTARGS ?= -race -coverprofile=profile.out -covermode=atomic

export GOCD_URL=$(SERVER)
export GOCD_SKIP_SSL_CHECK=1

default: install

build:
	go build -o ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

# ###########################
# Helpers
# ###########################
fmt:
	go fmt $(GOFMT_FILES)

# ###########################
# Testing
# ###########################
fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

test: fmtcheck
	TF_ACC=1 TESTARGS="$(TESTARGS)" bash ./scripts/go-test.sh

testacc: provision-test-gocd
	bash scripts/wait-for-test-server.sh
	TF_ACC=1 $(MAKE) test

provision-test-gocd:
	cp godata/default.gocd.config.xml godata/server/config/cruise-config.xml
	docker-compose build --build-arg UID=$(shell id -u) gocd-server
	docker-compose up -d

report_coverage:
	curl -s https://codecov.io/bash | bash -

teardown-test-gocd:
	rm -f godata/server/config/cruise-config.xml
	docker-compose down

cleanup: teardown-test-gocd 