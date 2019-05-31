package main

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/bech32"
	"os"
)

func main() {
	_, bz, err := bech32.DecodeAndConvert(os.Args[1])
	if err != nil {
		panic(err)
	}
	addr, err := bech32.ConvertAndEncode("xrn:", bz)
	if err != nil {
		panic(err)
	}
	fmt.Println(addr)
}
