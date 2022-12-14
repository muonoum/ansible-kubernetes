junkbox := $(shell find apps/junkbox -type f -name '*.go')

library/junkbox: apps/junkbox/go.mod apps/junkbox/go.sum $(junkbox)
	cd apps/junkbox && go build -v -o ../../$@ ./cmd/junkbox
