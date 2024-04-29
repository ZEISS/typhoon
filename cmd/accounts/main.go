package main

import (
	"github.com/zeiss/typhoon/cmd/accounts/cmd"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		panic(err)
	}
}
