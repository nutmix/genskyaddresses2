package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/skycoin/skycoin/src/cipher"
	bip39 "github.com/skycoin/skycoin/src/cipher/go-bip39"
)

func main() {
	const SEEDS_FILE = "seeds.csv"
	const ADDR_FILE = "addresses.txt"

	n := flag.Int64("n", 1, "number of address need to generate")
	entropy := flag.Int("e", 128, "bip39 entropy, can be 128, 256")
	outputPath := flag.String("path", "./", "file output path e.g. ~/ or ./")

	flag.Parse()

	fSeeds := openFile(*outputPath, SEEDS_FILE)
	fAddresses := openFile(*outputPath, ADDR_FILE)

	for i := int64(0); i < *n; i++ {
		seed, err := bip39Seed(*entropy)
		if err != nil {
			panic(err)
		}
		_, seckey := cipher.GenerateDeterministicKeyPair([]byte(seed))
		addr := cipher.AddressFromSecKey(seckey)
		//		fmt.Printf("\"%s\" \"%s\"\n", addr, seed)
		fSeeds.WriteString(fmt.Sprintf("\"%s\",\"%s\"\n", addr, seed))
		fAddresses.WriteString(fmt.Sprintf("\"%s\",\n", addr))
	}

	fSeeds.Close()
	fAddresses.Close()
}

func bip39Seed(n int) (string, error) {
	entropy, err := bip39.NewEntropy(n)
	if err != nil {
		return "", fmt.Errorf("generate bip39 entropy failed, err:%v", err)
	}

	seed, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("generate bip39 seed failed, err:%v", err)
	}

	return seed, nil
}

/** openfile will panic if cant create file **/
func openFile(outputPath string, name string) *os.File {
	f, err := os.Create(outputPath + name)
	if err != nil {
		fmt.Println("could not create file:" + outputPath + name)
		panic(err)
	}
	return f
}
