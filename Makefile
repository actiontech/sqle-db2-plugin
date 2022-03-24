override DOCKER         		= $(shell which docker)
override GIT_VERSION    		= $(shell git rev-parse --abbrev-ref HEAD)${CUSTOM} $(shell git rev-parse HEAD)
override PROJECT_NAME 			= sqle-db2-plugin
override DATABASE_NAME			?=DB2
override LDFLAGS 				= -ldflags "-X 'main.version=\"${GIT_VERSION}\"'"
override GOBIN					= ${shell pwd}/bin

GO_COMPILER_IMAGE ?= golang:1.16

# Copy from SQLE
PROJECT_VERSION = $(shell if [ "$$(git tag --points-at HEAD | tail -n1)" ]; then git tag --points-at HEAD | tail -n1 | sed 's/v\(.*\)/\1/'; else git rev-parse --abbrev-ref HEAD | sed 's/release-\(.*\)/\1/' | tr '-' '\n' | head -n1; fi)

default: install

install:
	go build -mod=vendor ${LDFLAGS} -o $(GOBIN)/$(PROJECT_NAME) ./

# This item is temporarily unavailable
docker_install:
	$(DOCKER) run -v $(shell pwd):/universe --rm $(GO_COMPILER_IMAGE) sh -c "go env -w GO111MODULE=on \
	&& go env -w GOPROXY=https://goproxy.cn,direct \
	&& go get -d github.com/ibmdb/go_ibm_db \
	&& cd /go/pkg/mod/github.com/ibmdb/go_ibm_db@v0.4.1/installer \
	&& go run setup.go \
	&& export DB2HOME=/go/pkg/mod/github.com/ibmdb/clidriver \
	&& export CGO_CFLAGS=-I/go/pkg/mod/github.com/ibmdb/clidriver/include \
	&& export CGO_LDFLAGS=-L/go/pkg/mod/github.com/ibmdb/clidriver/lib \
	&& export LD_LIBRARY_PATH=/go/pkg/mod/github.com/ibmdb/clidriver/lib \
	&& apt update \
	&& apt install -y libxml2\
	&& cd /universe \
	&& make install $(MAKEFLAGS)"

upload:
	curl -T $(GOBIN)/$(PROJECT_NAME) ftp://$(RELEASE_FTPD_HOST)/actiontech-sqle/plugins/$(PROJECT_VERSION)/$(PROJECT_NAME) --ftp-create-dirs