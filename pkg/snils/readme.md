SNILS CheckAndCount
================================

// Realising Algorithm of checking the control number of the Insurance number (SNILS)

According to http://www.consultant.ru/document/cons_doc_LAW_142584/1d9155a863a5949b14b95ecbb536aa84856a2a2e/

Manual check SNILS http://portal.fss.ru:8585/fss/lgot/snils

Play code https://go.dev/play/p/IMrxtpMOagJ

## Example

```go
package main
import (
  "fmt"
  
  "github.com/vasiunov/features/pkg/snils"
)
func main() {
      
	snils1 := Snils("11223344595")                               // true <nil>
	snils2 := Snils("112-233-445 95")                            // true <nil>
	snils3 := Snils("112 - - -233 - - - -445 - - - -95")         // true <nil>
	snils4 := Snils("112 - - -233")                              // false invalid SNILS format
	snils5 := Snils("112 - - -233dd1d31dc - - - -445 - - - -95") // false invalid SNILS format
	snils6 := Snils("112 - - -233 - - - -445 - - - -96")         // false <nil>

	fmt.Println(checkSnils(snils1))
	fmt.Println(checkSnils(snils2))
	fmt.Println(checkSnils(snils3))
	fmt.Println(checkSnils(snils4))
	fmt.Println(checkSnils(snils5))
	fmt.Println(checkSnils(snils6))
	
}
```