package lnd

import (
	"fmt"
	"testing"
)

func TestDescribePath(t *testing.T) {
	d := TestDataDescribePath()
	data, err := UnmarshalDescribePath(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
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
		return
	}
	fmt.Printf("network data: %v \n", data)
}

func TestGetInfo(t *testing.T) {
	d := TestDataGetInfo()
	data, err := UnmarshalGetInfo(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		return
	}
	fmt.Printf("get info : %v \n", data)
}

func TestGetNodeInfo(t *testing.T) {
	d := TestDataGetNodeInfo()
	data, err := UnmarshalGetNodeInfo(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		return
	}
	fmt.Printf("get node info : %v \n", data)
}


func TestGetChannelInfo(t *testing.T) {
	d := TestDataGetChannelInfo()
	data, err := UnmarshalGetChannelInfo(d)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		return
	}
	fmt.Printf("get channel info : %v \n", data)
}
