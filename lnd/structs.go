package lnd

import "encoding/json"

const (
	AddressNetwork = "tcp"
)

type Nodes struct {
	LastUpdate int    `json:"last_update"`
	PubKey     string `json:"pub_key"`
	Alias      string `json:"alias"`
	Addresses  []struct {
		Network string `json:"network"`
		Addr    string `json:"addr"`
	} `json:"addresses"`
	Color string `json:"color"`
}

type Edge struct {
	ChannelID   string `json:"channel_id"`
	ChanPoint   string `json:"chan_point"`
	LastUpdate  int    `json:"last_update"`
	Node1Pub    string `json:"node1_pub"`
	Node2Pub    string `json:"node2_pub"`
	Capacity    string `json:"capacity"`
	Node1Policy struct {
		TimeLockDelta    int    `json:"time_lock_delta"`
		MinHtlc          string `json:"min_htlc"`
		FeeBaseMsat      string `json:"fee_base_msat"`
		FeeRateMilliMsat string `json:"fee_rate_milli_msat"`
		Disabled         bool   `json:"disabled"`
		MaxHtlcMsat      string `json:"max_htlc_msat"`
	} `json:"node1_policy"`
	Node2Policy struct {
		TimeLockDelta    int    `json:"time_lock_delta"`
		MinHtlc          string `json:"min_htlc"`
		FeeBaseMsat      string `json:"fee_base_msat"`
		FeeRateMilliMsat string `json:"fee_rate_milli_msat"`
		Disabled         bool   `json:"disabled"`
		MaxHtlcMsat      string `json:"max_htlc_msat"`
	} `json:"node2_policy"`
}

type DescribePath struct {
	Nodes []Nodes `json:"nodes"`
	Edges []Edge  `json:"edges"`
}

func UnmarshalDescribePath(str string) (dp DescribePath, err error) {
	err = json.Unmarshal([]byte(str), &dp)
	return
}

type ListPeer struct {
	PubKey    string `json:"pub_key"`
	Address   string `json:"address"`
	BytesSent string `json:"bytes_sent"`
	BytesRecv string `json:"bytes_recv"`
	SatSent   string `json:"sat_sent"`
	SatRecv   string `json:"sat_recv"`
	Inbound   bool   `json:"inbound"`
	PingTime  string `json:"ping_time"`
	SyncType  string `json:"sync_type"`
}

type ListPeers struct {
	Peers []ListPeer `json:"peers"`
}

func UnmarshalListPeers(str string) (lp ListPeers, err error) {
	err = json.Unmarshal([]byte(str), &lp)
	return
}

type NetworkInfo struct {
	GraphDiameter        int     `json:"graph_diameter"`
	AvgOutDegree         float64 `json:"avg_out_degree"`
	MaxOutDegree         int     `json:"max_out_degree"`
	NumNodes             int     `json:"num_nodes"`
	NumChannels          int     `json:"num_channels"`
	TotalNetworkCapacity string  `json:"total_network_capacity"`
	AvgChannelSize       float64 `json:"avg_channel_size"`
	MinChannelSize       string  `json:"min_channel_size"`
	MaxChannelSize       string  `json:"max_channel_size"`
	MedianChannelSizeSat string  `json:"median_channel_size_sat"`
}

func UnmarshalNetworkInfo(str string) (ni NetworkInfo, err error) {
	err = json.Unmarshal([]byte(str), &ni)
	return
}

type GetInfo struct {
	Version             string `json:"version"`
	IdentityPubkey      string `json:"identity_pubkey"`
	Alias               string `json:"alias"`
	NumPendingChannels  int    `json:"num_pending_channels"`
	NumActiveChannels   int    `json:"num_active_channels"`
	NumInactiveChannels int    `json:"num_inactive_channels"`
	NumPeers            int    `json:"num_peers"`
	BlockHeight         int    `json:"block_height"`
	BlockHash           string `json:"block_hash"`
	BestHeaderTimestamp int    `json:"best_header_timestamp"`
	SyncedToChain       bool   `json:"synced_to_chain"`
	Testnet             bool   `json:"testnet"`
	Chains              []struct {
		Chain   string `json:"chain"`
		Network string `json:"network"`
	} `json:"chains"`
	Uris interface{} `json:"uris"`
}

func UnmarshalGetInfo(str string) (gi GetInfo, err error) {
	err = json.Unmarshal([]byte(str), &gi)
	return
}

type GetNodeInfo struct {
	Node struct {
		LastUpdate int    `json:"last_update"`
		PubKey     string `json:"pub_key"`
		Alias      string `json:"alias"`
		Addresses  []struct {
			Network string `json:"network"`
			Addr    string `json:"addr"`
		} `json:"addresses"`
		Color string `json:"color"`
	} `json:"node"`
	NumChannels   int    `json:"num_channels"`
	TotalCapacity string `json:"total_capacity"`
}

func UnmarshalGetNodeInfo(str string) (gni GetNodeInfo, err error) {
	err = json.Unmarshal([]byte(str), &gni)
	return
}

func UnmarshalGetChannelInfo(str string) (gni Edge, err error) {
	err = json.Unmarshal([]byte(str), &gni)
	return
}
