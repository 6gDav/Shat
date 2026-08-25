package main

import (
	"Shat/server/logic"
	"embed"
)

//go:embed all:page/build
var staticFS embed.FS

func init() {
	logic.SetStaticAssets(staticFS)
}
