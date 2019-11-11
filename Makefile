GOCMD=go
GOBUILD=$(GOCMD) build

OUTFILE=build/penunse

all: build
build:
	$(GOBUILD) -o $(OUTFILE)
