package main

type config struct {
	BinPath  string              `json:"bin_path"`
	Name     string              `json:"name"`
	Target   map[string][]string `json:"target"`
	GOGC     int                 `json:"gogc"`
	ToPlugin bool                `json:"to_plugin"`
	AutoRun  bool                `json:"auto_run"`
	Module   bool                `json:"module"`
}
