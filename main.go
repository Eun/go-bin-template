package main

import (
	"fmt"

	"github.com/Eun/go-bin-template/cmd/root"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	root.Version(fmt.Sprintf("%s %s %s", version, commit, date))
	root.Execute()
}
