package model

import "github.com/benjoe1126/atui/pkg/argo"

type Model struct {
	choices  []argo.Component
	cursor   int
	selected map[int]argo.Component
}
