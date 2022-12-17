.PHONY: clean dist-clean

arch := arm64
hash = $(shell cd apiserver-network-proxy && git rev-parse --short HEAD)
prefix := ../roles/konnectivity-server/files/konnectivity-server
proxy_server = $(prefix)-$(hash)-$(arch)

$(proxy_server): apiserver-network-proxy
	@echo "Building $(shell basename $(proxy_server))"
	rm -f apiserver-network-proxy/bin/proxy-server
	arch=$(arch) GOOS=linux GOARCH=$(arch) make -C apiserver-network-proxy bin/proxy-server
	mv apiserver-network-proxy/bin/proxy-server $(proxy_server)

clean:
	rm -f $(proxy_server)

dist-clean: clean
	rm -rf apiserver-network-proxy

apiserver-network-proxy:
	git clone https://github.com/kubernetes-sigs/apiserver-network-proxy.git
