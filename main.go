package main

import (
	"github.com/cwntr/go-stakenet/core"
	"github.com/cwntr/go-stakenet/explorer"
	"github.com/cwntr/go-stakenet/lnd"
)

func main() {
	//guarantee export
	_ = core.MasternodeItem{}
	_ = explorer.RawBlock{}
	_ = lnd.Nodes{}
}
