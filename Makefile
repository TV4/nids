all: test

test:
	go test -v ./...

fuzz:
	go-fuzz -bin=./nids-fuzz.zip -workdir=workdir

fuzz-build:
	go-fuzz-build github.com/TV4/nids

fuzz-clean:
	rm -rf nids-fuzz.zip workdir
