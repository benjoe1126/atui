package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/benjoe1126/atui/pkg/argo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	res := schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "applications"}
	list, err := clientset.Resource(res).Namespace("kszk-argocd").List(context.Background(), metav1.ListOptions{Limit: 4})
	if err != nil {
		panic(err.Error())
	}
	var cl argo.Application
	if err := runtime.DefaultUnstructuredConverter.FromUnstructuredWithValidation(list.Items[0].UnstructuredContent(), &cl, false); err != nil {
		panic(err.Error())
	}
	fmt.Println(cl.Spec.Source)
}
