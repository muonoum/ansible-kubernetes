.PHONY: crds system clean

junkbox := ../../../library/junkbox
version := v0.6.0-rc2
base_url := https://github.com/kubernetes-sigs/gateway-api/releases/download/

crds system:
	make -C ../../..

	curl -s -L $(base_url)/$(version)/standard-install.yaml | \
		$(junkbox) split-yaml --overwrite \
			--default 'system/{{.Kind}}--{{.Metadata.Name}}.yaml' \
			--kind 'CustomResourceDefinition=crds/{{.Metadata.Name}}.yaml' -

clean:
	rm -rf crds
	rm -rf system
