.PHONY: crds operator clean

junkbox := ../../../library/junkbox
version := v0.61.1
base_url := https://github.com/prometheus-operator/prometheus-operator/releases/download

crds operator:
	make -C ../../..

	curl -s -L $(base_url)/$(version)/bundle.yaml | \
		$(junkbox) split-yaml --overwrite \
			--default 'operator/{{.Kind}}--{{.Metadata.Name}}.yaml' \
			--kind 'CustomResourceDefinition=crds/{{.Metadata.Name}}.yaml' -

clean:
	rm -rf crds
	rm -rf operator
