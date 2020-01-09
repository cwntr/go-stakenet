package xsn

import (
	"testing"
	"fmt"
)

func TestXSNMerchantParse(t *testing.T) {
	merchants := ParseXSNMerchants(TestXSNMerchantList())
	x := merchants[0]

	if x.ID != "65643362343838323665646639656362303333646265633335623036356234613965366430356435" {
		fmt.Printf("[err on id]- full: %v \n", x)
		t.Fail()
		return
	}
	if x.Status != MNStatusEnabled {
		fmt.Printf("[err on status]- full: %v \n", x)
		t.Fail()
		return
	}
	if x.Protocol != "70209" {
		fmt.Printf("[err on protocol]- full: %v \n", x)
		t.Fail()
		return
	}
	if x.MerchantAddress != "Xv7CGGg4ozrFLHXYqtxoGwBnsRTdyfaX3G" {
		fmt.Printf("[err on merch-addr]- full: %v \n", x)
		t.Fail()
		return
	}
	if x.HashTPoSContractTx != "4fe3e1c0db4f276d08ac9100f4591413a9e200152e4b096665f9057a5bdd9fb5" {
		fmt.Printf("[err on hash-contract]- full: %v \n", x)
		t.Fail()
		return
	}
	if x.LastSeen != 1575452671 {
		fmt.Printf("[err on last-seen]- full: %v \n", x)
		t.Fail()
		return
	}
	if x.ActiveSeconds != 912259 {
		fmt.Printf("[err on active-seconds]- full: %v \n", x)
		t.Fail()
		return
	}
	if x.IP != "78.141.218.19:62583" {
		fmt.Printf("[err on ip]- full: %v \n", x)
		t.Fail()
		return
	}
}
