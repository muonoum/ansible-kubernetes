journalctl += -u metrics-server.service
systemctl += metrics-server.service
paths += /etc/metrics-server /var/lib/metrics-server /etc/systemd/system/metrics-server.service
