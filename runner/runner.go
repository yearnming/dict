package runner

type Runner struct {
	options *Options
}

func NewRunner(options *Options) (*Runner, error) {

	// 创建一个新的 Runner 实例，并且初始化它的 options 字段
	runner := &Runner{options: options}

	return runner, nil
}

func (r *Runner) Run() error {

	return nil
}
