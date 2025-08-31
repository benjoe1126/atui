package argo

import (
	"github.com/benjoe1126/atui/pkg/view"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AppConfig struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}

func (a AppConfig) Id() string {
	//TODO implement me
	panic("implement me")
}

func (a AppConfig) Name() string {
	//TODO implement me
	panic("implement me")
}

func (a AppConfig) View() view.View {
	//TODO implement me
	panic("implement me")
}

func (a AppConfig) Edit() error {
	//TODO implement me
	panic("implement me")
}

func (a AppConfig) Delete() error {
	//TODO implement me
	panic("implement me")
}

func (a AppConfig) SubComponents() []Component {
	//TODO implement me
	panic("implement me")
}
