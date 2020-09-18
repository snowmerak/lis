package main

type config struct {
	BinPath  string              `json:"bin_path"`
	BinName  string              `json:"bin_name"`
	Target   map[string][]string `json:"target"`
	GOGC     int                 `json:"gogc"`
	IsPlugin bool                `json:"is_plugin"`
	WillRun  bool                `json:"will_run"`
}
