journalctl += -u haproxy.service
systemctl += haproxy.service
paths += /etc/haproxy /var/lib/haproxy /etc/systemd/system/haproxy.service

.PHONY: reload
reload:
	echo reload |  socat /var/run/haproxy/master -

.PHONY: show-proc
show-proc:
	echo 'show proc' |  socat /var/run/haproxy/master -

.PHONY: show-peers
show-peers:
	echo 'show peers' |  socat /var/run/haproxy/admin -
