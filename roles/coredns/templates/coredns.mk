journalctl += -u coredns.service
systemctl += coredns.service
paths += /etc/coredns /var/lib/coredns /etc/systemd/system/coredns.service
