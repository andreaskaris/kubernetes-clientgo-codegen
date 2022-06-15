package main

import (
	"os"
	"path/filepath"
	"time"

	"example.com/m/pkg/generated/clientset/versioned"
	"example.com/m/pkg/generated/informers/externalversions"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := filepath.Join("~", ".kube", "config")
		if envvar := os.Getenv("KUBECONFIG"); len(envvar) > 0 {
			kubeconfig = envvar
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			klog.Fatal(err)
		}
	}

	f, err := versioned.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}
	sif := externalversions.NewSharedInformerFactory(f, 30*time.Second)
	informer := sif.Foo().V1().Foos()
	informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				klog.Infof("Added %s", key)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err == nil {
				klog.Infof("Updated %s", key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				klog.Infof("Deleted %s", key)
			}
		},
	})
	sif.Start(wait.NeverStop)
	sif.WaitForCacheSync(wait.NeverStop)
	ret, err := informer.Lister().Foos("default").List(labels.Everything())
	if err != nil {
		klog.Error(err)
	}
	for k, v := range ret {
		klog.Infof("%d -> %v", k, v)
		klog.Info("===")
	}

}
