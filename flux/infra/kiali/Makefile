junkbox := ../../../library/junkbox

all: operator server

operator: $(junkbox)
	make -C ../../..

	KUBECONFIG=../../../private/admin.kubeconfig \
		helm template kiali-operator kiali-operator \
		  --repo https://kiali.org/helm-charts \
		  --namespace kiali-operator --create-namespace \
		  --set cr.create=true \
		  --set cr.namespace=istio-system | \
				$(junkbox) split-yaml --overwrite \
					--default 'operator/{{.Kind}}--{{.Metadata.Name}}.yaml' \
					--kind 'CustomResourceDefinition=crds/{{.Metadata.Name}}.yaml' -

server: $(junkbox)
	make -C ../../..

	KUBECONFIG=../../../private/admin.kubeconfig \
		helm template kiali-server kiali-server \
		  --repo https://kiali.org/helm-charts \
		  --namespace istio-system --create-namespace \
		  --set auth.strategy=anonymous | \
				$(junkbox) split-yaml --overwrite \
					--default 'server/{{.Kind}}--{{.Metadata.Name}}.yaml' \
					--kind 'CustomResourceDefinition=crds/{{.Metadata.Name}}.yaml' -

clean:
	rm -rf operator
	rm -rf server
