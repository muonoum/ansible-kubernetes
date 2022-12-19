dashboard := parts/dashboard/gren.json $(shell find parts/dashboard/src)
server := parts/server/go.mod parts/server/go.sum \
	$(shell find parts/server -name *.go)
embed := parts/server/cmd/dashboard/embed
bundler := parts/bundler/go.mod parts/bundler/go.sum \
	$(shell find parts/bundler -name *.go)

all: build/dashboard build/reader

build:
	mkdir -p build

build/bundler: Makefile $(bundler) | build
	cd parts/bundler && go build -v -o ../../$@ ./cmd/bundler

build/dashboard: Makefile $(embed) $(server) | build
	cd parts/server && \
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
	go build -ldflags '-w' -v -o ../../$@ ./cmd/dashboard

build/reader: Makefile $(server) | build
	cd parts/server && \
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
	go build -ldflags '-w' -v -o ../../$@ ./cmd/reader

$(embed): Makefile build/bundler $(dashboard) | build
	rm -rf $@ && mkdir -p $@
	cp parts/dashboard/src/index.html $@
	cd parts/dashboard && ../../build/bundler \
	--output ../../$@ src/app.css src/app.js

clean:
	-rm -rf build
	-rm -rf $(embed)

dist-clean: clean
	-rm -rf parts/dashboard/.gren

.PHONY: run-reader
run-reader:
	cd parts/server && \
	go run -v ./cmd/reader \
	--ca-cert=../../../../setup/root.crt start \
	--log-format=console --address=127.0.0.1:3210 \
	--source-url="https://$(LB)/?stats;csv"

.PHONY: run-dashboard
run-dashboard: $(embed)
	cd parts/server && \
	go run -v ./cmd/dashboard \
	--log-format=console --address=:3310 \
	--source=http://localhost:3210/

.PHONY: reader-image
reader-image:
	podman build -f Containerfile.reader -t reader

.PHONY: dashboard-image
dashboard-image:
	podman build -f Containerfile.dashboard -t dashboard

.PHONY: images
images: reader-image dashboard-image

reader_tag := reader-9f6bbed6-ed83-44bf-a823-21e825559d75
dashboard_tag := dashboard-d89375e4-de80-4287-a04c-739f148e07d6

.PHONY: push
push: images
	podman tag localhost/reader ttl.sh/$(reader_tag)
	podman push ttl.sh/$(reader_tag)
	podman tag localhost/dashboard ttl.sh/$(dashboard_tag)
	podman push ttl.sh/$(dashboard_tag)

