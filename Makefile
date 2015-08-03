all: test

test:
	go test -v ./...

fuzz:
	go-fuzz -bin=./nid-fuzz.zip -workdir=workdir

fuzz-build:
	go-fuzz-build github.com/TV4/nid

fuzz-clean:
	rm -rf nid-fuzz.zip workdir
