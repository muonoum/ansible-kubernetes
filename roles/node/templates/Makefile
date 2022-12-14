include /root/Makefile.d/*.mk

all:

stop:
	systemctl daemon-reload
	systemctl disable $(systemctl)
	systemctl stop $(systemctl)

restart:
	systemctl daemon-reload
	systemctl restart $(systemctl)

journal:
	journalctl -f $(journalctl)

clean: stop $(clean)
	rm -rf $(paths)

dist-clean: clean
	rm /root/Makefile
	rm -rf /root/Makefile.d
