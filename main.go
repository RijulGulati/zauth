package main

import (
	"github.com/rijulgulati/zauth/internal/args"
	"github.com/rijulgulati/zauth/internal/common"
)

func main() {
	zc := &common.ZAuthCommon{}
	args.ParseArgs(zc)
}
