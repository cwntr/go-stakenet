package core

import (
	"regexp"
	"strings"

	"strconv"

	"github.com/cwntr/go-stakenet/common"
)

const (
	//Masternode CLI commands
	ArgumentMasternodeList           = "masternodelist"
	ArgumentMasternodeListOptionFull = "full"

	MNStatusEnabled          = "ENABLED"
	MNStatusNewStartRequired = "NEW_START_REQUIRED"

	//Governance CLI commands
	ArgumentGovernanceObject         = "gobject"
	ArgumentGovernanceObjectCount    = "count"
	ArgumentGovernanceObjectList     = "list"
	ArgumentGovernanceObjectAll      = "all"
	ArgumentGovernanceObjectGet      = "get"
	ArgumentGovernanceObjectGetVotes = "getvotes"
	ArgumentGovernanceInfo           = "getgovernanceinfo"

	//Merchant CLI commands
	ArgumentMerchantNodeList     = "merchantnodelist"
	ArgumentMerchantNodeListFull = "full"

	DateTimeFmt = "2006-01-02 15:04:05"
	DateFmt     = "2006-01-02"
)

//status protocol payee lastseen activeseconds lastpaidtime lastpaidblock IP
type MasternodeItem struct {
	Output        string `json:"output"`
	Status        string `json:"status"`
	Protocol      string `json:"protocol"`
	Address       string `json:"address"`
	LastSeen      int    `json:"lastSeen"`
	ActiveSeconds int    `json:"activeSeconds"`
	LastPaidTime  int    `json:"lastPaidTime"`
	LastPaidBlock int    `json:"lastPaidBlock"`
	IP            string `json:"ip"`
}

func GetXSNMNList(cli string, isTestRequest bool) (xsnItems []MasternodeItem, err error) {
	var data string
	if isTestRequest {
		data, err = GetTestMNLIST(), nil
	} else {
		data, err = common.ExecCLI(cli, ArgumentMasternodeList, ArgumentMasternodeListOptionFull)
	}

	if err != nil {
		return []MasternodeItem{}, err
	}

	var mnRegex = regexp.MustCompile(`(?m)\s+"COutPoint\((\w+), \d\)"\:\s"((\w+|\s+\w+))\s(\d+)\s+(\w+)\s(\d+)(\s+)(\d+)\s+(\d+)\s+(\d+)\s+(\w+.\w+.\w+.\w+:\w+)`)

	rows := strings.Split(data, `",`)

	for _, r := range rows {
		if strings.Contains(r, MNStatusEnabled) || strings.Contains(r, MNStatusNewStartRequired) {
			x := MasternodeItem{}

			for _, match := range mnRegex.FindAllStringSubmatch(r, -1) {
				x.Output = match[1]
				x.Status = strings.TrimSpace(match[2])
				x.Protocol = match[4]
				x.Address = match[5]
				x.LastSeen, _ = strconv.Atoi(match[6])
				x.ActiveSeconds, _ = strconv.Atoi(match[8])
				x.LastPaidTime, _ = strconv.Atoi(match[9])
				x.LastPaidBlock, _ = strconv.Atoi(match[10])
				x.IP = match[11]
			}
			xsnItems = append(xsnItems, x)
		}
	}
	return xsnItems, nil
}

func GetXSNGovernanceTotals(cli string, isTestRequest bool) (gov XSNGovTotal, err error) {
	var data string
	if isTestRequest {
		data, err = TestXSNGovernanceTotals(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectCount)
	} else {
		data, err = common.ExecCLI(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectCount)
	}
	return GovTotalToObject(data), err
}

func GetXSNGovInfo(cli string, isTestRequest bool) (gov XSNGovernanceInfo, err error) {
	var data string
	if isTestRequest {
		data, err = TestXSNGovInfo(cli, ArgumentGovernanceInfo)
	} else {
		data, err = common.ExecCLI(cli, ArgumentGovernanceInfo)
	}
	return GovInfoToObject(data), err
}

func GetXSNGovObjects(cli string, isTestRequest bool) (XSNGovernanceList, error) {
	var data string
	var err error

	if isTestRequest {
		data, err = TestXSNGovObjects(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectList, ArgumentGovernanceObjectAll)
	} else {
		data, err = common.ExecCLI(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectList, ArgumentGovernanceObjectAll)
	}
	if err != nil {
		return XSNGovernanceList{}, err
	}
	return GovObjectsToList(data)
}

func GetXSNGovObjectsLimited(cli string, limit int, isTestRequest bool) (XSNGovernanceList, error) {
	list, err := GetXSNGovObjects(cli, isTestRequest)
	if err != nil {
		return list, err
	}
	var limitedList XSNGovernanceList
	cnt := 1
	for _, item := range list.Objects {
		if cnt <= limit {
			limitedList.Objects = append(limitedList.Objects, item)
		}
		cnt++
	}
	return limitedList, nil
}

func GetXSNGovObject(cli string, hash string, isTestRequest bool, mnList []MasternodeItem) (XSNGovObjectWithVotes, error) {
	var objStr string
	var err error

	//get full governance object
	if isTestRequest {
		objStr, err = TestXSNFullGovObject(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectGet, hash)
	} else {
		objStr, err = common.ExecCLI(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectGet, hash)
	}
	if err != nil {
		return XSNGovObjectWithVotes{}, err
	}
	obj, err := ParseXSNFullGovObject(objStr)
	if err != nil {
		return XSNGovObjectWithVotes{}, err
	}

	if len(mnList) == 0 {
		//fetch current MN list
		mn, err := GetXSNMNList(cli, isTestRequest)
		if err != nil {
			return XSNGovObjectWithVotes{}, err
		}
		mnList = mn
	}

	//fetch current votes for full governance object hash
	var votesStr string
	if isTestRequest {
		votesStr, err = TestXSNGovObjectVotes(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectGetVotes, hash)
	} else {
		votesStr, err = common.ExecCLI(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectGetVotes, hash)
	}
	if err != nil {
		return XSNGovObjectWithVotes{}, err
	}

	//match current votes with current masternode list
	votes, err := ParseVotes(votesStr, mnList)
	if err != nil {
		return XSNGovObjectWithVotes{}, err
	}
	return XSNGovObjectWithVotes{Votes: votes, Object: obj}, nil
}

func GetXSNMerchants(cli string, isTestRequest bool) ([]XSNMerchantItem, error) {
	var data string
	var err error
	if isTestRequest {
		data, err = TestXSNMerchantListCLI(cli, ArgumentMerchantNodeList, ArgumentMerchantNodeListFull)
	} else {
		data, err = common.ExecCLI(cli, ArgumentMerchantNodeList, ArgumentMerchantNodeListFull)
	}
	if err != nil {
		return []XSNMerchantItem{}, err
	}
	return ParseXSNMerchants(data), nil
}
