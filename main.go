package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/chickenzord/envprompt/profile"
	"github.com/fatih/color"
)

var (
	profiles = map[string]profile.Profile{
		"chef": *profile.NewChefProfile("c"),
		"kube": *profile.NewKubeProfile("k"),
	}
)

func renderPrompt(p profile.Profile) string {
	label := p.Label
	val := p.GetValue()

	if val == "" || val == p.Default {
		return ""
	}

	c := color.New()
	if strings.Contains(val, "-prd") {
		c = c.Add(color.FgRed)
	}

	return fmt.Sprintf("%s:%s", label, c.Sprintf(val))
}

func main() {
	keys := []string{}
	for key, _ := range profiles {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	parts := []string{}
	for _, k := range keys {
		// p :=
		parts = append(parts, renderPrompt(profiles[k]))
	}

	prompt := strings.TrimSpace(strings.Join(parts, " "))
	fmt.Println(prompt)
}
