# go-slices v1.0.0

Ini adalah package bahasa pemrograman go untuk manajemen golang slices.

### func Shift()
```go
func (s Slices) Shift() Slices
```

### func Pop()
```go
func (s Slices) Pop() Slices
```
### func Remove(e)
```go
func (s Slices) Remove(element int) Slices
```

### Instalasi

`go get -u github.com/michaelwp/go-slices`

### Contoh penggunaan

```go 
package main

import (
	"fmt"
	goslices "github.com/michaelwp/go-slices"
)

func main()  {
	var slices = goslices.Slices{1, 2, 3, 4, 5, 6}

	fmt.Println(slices.Shift())
	fmt.Println(slices.Pop())
	fmt.Println(slices.Remove(2))
}
```

#### Hasil

```text
[2, 3, 4, 5, 6]
[1, 2, 3, 4, 5]
[1, 2, 4, 5, 6]
```