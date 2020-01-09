test/explorer:
	go test ./explorer/...

test/wallet:
	go test ./wallet/...

test/client:
	go test ./client/...

.PHONY: test
test: test/client test/wallet