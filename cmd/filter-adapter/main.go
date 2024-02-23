package main

import (
	"github.com/zeiss/typhoon/pkg/routing/adapter/common/sharedmain"
	"github.com/zeiss/typhoon/pkg/routing/adapter/filter"
)

func main() {
	sharedmain.MainWithController(filter.NewEnvConfig, filter.NewController, filter.NewAdapter)
}
