all: run

build:
	go-fuzz-build github.com/TV4/nid

clean:
	rm -rf nid-fuzz.zip workdir

run:
	go-fuzz -bin=./nid-fuzz.zip -workdir=workdir
