.DEFAULT_GOAL := all

clean:
	rm -f tipjar.zip
	rm -f tipjar

build:
	GOOS=linux go build

pack:
	upx -1 tipjar

zip:
	zip tipjar.zip tipjar

all:
	$(MAKE) clean
	$(MAKE) build
	$(MAKE) pack
	$(MAKE) zip

upload:
	/usr/local/bin/aws --profile personal lambda update-function-code --function-name TipJar --zip-file fileb://$(shell pwd)/tipjar.zip

prep:
	go get

site:
	$(MAKE) clean
	$(MAKE) prep
	$(MAKE) build
	$(MAKE) pack
	mv tipjar functions/payment
