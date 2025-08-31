package argo

import (
	"github.com/benjoe1126/atui/pkg/view"
)

type Component interface {
	Id() string
	Name() string
	View() view.View
	Edit() error
	Delete() error
	SubComponents() []Component
}

type ArgoComponentType int

const (
	APPLICATION ArgoComponentType = iota
	APPLICATION_SET
	CONFIG
	INVALID
)
