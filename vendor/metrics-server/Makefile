.PHONY: all clean dist-clean

arch := arm64
metrics_server_hash = $(shell cd metrics-server 2>/dev/null && git rev-parse --short HEAD || echo "unknown")
metrics_server_prefix := ../../roles/metrics-server/files/metrics-server
metrics_server = $(metrics_server_prefix)-$(metrics_server_hash)-$(arch)
targets = $(metrics_server)

all: $(targets)

$(metrics_server): metrics-server
	@echo $(shell basename $(metrics_server))
	rm -f metrics-server/metrics-server
	ARCH=$(arch) GOOS=linux GOARCH=$(arch) make -C metrics-server metrics-server
	mv metrics-server/metrics-server $(metrics_server)

metrics-server:
	git clone https://github.com/kubernetes-sigs/metrics-server.git

clean:
	rm -f $(targets)

dist-clean: clean
	rm -rf metrics-server


