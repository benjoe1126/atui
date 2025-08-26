package model

import (
	"log"

	"github.com/benjoe1126/atui/pkg/argo"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	choices  []argo.Component
	cursor   int
	selected map[int]argo.Component
}

func (m Model) Init() tea.Cmd {
	m.choices = []argo.Component{
		&argo.Application{},
	}
	m.selected = make(map[int]argo.Component)
	m.selected[0] = m.choices[0]
	return tea.Batch()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		log.Println(msg.String())
		switch msg.Type {
		case tea.KeyUp:
			log.Println("pressed up")
			m.cursor++
		case tea.KeyDown:
			m.cursor--
		case tea.KeyEscape:
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return "hihi"
}

func New() *Model {
	return &Model{
		selected: make(map[int]argo.Component),
		cursor:   0,
		choices:  []argo.Component{},
	}
}
