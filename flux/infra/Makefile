.PHONY: all clean

all:
	make -C external-secrets
	make -C flux
	make -C kiali
	make -C metallb

clean:
	make -C external-secrets clean
	make -C flux clean
	make -C kiali clean
	make -C metallb clean
