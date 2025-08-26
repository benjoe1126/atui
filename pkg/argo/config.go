package argo

import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type AppConfig struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}
