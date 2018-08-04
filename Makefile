#go parameters
GOCMD=go
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
