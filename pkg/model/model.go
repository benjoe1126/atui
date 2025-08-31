package model

import (
	"log"
	"time"

	"github.com/benjoe1126/atui/pkg/argo"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ModelStage int

const (
	TABLE ModelStage = iota
	EDIT
	YAML
	FILTER
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type Model struct {
	argoType  argo.ArgoComponentType
	namespace string
	keys      *argoKeyMap
	stage     ModelStage
	choices   []argo.Component
	table     table.Model
	cursor    int
}

type tickMsg struct{}

func tickEvery(ms time.Duration) tea.Cmd {
	return tea.Every(2*time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func noChange[T argo.Component](src []argo.Component, new []T) bool {
	if len(src) != len(new) {
		return false
	}
	if len(src) == 0 {
		return true
	}
	for i, _ := range src {
		if src[i].Id() != new[i].Id() {
			return false
		}
	}
	return true
}

func updateTableRows(choices []argo.Component, columns []table.Column) table.Model {
	if choices == nil || len(choices) == 0 {
		return table.New(
			table.WithFocused(true),
		)
	}
	rows := make([]table.Row, len(choices))
	for _, a := range choices {
		rows = append(rows, a.View().TableRowView())
	}
	return table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
}

func (m *Model) Init() tea.Cmd {
	m.choices = make([]argo.Component, 0)
	m.keys = NewArgoKeyMap()
	m.stage = TABLE
	m.argoType = argo.APPLICATION
	tmp := &argo.Application{}
	m.table = table.New(
		table.WithColumns(tmp.View().TableColumns()),
		table.WithFocused(true),
	)
	return tea.Batch(tickEvery(2 * time.Second))
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	var cmd tea.Cmd
	var componentSchema argo.Component
	switch m.argoType {
	case argo.APPLICATION:
		componentSchema = &argo.Application{}
	case argo.APPLICATION_SET:
		componentSchema = &argo.ApplicationSet{}
	case argo.CONFIG:
		componentSchema = &argo.AppConfig{}

	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.stage == FILTER {
			break
		}
		switch msg.Type {
		case tea.KeyUp:
			m.cursor--
			m.cursor = max(m.cursor, 0)
			if m.stage == TABLE {
				m.table, cmd = m.table.Update(msg)
				cmds = append(cmds, cmd)
			}
		case tea.KeyDown:
			m.cursor++
			m.cursor = min(m.cursor, len(m.choices)-1)
			if m.stage == TABLE {
				m.table, cmd = m.table.Update(msg)
				cmds = append(cmds, cmd)
			}
		case tea.KeyEscape:
			return m, tea.Quit
		default:
			if key.Matches(msg, m.keys.deleteResource) {
				if m.cursor >= 0 && m.choices[m.cursor] != nil {
					if err := m.choices[m.cursor].Delete(); err != nil {
						log.Println("error deleting resource:", err)
						break
					}
					tmp := make([]argo.Component, max(0, len(m.choices)-1))
					copy(tmp, m.choices[0:m.cursor])
					copy(tmp, m.choices[m.cursor+1:])
					m.choices = tmp
					m.table = updateTableRows(m.choices, componentSchema.View().TableColumns())
					m.cursor = 0
				}
			}
		}

	case tickMsg:
		if m.stage != TABLE {
			break
		}
		cmds = append(cmds, tickEvery(2*time.Second))
		switch m.argoType {
		case argo.APPLICATION:
			tmp, err := argo.ListApplications("", v1.ListOptions{})
			if err != nil || tmp == nil || noChange(m.choices, tmp) {
				break
			}
			m.choices = make([]argo.Component, 0, len(tmp))
			for _, a := range tmp {
				m.choices = append(m.choices, a)
			}
			m.table = updateTableRows(m.choices, componentSchema.View().TableColumns())

		}
	}
	if len(cmds) > 0 {
		return m, tea.Batch(cmds...)
	}
	return m, nil
}

func (m *Model) View() string {
	if len(m.choices) > 0 {
		return baseStyle.Render(m.table.View())
	}
	return "dumbass"
}

func New() *Model {
	return &Model{
		namespace: "default",
		cursor:    0,
		choices:   []argo.Component{},
	}
}
