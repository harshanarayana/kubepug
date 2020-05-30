package kubepug

import (
	k8sinput "github.com/rikatz/kubepug/pkg/kubepug/input/k8s"
	"github.com/rikatz/kubepug/pkg/schema"
)

// ListDeprecated lists the deprecated objects from a YamlInput file
func (i YamlInput) ListDeprecated() (deprecatedapis []schema.DeprecatedAPI) {
	return deprecatedapis
}

// ListDeleted lists the non-existend objects in some K8s version from a YamlInput file
func (i YamlInput) ListDeleted() (deletedapis []schema.DeletedAPI) {
	return deletedapis
}

// ListDeprecated lists the deprecated objects from a Kubernetes cluster
func (i K8sInput) ListDeprecated() (deprecatedapis []schema.DeprecatedAPI) {

	deprecatedapis = k8sinput.GetDeprecated(i.K8sapi, i.K8sconfig)
	return deprecatedapis

}

// ListDeleted lists the non-existend objects in some K8s version from a Kubernetes cluster
func (i K8sInput) ListDeleted() (deletedapis []schema.DeletedAPI) {

	deletedapis = k8sinput.GetDeleted(i.K8sapi, i.K8sconfig)
	return deletedapis
}
