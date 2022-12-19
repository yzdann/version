Version package inspired by [hashicorp/vault/version](https://github.com/hashicorp/vault/blob/main/version/version.go)

## Usage
```go
package main

import (
        "fmt"

        "github.com/yzdann/version"
)

func main() {
        version.Version = "1.1.1"
        version.GitCommit = "abcd"
        version.BuildDate = "1970-01-01T01:01:01Z"
        v := version.GetVersion().FullVersionNumber()
        fmt.Print(v)
}
```

### Run
```bash
$ go run ./main.go
v1.1.1 (abcd), built 1970-01-01T01:01:01Z
```