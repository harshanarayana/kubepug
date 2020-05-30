package kubepug

import (
	"github.com/rikatz/kubepug/pkg/parser"
	"github.com/rikatz/kubepug/pkg/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// Deprecator implements an interface for reading some sort of Input and comparing against the
// map of Kubernetes APIs to check if there's some Deprecated or Deleted
type Deprecator interface {
	ListDeprecated() []schema.DeprecatedAPI
	ListDeleted() []schema.DeletedAPI
}

// YamlInput defines a struct that will be used when comparing APIs against a YAML Input file
type YamlInput struct {
	File   string
	K8sapi *parser.KubernetesAPIs
}

// K8sInput defines a struct that will be used when comparing APIs against a K8s Cluster
type K8sInput struct {
	K8sconfig *genericclioptions.ConfigFlags
	K8sapi    parser.KubernetesAPIs
	Apiwalk   bool
}
