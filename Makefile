#go parameters
GOCMD=CGO_ENABLED=0 vgo
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

all:
	test build

build:
	$(GOBUILD)

test:
	$(GOTEST) -v

clean:
	$(GOCLEAN)
