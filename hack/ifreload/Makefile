CFLAGS=-I/usr/include/libnl3 -O3 -Wall -W
LDLIBS=-lnl-3 -lnl-route-3 -s

ifreload:

ifreload.rhel8: build
	docker build -f Dockerfile.ifreload -t ifreload .

build: tools
	docker build --no-cache -v $(PWD):/build -f Dockerfile.build .

tools:
	docker build -f Dockerfile.tools -t tools .

clean:
	rm -f ifreload

.PHONY: ifreload.rhel8 build tools clean
