.PHONY: all clean dist-clean

arch := arm64

proxy_hash = $(shell cd apiserver-network-proxy 2>/dev/null && git rev-parse --short HEAD || echo "unknown")
proxy_server_prefix := ../../roles/konnectivity-server/files/konnectivity-server
proxy_agent_prefix := ../../roles/konnectivity-agent/files/konnectivity-agent
proxy_server = $(proxy_server_prefix)-$(proxy_hash)-$(arch)
proxy_agent = $(proxy_agent_prefix)-$(proxy_hash)-$(arch)
targets = $(proxy_server) $(proxy_agent)

all: $(targets)

$(proxy_server): apiserver-network-proxy
	@echo $(shell basename $(proxy_server))
	rm -f apiserver-network-proxy/bin/proxy-server
	arch=$(arch) GOOS=linux GOARCH=$(arch) make -C apiserver-network-proxy bin/proxy-server
	mv apiserver-network-proxy/bin/proxy-server $(proxy_server)

$(proxy_agent): apiserver-network-proxy
	@echo $(shell basename $(proxy_agent))
	rm -f apiserver-network-proxy/bin/proxy-agent
	arch=$(arch) GOOS=linux GOARCH=$(arch) make -C apiserver-network-proxy bin/proxy-agent
	mv apiserver-network-proxy/bin/proxy-agent $(proxy_agent)

apiserver-network-proxy:
	git clone https://github.com/kubernetes-sigs/apiserver-network-proxy.git

clean:
	rm -f $(targets)

dist-clean: clean
	rm -rf apiserver-network-proxy


