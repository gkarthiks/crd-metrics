package utils

import (
	"context"
	"fmt"
	"github.com/gkarthiks/k8s-discovery"
	"github.com/pterm/pterm"
	apixV1client "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"strconv"
)

var K8s *discovery.K8s

func init() {
	K8s, _ = discovery.NewK8s()
}

func PrintList() {
	apiXClient := apixV1client.NewForConfigOrDie(K8s.RestConfig)
	crds, _ := apiXClient.CustomResourceDefinitions().List(context.Background(), v1.ListOptions{})

	var dataToTerminal [][]string
	header := []string{"SNO", "CRD Kind", "CRD Name", "Total CRs", "Namespaced", "Group", "Version", "Kind"}
	dataToTerminal = append(dataToTerminal, header)
	for idx, crd := range crds.Items {
		currentRecord := []string{
			strconv.Itoa(idx + 1),
			crd.Name,
			crd.Spec.Names.Singular,
			"",
			string(crd.Spec.Scope),
			crd.Spec.Group,
			crd.Spec.Versions[0].Name,
			crd.Spec.Names.Kind,
		}
		dataToTerminal = append(dataToTerminal, currentRecord)
	}
	pterm.DefaultTable.WithHasHeader(true).WithData(dataToTerminal).Render()

	client := dynamic.NewForConfigOrDie(K8s.RestConfig)
	//fmt.Println(err)
	var (
		runtimeClassGVR = schema.GroupVersionResource{
			Group:    "capsule.clastix.io",
			Version:  "v1alpha1",
			Resource: "CapsuleConfiguration",
		}
		//runtimeGVK = schema.GroupVersionKind{
		//	Group:   "capsule.clastix.io",
		//	Version: "v1alpha1",
		//	Kind:    "CapsuleConfiguration",
		//}
	)

	res := client.Resource(runtimeClassGVR)
	list, err := res.List(context.Background(), v1.ListOptions{})
	fmt.Println("========================list============")
	fmt.Println(list)
	fmt.Println("========================list============")
	fmt.Println(err)
}
