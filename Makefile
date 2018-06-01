.DEFAULT_GOAL := all

clean:
	rm -f tipjar.zip
	rm -f tipjar

build:
	GOOS=linux go build

zip:
	zip tipjar.zip tipjar

all:
	$(MAKE) clean
	$(MAKE) build
	$(MAKE) zip

