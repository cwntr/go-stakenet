package explorer

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func TestAPI_TransactionById(t *testing.T) {
	testTxId := "1f39356a0db916a50dc303c20622edf107fae09d37c949c26ee790853e74e983"
	client := NewXSNExplorerAPIClient(nil)
	tx, err := client.GetTransactionById(testTxId)
	if err != nil {
		t.Fail()
		return
	}
	if tx.ID != testTxId {
		t.Fail()
		return
	}
}

func TestAPI_RawTransactionById(t *testing.T) {
	testTxId := "1f39356a0db916a50dc303c20622edf107fae09d37c949c26ee790853e74e983"
	client := NewXSNExplorerAPIClient(nil)
	tx, err := client.GetRawTransactionById(testTxId)
	if err != nil {
		t.Fail()
		return
	}
	if tx.Txid != testTxId {
		t.Fail()
		return
	}
	if tx.Blockhash != "c9a89cddc18c67969c0dfa26efeb78d5350435c30bb4d27c9aee18fa02cf7a90" {
		t.Fail()
		return
	}
}

func TestAPI_GetAddress(t *testing.T) {
	testAddress := "7iacT67sNWkm1JKW1dBwpEN4BA7ZsdpPMW"
	client := NewXSNExplorerAPIClient(nil)
	addr, err := client.GetAddress(testAddress)
	if err != nil {
		t.Fail()
		return
	}
	if addr.Address != testAddress {
		t.Fail()
		return
	}
}

func TestAPI_GetAddressTransactionsV2(t *testing.T) {
	testAddress := "XgzfcNUMNomWrPMbbsgX6MqaU47fvVZKbX"
	client := NewXSNExplorerAPIClient(nil)
	q := url.Values{}
	q.Set("limit", "1")
	addr, err := client.GetAddressTransactionsV2(testAddress, q)
	if err != nil {
		t.Fail()
		return
	}
	if len(addr.Data) != 1 {
		t.Fail()
		return
	}
}

func TestAPI_GetAddressUTXOs(t *testing.T) {
	testAddress := "Xi2xEzvFkPtuK459PYBaangiwbyGHgMHYh"
	client := NewXSNExplorerAPIClient(nil)
	utxos, err := client.GetAddressUTXOs(testAddress)
	if err != nil {
		t.Fail()
		return
	}
	if len(utxos) > 0 {
		if utxos[0].Address != testAddress {
			t.Fail()
			return
		}
	}
}

func TestAPI_GetLatestBlocks(t *testing.T) {
	client := NewXSNExplorerAPIClient(nil)
	blocks, err := client.GetLatestBlocks()
	if err != nil {
		t.Fail()
		return
	}

	if len(blocks) < 5 {
		t.Fail()
		return
	}
	for _, b := range blocks {
		if b.Confirmations == 0 {
			t.Fail()
			return
		}
		if b.Hash == "" {
			t.Fail()
			return
		}
		if b.PreviousBlockhash == "" {
			t.Fail()
			return
		}
	}
}

func TestAPI_BlockByHash(t *testing.T) {
	testBlockHash := "b438e6f08812c9591c6f886ef0ba2ce873659a3797b32ba78e2a160a9b43aacf"
	client := NewXSNExplorerAPIClient(nil)
	blockInfo, err := client.GetBlockByQuery(testBlockHash)
	if err != nil {
		t.Fail()
		return
	}

	if blockInfo.Block.Hash != testBlockHash {
		t.Fail()
		return
	}

	if blockInfo.Rewards.Coinstake.Value != 9 {
		t.Fail()
		return
	}

	if blockInfo.Rewards.Masternode.Value != 9 {
		t.Fail()
		return
	}
	if len(blockInfo.Block.Transactions) == 0 {
		t.Fail()
		return
	}
}

func TestAPI_RawBlockByHash(t *testing.T) {
	testBlockHash := "b438e6f08812c9591c6f886ef0ba2ce873659a3797b32ba78e2a160a9b43aacf"
	client := NewXSNExplorerAPIClient(nil)
	rawBlock, err := client.GetRawBlocksByQuery(testBlockHash)
	if err != nil {
		t.Fail()
		return
	}

	if rawBlock.Hash != testBlockHash {
		t.Fail()
		return
	}

	if rawBlock.Confirmations == 0 {
		t.Fail()
		return
	}
	if len(rawBlock.Tx) == 0 {
		t.Fail()
		return
	}
	if rawBlock.Previousblockhash == "" {
		t.Fail()
		return
	}
}

func TestAPI_BlockTransactionsByHash(t *testing.T) {
	testBlockHash := "b438e6f08812c9591c6f886ef0ba2ce873659a3797b32ba78e2a160a9b43aacf"
	client := NewXSNExplorerAPIClient(nil)
	tx, err := client.GetBlocksTransactionsByHash(testBlockHash, url.Values{})
	if err != nil {
		t.Fail()
		return
	}
	if len(tx.Data) == 0 {
		t.Fail()
		return
	}
	for _, txi := range tx.Data {
		if txi.ID == "" || txi.Blockhash == "" {
			t.Fail()
			return
		}
	}
}

func TestAPI_BlockTransactionsByHashWithQuery(t *testing.T) {
	testBlockHash := "b438e6f08812c9591c6f886ef0ba2ce873659a3797b32ba78e2a160a9b43aacf"
	client := NewXSNExplorerAPIClient(nil)
	q := url.Values{}
	q.Set("limit", "2")
	q.Set("offset", "1")
	q.Set("orderBy", "time")
	tx, err := client.GetBlocksTransactionsByHash(testBlockHash, q)
	if err != nil {
		t.Fail()
		return
	}
	if len(tx.Data) != 2 {
		t.Fail()
		return
	}
	if tx.Data[0].ID != "07bd3cc7da380aa293a5271d7843034e902872814b09e4cf00397abc656c1ce4" ||
		tx.Data[0].Blockhash != testBlockHash {
		t.Fail()
		return
	}
	if tx.Data[1].ID != "491706404c7eadbbff3c68fc21c1be7bca134ad3b9502b626320549f001267db" ||
		tx.Data[1].Blockhash != testBlockHash {
		t.Fail()
		return
	}
}

func TestAPI_Stats(t *testing.T) {
	client := NewXSNExplorerAPIClient(nil)
	s, err := client.GetStats()
	if err != nil {
		t.Fail()
		return
	}
	if s.Blocks < 353000 {
		t.Fail()
		return
	}
	if s.Masternodes == 0 {
		t.Fail()
		return
	}
}

func TestAPI_GetRewardsSummary(t *testing.T) {
	client := NewXSNExplorerAPIClient(nil)
	summary, err := client.GetRewardsSummary()
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	if summary.AverageReward < 0 {
		t.Fail()
		return
	}
	if summary.AverageInput < 0 {
		t.Fail()
		return
	}
	if summary.MedianInput < 0 {
		t.Fail()
		return
	}
	if summary.AveragePoSInput < 0 {
		t.Fail()
		return
	}
	if summary.AverageTPoSInput < 0 {
		t.Fail()
		return
	}
	if summary.MedianWaitTime < 0 {
		t.Fail()
		return
	}
}

/*
func TestAPI_GetPrices(t *testing.T) {
	client := NewXSNExplorerAPIClient(nil)
	prices, err := client.GetPrices()
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	if prices.USD <= 0 {
		fmt.Println("b")
		t.Fail()
		return
	}
	if prices.EUR <= 0 {
		fmt.Println("c")
		t.Fail()
		return
	}
}
*/
func TestAPI_Balance(t *testing.T) {
	client := NewXSNExplorerAPIClient(nil)
	q := url.Values{}
	q.Set("limit", "13")
	tx, err := client.GetBalances(q)
	if err != nil {
		t.Fail()
		return
	}
	if len(tx.Data) != 13 {
		t.Fail()
		return
	}
}

func TestAPI_GetMasternodes(t *testing.T) {
	client := NewXSNExplorerAPIClient(nil)
	q := url.Values{}
	q.Set("limit", "2")
	q.Set("offset", "1")
	q.Set("orderBy", "activeSeconds:desc")
	mns, err := client.GetMasternodes(q)
	if err != nil {
		t.Fail()
		return
	}
	if mns.Offset != 1 || mns.Limit != 2 || len(mns.Data) != 2 {
		t.Fail()
		return
	}

	ip := strings.Split(mns.Data[0].IP, ":")[0] //take first IP found
	m, mErr := client.GetMasternodeByIp(ip)
	if mErr != nil {
		t.Fail()
		return
	}
	if m.Status == "" {
		t.Fail()
		return
	}
}
