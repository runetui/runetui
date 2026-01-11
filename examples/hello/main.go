package main

import (
	"log"

	"github.com/runetui/runetui"
)

func main() {
	app := runetui.New(func() runetui.Component {
		return runetui.Box(
			runetui.BoxProps{
				Direction: runetui.Column,
				Padding:   runetui.SpacingAll(2),
				Border:    runetui.BorderSingle,
			},
			runetui.Text("Hello, RuneTUI!", runetui.TextProps{Bold: true}),
			runetui.Text("Press Ctrl+C to quit"),
		)
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
