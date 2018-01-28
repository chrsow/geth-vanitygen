# go-ethereum-vanitygen
Vanitygen for Ethereum. Supports for specific prefix or suffix. 
The purposes are for learning Golang and generating Ethereum address for PoC of short address attack on Ethereum.

## Install
```
go get https://github.com/chrsow/geth-vanitygen
```

## Features
`prefix` searchig address for specific prefix

`suffix` searchig address for specific suffix

## Example
```
// search for eth. address with prefix 555
# geth-vanitygen -p 555
[+] Address with prefix 555 found.
Address: 0x555.....
PrivateKey: 3b.....

// search for eth. address with suffix 00
# geth-vanitygen -s 00
[+] Address with suffix 00 found.
Address: 0x2f.....00
PrivateKey: 81.....

// no option, search for default eth. address
Address: 0x25.....
PrivateKey: 5f.....

Address: 0x49.....
PrivateKey: e6.....
```