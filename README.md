# kube

    export cluster=kyuubee
    export vault="-address https://vault:8200 -ca-cert setup/root.crt"

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

    for n in {1..3}; do
        vault operator unseal $vault $(op read op://$cluster/vault/unseal${n})
    done

    vault login $vault $(op read op://$cluster/vault/token)

    vault kv put $vault kv/flux-system/repo \
        identity="$(op read 'op://$cluster/ssh/private key')" \
        identity.pub="$(op read 'op://$cluster/ssh/public key')" \
        known_hosts="$(ssh-keyscan -t ecdsa github.com)"

    kubectl create ns external-secrets

    kubectl -n external-secrets create secret generic vault-token \
        --from-literal=token=$(op read op://$cluster/vault/token)

    kubectl apply -k flux/infra/external-secrets/crds
    kubectl apply -k flux/clusters/$cluster/infra/external-secrets

## istio

    istioctl install --set profile=minimal

## flux

    kubectl apply -k flux/infra/flux/crds
    kubectl apply -k flux/infra/flux/system

## cluster

    kubectl apply -k flux/clusters/$cluster
