package core

import (
	"fmt"
	"testing"
)

func TestGovernanceTotals(t *testing.T) {
	govTotal := GovTotalToObject(getTestGovernanceTotals())
	if govTotal.Objects != 15 {
		fmt.Printf("govTotal.Objects: %v", govTotal.Objects)
		t.Fail()
		return
	}
	if govTotal.Proposals != 12 {
		fmt.Printf("govTotal.Proposals: %v", govTotal.Proposals)
		t.Fail()
		return
	}
	if govTotal.Votes != 3272 {
		fmt.Printf("govTotal.Proposals: %v", govTotal.Proposals)
		t.Fail()
		return
	}
}

func TestGetFullVoteHistory(t *testing.T) {
	votes, err := GetFullVoteHistory("", true)
	fmt.Println(len(votes))
	if err != nil {
		fmt.Printf("err: %v", err)
		t.Fail()
		return
	}
	if len(votes) != 8075 {
		t.Fail()
		return
	}
}

func TestVotes(t *testing.T) {
	str := getTestVotes()
	votes, err := ParseVotes(str, []MasternodeItem{})
	if err != nil {
		t.Fail()
		return
	}
	vote := votes.Votes[0]
	if vote.Epoch != 1535305644 {
		fmt.Printf("[err on Epoch]- full: %v \n", vote)
		t.Fail()
		return
	}
	if vote.Vote != "YES" {
		fmt.Printf("[err on Vote]- full: %v \n", vote)
		t.Fail()
		return
	}
	if vote.Type != "DELETE" {
		fmt.Printf("[err on DELETE]- full: %v \n", vote)
		t.Fail()
		return
	}
	if vote.Output != "2cbc606fe775f513d2a4f20d87c375f3d28d5bc47680b991b22e0fd963c7ab2c" {
		fmt.Printf("[err on Output]- full: %v \n", vote)
		t.Fail()
		return
	}
	if vote.Addr != "" {
		fmt.Printf("[err on Addr]- full: %v \n", vote)
		t.Fail()
		return
	}
}

func getTestGovernanceTotals() string {
	return `Governance Objects: 15 (Proposals: 12, Triggers: 3, Watchdogs: 0/0, Other: 0; Erased: 2), Votes: 3272`
}

func getTestVotes() string {
	return ` "272b0dc8063ef4f6022af6c9cedb28da036814e5c2d31ace5c6d9b3b5097b728": "CTxIn(COutPoint(2cbc606fe775f513d2a4f20d87c375f3d28d5bc47680b991b22e0fd963c7ab2c, 0), scriptSig=):1535305644:YES:DELETE",
  "515ea3b3dea314ec38ae836e2b9545ca6e03bf78a89476e8a2c8dd0aed651bf7": "CTxIn(COutPoint(918247cff0be955c8c770c52033e7e1fa93b7a6b82aadc0ace24b0f6b3e05416, 1), scriptSig=):1535305271:YES:DELETE",
  "121e9bcb486df6bf701fabb5526bb6288b05a86fd33d8636d58822698db1f35d": "CTxIn(COutPoint(31eb4001098c40f1f712c4a36defa55abae40d614726f324769e31242e7584b9, 0), scriptSig=):1535304955:YES:DELETE",
  "6e205a51e273e76ec149ebeb99a58d8baac071fdeb5f1c8b6e600104756472fa": "CTxIn(COutPoint(5f75a098a5d547a7599a150d6980db3e55485977d49535addb25e1bbbe96fd27, 1), scriptSig=):1531091420:YES:FUNDING",
  "d83440adf6b2c465abf38ade1cba04694e06f4a2e3d49a5fe03981ad624b04b9": "CTxIn(COutPoint(0403f900a4e52383365d2dbe0e238f914ad7d6aaf2327cab369aab1066d09f49, 1), scriptSig=):1530626702:YES:FUNDING",`
}
