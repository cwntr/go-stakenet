# go-stakenet

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelmota/cwntr/go-crypto-tools/LICENSE.md)
[![Build Status](https://travis-ci.org/cwntr/go-stakenet.svg?branch=master)](https://travis-ci.org/cwntr/go-stakenet)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/b7b3a89480af4de797166377948137ef)](https://www.codacy.com/app/cwntr/go-stakenet?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cwntr/go-stakenet&amp;utm_campaign=Badge_Grade)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](#contributing)


A collection of different clients for Stakenet (XSN) techologies. Find out more about Stakenet: [https://stakenet.io/](https://stakenet.io)

You can see all the clients active and in use on [https://stakenet.info](https://stakenet.info) 

## Installation

requirement: go v1.11+

``` > go get -u github.com/cwntr/go-stakenet```

## Stakenet (XSN) Clients implemented

| What | Name                 | Reference                                                                                |  Date    |
|------|----------------------|------------------------------------------------------------------------------------------|---------|
| CLI  | XSN Core Wallet      | [github.com/X9Developers/XSN](https://github.com/X9Developers/XSN)                       | 2020-01 |
| API  | XSN Block Explorer   | [github.com/X9Developers/block-explorer](https://github.com/X9Developers/block-explorer) | 2020-01 |
| CLI  | XSN Lightning Wallet | [github.com/lightningnetwork](https://github.com/lightningnetwork)                       | 2020-01 |


## Usage Explorer Client

```
import (
	"fmt"

	"github.com/cwntr/go-stakenet/explorer"
	"github.com/cwntr/go-stakenet/tools"
)

func testExplorer() {
	// no parameter will do on-fly-requests and responses without caching
	e := explorer.NewXSNExplorerAPIClient(nil)
	stats, err := e.GetStats()
	if err != nil {
		return
	}
	fmt.Printf("stats: %v\n", stats)



	// with recorder pointer parameter will locally store request and response pairs. This should only be used for responses
	// that will not change. e.g. get all details of a block, since a block is not gonna change.
	recorderPath := "records/xsn_block/%s"
	blockHash := "bf069bd8e1ce427c3dd7adf1aacc907051536210351bb8abcc76325486bce61d"
	blockRec, err := tools.CreateRecorder(fmt.Sprintf(recorderPath, blockHash))
	if err != nil {
		return
	}

	e2 := explorer.NewXSNExplorerAPIClient(blockRec)
	blockData, err := e2.GetBlockByQuery(blockHash)
	blockRec.Stop() //flush

	fmt.Printf("blockData: %v\n", blockData)
}
```
