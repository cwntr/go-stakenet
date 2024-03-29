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
		LastUpdate       int    `json:"last_update"`
	} `json:"node1_policy"`
	Node2Policy struct {
		TimeLockDelta    int    `json:"time_lock_delta"`
		MinHtlc          string `json:"min_htlc"`
		FeeBaseMsat      string `json:"fee_base_msat"`
		FeeRateMilliMsat string `json:"fee_rate_milli_msat"`
		Disabled         bool   `json:"disabled"`
		MaxHtlcMsat      string `json:"max_htlc_msat"`
		LastUpdate       int    `json:"last_update"`
	} `json:"node2_policy"`
}

type DescribePath struct {
	Nodes []Nodes `json:"nodes"`
	Edges []Edge  `json:"edges"`
}

type ListChainTxns struct {
	Transactions []struct {
		TxHash           string   `json:"tx_hash"`
		Amount           string   `json:"amount"`
		NumConfirmations int      `json:"num_confirmations"`
		BlockHash        string   `json:"block_hash"`
		BlockHeight      int      `json:"block_height"`
		TimeStamp        string   `json:"time_stamp"`
		TotalFees        string   `json:"total_fees"`
		DestAddresses    []string `json:"dest_addresses"`
		RawTxHex         string   `json:"raw_tx_hex"`
	} `json:"transactions"`
}

type PendingChannels struct {
	TotalLimboBalance           string        `json:"total_limbo_balance"`
	PendingOpenChannels         []interface{} `json:"pending_open_channels"`
	PendingClosingChannels      []interface{} `json:"pending_closing_channels"`
	PendingForceClosingChannels []interface{} `json:"pending_force_closing_channels"`
	WaitingCloseChannels        []interface{} `json:"waiting_close_channels"`
}

func UnmarshalDescribePath(str string) (dp DescribePath, err error) {
	err = json.Unmarshal([]byte(str), &dp)
	return
}

type WalletBalance struct {
	TotalBalance       string `json:"total_balance"`
	ConfirmedBalance   string `json:"confirmed_balance"`
	UnconfirmedBalance string `json:"unconfirmed_balance"`
}

func UnmarshalWalletBalance(str string) (dp WalletBalance, err error) {
	err = json.Unmarshal([]byte(str), &dp)
	return
}

type ListChannels struct {
	Channels []struct {
		Active                bool          `json:"active"`
		RemotePubkey          string        `json:"remote_pubkey"`
		ChannelPoint          string        `json:"channel_point"`
		ChanID                string        `json:"chan_id"`
		Capacity              string        `json:"capacity"`
		LocalBalance          string        `json:"local_balance"`
		RemoteBalance         string        `json:"remote_balance"`
		CommitFee             string        `json:"commit_fee"`
		CommitWeight          string        `json:"commit_weight"`
		FeePerKw              string        `json:"fee_per_kw"`
		UnsettledBalance      string        `json:"unsettled_balance"`
		TotalSatoshisSent     string        `json:"total_satoshis_sent"`
		TotalSatoshisReceived string        `json:"total_satoshis_received"`
		NumUpdates            string        `json:"num_updates"`
		PendingHtlcs          []interface{} `json:"pending_htlcs"`
		CsvDelay              int           `json:"csv_delay"`
		Private               bool          `json:"private"`
		Initiator             bool          `json:"initiator"`
		ChanStatusFlags       string        `json:"chan_status_flags"`
		LocalChanReserveSat   string        `json:"local_chan_reserve_sat"`
		RemoteChanReserveSat  string        `json:"remote_chan_reserve_sat"`
		StaticRemoteKey       bool          `json:"static_remote_key"`
		//CommitmentType        string        `json:"commitment_type"`
		Lifetime              string        `json:"lifetime"`
		Uptime                string        `json:"uptime"`
		CloseAddress          string        `json:"close_address"`
		PushAmountSat         string        `json:"push_amount_sat"`
		ThawHeight            int           `json:"thaw_height"`
	} `json:"channels"`
}

func UnmarshalListChannels(str string) (dp ListChannels, err error) {
	err = json.Unmarshal([]byte(str), &dp)
	return
}

type HTLC struct {
	ChanID        string `json:"chan_id"`
	HtlcIndex     string `json:"htlc_index"`
	AmtMsat       string `json:"amt_msat"`
	AcceptHeight  int    `json:"accept_height"`
	AcceptTime    string `json:"accept_time"`
	ResolveTime   string `json:"resolve_time"`
	ExpiryHeight  int    `json:"expiry_height"`
	State         string `json:"state"`
	CustomRecords struct {
	} `json:"custom_records"`
	MppTotalAmtMsat string `json:"mpp_total_amt_msat"`
}

type Invoice struct {
	Memo            string        `json:"memo"`
	RPreimage       interface{}   `json:"r_preimage"`
	RHash           string        `json:"r_hash"`
	Value           string        `json:"value"`
	ValueMsat       string        `json:"value_msat"`
	Settled         bool          `json:"settled"`
	CreationDate    string        `json:"creation_date"`
	SettleDate      string        `json:"settle_date"`
	PaymentRequest  string        `json:"payment_request"`
	DescriptionHash interface{}   `json:"description_hash"`
	Expiry          string        `json:"expiry"`
	FallbackAddr    string        `json:"fallback_addr"`
	CltvExpiry      string        `json:"cltv_expiry"`
	RouteHints      []interface{} `json:"route_hints"`
	Private         bool          `json:"private"`
	AddIndex        string        `json:"add_index"`
	SettleIndex     string        `json:"settle_index"`
	AmtPaid         string        `json:"amt_paid"`
	AmtPaidSat      string        `json:"amt_paid_sat"`
	AmtPaidMsat     string        `json:"amt_paid_msat"`
	State           string        `json:"state"`
	Htlcs           []HTLC  `json:"htlcs"`
	Features struct {
		Num9 struct {
			Name       string `json:"name"`
			IsRequired bool   `json:"is_required"`
			IsKnown    bool   `json:"is_known"`
		} `json:"9"`
		Num15 struct {
			Name       string `json:"name"`
			IsRequired bool   `json:"is_required"`
			IsKnown    bool   `json:"is_known"`
		} `json:"15"`
		Num17 struct {
			Name       string `json:"name"`
			IsRequired bool   `json:"is_required"`
			IsKnown    bool   `json:"is_known"`
		} `json:"17"`
	} `json:"features"`
	IsKeysend bool `json:"is_keysend"`

}

type ListInvoices struct {
	Invoices []Invoice  `json:"invoices"`
	LastIndexOffset  string `json:"last_index_offset"`
	FirstIndexOffset string `json:"first_index_offset"`
}

func UnmarshalListInvoices(str string) (invoices ListInvoices, err error) {
	err = json.Unmarshal([]byte(str), &invoices)
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
	NumZombieChans       string  `json:"num_zombie_chans"`
}

func UnmarshalNetworkInfo(str string) (ni NetworkInfo, err error) {
	err = json.Unmarshal([]byte(str), &ni)
	return
}

type GetInfo struct {
	Version             string `json:"version"`
	CommitHash          string `json:"commit_hash"`
	IdentityPubkey      string `json:"identity_pubkey"`
	Alias               string `json:"alias"`
	Color               string `json:"color"`
	NumPendingChannels  int    `json:"num_pending_channels"`
	NumActiveChannels   int    `json:"num_active_channels"`
	NumInactiveChannels int    `json:"num_inactive_channels"`
	NumPeers            int    `json:"num_peers"`
	BlockHeight         int    `json:"block_height"`
	BlockHash           string `json:"block_hash"`
	BestHeaderTimestamp string `json:"best_header_timestamp"`
	SyncedToChain       bool   `json:"synced_to_chain"`
	SyncedToGraph       bool   `json:"synced_to_graph"`
	Testnet             bool   `json:"testnet"`
	Chains              []struct {
		Chain   string `json:"chain"`
		Network string `json:"network"`
	} `json:"chains"`
	Uris []interface{} `json:"uris"`
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
	NumChannels   int           `json:"num_channels"`
	TotalCapacity string        `json:"total_capacity"`
	Channels      []interface{} `json:"channels"`
}

func UnmarshalGetNodeInfo(str string) (gni GetNodeInfo, err error) {
	err = json.Unmarshal([]byte(str), &gni)
	return
}

func UnmarshalGetChannelInfo(str string) (gni Edge, err error) {
	err = json.Unmarshal([]byte(str), &gni)
	return
}

func UnmarshalListChainTxns(str string) (lctx ListChainTxns, err error) {
	err = json.Unmarshal([]byte(str), &lctx)
	return
}

func UnmarshalPendingChannels(str string) (lctx PendingChannels, err error) {
	err = json.Unmarshal([]byte(str), &lctx)
	return
}
