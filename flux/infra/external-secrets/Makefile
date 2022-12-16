.PHONY: crds operator clean

junkbox := ../../../library/junkbox

crds operator:
	make -C ../../..

	KUBECONFIG=../../../private/admin.kubeconfig \
		helm template external-secrets external-secrets \
			--repo https://charts.external-secrets.io \
			--namespace external-secrets --create-namespace \
			--set installCRDs=true | \
				$(junkbox) split-yaml --overwrite \
					--default 'operator/{{.Kind}}--{{.Metadata.Name}}.yaml' \
					--kind 'CustomResourceDefinition=crds/{{.Metadata.Name}}.yaml' -

clean:
	rm -rf operator
	rm -rf crds
