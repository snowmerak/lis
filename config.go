package main

type config struct {
	BinPath  string              `json:"bin_path" yaml:"bin_path"`
	Name     string              `json:"name" yaml:"name"`
	Target   map[string][]string `json:"target" yaml:"target"`
	GOGC     int                 `json:"gogc" yaml:"gogc"`
	ToPlugin bool                `json:"to_plugin" yaml:"to_plugin"`
	Module   bool                `json:"module" yaml:"module"`
	Etc      []string            `json:"etc" yaml:"etc"`
}
