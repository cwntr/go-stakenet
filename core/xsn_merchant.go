package core

import (
	"regexp"
	"strconv"
	"strings"
)

type XSNMerchantItem struct {
	ID                 string
	Status             string
	Protocol           string
	MerchantAddress    string
	HashTPoSContractTx string
	LastSeen           int
	ActiveSeconds      int
	IP                 string
}

func ParseXSNMerchants(data string) []XSNMerchantItem {
	re := regexp.MustCompile(`(?m)\s+"(\w+)":\s+"\s+(\w+)\s+(\d+)\s+(\w+)\s+(\w+)\s+(\d+)\s+(\d+)\s(.*?)"`)
	rows := strings.Split(data, `,`)
	var mList []XSNMerchantItem
	for _, r := range rows {
		if !strings.Contains(r, MNStatusEnabled) {
			continue
		}
		x := XSNMerchantItem{}
		for _, match := range re.FindAllStringSubmatch(r, -1) {
			x.ID = match[1]
			x.Status = match[2]
			x.Protocol = match[3]
			x.MerchantAddress = match[4]
			x.HashTPoSContractTx = match[5]
			x.LastSeen, _ = strconv.Atoi(match[6])
			x.ActiveSeconds, _ = strconv.Atoi(match[7])
			x.IP = match[8]
			mList = append(mList, x)
		}
	}
	return mList
}
