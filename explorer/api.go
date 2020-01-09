package explorer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dnaeon/go-vcr/recorder"
)

const (
	APIUrl = "https://xsnexplorer.io/api"

	//API
	GetTransactionById    = "/transactions/%s"
	GetRawTransactionById = "/transactions/%s/raw"

	GetAddress      = "/addresses/%s"
	GetAddressUTXOs = "/addresses/%s/utxos"

	GetBlocksLatest             = "/blocks"
	GetBlocksByQuery            = "/blocks/%s"
	GetRawBlocksByQuery         = "/blocks/%s/raw"
	GetBlocksTransactionsByHash = "/v2/blocks/%s/transactions"

	GetStats          = "/stats"
	GetRewardsSummary = "/rewards-summary"
	GetRewardsPrices  = "/prices"

	GetBalances = "/v2/balances"

	GetMasternodes     = "/masternodes"
	GetMasternodesByIP = "/masternodes/%s"

	//V2
	GetAddressTransactionsV2 = "/v2/addresses/%s/transactions"
)

//Github repo: https://github.com/X9Developer
//API doc:  https://github.com/X9Developers/block-explorer/blob/master/server/conf/routes
//Explorer: https://xsnexplorer.io/api
type API struct {
	baseUrl string
	c       *http.Client
}

//NewXSNExplorerAPIClient initializes the client with the given base-url. Injects recorder if passed.
func NewXSNExplorerAPIClient(rec *recorder.Recorder) *API {
	m := API{}
	if rec != nil {
		m.c = &http.Client{Transport: rec}
	} else {
		m.c = &http.Client{}
	}
	m.baseUrl = APIUrl
	return &m
}

func (x *API) GetPort() string {
	return "62583"
}

func (x *API) GetLink() string {
	return APIUrl
}

func (x *API) GetTransactionById(txId string) (t Transaction, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetTransactionById, txId)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &t)
	return
}

func (x *API) GetRawTransactionById(txId string) (t RawTransaction, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetRawTransactionById, txId)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &t)
	return
}

func (x *API) GetAddress(address string) (a Address, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetAddress, address)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &a)
	return
}

//address: String, offset: Int ?= 0, limit: Int ?= 10, 	orderBy: String ?= ""
func (x *API) GetAddressTransactionsV2(address string, query url.Values) (a AddressTransactionsV2, err error) {
	url := fmt.Sprintf("%s%s?%s", x.baseUrl, fmt.Sprintf(GetAddressTransactionsV2, address), query.Encode())
	body, err := x.get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &a)
	return
}

func (x *API) GetAddressUTXOs(address string) (a AddressUTXOs, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetAddressUTXOs, address)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &a)
	return
}

func (x *API) GetLatestBlocks() (b Blocks, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetBlocksLatest))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &b)
	return
}

func (x *API) GetBlockByQuery(hash string) (b BlockInfo, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetBlocksByQuery, hash)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &b)
	return
}

func (x *API) GetRawBlocksByQuery(hash string) (b RawBlock, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetRawBlocksByQuery, hash)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &b)
	return
}

// blockhash: String, offset: Int ?= 0, limit: Int ?= 10, orderBy: String ?= "")
func (x *API) GetBlocksTransactionsByHash(blockhash string, query url.Values) (b BlockTransactions, err error) {
	body, err := x.get(fmt.Sprintf("%s%s?%s", x.baseUrl, fmt.Sprintf(GetBlocksTransactionsByHash, blockhash), query.Encode()))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &b)
	return
}

func (x *API) GetStats() (s Stats, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetStats))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &s)
	return
}

func (x *API) GetRewardsSummary() (s RewardSummary, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetRewardsSummary))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &s)
	return
}

//(offset: Int ?= 0, limit: Int ?= 10, orderBy: String ?= ""
func (x *API) GetBalances(query url.Values) (b Balance, err error) {
	body, err := x.get(fmt.Sprintf("%s%s?%s", x.baseUrl, GetBalances, query.Encode()))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &b)
	return
}

//offset: Int ?= 0, limit: Int ?= 10, orderBy: String ?= "
func (x *API) GetMasternodes(query url.Values) (m Masternodes, err error) {
	body, err := x.get(fmt.Sprintf("%s%s?%s", x.baseUrl, GetMasternodes, query.Encode()))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &m)
	return
}

func (x *API) GetMasternodeByIp(ip string) (m Masternode, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetMasternodesByIP, ip)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &m)
	return
}

func (x *API) get(url string) ([]byte, error) {
	res, err := x.c.Get(url)
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
