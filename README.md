# TMPhone
Inspired by KNphone and MLPhone repository, TMPhone is a phonetic algorithm for indexing Tamizh words by their pronounciation like Metaphone for English. Similar to KNPhone and MLPhonem this algorithm generates three romanized phonetic keys (hashes) of varying phonetic affinities for a given Tamizh word. This package is implemented in go.

- `key0` = a broad phonetic hash comparable to a Metaphone key that doesn't account for hard sounds and phonetic modifiers
- `key1` = is a slightly more inclusive hash that accounts for hard sounds.
- `key2` = highly inclusive and narrow hash that accounts for hard sounds and phonetic modifiers.

### Examples

| Word       | Pronunciation | key0    | key1    | key2      |
| ---------- | ------------- | ------- | ------- | --------- |
| திங்கள்      | tiṅkaḷ        | THNKL   | THNKL1  | TH4NKL1   |
| தண்ணீர்     | thaṇṇīr       | THNR    | THN2R   | THN24R    |
| சிப்பாய்      | sippāy        | SPY     | SP2Y    | S4P23Y    |
| வௌவால்    | vauvāl        | VVL     | VVL     | V9V3L     |

### Go implementation

Install the package:
`go get -u github.com/knadh/knphone`

```go
package main

import (
	"fmt"

	"github.com/knadh/knphone"
)

func main() {
	k := knphone.New()
	fmt.Println(k.Encode("ಅಂಕೆಸಂಖ್ಯೆ"))
	fmt.Println(k.Encode("ಊಱಿಸಾಱು"))
}

```

License: GPLv3