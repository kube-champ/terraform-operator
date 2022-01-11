package kube

import (
	"context"
	"os"
	"path/filepath"

	"github.com/kube-champ/terraform-operator/pkg/utils"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var ClientSet kubernetes.Interface

func CreateK8SConfig() (*rest.Config, error) {
	l := log.FromContext(context.Background())
	dir, err := os.Getwd()

	if err != nil {
		l.Error(err, "could not retreive currect directory")
		return nil, err
	}

	kubeconfigPath := filepath.Join(dir, "kubeconfig")

	var clientset *kubernetes.Clientset
	var config *rest.Config

	if utils.FileExists(kubeconfigPath) {
		if config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath); err != nil {
			l.Error(err, "failed to create K8s config from kubeconfig")
			return nil, err
		}
	} else {
		if config, err = rest.InClusterConfig(); err != nil {
			l.Error(err, "Failed to create in-cluster k8s config")
			return nil, err
		}
	}

	clientset, err = kubernetes.NewForConfig(config)

	if err != nil {
		l.Error(err, "Failed to create K8s clientset")
	}

	ClientSet = clientset

	return config, nil
}
