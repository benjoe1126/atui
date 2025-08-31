package argo

import (
	"context"

	"github.com/benjoe1126/atui/pkg/kube"
	"github.com/benjoe1126/atui/pkg/view"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ApplicationSet struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty"`
	Spec          ApplicationSetSpec   `json:"spec,omitempty"`
	Status        ApplicationSetStatus `json:"status,omitempty"`
}

func (a *ApplicationSet) Id() string {
	//TODO implement me
	panic("implement me")
}

func (a *ApplicationSet) View() view.View {
	//TODO implement me
	panic("implement me")
}

func (a *ApplicationSet) SubComponents() []Component {
	//TODO implement me
	panic("implement me")
}

type ApplicationSetSpec struct {
	Generators []ApplicationSetGenerator `json:"generators,omitempty"`
	Template   ApplicationSetTemplate    `json:"template,omitempty"`
	SyncPolicy *ApplicationSetSyncPolicy `json:"syncPolicy,omitempty"`
}

type ApplicationSetGenerator struct {
	List     *ListGenerator    `json:"list,omitempty"`
	Cluster  *ClusterGenerator `json:"cluster,omitempty"`
	Git      *GitGenerator     `json:"git,omitempty"`
	Matrix   *MatrixGenerator  `json:"matrix,omitempty"`
	Merge    *MergeGenerator   `json:"merge,omitempty"`
	Clusters *ClusterGenerator `json:"clusters,omitempty"` // alias
}

type ListGenerator struct {
	Elements []map[string]string `json:"elements,omitempty"`
}

type ClusterGenerator struct {
	Selector *v1.LabelSelector `json:"selector,omitempty"`
	Values   map[string]string `json:"values,omitempty"`
}

type GitGenerator struct {
	RepoURL             string                      `json:"repoURL,omitempty"`
	Revision            string                      `json:"revision,omitempty"`
	Directories         []GitDirectoryGeneratorItem `json:"directories,omitempty"`
	Files               []GitFileGeneratorItem      `json:"files,omitempty"`
	Template            *ApplicationSetTemplate     `json:"template,omitempty"`
	RequeueAfterSeconds *int64                      `json:"requeueAfterSeconds,omitempty"`
}

type GitDirectoryGeneratorItem struct {
	Path string `json:"path,omitempty"`
}

type GitFileGeneratorItem struct {
	Path string `json:"path,omitempty"`
}

type MatrixGenerator struct {
	Generators []ApplicationSetGenerator `json:"generators,omitempty"`
}

type MergeGenerator struct {
	Generators []ApplicationSetGenerator `json:"generators,omitempty"`
	MergeKeys  []string                  `json:"mergeKeys,omitempty"`
}

// Template for each generated Application
type ApplicationSetTemplate struct {
	ApplicationMetadata ApplicationSetTemplateMeta `json:"metadata,omitempty"`
	Spec                ApplicationSpec            `json:"spec,omitempty"`
}

type ApplicationSetTemplateMeta struct {
	Name        string            `json:"name,omitempty"`
	Namespace   string            `json:"namespace,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type ApplicationSetSyncPolicy struct {
	PreserveResourcesOnDeletion bool `json:"preserveResourcesOnDeletion,omitempty"`
}

type ApplicationSetStatus struct {
	Conditions []Condition `json:"conditions,omitempty"`
}

func (a *ApplicationSet) Name() string {
	return a.GetName()
}

func (a *ApplicationSet) Edit() error {
	//TODO implement me
	panic("implement me")
}

func (a *ApplicationSet) Delete() error {
	//TODO implement me
	panic("implement me")
}

func (a *ApplicationSet) GetCreatedApplications() ([]Application, error) {
	resKind := schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "applications"}
	client, err := kube.NewClient()
	if err != nil {
		return nil, err
	}
	apps, err := client.Resource(resKind).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	ret := make([]Application, 0)
	for _, app := range apps.Items {
		if len(app.GetOwnerReferences()) > 0 && app.GetOwnerReferences()[0].Name == a.Name() {
			var appManifest Application
			if err := runtime.DefaultUnstructuredConverter.FromUnstructuredWithValidation(app.UnstructuredContent(), &appManifest, false); err != nil {
				return nil, err
			}
			ret = append(ret, appManifest)
		}
	}
	return ret, nil
}
