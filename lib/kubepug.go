package lib

import (
	"github.com/rikatz/kubepug/pkg/kubepug"
	"github.com/rikatz/kubepug/pkg/parser"
	"github.com/rikatz/kubepug/pkg/schema"
	"github.com/rikatz/kubepug/pkg/utils"
	log "github.com/sirupsen/logrus"

	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// Config configuration object for Kubepug
// configurations for kubernetes and for kubepug functionality
type Config struct {
	K8sVersion      string
	ForceDownload   bool
	APIWalk         bool
	SwaggerDir      string
	ShowDescription bool
	Input           string
	ConfigFlags     *genericclioptions.ConfigFlags
}

// Kubepug struct to be used
type Kubepug struct {
	Config Config
}

// NewKubepug returns a new kubepug library
func NewKubepug(config Config) *Kubepug {
	return &Kubepug{Config: config}
}

// GetDeprecated returns the list of deprecated APIs
func (k *Kubepug) GetDeprecated() (result *schema.Result, err error) {
	log.Debugf("Populating the KubernetesAPI map from swagger.json")

	var KubernetesAPIs parser.KubernetesAPIs = make(parser.KubernetesAPIs)

	log.Infof("Downloading the swagger.json file")
	swaggerfile, err := utils.DownloadSwaggerFile(k.Config.K8sVersion, k.Config.SwaggerDir, k.Config.ForceDownload)

	if err != nil {
		return &schema.Result{}, err
	}

	log.Infof("Populating the Deprecated Kubernetes APIs Map")
	err = KubernetesAPIs.PopulateKubeAPIMap(swaggerfile)

	if err != nil {
		return &schema.Result{}, err
	}

	log.Debugf("Kubernetes APIs Populated: %#v", KubernetesAPIs)

	result = k.getResults(KubernetesAPIs)

	return result, nil
}

func (k *Kubepug) getResults(kubeapis parser.KubernetesAPIs) (result *schema.Result) {
	var inputMode kubepug.Deprecator
	if k.Config.Input != "" {
		inputMode = kubepug.YamlInput{
			File:   k.Config.Input,
			K8sapi: &kubeapis,
		}
	} else {
		inputMode = kubepug.K8sInput{
			K8sconfig: k.Config.ConfigFlags,
			K8sapi:    kubeapis,
			Apiwalk:   k.Config.APIWalk,
		}
	}
	results := kubepug.GetDeprecations(inputMode)
	return &results

}
