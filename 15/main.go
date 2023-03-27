/*
Данный фрагмент кода может привести к тому, что создается огромная строка v размером в 1 МБ,
а затем в переменную justString записываются первые 100 символов из этой строки.
 Таким образом, выделяется лишняя память на хранение огромной строки,
 которая после этого никогда не будет использоваться.

Чтобы исправить эту проблему, можем использовать bufio.Reader и читать строку по фрагментам вместо того,
 чтобы загружать ее целиком в память.

*/

package main

import (
	"bufio"
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	rdr := bufio.NewReader(strings.NewReader(v))
	justString, _ = rdr.ReadString('\n')
}

func main() {
	someFunc()
	fmt.Println(justString)
}
