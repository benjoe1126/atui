package model

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type argoKeyMap struct {
	editResource   key.Binding
	deleteResource key.Binding
	viewResource   key.Binding
	syncResource   key.Binding
	selectResource key.Binding
}

func NewArgoKeyMap() *argoKeyMap {
	return &argoKeyMap{
		editResource: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "edit resource"),
		),
		deleteResource: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "delete resource"),
		),
		viewResource: key.NewBinding(
			key.WithKeys("v"),
			key.WithHelp("v", "view resource"),
		),
		syncResource: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "sync resource"),
		),
		selectResource: key.NewBinding(
			key.WithKeys(tea.KeySpace.String()),
			key.WithHelp("space", "select resource"),
		),
	}
}
