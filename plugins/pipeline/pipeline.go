package pipeline

func Filter(data interface{}) interface{} {
	// should call plugins of type PIPELINE
	return RandomAnimator(data)
}

func init() {}
