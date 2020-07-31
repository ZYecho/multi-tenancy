package blocknsquota

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/multi-tenancy/benchmarks/kubectl-mtb/bundle/box"
	"sigs.k8s.io/multi-tenancy/benchmarks/kubectl-mtb/pkg/benchmark"
	"sigs.k8s.io/multi-tenancy/benchmarks/kubectl-mtb/test"
	"sigs.k8s.io/multi-tenancy/benchmarks/kubectl-mtb/test/utils"
)

var b = &benchmark.Benchmark{

	PreRun: func(tenantNamespace string, kclient, tclient *kubernetes.Clientset) error {

		return nil
	},

	Run: func(tenantNamespace string, kclient, tclient *kubernetes.Clientset) error {
		resources := []utils.GroupResource{
			{
				APIGroup: "",
				APIResource: metav1.APIResource{
					Name: "resourcequotas",
				},
			},
		}
		verbs := []string{"create", "update", "patch", "delete", "deletecollection"}
		for _, resource := range resources {
			for _, verb := range verbs {
				access, msg, err := utils.RunAccessCheck(tclient, tenantNamespace, resource, verb)
				if err != nil {
					fmt.Println(err.Error())
				}
				if access {
					return fmt.Errorf(msg)
				}
			}
		}
		return nil
	},
}

func init() {
	// Get the []byte representation of a file, or an error if it doesn't exist:
	err := b.ReadConfig(box.Get("block_ns_quota/config.yaml"))
	if err != nil {
		fmt.Println(err)
	}

	test.BenchmarkSuite.Add(b)
}
