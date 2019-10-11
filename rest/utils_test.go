package rest

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"testing"
)

func TestSelectUtxos(t *testing.T) {
	val := 101
	cli := NewRestCli("http://172.168.3.77:18443", "test", "test", "138.91.6.125:50071")
	utxos, err := cli.ListUnspent(1, 710, "mjEoyyCPsLzJ23xMX6Mti13zMyN36kzn57")
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}

	selected, sum, err := SelectUtxos(utxos, int64(val*btcutil.SatoshiPerBitcoin))
	if err != nil {
		t.Fatalf("failed to select: %v", err)
	}

	for _, v := range selected {
		fmt.Printf("%v", v)
	}
	fmt.Printf("sum is %d, val is %d\n", sum, int64(val*btcutil.SatoshiPerBitcoin))
}

func TestSth(t *testing.T) {
	str := "76a91428d2e8cee08857f569e5a1b147c5d5e87339e08188ac"
	b, _ := hex.DecodeString(str)
	s, err := txscript.DisasmString(b)
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Println(s)
}
