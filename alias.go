package main

import alias "github.com/linuxsuren/go-cli-alias/pkg"

func getAliasList() []alias.Alias {
	return []alias.Alias{
		{Name: "cm", Command: "checkout master"},
	}
}
