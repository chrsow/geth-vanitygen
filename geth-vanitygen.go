package main

import (
	"flag"
	"fmt"
	"os"
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

// searchPrefix: search eth address with prefix `word`
func searchPrefix(word string) {

}

// searchSuffix: search eth address with suffix `word`

func searchSuffix(word string) {

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
		fmt.Println("[-] Search for prefix and suffix at the same time is not support now.")
		os.Exit(0)
	} else if prefix.set {
		word = prefix.value
		fmt.Println(word)
	} else if suffix.set {
		word = suffix.value
		fmt.Println(word)
	} else {

	}

}
