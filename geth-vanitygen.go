package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

// == BEGIN: FLAG PARSE CHECKING ==
// stringFlag:
type stringFlag struct {
	set   bool
	value string
}

// Set: for for checking an existance of the flag
func (sf *stringFlag) Set(x string) error {
	sf.value = x
	sf.set = true
	return nil
}

//
func (sf *stringFlag) String() string {
	return sf.value
}

// == END: FLAG PARSE CHECKING ==

func genAccount() string {
	// 1. generate private key, ECDSA(private key)  => public key
	key, _ := crypto.GenerateKey()
	pubKey := key.PublicKey
	// 2. public key => address
	address := crypto.PubkeyToAddress(pubKey)
	addressHex := hex.EncodeToString(address[:])
	return addressHex
}

// searchPrefix: search eth address with prefix `word`
func searchPrefix(word string) (string, string) {
	address := genAccount()
	fmt.Println(address)
	return "123", "12414"
}

// searchSuffix: search eth address with suffix `word`
func searchSuffix(word string) (string, string) {
	address := genAccount()
	fmt.Println(address)
	return "", ""
}

func deafaultSearch(word string) (string, string) {
	address := genAccount()
	fmt.Println(address)
	return "", ""
}

func foundAddress(addr string, privKey string) {
	fmt.Printf("Address: 0x%s\n", addr)
	fmt.Printf("PrivateKey: 0x%s\n", privKey)
}

// prefix, suffix from cli
var prefix stringFlag
var suffix stringFlag

func init() {
	flag.Var(&prefix, "p", "")
	flag.Var(&suffix, "s", "")
}

func main() {
	var word string
	flag.Parse()
	if prefix.set && suffix.set {
		fmt.Println("[-] Search for prefix and suffix at the same time is not supported now.")
		os.Exit(1)
	} else if prefix.set { // Prefix searching
		word = prefix.value
		addr, privKey := searchPrefix(word)
		fmt.Printf("[+] Address with prefix %s found.\n", word)
		foundAddress(addr, privKey)
	} else if suffix.set { // Suffix searching
		word = suffix.value
		addr, privKey := searchSuffix(word)
		fmt.Printf("[+] Address with suffix %s found.\n", word)
		foundAddress(addr, privKey)
	} else { // Default searching
		// panic(fmt.Sprintf("[-] "))
	}
}
