package core

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestRegex(t *testing.T) {
	data := GetTestMNLIST()

	var mnRegex = regexp.MustCompile(`(?m)\s+"COutPoint\((\w+), \d\)"\:\s"((\w+|\s+\w+))\s(\d+)\s+(\w+)\s(\d+)(\s+(\d+)\s+)(\d+)\s+(\d+)\s+(\w+.\w+.\w+.\w+:\w+)`)

	rows := strings.Split(data, `",`)

	for _, r := range rows {
		for _, match := range mnRegex.FindAllStringSubmatch(r, -1) {
			fmt.Printf("match1 : %v \n", match[1])
			fmt.Printf("match2 : %v \n", match[2])
			fmt.Printf("match3 : %v \n", match[3])
			fmt.Printf("match4 : %v \n", match[4])
			fmt.Printf("match5 : %v \n", match[5])
			fmt.Printf("match6 : %v \n", match[6])
			fmt.Printf("match7 : %v \n", match[7])
			fmt.Printf("match8 : %v \n", match[8])
			fmt.Printf("match9 : %v \n", match[9])
			fmt.Printf("match10 : %v \n", match[10])
			fmt.Printf("match11 : %v \n", match[11])

		}
	}
}

func TestXSNMNParse(t *testing.T) {
	data, err := GetXSNMNList("test", true)
	if err != nil {
		t.Fail()
		return
	}
	testMN := data[0]
	if testMN.IP != "149.28.141.200:62583" {
		t.Fail()
		return
	}
	if testMN.Protocol != "70209" {
		t.Fail()
		return
	}
	if testMN.Output != "6067a102d131ea27fd4a3fb8c6b189e9f809a42398457df8a5c5f4ba9e650200" {
		t.Fail()
		return
	}
	if testMN.Address != "XypnJqfangXpSFVe3wYcKjQb4UzkNVjE63" {
		t.Fail()
		return
	}
}
