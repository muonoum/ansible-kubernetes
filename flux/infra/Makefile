.PHONY: all clean

all:
	make -C external-secrets
	make -C flux
	make -C gateway-api
	make -C metallb
	make -C prometheus

clean:
	make -C external-secrets clean
	make -C flux clean
	make -C gateway-api clean
	make -C metallb clean
	make -C prometheus clean
