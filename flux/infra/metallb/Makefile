.PHONY: operator/crds operator/system standalone/crds standalone/system clean

junkbox := ../../../library/junkbox

base_url := https://raw.githubusercontent.com/metallb
operator_version := v0.13.9
standalone_version := v0.13.7

all: operator/crds operator/system standalone/crds standalone/system

operator/crds operator/system:
	make -C ../../..

	curl -sL $(base_url)/metallb-operator/$(operator_version)/bin/metallb-operator.yaml | \
			$(junkbox) split-yaml --overwrite \
				--default 'operator/operator/{{.Kind}}--{{.Metadata.Name}}.yaml' \
				--kind 'CustomResourceDefinition=operator/crds/{{.Metadata.Name}}.yaml' -

standalone/crds standalone/system:
	make -C ../../..

	curl -sL $(base_url)/metallb/$(standalone_version)/config/manifests/metallb-native.yaml | \
			$(junkbox) split-yaml --overwrite \
				--default 'standalone/system/{{.Kind}}--{{.Metadata.Name}}.yaml' \
				--kind 'CustomResourceDefinition=standalone/crds/{{.Metadata.Name}}.yaml' -

clean:
	rm -rf standalone
	rm -rf operator
