# kube

## deploy

    make -C vendor

    ansible-playbook deploy.yaml
    ansible-playbook admin.yaml

    kubectl apply -k flux/infra/kube-system/kube-router
    ansible-playbook cni.yaml

    kubectl apply -k flux/infra/kube-system/coredns
    kubectl apply -k flux/infra/kube-system/konnectivity
    kubectl apply -k flux/infra/kube-system/metrics-server

## external-secrets

    for KEY in $(seq 1 3); do
        vault operator unseal -address https://vault:8200 -ca-cert setup/root.crt \
            $(op read op://kyuubee/vault/unseal${KEY})
    done

    vault login -address https://vault:8200 -ca-cert setup/root.crt \
        $(op read op://kyuubee/vault/token)

    vault kv put -address https://vault:8200 -ca-cert setup/root.crt \
        kv/flux-system/repo \
        identity="$(op read 'op://kyuubee/ssh/private key')" \
        identity.pub="$(op read 'op://kyuubee/ssh/public key')" \
        known_hosts="$(ssh-keyscan -t ecdsa github.com)"

    kubectl create ns external-secrets

    kubectl -n external-secrets create secret generic vault-token \
        --from-literal=token=$(op read op://kyuubee/vault/token)

    kubectl apply -k flux/infra/external-secrets/crds
    kubectl apply -k flux/clusters/kyuubee/infra/external-secrets

## istio

    istioctl install --set profile=minimal

## flux

    kubectl apply -k flux/infra/flux/crds
    kubectl apply -k flux/infra/flux/system
    kubectl apply -f flux/clusters/kyuubee/cluster-repo-secret.yaml
    kubectl apply -f flux/clusters/kyuubee/cluster-repo.yaml
    kubectl apply -f flux/clusters/kyuubee/cluster.yaml
