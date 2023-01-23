SNILS CheckAndCount
================================

Realising Algorithm for generating the control number of the Insurance number (SNILS)

According to http://www.consultant.ru/document/cons_doc_LAW_142584/1d9155a863a5949b14b95ecbb536aa84856a2a2e/

Manual check SNILS http://portal.fss.ru:8585/fss/lgot/snils

## Example

```go
package main
import (
  "fmt"
  
  "github.com/vasiunov/features/pkg/snils"
)
func main() {
       
	fmt.Println(snils.CheckAndCount("112-233-445 95")) // true <nil>
	fmt.Println(snils.CheckAndCount("112-243-445 00")) // true <nil>
	fmt.Println(snils.CheckAndCount("112-243-446 00")) // true <nil>
	fmt.Println(snils.CheckAndCount("112-243-447 04")) // false <nil>
	fmt.Println(snils.CheckAndCount("112-243-447 0w")) // false недопустимый символ в СНИЛС
	fmt.Println(snils.CheckAndCount("112-243-447 0")) // false неверная длинна СНИЛС
	fmt.Println(snils.CheckAndCount("112-243-447 023")) // false неверная длинна СНИЛС
	fmt.Println(snils.CheckAndCount("112-243-447 0ц")) // false неверная длинна СНИЛС
	
}
```