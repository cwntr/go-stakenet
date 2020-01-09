test/explorer:
	go test ./core/...

test/wallet:
	go test ./explorer/...

test/lnd:
	go test ./lnd/...

.PHONY: test
test: test/explorer test/wallet test/lnd