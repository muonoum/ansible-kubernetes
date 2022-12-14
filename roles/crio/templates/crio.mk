journalctl += -u crio.service
systemctl += crio.service
paths += /etc/crio /etc/systemd/system/crio.service

pods:
	@crictl pods
