journalctl += -u kube-router.service
systemctl += kube-router.service
paths += /etc/kube-router /var/lib/kube-router /etc/systemd/system/kube-router.service
