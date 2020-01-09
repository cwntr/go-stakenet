package main

import (
	"github.com/cwntr/go-stakenet/client"
	"github.com/cwntr/go-stakenet/explorer/xsn"
	walletxsn "github.com/cwntr/go-stakenet/wallet/xsn"
)

func main() {
	//guarantee export
	_ = xsn.API{}
	_ = client.CMCInfoMap{}
	_ = walletxsn.MasternodeItem{}
}