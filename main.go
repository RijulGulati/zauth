package main

import (
	"github.com/grijul/zauth/internal/args"
	"github.com/grijul/zauth/internal/common"
)

func main() {
	zc := &common.ZAuthCommon{}
	args.ParseArgs(zc)
}
