journalctl += -u kube-proxy.service
systemctl += kube-proxy.service
paths += /etc/kube-proxy /var/lib/kube-proxy /etc/systemd/system/kube-proxy.service
