package profile

import (
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type kubeContextCollector struct{}

func (c *kubeContextCollector) Collect() string {
	var kubeconfig string
	if val, ok := os.LookupEnv("KUBECONFIG"); ok {
		kubeconfig = val
	} else {
		home := os.Getenv("HOME")
		kubeconfig = path.Join(home, ".kube", "config")
	}

	f, err := os.Open(kubeconfig)
	defer f.Close()
	if err != nil {
		return ""
	}

	var cfg map[string]interface{}
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return ""
	}

	if val, ok := cfg["current-context"]; ok {
		return val.(string)
	}

	return ""
}

func NewKubeProfile(label string) *Profile {
	return &Profile{
		Label:   label,
		Default: "",
		Collectors: []Collector{
			&kubeContextCollector{},
		},
	}
}
