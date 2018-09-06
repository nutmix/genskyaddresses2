# genskyaddresses2
generates a file of addresses in a format to be pasted, and a csv file of the same addresses with their seeds.

The address file is called "addresses.txt" and looks like this:

    "yKUF4QDBmFF17hqhvc67SDqbXsaCamQWtA",
    "8cRizt43QqDQPrAdnTKpd4uBuvg13mAsY3",
    "2ibaTMkRKfGh3H18gKieJV7jhDktkcpZc2K",
etc.

This is ready to paste into the fiber.toml file.

The seed file is called seeds.csv and the contents look like this:

    "yKUF4QDBmFF17hqhvc67SDqbXsaCamQWtA","affair time fragile rival vapor betray sponsor cricket victory fish turtle twice"
    "8cRizt43QqDQPrAdnTKpd4uBuvg13mAsY3","idea extend fresh effort kind gauge trick hawk eight dry design vital"
    "2ibaTMkRKfGh3H18gKieJV7jhDktkcpZc2K","leader easily birth evoke beauty drastic judge layer warfare bench obtain tooth"

This allows the coin owner to give one or more seeds to the new coin owners. The new coin owners can import the seed(s) into the skycoin wallet.

# installation

    go get github.com/nutmix/genskyaddresses2

# Usage:

    $ genskyaddresses2 -help
    Usage of genskyaddresses2:
    -e int
            bip39 entropy, can be 128, 256 (default 128)
    -n int
            number of address need to generate (default 1)
    -path string
            file output path e.g. ~/ or ./ (default "./")


e.g.

    $ genskyaddresses2 -n 100
    $ ls
    seeds.csv
    addresses.txt


Note, this is a modified version of:

    https://github.com/iketheadore/genskyaddress
