package lnd

import (
	"fmt"

	"github.com/cwntr/xsn-tech/pkg/common"
)

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
