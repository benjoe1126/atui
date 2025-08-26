package argo

import (
	"github.com/benjoe1126/atui/pkg/view"
)

type Component interface {
	Name() string
	View() view.View
	Edit() error
	Delete() error
	SubComponents() []Component
}
