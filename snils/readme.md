SNILS checkAndCount
================================

Realising Algorithm for generating the control number of the Insurance number (SNILS)
According to http://www.consultant.ru/document/cons_doc_LAW_142584/1d9155a863a5949b14b95ecbb536aa84856a2a2e/
Manual check SNILS http://portal.fss.ru:8585/fss/lgot/snils

## Example

```go
package main
import (
  "fmt"
  
  "github.com/vasiunov/features"
)
func main() {
       
	fmt.Println(snils.checkAndCount("112-233-445 95"))
	fmt.Println(snils.checkAndCount("112-233-446 96"))
	fmt.Println(snils.checkAndCount("112-243-445 00"))
	fmt.Println(snils.checkAndCount("112-243-446 00"))
	fmt.Println(snils.checkAndCount("112-243-447 01"))
	fmt.Println(snils.checkAndCount("112-243-447 0w"))
	fmt.Println(snils.checkAndCount("112-243-447 0"))
	fmt.Println(snils.checkAndCount("112-243-447 023"))
	fmt.Println(snils.checkAndCount("112-243-447 0ц"))
	
}
```