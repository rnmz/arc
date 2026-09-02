package config

import (
	"os"
	"path"

	"github.com/bytedance/sonic"
)

type Plan struct {
	Title string `json:"title"`
	Size  int    `json:"size"`
}

func LoadPlans() ([]Plan, error) {
	file := path.Join(".", "config", "plans.json")
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var plans []Plan
	if err := sonic.Unmarshal(data, &plans); err != nil {
		return nil, err
	}
	return plans, nil
}
