package core

const (
	Loc          = "loc"
	Tasks        = "tasks"
	Handlers     = "handlers"
	Complexity   = "mcc"
	Vars         = "vars"
	Defaults     = "defaults"
	GroupVars    = "group_vars"
	HostVars     = "host_vars"
	Dependencies = "dependencies"
	Rloc         = "rloc"
	CommentLines = "comment_lines"
	MaxNL        = "max_nesting_level"
	Files        = "files"
	Templates    = "templates"
	StaticFiles  = "static_files"
	Plugins      = "plugins"
	Roles        = "roles"
	Plays        = "plays"
	CustomFacts  = "custom_facts"
	TaggedTasks  = "tagged_tasks"
)

var MetricNames = [...]string{
	Loc,
	Tasks,
	Handlers,
	Complexity,
	Vars,
	Defaults,
	GroupVars,
	HostVars,
	Dependencies,
	Rloc,
	CommentLines,
	MaxNL,
	Files,
	Templates,
	StaticFiles,
	Plugins,
	Roles,
	Plays,
	CustomFacts,
	TaggedTasks,
}

func IsValidMetric(metric string) bool {
	for _, b := range MetricNames {
		if b == metric {
			return true
		}
	}
	return false
}

func DescribeMetric(metric string) string {
	switch metric {
	case Loc:
		return "Number of code lines in file (agg: sum)"
	case Tasks:
		return "Number of tasks in file (agg: sum)"
	case Handlers:
		return "Number of handlers in file (agg: sum)"
	case Complexity:
		return "Complexity of file (agg: sum, see documentation)"
	case Vars:
		return "Number of variables in role (agg: sum)"
	case Defaults:
		return "Number of defaults in role (agg: sum)"
	case GroupVars:
		return "Number of group Vars in file (agg: sum)"
	case HostVars:
		return "Number of host Vars in file (agg: sum)"
	case Dependencies:
		return "Number of dependencies to roles in roles (agg: sum)"
	case Rloc:
		return "Number of non-blank, non-comment lines in file (agg: sum)"
	case CommentLines:
		return "Number of comment-only lines in file (agg: sum)"
	case Files:
		return "Number of files in directory and sub-directories (agg: sum)"
	case Templates:
		return "Number of templates in role (agg: sum)"
	case StaticFiles:
		return "Number of static files in role (agg: sum)"
	case Plugins:
		return "Number of plugins in role / ansible project"
	case Roles:
		return "Number of roles in ansible project (agg: sum)"
	case Plays:
		return "Number of plays in ansible project (agg: sum)"
	case CustomFacts:
		return "Number of custom facts in file (agg: sum)"
	case MaxNL:
		return "Maximal nesting level (agg: max)"
	case TaggedTasks:
		return "Number of tasks that have tags (agg: sum)"
	default:
		return "Unknown metric"
	}
}
