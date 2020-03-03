package ansible

type dependencies struct {
	val int
}

func (metric dependencies) description() string {
	return "Number of dependencies"
}

func (metric dependencies) name() string {
	return "dependencies"
}

func (metric dependencies) value() int {
	return metric.val
}

type DependencyCalculator struct {
}

func (calculator DependencyCalculator) analyzeContent(content string) dependencies {
	return dependencies{
		val: len(ReadMetaString(content).AnsibleDependencies),
	}
}
