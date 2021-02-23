# Bishop
Maybe it's not a Queen yet but it's still something something better than pawn to write SA:MP gamemodes. Do it with Go!

# What is this anyway?
This is a wrapper for [sampgdk](https://github.com/Zeex/sampgdk) adjusted for go users.
All the files in the root directory (except `README.md` and `go.mod` are generated with `contrib/generate.sh` script.
Basically, swig and some code generation.

# TODO
- [ ] Windows
- [ ] Argument names in `bishop.go` file (SWIG lacks such functionallity)

# Example
``` go
package main

import (
	"fmt"

	"github.com/Savolro/bishop"
)

func init() {
	bishop.CallbackOnGameModeInit = func() bool {
		// TODO: s/worst/best
		fmt.Println("Hello from the worst gamemode ever created!")
		return true
	}
}

func main() {}
```

# Building
## Linux

``` sh
GOARCH=386 CGO_ENABLED=1 go build -buildmode=c-shared -o yourgamemode.so
```

# How to use
1. Download SA:MP server from [here](https://www.sa-mp.com/download.php). You can delete all filterscripts, scriptfiles, gamemodes, etc.
1. Create empty gamemode `.amx` file to `gamemodes` sub-directory
1. Create sub-directory `plugins` and place yourgamemode.so file in there
1. Configure your `server.cfg` file
1. Add such line to your server.cfg (on Linux):
```
plugins yourgamemode.so
```
