journalctl += -u kube-controller-manager.service
systemctl += kube-controller-manager.service
paths += /etc/kube-controller-manager /var/lib/kube-controller-manager /etc/systemd/system/kube-controller-manager.service
