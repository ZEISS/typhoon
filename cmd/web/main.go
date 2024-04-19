package main

import (
	"github.com/zeiss/typhoon/cmd/web/cmd"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		panic(err)
	}
}
