package argo

type Component interface {
	Name() string
	View() string
	Edit() error
	Delete() error
}
