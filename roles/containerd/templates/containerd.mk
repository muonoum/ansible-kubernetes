journalctl += -u containerd.service
systemctl += containerd.service
paths += /etc/containerd /etc/systemd/system/containerd.service

pods:
	@crictl pods
