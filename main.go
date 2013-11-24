package main

import (
	. "github.com/lxn/walk/declarative"
)

func main() {
	mw := new(drawWindow)

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Walk Drawing Example",
		MinSize:  Size{320, 240},
		Size:     Size{800, 600},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			CustomWidget{
				AssignTo:            &mw.paintWidget,
				ClearsBackground:    true,
				InvalidatesOnResize: true,
				Paint:               mw.drawStuff,
			},
		},
	}).Run(); err != nil {
		panic(err)
	}
}
