## Base64

Пакет для работы с кодировкой base64.

```go
package main

import (
	"fmt"
	"github.com/web-vovan/base64"
)

func main() {
	// Кодирование
	fmt.Println(base64.Encode("hello, world!")) // aGVsbG8sIHdvcmxkIQ==

	// Декодирование
	fmt.Println(base64.Decode("aGVsbG8sIHdvcmxkIQ==")) // hello, world!
}

```