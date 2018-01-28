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
.....

// search for eth. address with suffix 00
# geth-vanitygen -s 00
```