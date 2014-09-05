SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

all: embed build

deps:
	go get github.com/astaxie/beego

embed:

build:

test:

install:

clean:

dpkg:

run:
