journalctl += -u kubelet.service
systemctl += kubelet.service
paths += /etc/kubelet /var/lib/kubelet /etc/systemd/system/kubelet.service
