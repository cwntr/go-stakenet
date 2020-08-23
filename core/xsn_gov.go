package core

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ProposalMap struct {
	Map map[string]XSNGovObject `json:"-"` //dynamic property names
}

//Actual Proposal
type XSNProposalDataType1 struct {
	Name           string `json:"name"`
	PaymentAddress string `json:"payment_address"`
	PaymentAmount  int    `json:"payment_amount"`
	Type           int    `json:"type"`
	URL            string `json:"url"`
	EndEpoch       int    `json:"end_epoch"`
	StartEpoch     int    `json:"start_epoch"`
	EndDateStr     string `json:"end_date"`
	StartDateStr   string `json:"start_date"`
}

//Trigger
type XSNProposalDataType2 struct {
	Type             int    `json:"type"`
	EventBlockHeight int    `json:"event_block_height"`
	PaymentAddresses string `json:"payment_addresses"`
	PaymentAmounts   string `json:"payment_amounts"`
}

type XSNGovObject struct {
	DataHex             string               `json:"dataHex"`
	DataString          string               `json:"dataString"`
	Data1               XSNProposalDataType1 `json:"data1"`
	Data2               XSNProposalDataType2 `json:"data2"`
	Hash                string               `json:"hash"`
	CollateralHash      string               `json:"collateralHash"`
	ObjectType          int                  `json:"objectType"`
	CreationTime        int                  `json:"creationTime"`
	CreationDate        string               `json:"creationDate"`
	SigningMasternode   string               `json:"signingMasternode"`
	AbsoluteYesCount    int                  `json:"absoluteYesCount"`
	YesCount            int                  `json:"yesCount"`
	NoCount             int                  `json:"noCount"`
	AbstainCount        int                  `json:"abstainCount"`
	FBlockchainValidity bool                 `json:"fBlockchainValidity"`
	IsValidReason       string               `json:"IsValidReason"`
	FCachedValid        bool                 `json:"fCachedValid"`
	FCachedFunding      bool                 `json:"fCachedFunding"`
	FCachedDelete       bool                 `json:"fCachedDelete"`
	FCachedEndorsed     bool                 `json:"fCachedEndorsed"`
}

type XSNGovObjectWithVotes struct {
	Object XSNGovFullObject     `json:"govObject"`
	Votes  XSNGovObjectVoteList `json:"votes"`
}

type XSNGovFullObject struct {
	DataHex        string               `json:"DataHex"`
	DataString     string               `json:"DataString"`
	Data1          XSNProposalDataType1 `json:"data1"`
	Data2          XSNProposalDataType2 `json:"data2"`
	Hash           string               `json:"Hash"`
	CollateralHash string               `json:"CollateralHash"`
	ObjectType     int                  `json:"ObjectType"`
	CreationTime   int                  `json:"CreationTime"`
	CreationDate   string               `json:"creationDate"`
	FundingResult  struct {
		AbsoluteYesCount int `json:"AbsoluteYesCount"`
		YesCount         int `json:"YesCount"`
		NoCount          int `json:"NoCount"`
		AbstainCount     int `json:"AbstainCount"`
	} `json:"FundingResult"`
	ValidResult struct {
		AbsoluteYesCount int `json:"AbsoluteYesCount"`
		YesCount         int `json:"YesCount"`
		NoCount          int `json:"NoCount"`
		AbstainCount     int `json:"AbstainCount"`
	} `json:"ValidResult"`
	DeleteResult struct {
		AbsoluteYesCount int `json:"AbsoluteYesCount"`
		YesCount         int `json:"YesCount"`
		NoCount          int `json:"NoCount"`
		AbstainCount     int `json:"AbstainCount"`
	} `json:"DeleteResult"`
	EndorsedResult struct {
		AbsoluteYesCount int `json:"AbsoluteYesCount"`
		YesCount         int `json:"YesCount"`
		NoCount          int `json:"NoCount"`
		AbstainCount     int `json:"AbstainCount"`
	} `json:"EndorsedResult"`
	FLocalValidity  bool   `json:"fLocalValidity"`
	IsValidReason   string `json:"IsValidReason"`
	FCachedValid    bool   `json:"fCachedValid"`
	FCachedFunding  bool   `json:"fCachedFunding"`
	FCachedDelete   bool   `json:"fCachedDelete"`
	FCachedEndorsed bool   `json:"fCachedEndorsed"`
}

type XSNGovernanceList struct {
	Objects []XSNGovObject `json:"objects"`
}

func (c XSNGovernanceList) Len() int      { return len(c.Objects) }
func (c XSNGovernanceList) Swap(i, j int) { c.Objects[i], c.Objects[j] = c.Objects[j], c.Objects[i] }
func (c XSNGovernanceList) Less(i, j int) bool {
	return c.Objects[i].CreationTime > c.Objects[j].CreationTime
}

type XSNGovTotal struct {
	Objects   int    `json:"objects"`
	Proposals int    `json:"proposals"`
	Triggers  int    `json:"triggers"`
	Watchdog  string `json:"watchdog"`
	Other     int    `json:"other"`
	Erased    int    `json:"erased"`
	Votes     int    `json:"votes"`
}

type XSNGovObjectVote struct {
	Epoch  int
	Date   string `json:"date"`
	Vote   string `json:"vote"`
	Type   string `json:"type"`
	Output string `json:"output"`
	Addr   string `json:"address"`
}

type XSNGovObjectVoteList struct {
	Votes []XSNGovObjectVote `json:"votes"`
}

func (c XSNGovObjectVoteList) Len() int      { return len(c.Votes) }
func (c XSNGovObjectVoteList) Swap(i, j int) { c.Votes[i], c.Votes[j] = c.Votes[j], c.Votes[i] }
func (c XSNGovObjectVoteList) Less(i, j int) bool {
	return c.Votes[i].Epoch > c.Votes[j].Epoch
}

type XSNGovernanceInfo struct {
	Governanceminquorum          int     `json:"governanceminquorum"`
	Masternodewatchdogmaxseconds int     `json:"masternodewatchdogmaxseconds"`
	Proposalfee                  float64 `json:"proposalfee"`
	Superblockcycle              int     `json:"superblockcycle"`
	Lastsuperblock               int     `json:"lastsuperblock"`
	Nextsuperblock               int     `json:"nextsuperblock"`
	Maxgovobjdatasize            int     `json:"maxgovobjdatasize"`
}

func GovObjectsToList(str string) (XSNGovernanceList, error) {
	var list XSNGovernanceList
	var p ProposalMap
	if err := json.Unmarshal([]byte(str), &p); err != nil {
		return list, err
	}
	if err := json.Unmarshal([]byte(str), &p.Map); err != nil {
		return list, err
	}

	for _, object := range p.Map {
		o := XSNGovObject{
			DataHex:             object.DataHex,
			DataString:          object.DataString,
			Data1:               object.Data1,
			Data2:               object.Data2,
			Hash:                object.Hash,
			CollateralHash:      object.CollateralHash,
			ObjectType:          object.ObjectType,
			CreationTime:        object.CreationTime,
			SigningMasternode:   object.SigningMasternode,
			AbsoluteYesCount:    object.AbsoluteYesCount,
			YesCount:            object.YesCount,
			NoCount:             object.NoCount,
			AbstainCount:        object.AbstainCount,
			FBlockchainValidity: object.FBlockchainValidity,
			IsValidReason:       object.IsValidReason,
			FCachedValid:        object.FCachedValid,
			FCachedFunding:      object.FCachedFunding,
			FCachedDelete:       object.FCachedDelete,
			FCachedEndorsed:     object.FCachedEndorsed,
		}
		i := int64(o.CreationTime)
		ut := time.Unix(i, 0)
		o.CreationDate = ut.Format(DateFmt)
		deserial := strings.Replace(o.DataString, `\"`, `"`, -1)
		if o.ObjectType == 1 {
			var re = regexp.MustCompile(`(?m)(\{.*\})`)
			for _, match := range re.FindAllStringSubmatch(deserial, -1) {
				objectStr := match[1]
				var type1 XSNProposalDataType1
				err := json.Unmarshal([]byte(objectStr), &type1)
				if err != nil {
					fmt.Printf("err unmarshalling: %v \n", err)
					continue
				} else {
					s := int64(type1.StartEpoch)
					tm := time.Unix(s, 0)

					type1.StartDateStr = tm.Format(DateFmt)
					e := int64(type1.EndEpoch)
					em := time.Unix(e, 0)
					type1.EndDateStr = em.Format(DateFmt)
					o.Data1 = type1
				}
			}
		}
		if o.ObjectType == 2 {
			var re = regexp.MustCompile(`(?m)(\{.*\})`)
			for _, match := range re.FindAllStringSubmatch(deserial, -1) {
				objectStr := match[1]
				var type2 XSNProposalDataType2
				err := json.Unmarshal([]byte(objectStr), &type2)
				if err != nil {
					fmt.Printf("err unmarshalling: %v \n", err)
					continue
				} else {
					o.Data2 = type2
				}
			}
		}
		list.Objects = append(list.Objects, o)
	}
	sort.Sort(list)
	return list, nil
}

func GovTotalToObject(str string) XSNGovTotal {
	var re = regexp.MustCompile(`(?m)Governance Objects: (\d+)\s\(Proposals: (\d+), Triggers: (\d+), Watchdogs: (\d+\/\d+), Other: (\d+); Erased: (\d+)\), Votes: (\d+)`)
	govTotal := XSNGovTotal{}
	for _, match := range re.FindAllStringSubmatch(str, -1) {
		govTotal.Objects, _ = strconv.Atoi(match[1])
		govTotal.Proposals, _ = strconv.Atoi(match[2])
		govTotal.Triggers, _ = strconv.Atoi(match[3])
		govTotal.Watchdog = match[4]
		govTotal.Other, _ = strconv.Atoi(match[5])
		govTotal.Erased, _ = strconv.Atoi(match[6])
		govTotal.Votes, _ = strconv.Atoi(match[7])
	}
	return govTotal
}

func GovInfoToObject(str string) XSNGovernanceInfo {
	var govInfo XSNGovernanceInfo
	err := json.Unmarshal([]byte(str), &govInfo)
	if err != nil {
		fmt.Printf("unable to unmarshall governance info")
	}
	return govInfo
}

const (
	MatchTypeDirect = "direct"
	MatchTypeGuess  = "guess"
)

type MatchVoteGrouped struct {
	Totals             MatchVoteCounter
	ProposalMatchVotes map[string]MatchVoteCounter
}

type MatchVoteCounter struct {
	DirectMatches     []MatchedVote
	GuessedMatches    []MatchedVote
	TotalDirectMatch  int
	TotalGuessedMatch int
}

type MatchedVoteList []MatchedVote

type MatchedVote struct {
	GovType      string
	GovName      string
	GovHash      string
	GovStartDate string
	GovEndDate   string
	GovAmount    int
	GovMatchType string //for epoch +-2
	Vote         XSNGovObjectVote
}

func GetVotesByMN(address string, cli string, isTestRequest bool) (MatchVoteGrouped, error) {
	var mVotes []MatchedVote
	l, err := GetXSNGovObjects(cli, isTestRequest)
	if err != nil {
		return MatchVoteGrouped{}, err
	}

	//fetch current MN list
	mn, err := GetXSNMNList(cli, isTestRequest)
	if err != nil {
		fmt.Printf("err fetching mn list: %v \n", err)
		return MatchVoteGrouped{}, err
	}
	for _, o := range l.Objects {
		objWithVotes, oErr := GetXSNGovObject(cli, o.Hash, isTestRequest, mn)
		if oErr != nil {
			return MatchVoteGrouped{}, err
		}
		foundDateTimeStr := ""
		for _, vote := range objWithVotes.Votes.Votes {
			if vote.Addr == address {
				if o.ObjectType == 1 {
					mVote := MatchedVote{
						Vote:         vote,
						GovType:      "Proposal",
						GovName:      o.Data1.Name,
						GovHash:      o.Hash,
						GovAmount:    o.Data1.PaymentAmount,
						GovStartDate: o.Data1.StartDateStr,
						GovEndDate:   o.Data1.EndDateStr,
						GovMatchType: MatchTypeDirect,
					}
					foundDateTimeStr = mVote.Vote.Date
					mVotes = append(mVotes, mVote)
				}
				if o.ObjectType == 2 {
					mVote := MatchedVote{
						Vote:         vote,
						GovType:      "Trigger",
						GovName:      " - ",
						GovHash:      o.Hash,
						GovStartDate: " - ",
						GovEndDate:   " - ",
						GovMatchType: MatchTypeDirect,
					}
					foundDateTimeStr = mVote.Vote.Date
					mVotes = append(mVotes, mVote)
				}
			}
		}
		if foundDateTimeStr != "" {
			for _, vote := range objWithVotes.Votes.Votes {
				if vote.Date == foundDateTimeStr && vote.Addr != address {
					if o.ObjectType == 1 {
						mVote := MatchedVote{
							Vote:         vote,
							GovType:      "Proposal",
							GovName:      o.Data1.Name,
							GovHash:      o.Hash,
							GovStartDate: o.Data1.StartDateStr,
							GovEndDate:   o.Data1.EndDateStr,
							GovMatchType: MatchTypeGuess,
						}
						foundDateTimeStr = mVote.Vote.Date
						mVotes = append(mVotes, mVote)
					}
					if o.ObjectType == 2 {
						mVote := MatchedVote{
							Vote:         vote,
							GovType:      "Trigger",
							GovName:      " - ",
							GovHash:      o.Hash,
							GovStartDate: " - ",
							GovEndDate:   " - ",
							GovMatchType: MatchTypeGuess,
						}
						foundDateTimeStr = mVote.Vote.Date
						mVotes = append(mVotes, mVote)
					}
				}
			}
		}
	}
	return GroupMatchVotes(mVotes)
}

func GetFullVoteHistory(cli string, isTestRequest bool) (MatchedVoteList, error) {
	var mVotes MatchedVoteList
	l, err := GetXSNGovObjects(cli, isTestRequest)
	if err != nil {
		return MatchedVoteList{}, err
	}

	//fetch current MN list
	mn, err := GetXSNMNList(cli, isTestRequest)
	if err != nil {
		fmt.Printf("err fetching mn list: %v \n", err)
		return MatchedVoteList{}, err
	}
	for _, o := range l.Objects {
		objWithVotes, oErr := GetXSNGovObject(cli, o.Hash, isTestRequest, mn)
		if oErr != nil {
			return MatchedVoteList{}, err
		}
		for _, vote := range objWithVotes.Votes.Votes {
			if o.ObjectType == 1 {
				mVote := MatchedVote{
					Vote:         vote,
					GovType:      "Proposal",
					GovName:      o.Data1.Name,
					GovHash:      o.Hash,
					GovAmount:    o.Data1.PaymentAmount,
					GovStartDate: o.Data1.StartDateStr,
					GovEndDate:   o.Data1.EndDateStr,
					GovMatchType: MatchTypeDirect,
				}
				if mVote.Vote.Addr == "" {
					mVote.Vote.Addr = " - not in current mn list -"
				}
				mVotes = append(mVotes, mVote)
			}
			if o.ObjectType == 2 {
				mVote := MatchedVote{
					Vote:         vote,
					GovType:      "Trigger",
					GovName:      " - ",
					GovHash:      o.Hash,
					GovStartDate: " - ",
					GovEndDate:   " - ",
					GovMatchType: MatchTypeDirect,
				}
				if mVote.Vote.Addr == "" {
					mVote.Vote.Addr = " - not in current mn list -"
				}
				mVotes = append(mVotes, mVote)
			}
		}
	}
	return mVotes, nil
}

func GroupMatchVotes(v []MatchedVote) (MatchVoteGrouped, error) {
	g := MatchVoteGrouped{}

	var directMatch int
	var guessedMatch int

	var directMatches []MatchedVote
	var guessedMatches []MatchedVote

	for _, vote := range v {
		if vote.GovMatchType == MatchTypeDirect {
			directMatches = append(directMatches, vote)
			directMatch++
		}
		if vote.GovMatchType == MatchTypeGuess {
			guessedMatches = append(guessedMatches, vote)
			guessedMatch++
		}
	}

	g.Totals.DirectMatches = directMatches
	g.Totals.GuessedMatches = guessedMatches
	g.Totals.TotalDirectMatch = directMatch
	g.Totals.TotalGuessedMatch = guessedMatch
	return g, nil
}

func ParseVotes(votestr string, mns []MasternodeItem) (XSNGovObjectVoteList, error) {
	var list XSNGovObjectVoteList

	rows := strings.Split(votestr, `",`)
	var re = regexp.MustCompile(`(?m)COutPoint\((\w+),\s+\d+\),\s+scriptSig=\):(\d+):(\w+):(\w+)`)

	for _, r := range rows {
		vote := XSNGovObjectVote{}
		for _, match := range re.FindAllStringSubmatch(r, -1) {
			var err error
			vote.Epoch, err = strconv.Atoi(match[2])
			if err != nil {
				return list, err
			}
			i := int64(vote.Epoch)
			it := time.Unix(i, 0)
			vote.Date = it.Format(DateTimeFmt)
			vote.Vote = match[3]
			vote.Type = match[4]
			vote.Output = match[1]
			if vote.Output != "" {
				for _, mn := range mns {
					if vote.Output == mn.Output {
						vote.Addr = mn.Address
					}
				}
			}
		}
		list.Votes = append(list.Votes, vote)
	}
	sort.Sort(list)
	return list, nil
}

func ParseXSNFullGovObject(str string) (XSNGovFullObject, error) {
	var obj XSNGovFullObject
	if err := json.Unmarshal([]byte(str), &obj); err != nil {
		return XSNGovFullObject{}, err
	}

	i := int64(obj.CreationTime)
	ut := time.Unix(i, 0)
	obj.CreationDate = ut.Format(DateTimeFmt)
	deserial := strings.Replace(obj.DataString, `\"`, `"`, -1)
	if obj.ObjectType == 1 {
		var re = regexp.MustCompile(`(?m)(\{.*\})`)
		for _, match := range re.FindAllStringSubmatch(deserial, -1) {
			objectStr := match[1]
			var type1 XSNProposalDataType1
			err := json.Unmarshal([]byte(objectStr), &type1)
			if err != nil {
				fmt.Printf("err unmarshalling: %v \n", err)
				continue
			} else {
				s := int64(type1.StartEpoch)
				tm := time.Unix(s, 0)

				type1.StartDateStr = tm.Format(DateTimeFmt)
				e := int64(type1.EndEpoch)
				em := time.Unix(e, 0)
				type1.EndDateStr = em.Format(DateTimeFmt)
				obj.Data1 = type1
			}
		}
	}
	if obj.ObjectType == 2 {
		var re = regexp.MustCompile(`(?m)(\{.*\})`)
		for _, match := range re.FindAllStringSubmatch(deserial, -1) {
			objectStr := match[1]
			var type2 XSNProposalDataType2
			err := json.Unmarshal([]byte(objectStr), &type2)
			if err != nil {
				fmt.Printf("err unmarshalling: %v \n", err)
				continue
			} else {
				obj.Data2 = type2
			}
		}
	}
	return obj, nil
}

func TestXSNGovernanceTotals(command string, args ...string) (string, error) {
	return `Governance Objects: 25 (Proposals: 24, Triggers: 1, Watchdogs: 0/0, Other: 0; Erased: 0), Votes: 6407`, nil
}

func TestXSNGovInfo(command string, args ...string) (string, error) {
	return `{
  "governanceminquorum": 10,
  "masternodewatchdogmaxseconds": 7200,
  "proposalfee": 5.00000000,
  "superblockcycle": 43200,
  "lastsuperblock": 475200,
  "nextsuperblock": 518400,
  "maxgovobjdatasize": 16384
}`, nil
}

func TestXSNFullGovObject(command string, args ...string) (string, error) {
	return `{
		"DataHex": "5b5b2270726f706f73616c222c7b22656e645f65706f6368223a313533313930323033312c226e616d65223a2278736e2d6465762d66756e642d6a756c79222c227061796d656e745f61646472657373223a22586d4639377547523436314e6a344e536e536871524a4b345664396d32314a7a5747222c227061796d656e745f616d6f756e74223a3135313130302c2273746172745f65706f6368223a313532393332363535312c2274797065223a312c2275726c223a2268747470733a2f2f78736e636f696e2e696f227d5d5d",
		"DataString": "[[\"proposal\",{\"end_epoch\":1531902031,\"name\":\"xsn-dev-fund-july\",\"payment_address\":\"XmF97uGR461Nj4NSnShqRJK4Vd9m21JzWG\",\"payment_amount\":151100,\"start_epoch\":1529326551,\"type\":1,\"url\":\"https://xsncoin.io\"}]]",
		"Hash": "4b5c9f706cc0a8b58e31bb764845fffbbfed0ed52f74086776cb9936c871a427",
		"CollateralHash": "e124bfb35cc8828cb740e97f587922c4c9463845af089e6ef27d3ed70620ed48",
		"ObjectType": 1,
		"CreationTime": 1530020640,
		"FundingResult": {
	"AbsoluteYesCount": 320,
	"YesCount": 320,
	"NoCount": 0,
	"AbstainCount": 0
	},
		"ValidResult": {
	"AbsoluteYesCount": 0,
	"YesCount": 0,
	"NoCount": 0,
	"AbstainCount": 0
	},
		"DeleteResult": {
	"AbsoluteYesCount": 3,
	"YesCount": 3,
	"NoCount": 0,
	"AbstainCount": 0
	},
		"EndorsedResult": {
	"AbsoluteYesCount": 0,
	"YesCount": 0,
	"NoCount": 0,
	"AbstainCount": 0
	},
		"fLocalValidity": true,
		"IsValidReason": "",
		"fCachedValid": true,
		"fCachedFunding": true,
		"fCachedDelete": false,
		"fCachedEndorsed": false
	}`, nil
}

func TestXSNGovObjectVotes(command string, args ...string) (string, error) {
	return `
{
  "65643362343838323665646639656362303333646265633335623036356234613965366430356435": "           ENABLED 70209 Xv7CGGg4ozrFLHXYqtxoGwBnsRTdyfaX3G 564b28288c0f5f269a1c3e149dd8abde00b147de85ef611dffa0d1b2072b01af 1598183686 11279510 78.141.218.19:62583",
  "37636264653365303737356561306364633432666364613231643234346363663539383764313933": "           ENABLED 70209 XpASAmCN6ufTb8QeBmY93rzRKGCC69u66P 29f1502a3bc482ea2550bfaca96e3ae802a934b7874c84e949470697ebd122e1 1598183749  3616229 178.128.70.200:62583",
  "64323435623435303431646465363263306336366463636538633361346365626461633761343836": "           ENABLED 70209 Xnxmn7zcjiSbzsEpvidLRSLp19xjYaXiNK 9b61db1b47b565c4ed5966719dd235aed7c93894c85e951d242dd57193b4c988 1598183675  3101293 149.28.148.29:62583",
  "33336266653533633632376630393333363764373531323862646265643762653433336364303563": "           ENABLED 70209 Xj9bTMMPk4v6u8UL1Pk7Rm7ufURqqh4SUj 429ff1214254f6d4c5609a7293f695a14e180128cb455cf970083319b3bacaa1 1598184075  1118736 51.83.134.197:62583",
  "66663637663261613136343131303030376665326561356463326438346363643362366537393362": "           ENABLED 70209 Xg7KAiZnCHyGykTEHk5NY6ZtRxLJKQ1CWK 8bed96afa074c0c5cbbb71da26537c4d3b35b8fa9ac3f8a762f1d5b5699c4e46 1598184058 17617172 149.28.90.238:62583",
  "64313166363037346637373361653063393139623732636435313165393037343739316664396135": "           ENABLED 70209 XqomRsYiUkECm2EYLhjyF4Y4ktD6w6Xhur d9a6827084c1ff7013602c826b526ddc229760195f3a140b551e6eacead6f514 1598184151  1112212 108.61.117.193:62583",
  "37616365366538656339353035366431613634653962646139656437306465646662653465376537": "           ENABLED 70209 Xwq3isnh4Nfr6NBUr1SZZdpdSXZVe969w4 4b810d83f321eb6f58dfb7c1c113cf52e3b3e90b18c1c79bae5ac720a0b1ff33 1598184022 35326389 45.77.192.175:62583",
  "33313165373631326664623532326231313131643064376139656562343930343732356139633438": "           ENABLED 70209 XhJmngaMHdbJfHA5fXdSLqgi5V5yUnsuQL 8bc14086c3567c834af5479f73a00ad2cdcb0b8e3109ec2f05c29c00b8b7eaf0 1598184066   674177 92.60.45.150:62583",
  "38363266666461343335653537636561396263393630333230306238303965313062346533326236": "           ENABLED 70209 XsJD5Qk1kgSf8KszmxeHojryWVByncb9ix 61ebddda53f4f406363315b8a3b76734085154f73b3837984c5628ba3810f8f9 1598183824   259663 206.81.28.44:62583",
  "65643763623436623666366131313739636662376130366533313262613362393034353332636633": "           ENABLED 70209 Xxrd8aF6XGa6BA1XpngYMd3gogWvGakozx f840a7cccbb7894f2009417d92cfcd75c34fb7411f983f212e7b9730a040ec6b 1598184122 11409506 108.61.103.169:62583",
  "62383334636662306263623561353661303266613139636133623239313530303433363139323331": "           ENABLED 70209 XfCxJJsNa8KV3AwaeU4YzyJpdbR5oeFKwy 7fad75e7e6080e920748f6d1074349b5776265f34bc17bab4279ba923af5765b 1598183802  1463683 95.179.194.65:62583",
  "66383233313535663634646563333965613661343566356234326463366663616137633738303334": "           ENABLED 70209 XfUTF1io28zFmtytcnrk9s3ExkxueNqTkh 9f01356f9ec16c872ddde52e08e1a61906b9288b5894a5c6845e478444aeb87e 1598183747  3804027 66.42.57.239:62583",
  "34633032636161393836623034313033623962373035383062333665343666626533666363656237": "           ENABLED 70209 XsSjT8JFSej25EEjCPuQKEa6GoJrypdKHn 4219fe626ecde633229ee6fa63935ef0501807635b66ebae908230f4aa0b3c06 1598183862  3713281 95.179.236.36:62583",
  "33313362393332383532646631313930613764613839643362323637363236353164306636666262": "           ENABLED 70209 XsmuEeGGt2TNLfyQibxLK1decurCpjh6of ebe89188a2ee11dc906614784f261c246867719fb7ef99f8c5c3dfbe04d75ef4 1598183799 19678721 54.37.74.101:62583",
  "65316264333730343939386234636635386436636533613964663762316133613066646364353636": "           ENABLED 70209 Xk4ax3dcKFuYXBp6tCyvu3qNre1okWnSVp 5b331d3aa466c4595a208a6264c29ac6c46557055071992be2cf7e0fd05c2d60 1598183654 17037687 107.191.62.138:62583",
  "39366633326136353366313163376465313237346564303931326531336132333139386662366337": "           ENABLED 70209 Xttpzr8Mn3cip94HZTCv4zmGXEqzyo8afg 42698b416c64c426ff37d9d843b96b3e8f9f44edfb723fc1f6417307dabdda17 1598183846  9506674 45.76.35.27:62583",
  "66616263316430326134393366636431356663646235363963343265623232396335386237346339": "           ENABLED 70209 Xu43GZtTDHyLSAuNvcLxtE3qhWVvpJtHLE faba97b9d1a4e4a4c2bec4d5b84e905c580d798b1e5cb4a6eb047d0e9613ef2c 1598184063  1818412 78.141.209.248:62583",
  "32363231636631613161393063366336363161326661613565393337366539326366333230383665": "           ENABLED 70209 Xkidyzg1GyaXBowFhaXZXVAzmaUYtR9FpS 15694067c2f99d4494c083652ae93b1428525cea4abfce5e300a5583878e6085 1598183954  4040307 45.32.237.12:62583",
  "64323638343336356363383864323561666463336230326462376633623839363366333532326662": "           ENABLED 70209 XyaiRDWbfmgAKhSiQK47Lc8x6TfFhmqici 19c8a18569981991a7b61a1a370c1946cc7c4eb704379e35eca3d3ef3cc4e947 1598184052 29255944 45.76.42.254:62583",
  "39653164323066386338663437363235376162626235303962613531346537376136356335623139": "           ENABLED 70209 XczvAzAPnqAa2hQL3ip1rJHmiBoZ6E5RGM 6170c1916b53f7c18353dc6271efa53bf17f43635de571976a9a7719fad76be0 1598183732   934396 104.238.173.22:62583",
  "35636631313064663237396664356231633535633031616634633133656238623965636632373639": "           ENABLED 70209 XkGrUWj6YPwsdmWm8TLTaX195cVLMnY2rk 2f83c2525a35ccc9d3b2cfd787b6943b78e2ab35372ea93a4f4a3fb2ee90c7b2 1598183730 18315069 206.81.2.218:62583",
  "31353866323339393533366534633261353661353232656163313736386230323932643563613464": "           ENABLED 70209 XhnArCJvsj81R6aD29qUiAzJEPt27AerXG 38a5be7c45d9d6079e7e1a9ddfd379a671dce08cc263e7e1c37f29a014698ad7 1598183773 12484671 92.60.45.18:62583",
  "36336430313433346263306334363434303735373936646230626563343237313939353165306239": "           ENABLED 70209 XsdfZgWCnbKp1tJGyEpLmstd1X54emCued 9e8ee1aeb948f02f85b94df0582ee03b73858a358c3f5b85d3f8413658629e57 1598184111 19203879 92.60.45.16:62583",
  "38343661353838316631366264356135636364666235313163366666346638666134656634653865": "           ENABLED 70209 XofJM522oEFyRZDZ4tcukyuxPvcDkqNkSi d1b992064411d40ad998e2cf1b18c0eb7f6c19810e3b039b428676a7d7d5e042 1598183932 19184713 92.60.45.17:62583",
  "32313839613261343936613830313464306162326139396163393331383138616532316136316330": "           ENABLED 70209 XtE3u4Rpy59P9RjW96s8DgtovcFubs9fpx 71396a903964376b22ac02933aa068d0287d11cc2fc36438e32bd164dea10a4a 1598183716 15477298 66.42.106.231:62583",
  "33323334353561386539303235326265323538623437336332633862356534303734623663356532": "           ENABLED 70209 XwMuPvjCDEhnG8sb3Uszcp8YnbSHgEzJjq 5bfa8f1cb24593d7a6efd3e2292a5c620ad4c96e38c35215be0281a0523b139d 1598183727   142808 92.60.45.23:62583",
  "32353263363135633662303932643438303832623839636562386563383234623537396362303536": "           ENABLED 70209 XibDWssbw82U5brXsu8o8sFDQb4UkJXnHw 518225357becba3b65963838ea4794982f990ff6ab31ee7830c38aa513235a0c 1598184007  2676817 165.22.198.139:62583",
  "37613761313563383232363661366138313137363466303361333062366530346366363161386138": "           ENABLED 70209 Xr4d4sXX4wHqsS3UYqi2RX7ixxgWZbytuH 36b90367fbe674e21a6d30477326ef0307831a548cf37d5b06306ada0c992431 1598183595 22071646 159.65.14.142:62583",
  "61363330383461323537353531326562346633323534383434313034396331653130626432613037": "           ENABLED 70209 XbLjm7yz49xzbybtD9zUPEwJj39cBeVsEo 581da85fa0b2778f0313f33fea04e8ccb036a8b8e368884d893a10fb914c26a6 1598183614 23644509 95.179.153.163:62583",
  "36313833353865333230666436636562656366646236353165633831383631303136633162313534": "           ENABLED 70209 XiQfXwy8GWNhssVsqvTBEQLzwk2JM6qp6X 6619b03e7b76aa1500274e9a5d6e0660c335c94cd5e5ab34435b7caad4f0776f 1598183595  1043800 92.60.45.149:62583",
  "33306164366163316664636639326234626531643137376335653435626533376130646435643333": "           ENABLED 70209 XfNSjwaSE2qYxEZ5cK2382SHyRcsMEFGne 31f9040f4754dce6b1af2d5bff9384cd9488e2517699a6e43a59a9544a4b334a 1598183974   702549 94.237.93.104:62583",
  "62343533393030633132666331373561333336623638653665333762393031313230646439393964": "           ENABLED 70209 Xq4AF1i7CSVmXeombDysJyZJkAWasZeG4z 595b7506bb914ecdc48c26453a2d1af236850d7f37a5288487f21a243f30ed6b 1598183999  1468445 94.237.54.125:62583",
  "36393263626262353735626365366539363465316237363061316637643134353436666334373466": "           ENABLED 70209 Xhv3T3YNetwfbxnt2obHZHcvCBH8c7Q1qh bc8b4b0aaaf6388dd7589497173f0917b823bd403297b17a01ed8bf3449b0c13 1598183640  1445527 92.60.45.22:62583",
  "63336133373161393839313463633737373966653436383034396532643138626334366332613132": "           ENABLED 70209 XcLtpSNp934KQ997FbEXxEWndDfwe1Ms1w b8456360d9ec47d91d1ec309ad2030c34ced3e24e5cb9168a6ff2ef48ec50ca0 1598183714  3712266 95.179.183.62:62583",
  "39303166613664616237383861626130303538303232333163333837333431666562633066356234": "           ENABLED 70209 XsBfs3N44tPjBAAKZyyXDHRSoz7uB9xhMN 75ffcd9ee3b7a2db5c07fa9b5c3fafc1f39e641b94b58f76a3923dd36181b45a 1598183602   938577 92.60.45.19:62583",
  "64323465366633336637333365383431653135383933626530636431353831396530633237333632": "           ENABLED 70209 XjfQjHSTkgB8ppuTtdTyiFfH9NvwTiWcfQ 41c7aa675b4ce518ee655b376f09f293da18a372ade188c89d07b3e941b8ed9d 1598183834  1626163 45.77.56.20:62583",
  "31353034363439653062326165326235613632333132313432363837353139316538346339363132": "           ENABLED 70209 XcP83mCjmQ9hxuDK3eHZEWmnGb45yXLHAQ 0fbfff04a05b402073ab3601fa8d521978fe151721034ad02b8e4e605873e30a 1598184052 23639512 92.60.45.20:62583",
  "35313331346434633135663061613437353166666261653665396363633139363035633234316533": "           ENABLED 70209 XwQTzjkDsb4GyA5pMekLiKcM6kUrVqYcWf 09a2c1ecd8eb063a3575c9e441af91edd8a44b883014cdf3162c85bfb02f3b21 1598183975  7622314 54.37.72.132:62583",
  "61376636373132313433646431613364616138393264373464376462363065373231656564323665": "           ENABLED 70209 Xknpr3QnddRN2zD54XN7qviRpuz8VTtjap 1a8295133c0ff249fa1ebab5a0ffb28d7347eb4fe95ae0171a101e98278bb8ee 1598183862  3868792 45.32.131.189:62583",
  "65323561363834343332666434366565336330323766396333306538333839396338373663326137": "           ENABLED 70209 XqysdtG8KuGUFGXHGTCaYksb4mpNbCfGB4 0856ffc5db3a0247fc40910420b02f642308ddbca158bd2b8601f025924534c1 1598183623 21930316 178.128.224.218:62583",
  "34363364663661626564616336343534333836306639346530383636303062346131663235393061": "           ENABLED 70209 XbdaLrS199UTjVGZYPy7GA1S511NaRvGvj 91b8f8527821c476791a5f69326b791043adf25739bc6764daee6d21579ce318 1598183863    12787 192.248.166.47:62583",
  "31643562386364343138363163346532323536623239313530636162663932336463333439333837": "           ENABLED 70209 Xo3hQEfUYroRhqLHMCMnb5C7ru6zMCHHUr de740f51481eef90481fabe87ce40e4a11d5d575600c2d3e54fb9c632a8b78b7 1598183956  7523437 80.240.19.26:62583",
  "63613237613166656362383465376338343537633730666136333338626132396465323333616664": "           ENABLED 70209 XymnSS1w9dDs4oKW8obPY9D1BjtXJ7Un9y d065f57e9cdb4dc10fd09510e756024b23ce9b1df4453c6f16e4da15468d07c5 1598183741 10467146 185.92.220.166:62583",
  "35646463396166303934366233643162636130373765633934373062363633316461343134653831": "           ENABLED 70209 XnUYkQbZ3CthVd3QZF5jdWDcbnxWwWSr62 f1ca075f04807f8aaf5cb12e9dfa54d22427e5ebc71fd5333727fa04cf10f143 1598183692  1463540 95.179.237.125:62583",
  "61643337366364353632636534356663383837366337626464353930656538646365626365613830": "           ENABLED 70209 XnSVXf4ZA2ERE1kFxxe3h3KG9Gqv1VZRLk 82d4d757dcb9f1961348e02943a0a1127f0c8426ba6c331a5a6ae214a3a0118f 1598183751  1466522 199.247.2.40:62583",
  "39303963626438316334323163373837663130656263623934343863346466643034626431353266": "           ENABLED 70209 XeyodgDpVHut6VHEHSNTL2NJiqgPhAcePE bc19fa74708c64b7bfc95e1017d582c51717f5d645dcd8ff820474ae8fde9456 1598183638   291009 192.248.174.185:62583",
  "32313263336531633463636635303134333165393163343166636463323961656166623232633962": "           ENABLED 70209 XpqL7f8PxWq1ADbyrY2M8GQbFGKgfoUTX4 b5ee7a655d7575b5c74955db93c77fefaf8fdadc8e408876d2f276097c538666 1598183854  2597563 92.60.45.28:62583",
  "62303035656565616631666638653361666562626636383965656539613135333237356338643136": "           ENABLED 70209 Xck63LVsyXecNGZGAzMAeWw5si5YFNhNHZ 0bfe7867e3bc749e6b3ddbd2bb4484dc6f9218b96e1ecb451bbb1508be0fc54e 1598183654   516740 108.61.164.170:62583",
  "63663362343130373132633032636636623563313132366366616463613962343639383339363065": "           ENABLED 70209 Xc1ybnhUhUYrcCjKCyBAiTnDp9ju894hFb ed9cc92cb50caf538371b21ae0f93d1e4e7411739bfe59ba75a5db1b3e96e4e6 1598183746   603711 80.240.22.149:62583",
  "37386665353166643934333265326537646161333437656436383636373563356532363066666433": "           ENABLED 70209 Xv1nLjejZXMMHQkuoZ5zSjU7CyeNDQwk8g 8d48668af15b76f2f1bb0f4c6971759e0348dfd80cabdadcef6deaaa726720a5 1598184154  3714081 95.179.129.236:62583",
  "39656539366461656137623039353932613439613964353030303037626436383739303931666334": "           ENABLED 70209 XtZqTLkj6MPbqD1PPwD1aJt6TTDrZFdPwW 5c91d525cd242f54868c550f26f5c418341f8b9f6b625888851486cb274ab5b6 1598183902 12437521 92.60.45.152:62583",
  "34333735303166303632353232333832303666633430373039356265356233336238313266396335": "           ENABLED 70209 XtjdKwwUdY6Ueg4DxEdbwxV82Z6vJF168a 2e352d3d38e57df690a4ebb3b8d08b97ac7155739dd24853e4af984ea98b82c6 1598183588  2411984 92.60.45.21:62583",
  "35363965343733376430613839663133393336316335333764313238303363303263333336353562": "           ENABLED 70209 Xj26ZChLL7ob6L7RuonhL5VVKjmZ3cmQZU 98d4373695856e1f1870d13f23075dc992d79e2ab6b6f44a2eb9b69ac371e318 1598183803   783877 92.60.45.25:62583",
  "30626165316131633364666336663764613435356434346630353463343833386236303965313035": "           ENABLED 70209 XbDvo84ao8Zt4TtB6BXbXHn9whmnvnFTy3 06586fae3777ff8e019ea2ec7a5dcce70cc8eb5216095edb298508462cfc594f 1598183723  1403185 192.248.145.101:62583",
  "35643835613634316538653235356234633834333233643439646136623736323138373630336536": "           ENABLED 70209 Xwf3PsSVifzb7JEJnk41CdHtkJjZWzkVQM edb9b2c8d548f405438e54ef209282e0b9a4b4aaddd966efe5585522591f054a 1598183813  8635803 45.63.84.43:62583",
  "65643134643632326366366530663861393330393266646139663764363632336564303834343863": "           ENABLED 70209 XoUVwFkvL1z5HDvJ8nkemhb9WvrwTYNA5j 420f08c57538af277c545ea585f3cb7141edc6530a5385771d9b6126b53010c7 1598183958  9848975 94.237.46.48:62583",
  "63386436346532646234376562333339323137343731323964356233353563666362366633306465": "           ENABLED 70209 XvwfsHFUuXriQELp5rnZSHr7tw3kDZ9hWa 0a2c16fe88065a44fadaa70fce69b7d0948e4e2e0af53e5d490804fffff84e89 1598183946  1297883 149.28.135.205:62583",
  "33653465303434323537343936393737663966626239323464393934623664346336636163643135": "           ENABLED 70209 Xcg8Yy4hEN581SLgrsmHdwMr3m44FYbM7V 68475958573e168fb43b45ab29efb6fa0fdeee475e391f3c9e8281ac3142a574 1598183901    12880 45.32.176.220:62583",
  "30353065383830633663656130653261386130393638643063623361643734373666316365666439": "           ENABLED 70209 XvZAucHr8BcyvDxDTyVqZngBwb2BgqGpDQ 9da68ad09fca4771e24698ee80241f0387a4bb990df8ab9b4448d68c36d4bd64 1598183616  3541327 209.250.244.131:62583",
  "66306366316338336333316536366330386465626164303833366337666264633638366464643562": "           ENABLED 70209 Xj4aakhPc1qgHCCDVj7ZUB3N5akx5ewMYC 93fd8f42d2d3304fb807dbe3d7af62d58499567bf8faf1b605f4f91e8d211731 1598183650  3532666 45.76.141.191:62583",
  "63303665373662613137353238376564333338373464343737663231653130346665633766646636": "           ENABLED 70209 XyCp5UL8oSTzgdHLNrDV8jwygPEsBpoeSe 85ad8c5f5c521eb291b00e7ceb9708492c22c82a2e9bb8eb682dccd046de5e9e 1598183677  2042385 95.179.201.159:62583",
  "63353933393535306366616561366163663037353961613830626164623265326332363966613963": "           ENABLED 70209 XpzsEBshbw52ErFkuMJrb8JyKVpBdh4ywW 7c1db0eeed6ec6850116085b688f0e1a485a685f8477b211085cc684dde51e9c 1598184134  7754739 140.82.39.52:62583",
  "63383038396635343331373631326339633435613638386262346531646238303862633061323239": "           ENABLED 70209 XeUzWLN2ZAswReeRws6UNvHmhLbo9MAMfh 7729e5f8c2e5c54f63577f4c130e73403ab92238b217df6d3c6a41693b5b3626 1598183696  2285727 144.202.68.234:62583",
  "30333232613461636638623163376461656364663235373861613033626564613231663462333063": "           ENABLED 70209 Xbr1WxGJiCNcYNMnq3poCJMt3BcHhLj2qP 72bec00c2bdc30b38b56a795873ee782b37be56321e87fe10af6f910fd3ccd78 1598183803  2084258 45.77.179.69:62583",
  "32646332663564373562343637353161363438343631636264373831326337636531646132666639": "           ENABLED 70209 XyQRR3TPgXMJaofg11ARkFeXDLRFeGGiCq 46022d17fd4091b697e374c39e1df5e137d91e7c506aafe6ce1f6b802e27cfb1 1598183801   307777 140.82.36.203:62583",
  "64663266613336633232666533613365393761336438343864616434353230386431663332373738": "           ENABLED 70209 XmeAnDRsoeRA6E4gA413KRw9PLRDNz6E95 0824d36dd70bb0eaa30c7c619e0eed95b2de03611befc7968ef82edf0d98b65a 1598183985 16982190 35.192.21.170:62583",
  "32373431303732663834623932383563333561616231383231343832646532616665656136363166": "           ENABLED 70209 XdYt54frzUPhT7vqWpJC2JZt3DW7PBBtsk 9acfd9989e38feaa0ac9d113cff25c62f1d1105a4834bf62aebb62565897e71d 1598183636  2406799 92.60.45.146:62583",
  "64633664373064643132313765383634623464646564353131333832653336363436323065303230": "           ENABLED 70209 Xdgfwwk9mB4NbPQWNqz7Wxa7Choz9BrQt7 ef92add830567677cd159fddf52721142f976f056efaca2f242fe91e44798a9e 1598183795  3715882 78.141.198.25:62583",
  "62613866323261356334343935643536353536306134373130376461356132376136396239623165": "           ENABLED 70209 XdUgWqnsgoyxheoKGMBLbEpKGCp82hqXbh 04b6c55e43f25cde710c353e054fe34ca8fa16a795a515f111fc0ae47357e63f 1598183950 18307126 95.179.154.119:62583",
  "38356364616634356133393136306635663464313934663836653330353435633739623338376636": "           ENABLED 70209 XyANd7szY9MPxjzbAxFCoSsx1VZWX4aRmq b506164a244e693e1ba81f68796f6fbe674f1cd1b7ae9293b0fc1d4d86b16a28 1598184167 19230808 92.60.45.24:62583",
  "36373535326132303533303532623161326466303930353535666565653239336439656233663231": "           ENABLED 70209 Xdiehuu8yoizh8xSjQpQDxWfWUriwizGkd 3cd357da3522ec6a844eb1e3994ee6007f823c04cfc3e4315dccf790c19bfe9f 1598184036   752792 51.178.171.9:62583",
  "35303531346266643264613938386534663131613332363866316436306266376563303930376265": "           ENABLED 70209 Xt1cemFL3NKbydi2ttdHBtNZgEVXL1SLGy 8cc28410ba146cce99287645cdc8b107c8e411cf92c32c3cb4ee3c25dc3afcb5 1598184103  1465555 149.248.62.156:62583",
  "34636639643833346362383838323064393066663036306432386639326131656562653335396563": "           ENABLED 70209 XxEYzDLCZt5r42SPPrtSjcNuaKn9QjHcQm 431c3832db9db38916e97933246f20d079f1473a61b01c7d3044b42a7a300b97 1598184043 22767726 45.32.144.2:62583",
  "65306361633164613463653664396537333139626237663362623061646365386531663737646663": "           ENABLED 70209 Xyhu22QQaijo98DpCk4gDq58mXuJ9EFFBS d453ed62fbdb17cbac560a583db8452ec0c48ce2e6dcd6177240d8cad71cde68 1598183886   377888 45.77.6.70:62583",
  "32363734353966646166306164313163373734313031393533663232623538366230366232383764": "           ENABLED 70209 Xn6cii5j4oMTtWDokw3WxtvqqMV2vkg3sy cc112af8f5d173d5b78e5500f29df6b2428169a4611915ed9cf97f2f70c3393c 1598183664  1350407 209.250.250.138:62583",
  "33383331326664326361383333306232373730373135366333333363323961353162303239656361": "           ENABLED 70209 XuABccCsmz6FK6nzHfCgQF6k8HRC74FiCf 98f1ac678b85048c18730c6df6391e53254691c74c3ecefd170edce31657dc2b 1598184182 20175376 92.60.45.148:62583",
  "37383338323935626563656465643530643934373932653063633762316465373836663132386231": "           ENABLED 70209 XrqaV1VMbwSq3894YGkya46gQ7SJQfsjCp 673a53f030d00490a0d77535bb3cd27a176915bc78f7345ac8f2512acb92e122 1598183728  4166579 45.76.132.176:62583",
  "30656633663165656232636134626232386465373766373434636537333936373330333338323739": "           ENABLED 70209 XmmKZv7boe5PZL6TbssFm7BjXSZCGjqGwG 18263817cb74ba97d08a94dbadf2ae4078b3eee3d937ef6923d99599c1a40697 1598183771  2997509 92.60.45.15:62583",
  "62336330613961663437613763373337303530376266306266636636323331383261656137303536": "           ENABLED 70209 XiZuDCg9D4rNSfiwDUn7skkEhxnzLYbZkH 8465443073051f8ed407f9f648f8ce8c463b24256d0b2bb3cb7e0c98776a0940 1598183587 14323003 78.141.217.113:62583",
  "64313838313134316236383864376232366237636662393330653736393136626463666363343632": "           ENABLED 70209 Xjh63117eqhgkZmq74VipHkoSqJ8obznqr ac894ccd76646c41bbab12c94fa77d5299d5670ddd23c6b38df72de73d4455ae 1598183895   199139 199.247.6.25:62583",
  "32303363653038646631353561636135656266626531316265353737623363363630626534613637": "           ENABLED 70209 Xk6zyHwKbmL7Rp7UxVRxV4EoVqaYLcg5y1 3c6f80250e13c17efa03cf981194a0775a09981bf75f5ce84968bdef82237349 1598183616  1445499 92.60.45.147:62583",
  "36383165613065623435373332353634303664656666653535303735616465343966306433363232": "           ENABLED 70209 XdojZTbTDmKJmoyqqTCx3YRiAjTpRTmw5Z 8b5c4b85456d05b5039e510b6c09e81a6436de60b0be8ef055f551afaa712fc2 1598183792  1465601 45.32.137.108:62583",
  "32653631323661316436353563633966353666396364303666323932656665386233626232353235": "           ENABLED 70209 Xe5G36oXdSFZR9xo7AP1CUacNYeET7WCNo a8a25422e87540135e0b87a4152a844558b3b06de29581060ebd940c6dfbef20 1598183753 15734955 54.37.74.98:62583",
  "30613338616564666338376139316133386666666465616336376636656661643631353930396231": "           ENABLED 70209 XrpvdoDCkycLiaMFkhEQ86GMAZ3UobA6ob d6921b9c321e63ba6f2e9884920d4b19466618ca43a58afe30f82b44d9e62bce 1598183940  4311861 155.138.139.241:62583",
  "33636234313136366664366363366139386133376262373037393139656538363937383862663438": "           ENABLED 70209 XhKVw3syTTXS48mGXWkRHEptXg7MPGJuA8 681bda37b0116d0777d7200f2d9c60ed242a2228ac46ceb5043e437a26c2fb40 1598183750 22086477 142.93.43.44:62583",
  "30346536333838373064353436346535333735343266366136386238393431383332633235356131": "           ENABLED 70209 XqPuMkNMU91gwUwxb2vkHrQmhMGWS9db2G 2b73c4713fa6b143d3301a45b0abe3cf026b1f9fed598027a05c143c78422f37 1598184006  1089555 95.179.179.59:62583",
  "66633265386431643331376661633062626537643933613735386531626132663363313936366230": "           ENABLED 70209 XrmZ4wzWAH8oBqfoJGnzgP3TmkxHbBJxGa 991cbe93ccb879e779c380d52e17a0081f1af177c74741f84b61a637170837a4 1598183801  9840765 173.249.14.252:62583",
  "33356266363132343564366333653064356662643361626661653863366165636561376338316461": "           ENABLED 70209 XvcCG44khz47zSS3kNoZCKh8Q1D1CirAkg 0981840a1d3ada85d346a6c62d6546bf6a5e4b662b97a4490d998003fd906441 1598184161   666006 95.179.218.107:62583",
  "61643837346331656433636362656536333833653265623332656265303030623564366661383632": "           ENABLED 70209 XjgVq85B4ShVtooz61riPAhFmRosYzseBk 107d2548f693c1d2379f0ae6eab546f18158c900bc04c9316c77d56df96026e3 1598184036  7133859 95.179.194.236:62583",
  "63353533653830376262663031313830343837336566393736663039663730353965626539366235": "           ENABLED 70209 XsEzitkfjSuiUawhLVDA4HxwVgbQmc6Xi6 e2372164ace797c9cefa5e94b15ddf851e51ecd2291b8b3b8f28378cf8a00ce0 1598183797 20177954 92.60.45.151:62583",
  "62373663373665383635353237613131333065343335373566656132663336626533323735356337": "           ENABLED 70209 XtrpKA5j431QSuG5FF2m5FeTnyWag8JBeW 141adc302c2c1bbe690eb3ee53c18bd4e0fc14807deeba3aa776d1edca19c232 1598183863    22757 139.180.131.137:62583",
  "64383832363937306430303335313265393866323864363961323636626233393133653638613732": "           ENABLED 70209 Xm8VFyWd8aCVVVyGozFG9h55Qnos1XbrN7 2d430564f1d33c09b20eb9db21bdc9b23df2d7755c20b48502ce826af8ec2dae 1598183674  3712194 95.179.199.191:62583",
  "61643463373765373338656235616435656563633134336633623137306563626537393636366530": "           ENABLED 70209 Xw9N6GL7QPyLP2Um4ZWh8eTH6D689ALF8D 23d4074d65f0920546066f577938142036c15a93e9d97019275bdf5aa8f4ae95 1598183859  4470862 95.216.100.13:62583",
  "37336132393261356436353830363736643438386163613037353461646537316131313535313761": "           ENABLED 70209 XmqbQWKDe1zCMxgSqnAcTrX6rG9yb8dMbV dce8e5a3c462ace3bee5e922003e85d527f42504a569b2288083726cb86b0dad 1598183622  7624977 51.68.191.98:62583",
  "32373838376534386238616665663232346335326265633232323236343636646137656639666339": "           ENABLED 70209 Xu4wFNQDmEETsbc1vWr2DnNimBv95yMsF2 200eb4472750c0845b780ab0c6fd26352f140f09aacca9b53b599a041ce2d8a0 1598184131  4457563 137.220.59.215:62583",
  "64316233656335343034623430636437643837366236353434393866393636383631303833353561": "           ENABLED 70209 XiupBJCRydXst9zQceP4iDraV6D9Ygv6WG e8124cda139dd599c600616789d61725a2e353017a567e02fd113807de0e2d6d 1598184022  1101235 92.60.45.27:62583",
  "35306338306163316233356133376535353564613632376262336137613830643730303135663432": "           ENABLED 70209 XgjnF28CLXnWQYGJMMzRzQrJdUPihG2e8G b5ea4dc6d87ada2beee36a36fe4c663a2e499c64fe08cb2d12b88b3014304f72 1598183967  1467856 108.61.167.184:62583",
  "30386639626233616637363733643166663163623438383165326339643465373730656339613565": "           ENABLED 70209 XjK4wZmdGpU3TBHjXYuPMS6EBbpZM7UTWB 955980f8830b2027633b4869a5ccc75dd94bc4e006ec34b1f77740aa499bf7f4 1598183926   925744 78.141.197.95:62583",
  "62653930663930393938316435373365636532366231333866313537313633626232666436393366": "           ENABLED 70209 XgU9Nis9snH67A43t9wuKTxaayNCMPAQhe 07bdfaa7d45f976094a67fc32f9c065732a6d5705f14ce5b895e70f634a886e9 1598183735   691182 217.69.4.220:62583",
  "35313261366661333932653332373561363434366336386232343062666365643438333638383665": "           ENABLED 70209 XkmHLXvZsDzm4r1gEA7Hkwc25mP62GbYHb a89be16bec642d744f4b905ee0eed3dadebb8365392f3ec164bacc508fb6c6d0 1598183937  1463803 209.250.255.116:62583",
  "34613162653263373166306361303363646161396362386135306261393033653136316166343533": "       PRE_ENABLED 70209 XiLkLfynUZTzP6dMBmtF2ZdUJL2y5sjcrx c2a494aaf7736803306b08825cdde2c3b6c65ba2f6f1ff12cc9159950fa7c7b2 1598181968        0 84.46.126.151:62583",
  "39323930376634343263646333306538323362643962646235396330393832303765636634323463": "           ENABLED 70209 Xhe5DufLj2wCn9KZ7UJ1TtKmU5mEh36wUc 7d0b7f981ae64c7064ec1537edeef934b6760c30ad48f25b77e04b4a678b4db4 1598184094  9634095 45.76.131.63:62583",
  "39633932656139303234643062383565633362646261653233393566613832303634663631646636": "           ENABLED 70209 Xy8BxGkNujvCmVAJ8PVK7fXdHWD4UXGooz eaea105551b50965989bcf9dc0c3821b21e2eadee3124e2e6cbe370e4382f426 1598183787  3679116 92.60.45.26:62583",
  "35623861613339613731613339346639663633353031323335356533633838333033346462333334": "           ENABLED 70209 XfVVmGZMSTFvR4rvxGiL2amFcD2kdc9w7X 5a537b06e3238aff61a45b0768270e078f3fbf548231b3d4c3adb12bb31e7b8e 1598184071 34369523 45.79.218.6:62583"
}
`, nil
}

func TestXSNMerchantListCLI(command string, arguments ...string) (string, error) {
	return TestXSNMerchantList(), nil
}
