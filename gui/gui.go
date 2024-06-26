package gui

import (
	gDraw "biliget/gui/draw"

	g "github.com/AllenDang/giu"
)

func RunGiuMain() {
	w := g.NewMasterWindow("Biliget", 800, 600, g.MasterWindowFlagsNotResizable)
	w.Run(loopMain)
}

func loopMain() {
	g.SingleWindow().Layout(
		g.PrepareMsgbox(),
		g.TabBar().TabItems(
			g.TabItem("下载"),
			g.TabItem("下载队列"),
			g.TabItem("设置"),
		),
		g.Custom(gDraw.DrawLogin),
	)
}
