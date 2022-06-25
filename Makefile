BINARY_NAME=mango
# These are the values we want to pass for VERSION and BUILD
# git tag 1.0.1
# git commit -am "One more change after the tags"
VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`

VERSION_LDFLAG=-X main.Version=${VERSION}
BUILD_LDFLAG=-X main.Build=${BUILD}

LDFLAGS=-ldflags "-w -s ${VERSION_LDFLAG} ${BUILD_LDFLAG}"

# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

# If the first argument is "run"...
ifeq (dev-run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

build:
	go build ${LDFLAGS} -o ${BINARY_NAME} main.go
	GOARCH=amd64 GOOS=darwin go build ${LDFLAGS} -o ${BINARY_NAME}-darwin-amd64 main.go
	GOARCH=arm64 GOOS=darwin go build ${LDFLAGS} -o ${BINARY_NAME}-darwin-arm64 main.go
	GOARCH=amd64 GOOS=linux  go build ${LDFLAGS} -o ${BINARY_NAME}-linux-amd64 main.go
	GOARCH=amd64 GOOS=window go build ${LDFLAGS} -o ${BINARY_NAME}-windows-amd64 main.go

dev-run:
	go run ${LDFLAGS} main.go ${RUN_ARGS}

run:
	./${BINARY_NAME} ${RUN_ARGS}

build_and_run: build run

# install linux amd64
install:
	go install ${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
	rm ${BINARY_NAME}-*

.PHONY: clean install
