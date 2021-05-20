package logx

// A Conf is a logging config.
type Conf struct {
	ServiceName         string `json:",optional"`
	Mode                string `json:",default=console,options=console|file|volume"`
	TimeFormat          string `json:",optional"`
	Path                string `json:",default=logs"`
	Level               string `json:",default=info,options=info|error|severe"`
	Compress            bool   `json:",optional"`
	KeepDays            int    `json:",optional"`
	StackCoolDownMillis int    `json:",default=100"`
}
