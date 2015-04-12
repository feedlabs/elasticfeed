package plugin

type Profiler struct {
	data map[string]string
}

func NewProfiler(data map[string]string) *Profiler {
	return &Profiler{data}
}
