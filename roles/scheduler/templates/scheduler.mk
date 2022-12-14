journalctl += -u kube-scheduler.service
systemctl += kube-scheduler.service
paths += /etc/kube-scheduler /var/lib/kube-scheduler /etc/systemd/system/kube-scheduler.service
