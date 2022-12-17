junkbox := ../../../library/junkbox
version := v0.61.1
base_url := https://github.com/prometheus-operator/prometheus-operator/releases/download

.PHONY: crds operator
crds operator:
	make -C ../../..

	curl -s -L $(base_url)/$(version)/bundle.yaml | \
		$(junkbox) split-yaml --overwrite \
			--default 'operator/{{.Kind}}--{{.Metadata.Name}}.yaml' \
			--kind 'CustomResourceDefinition=crds/{{.Metadata.Name}}.yaml' -

.PHONY: clean
clean:
	rm -rf crds
	rm -rf operator
