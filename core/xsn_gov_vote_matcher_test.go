package xsn

import (
	"fmt"
	"testing"
)

func TestGetVotesByMN(t *testing.T) {
	//tstAddr := "XwUFMpHMRjk8A7XpNT2EkRXxZKahfyfnhn" // owner with lots of masternodes
	//tstAddr := "Xfg33S8bvaXcek9maDm2v9HbxtnwRARZU8"
	tstAddr := "Xn6hrcuk1YSDw6Xqa8b5DeMa2ZKYnAMxAM"
	cli := "/home/wintan/Desktop/QTs/xsn-cli"
	votes, err := GetVotesByMN(tstAddr, cli, false)
	fmt.Printf("guessed votes: %v \n", votes.Totals.TotalGuessedMatch)
	fmt.Printf("direct votes: %v \n", votes.Totals.TotalDirectMatch)
	fmt.Printf("direct votes objects:  %+v \n", votes.Totals.DirectMatches)
	fmt.Printf("err: %v \n", err)
}
