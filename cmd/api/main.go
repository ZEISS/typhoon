package main

import (
	"github.com/zeiss/typhoon/cmd/api/cmd"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		panic(err)
	}
}
