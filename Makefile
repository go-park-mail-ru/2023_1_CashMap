all: build check test

build:
	./scripts/build.sh

check: build
	./scripts/linters.sh

test: build
	./scripts/test.sh