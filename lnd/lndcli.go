package lnd

import (
	"fmt"

	"github.com/cwntr/go-stakenet/common"
)

func GetWalletBalance(cliPath string, isTest bool, options ...string) (dp WalletBalance, err error) {
	str := ""
	if isTest {
		str = TestDataGetWalletBalance()
	} else {
		options = append(options, "walletbalance")
		str, err = common.ExecCLI(cliPath, options...)
		if err != nil {
			fmt.Printf("err: %v \n", err)
			return
		}
	}
	dp, err = UnmarshalWalletBalance(str)
	return
}

func GetListChannels(cliPath string, isTest bool, options ...string) (dp ListChannels, err error) {
	str := ""
	if isTest {
		str = TestDataListChannels()
	} else {
		options = append(options, "listchannels")
		str, err = common.ExecCLI(cliPath, options...)
		if err != nil {
			fmt.Printf("err: %v \n", err)
			return
		}
	}
	dp, err = UnmarshalListChannels(str)
	return
}

func GetDescribePath(cliPath string, isTest bool, options ...string) (dp DescribePath, err error) {
	str := ""
	if isTest {
		str = TestDataDescribePath()
	} else {
		options = append(options, "describegraph")
		str, err = common.ExecCLI(cliPath, options...)
		if err != nil {
			fmt.Printf("err: %v \n", err)
			return
		}
	}
	dp, err = UnmarshalDescribePath(str)
	return
}

func GetListPeers(cliPath string, isTest bool, options ...string) (dp ListPeers, err error) {
	str := ""
	if isTest {
		str = TestDataListPeers()
	} else {
		options = append(options, "listpeers")
		str, err = common.ExecCLI(cliPath, options...)
		if err != nil {
			return
		}
	}
	dp, err = UnmarshalListPeers(str)
	return
}

func GetNetworkInfo(cliPath string, isTest bool, options ...string) (ni NetworkInfo, err error) {
	str := ""
	if isTest {
		str = TestDataGetNetworkInfo()
	} else {
		options = append(options, "getnetworkinfo")
		str, err = common.ExecCLI(cliPath, options...)
		if err != nil {
			return
		}
	}
	ni, err = UnmarshalNetworkInfo(str)
	return
}

func GetGetInfo(cliPath string, isTest bool, options ...string) (gi GetInfo, err error) {
	str := ""
	if isTest {
		str = TestDataGetInfo()
	} else {
		options = append(options, "getinfo")
		str, err = common.ExecCLI(cliPath, options...)
		if err != nil {
			return
		}
	}
	gi, err = UnmarshalGetInfo(str)
	return
}

func GetGetNodeInfo(cliPath string, pubKey string, isTest bool, options ...string) (gni GetNodeInfo, err error) {
	str := ""
	if isTest {
		str = TestDataGetNodeInfo()
	} else {
		options = append(options, "getnodeinfo")
		options = append(options, pubKey)
		str, err = common.ExecCLI(cliPath, options...)
		if err != nil {
			return
		}
	}
	gni, err = UnmarshalGetNodeInfo(str)
	return
}

func GetChannelInfo(cliPath string, channelId string, isTest bool, options ...string) (gni Edge, err error) {
	str := ""
	if isTest {
		str = TestDataGetChannelInfo()
	} else {
		options = append(options, "getchaninfo")
		options = append(options, channelId)
		str, err = common.ExecCLI(cliPath, options...)
		if err != nil {
			return
		}
	}
	gni, err = UnmarshalGetChannelInfo(str)
	return
}

func GetListChainTxns(cliPath string, isTest bool, options ...string) (lct ListChainTxns, err error) {
	str := ""
	if isTest {
		str = TestListChainTxns()
	} else {
		options = append(options, "listchaintxns")
		str, err = common.ExecCLI(cliPath, options...)
		if err != nil {
			return
		}
	}
	lct, err = UnmarshalListChainTxns(str)
	return
}

func GetPendingChannels(cliPath string, isTest bool, options ...string) (lct PendingChannels, err error) {
	str := ""
	if isTest {
		str = TestDataPendingChannels()
	} else {
		options = append(options, "pendingchannels")
		str, err = common.ExecCLI(cliPath, options...)
		if err != nil {
			return
		}
	}
	lct, err = UnmarshalPendingChannels(str)
	return
}

func GetListPendingInvoices(cliPath string, isTest bool, options ...string) (lct ListInvoices, err error) {
	str := ""
	if isTest {
		str = TestDataPendingChannels()
	} else {
		options = append(options, "listinvoices")
		options = append(options, "--pending_only")
		str, err = common.ExecCLI(cliPath, options...)
		if err != nil {
			return
		}
	}
	lct, err = UnmarshalListInvoices(str)
	return
}

//Pass payment hash to cancel invoice
func CancelInvoice(cliPath string, rhash string, options ...string) (err error) {
	options = append(options, "cancelinvoice")
	options = append(options, rhash)
	_ , err = common.ExecCLI(cliPath, options...)
	if err != nil {
		return
	}
	return
}