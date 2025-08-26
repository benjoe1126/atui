package argo

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ApplicationView struct {
	app *Application
}

func (a *ApplicationView) TableRowView() table.Row {
	return table.Row{
		a.app.GetName(),
		a.app.GetNamespace(),
	}
}

func (a *ApplicationView) EditView() textarea.Model {
	//TODO implement me
	panic("implement me")
}

func (a *ApplicationView) ArgoView() string {
	//TODO implement me
	panic("implement me")
}

type Application struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec          ApplicationSpec   `json:"spec,omitempty"`
	Status        ApplicationStatus `json:"status,omitempty"`
}

type ApplicationSpec struct {
	Project     string                 `json:"project,omitempty"`
	Source      ApplicationSource      `json:"source,omitempty"`
	Destination ApplicationDestination `json:"destination,omitempty"`
	SyncPolicy  *SyncPolicy            `json:"syncPolicy,omitempty"`
}

type ApplicationSource struct {
	RepoURL        string `json:"repoURL,omitempty"`
	Path           string `json:"path,omitempty"`
	TargetRevision string `json:"targetRevision,omitempty"`
	Chart          string `json:"chart,omitempty"`
}

type ApplicationDestination struct {
	Server    string `json:"server,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type SyncPolicy struct {
	Automated   *AutomatedSyncPolicy `json:"automated,omitempty"`
	SyncOptions []string             `json:"syncOptions,omitempty"`
}

type AutomatedSyncPolicy struct {
	Prune    bool `json:"prune,omitempty"`
	SelfHeal bool `json:"selfHeal,omitempty"`
}

type ApplicationStatus struct {
	Health       HealthStatus `json:"health,omitempty"`
	Sync         SyncStatus   `json:"sync,omitempty"`
	Conditions   []Condition  `json:"conditions,omitempty"`
	ReconciledAt *v1.Time     `json:"reconciledAt,omitempty"`
}

type HealthStatus struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type SyncStatus struct {
	Status    string   `json:"status,omitempty"`
	Revision  string   `json:"revision,omitempty"`
	Revisions []string `json:"revisions,omitempty"`
}

type Condition struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

func (app *Application) SetAutoSync(bool) {
	//TODO implement me
	panic("implement me")
}

func (app *Application) Sync() error {
	//TODO implement me
	panic("implement me")
}

func (app *Application) Delete() error {
	//TODO implement me
	panic("implement me")
}

func (app *Application) Edit() error {
	//TODO implement me
	panic("implement me")
}

func (app *Application) View() string {
	return fmt.Sprintf("%v", app)
}

func (app *Application) Name() string {
	return "application"
}
