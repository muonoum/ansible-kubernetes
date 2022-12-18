.PHONY: all clean dist-clean

all:
	make -C metrics-server
	make -C apiserver-network-proxy

clean:
	make -C metrics-server clean
	make -C apiserver-network-proxy clean

dist-clean:
	make -C metrics-server dist-clean
	make -C apiserver-network-proxy dist-clean
