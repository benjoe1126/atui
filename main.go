package main

import (
	"context"
	"fmt"

	"github.com/benjoe1126/atui/pkg/argo"
	"github.com/benjoe1126/atui/pkg/kube"
	"github.com/benjoe1126/atui/pkg/model"
	tea "github.com/charmbracelet/bubbletea"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

func main() {
	// create the clientset
	clientset, err := dynamic.NewForConfig(kube.Kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	res := schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "applications"}
	list, err := clientset.Resource(res).Namespace("argo").List(context.Background(), metav1.ListOptions{Limit: 4})
	if err != nil {
		panic(err.Error())
	}
	var cl argo.Application
	if err := runtime.DefaultUnstructuredConverter.FromUnstructuredWithValidation(list.Items[0].UnstructuredContent(), &cl, false); err != nil {
		panic(err.Error())
	}
	var appset argo.ApplicationSet
	res = schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "applicationsets"}
	list, err = clientset.Resource(res).Namespace("argo").List(context.Background(), metav1.ListOptions{Limit: 4})
	if err := runtime.DefaultUnstructuredConverter.FromUnstructuredWithValidation(list.Items[0].UnstructuredContent(), &appset, false); err != nil {
		panic(err.Error())
	}
	fmt.Println(appset.GetCreatedApplications())
	m := model.New()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		panic(err)
	}
}
