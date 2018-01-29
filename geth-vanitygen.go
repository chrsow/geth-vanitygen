package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"regexp"

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

//validate word
func validateWord(word string) {
	r, _ := regexp.MatchString(`^[0-9a-fA-F]+$`, word)
	if !r {
		fmt.Printf("[-] %s: is not a valid hexadecimal.\n", word)
		os.Exit(1)
	} else if len(word) > 40 {
		fmt.Println("[-] You can't generate matching Ethereum address for more than 40 characters (20 bytes).")
		os.Exit(1)
	}
}

func generateAccount() string {
	// 1. generate private key, ECDSA(private key)  => public key
	key, _ = crypto.GenerateKey()
	pubKey := key.PublicKey
	// 2. public key => address
	address := crypto.PubkeyToAddress(pubKey)
	addressHex := hex.EncodeToString(address[:])
	return addressHex
}

// TODO: add time being used so user will know how long they spend time on the genearation
// TODO: add some cli emo or loading bar
// searchAddress: search eth address
func searchAddress(word string, searchType string) (string, string) {
	length := len(word)
	found := false
	var address string

	switch searchType {
	case "prefix":
		for !found {
			address = generateAccount()
			if address[:length] == word {
				found = true
			}
		}
	case "suffix":
		for !found {
			address = generateAccount()
			if address[40-length:] == word {
				found = true
			}
		}
	default:
		address = generateAccount()
	}

	privateKey := hex.EncodeToString(crypto.FromECDSA(key))
	return address, privateKey
}

func foundAddress(address string, privateKey string) {
	fmt.Printf("Address: 0x%s\n", address)
	fmt.Printf("PrivateKey: %s\n\n", privateKey)
}

// prefix, suffix from cli
var prefix stringFlag
var suffix stringFlag
var key *ecdsa.PrivateKey

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
		validateWord(word)
		address, privateKey := searchAddress(word, "prefix")
		fmt.Printf("[+] Address with prefix \"%s\" found.\n", word)
		foundAddress(address, privateKey)
	} else if suffix.set { // Suffix searching
		word = suffix.value
		validateWord(word)
		address, privateKey := searchAddress(word, "suffix")
		fmt.Printf("[+] Address with suffix \"%s\" found.\n", word)
		foundAddress(address, privateKey)
	} else { // Default searching
		address, privateKey := searchAddress(word, "")
		fmt.Println("[+] Generate Ethereum address")
		foundAddress(address, privateKey)
	}
}
