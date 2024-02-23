package main

import (
	"github.com/zeiss/typhoon/pkg/routing/adapter/common/sharedmain"
	"github.com/zeiss/typhoon/pkg/routing/adapter/splitter"
)

func main() {
	sharedmain.MainWithController(splitter.NewEnvConfig, splitter.NewController, splitter.NewAdapter)
}
