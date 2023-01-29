
VERSION    = `cat VERSION`


LDFLAGS    = "\
    -X 'github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common.version=${VERSION}'"

all:

fmt:
	go fmt ./zenlayercloud/... ./example/...

project:
	go build -ldflags ${LDFLAGS}

test:

