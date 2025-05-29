GO=go
GOBUILD=${GO} build
GOVARS=GOOS=linux GOARCH=amd64 CGO_ENABLED=0
GOLDFLAGS="-ldflags=-s -w"

export

.PHONY: all build deploy output clean

all: deploy

# NOTE: The go binary must be called "bootstrap" for AWS Lambda to find it.
build: bin/bootstrap api.zip

%.zip:
	zip -j bin/$@ bin/bootstrap

stage-%: build
	serverless deploy --stage $(*)

output:
	mkdir -p ./bin

clean:
	$(RM) -rf ./bin

bin/%:
	$(MAKE) -C cmd/$(*) $(@)

undeploy:
	serverless remove

deploy: build
	serverless deploy
