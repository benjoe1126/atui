package argo

import (
	"context"
	"fmt"

	"github.com/benjoe1126/atui/pkg/kube"
	"github.com/benjoe1126/atui/pkg/view"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ApplicationView struct {
	app *Application
}

func (a ApplicationView) TableRowView() table.Row {
	return table.Row{
		a.app.GetName(),
		a.app.Namespace,
		a.app.Status.Sync.Status,
		a.app.Status.Health.Status,
		a.app.Status.Sync.Revision,
		a.app.Spec.Project,
	}
}
func (a ApplicationView) TableColumns() []table.Column {
	return []table.Column{
		table.Column{Title: "Name", Width: 20},
		table.Column{Title: "Namespace", Width: 15},
		table.Column{Title: "Status", Width: 15},
		table.Column{Title: "Health", Width: 15},
		table.Column{Title: "Revision", Width: 15},
		table.Column{Title: "Project", Width: 20},
	}
}

func (a ApplicationView) EditView() textarea.Model {
	//TODO implement me
	panic("implement me")
}

func (a ApplicationView) ArgoView() string {
	return fmt.Sprint(a.app.Spec)
}

type Application struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec          ApplicationSpec   `json:"spec,omitempty"`
	Status        ApplicationStatus `json:"status,omitempty"`
}

func (app *Application) Id() string {
	return fmt.Sprintf("%v", app.GetUID())
}

func (app *Application) View() view.View {
	return ApplicationView{app: app}
}

func (app *Application) SubComponents() []Component {
	//TODO implement me
	panic("implement me")
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
	client, err := kube.NewClient()
	if err != nil {
		return fmt.Errorf("create kube client failed during delete: %v", err)
	}
	res := schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "applications"}
	if err := client.Resource(res).Namespace(app.Namespace).Delete(context.Background(), app.GetName(), v1.DeleteOptions{}); err != nil {
		return fmt.Errorf("delete application failed during delete: %v", err)
	}
	return nil
}

func (app *Application) Edit() error {
	//TODO implement me
	panic("implement me")
}

func (app *Application) Name() string {
	return "application"
}

func ListApplications(namespace string, filter v1.ListOptions) ([]*Application, error) {
	client, err := kube.NewClient()
	if err != nil {
		return nil, err
	}
	res := schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "applications"}
	ul, err := client.Resource(res).Namespace(namespace).List(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var apps []*Application
	for _, item := range ul.Items {
		var app Application
		if err = runtime.DefaultUnstructuredConverter.FromUnstructuredWithValidation(item.UnstructuredContent(), &app, false); err != nil {
			return nil, err
		}
		apps = append(apps, &app)
	}
	return apps, nil
}
