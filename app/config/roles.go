package config

import (
	"os"
	"path"

	"github.com/bytedance/sonic"
)

type Role struct {
	Title string `json:"title"`
	Size  string `json:"size"`
}

func LoadRoles() ([]Role, error) {
	file := path.Join(".", "config", "roles.json")
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var roles []Role
	if err := sonic.Unmarshal(data, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}
