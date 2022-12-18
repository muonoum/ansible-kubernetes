.PHONY: clean

junkbox := ../../../../library/junkbox
version := v0.6.2
url := https://github.com/kubernetes-sigs/metrics-server/releases/download/$(version)/components.yaml

$(version):
	make -C ../../../..

	curl -Ls $(url) | \
		$(junkbox) split-yaml --overwrite \
			--default "$(version)/{{.Kind}}--{{.Metadata.Name}}.yaml" -

clean:
	rm -rf $(version)
