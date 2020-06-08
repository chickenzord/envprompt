package profile

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type chefEnvCollector struct{}

func (c *chefEnvCollector) Collect() string {
	return os.Getenv("CHEF_PROFILE")
}

type chefFileCollector struct{}

func (c *chefFileCollector) Collect() string {
	home := os.Getenv("HOME")
	bytes, err := ioutil.ReadFile(path.Join(home, ".chef", "context"))
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(bytes))
}

func NewChefProfile(label string) *Profile {
	return &Profile{
		Label:   label,
		Default: "default",
		Collectors: []Collector{
			&chefEnvCollector{},
			&chefFileCollector{},
		},
	}
}
