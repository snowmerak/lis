package main

type BuildConfig struct {
	BinPath  string              `json:"bin_path" yaml:"bin_path"`
	Name     string              `json:"name" yaml:"name"`
	Target   map[string][]string `json:"target" yaml:"target"`
	GOGC     int                 `json:"gogc" yaml:"gogc"`
	ToPlugin bool                `json:"to_plugin" yaml:"to_plugin"`
	Module   bool                `json:"module" yaml:"module"`
	Etc      []string            `json:"etc" yaml:"etc"`
}

type TestConfig struct {
	Targets []struct {
		Package string   `json:"package" yaml:"package"`
		Test    []string `json:"test" yaml:"test"`
	} `json:"targets" yaml:"targets"`
	TestFlags []string `json:"flags" yaml:"flags"`
}

type BenchConfig struct {
	Targets []struct {
		Package string   `json:"package" yaml:"package"`
		Bench   []string `json:"bench" yaml:"bench"`
	} `json:"targets" yaml:"targets"`
	BenchFlags []string `json:"flags" yaml:"flags"`
}
