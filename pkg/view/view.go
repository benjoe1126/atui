package view

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
)

type View interface {
	TableRowView() table.Row
	EditView() textarea.Model
	ArgoView() string
	TableColumns() []table.Column
}
