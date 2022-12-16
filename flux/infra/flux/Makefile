.PHONY: crds system clean

junkbox := ../../../library/junkbox

crds system:
	make -C ../../..

	flux install --export | \
				$(junkbox) split-yaml --overwrite \
					--default 'system/{{.Kind}}--{{.Metadata.Name}}.yaml' \
					--kind 'CustomResourceDefinition=crds/{{.Metadata.Name}}.yaml' -

clean:
	rm -rf system
	rm -rf crds
