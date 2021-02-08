package lnd

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestDescribePath(t *testing.T) {
	d := TestDataDescribePath()
	data, err := UnmarshalDescribePath(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	for _, node := range data.Nodes {
		for _, addr := range node.Addresses {
			fmt.Printf("network: %s | addr: %s | node: %s | alias: %s \n", addr.Network, addr.Addr, node.PubKey, node.Alias)
		}
	}

	for _, edge := range data.Edges {
		fmt.Printf("edge: channelID: %s\n", edge.ChannelID)
	}
}

func TestListPeers(t *testing.T) {
	d := TestDataListPeers()
	data, err := UnmarshalListPeers(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	for _, peer := range data.Peers {
		fmt.Printf("peer: %v \n", peer)
	}
}

func TestNetworkInfo(t *testing.T) {
	d := TestDataGetNetworkInfo()
	data, err := UnmarshalNetworkInfo(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	fmt.Printf("network data: %v \n", data)
}

func TestGetInfo(t *testing.T) {
	d := TestDataGetInfo()
	data, err := UnmarshalGetInfo(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	fmt.Printf("get info : %v \n", data)
}

func TestGetNodeInfo(t *testing.T) {
	d := TestDataGetNodeInfo()
	data, err := UnmarshalGetNodeInfo(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	fmt.Printf("get node info : %v \n", data)
}

func TestGetChannelInfo(t *testing.T) {
	d := TestDataGetChannelInfo()
	data, err := UnmarshalGetChannelInfo(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	fmt.Printf("get channel info : %v \n", data)
}

func TestWalletBalance(t *testing.T) {
	d := TestDataGetWalletBalance()
	data, err := UnmarshalWalletBalance(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	fmt.Printf("get channel info : %v \n", data)
}

func TestGetListChannels(t *testing.T) {
	d := TestDataListChannels()
	data, err := UnmarshalListChannels(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	fmt.Printf("get channel info : %v \n", data)
}

func TestUnmarshalListChainTxns(t *testing.T) {
	d := TestListChainTxns()
	data, err := UnmarshalListChainTxns(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	fmt.Printf("get chain txns : %v \n", data)
}

func TestUnmarshalGetPendingChannel(t *testing.T) {
	d := TestDataPendingChannels()
	data, err := UnmarshalPendingChannels(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	fmt.Printf("pending channels : %v \n", data)
}

func TestUnmarshalListInvoices(t *testing.T) {
	d := TestInvoices()
	data, err := UnmarshalListInvoices(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	gracePeriod := 1
	graceTs := time.Now().Add(time.Duration(-gracePeriod) * time.Second)
	for i, c := range data.Invoices {
		for _, h := range c.Htlcs {
			ts, _ := strconv.ParseInt(c.CreationDate, 10, 64)
			invoiceTs := time.Unix(ts, 0)
			fmt.Printf("invoices#%d: chanId: %s | created-at: %s | amt: %s sats\n", i , h.ChanID, time.Unix(ts, 0).Format(time.RFC3339), c.Value)
			if invoiceTs.Before(graceTs) {
				fmt.Printf("invoices#%d: can cancel with hash: %s \n", i, c.RHash)
			}
		}
	}
}
