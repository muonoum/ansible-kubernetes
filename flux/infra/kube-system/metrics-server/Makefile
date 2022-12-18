.PHONY: clean

version := v0.6.2

base := https://github.com/kubernetes-sigs/metrics-server/releases/download
url := $(base)/$(version)/components.yaml
junkbox := ../../../../library/junkbox

$(version):
	make -C ../../../..

	curl -Ls $(url) | \
		$(junkbox) split-yaml --overwrite \
			--default "$(version)/{{.Kind}}--{{.Metadata.Name}}.yaml" -

clean:
	rm -rf $(version)
