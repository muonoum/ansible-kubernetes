journalctl += -u etcd.service
systemctl += etcd.service
paths += /etc/etcd /var/lib/etcd /etc/systemd/system/etcd.service

etcd-members:
	@etcdctl --cacert /etc/etcd/root.crt \
			--cert /etc/etcd/peer.crt \
			--key /etc/etcd/peer.key \
			--endpoints=https://etcd:2379 \
		member list
