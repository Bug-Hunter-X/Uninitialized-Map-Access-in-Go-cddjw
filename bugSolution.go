```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["key"] = 10 // Now this is safe
    fmt.Println(m["key"])
}
```