package main

import (
  "fmt"
)

func main() {
  fmt.Println("hello")
}
/*
workspace
bin #executable binaries
pkg #compile object files
src #source code
we create just the src directory, the rest is maintain be the golang(bin and pkg)
we set the gopath for the workspace
export GOPATH=$HOME/go/get-started/
also the bin directory
export PATH=$PATH:GOPATH/bin
choose a namespace (github.com/claudiug) as gihub and username
mkdir -p src/github.com/claudiug/hello

*/

