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
		return "Number of code lines in File (agg: sum)"
	case Tasks:
		return "Number of tasks in File (agg: sum)"
	case Handlers:
		return "Number of Handlers in File (agg: sum)"
	case Complexity:
		return "Complexity of File (agg: sum, see documentation)"
	case Vars:
		return "Number of Variables in Role (agg: sum)"
	case Defaults:
		return "Number of Defaults in Role (agg: sum)"
	case GroupVars:
		return "Number of Group Vars in File (agg: sum)"
	case HostVars:
		return "Number of Host Vars in File (agg: sum)"
	case Dependencies:
		return "Number of Dependencies to Roles in Roles (agg: sum)"
	case Rloc:
		return "Number of non-blank, non-comment lines in File (agg: sum)"
	case CommentLines:
		return "Number of comment-only Lines in File (agg: sum)"
	case Files:
		return "Number of Files in Directory and Sub-Directories (agg: sum)"
	case Templates:
		return "Number of Templates in Role (agg: sum)"
	case StaticFiles:
		return "Number of static Files in Role (agg: sum)"
	case Plugins:
		return "Number of Plugins in Role / Ansible Project"
	case Roles:
		return "Number of Roles in Ansible Project (agg: sum)"
	case Plays:
		return "Number of Plays in Ansible Project (agg: sum)"
	case CustomFacts:
		return "Number of Custom Facts in File (agg: sum)"
	case MaxNL:
		return "Maximal Nesting Level (agg: max)"
	case TaggedTasks:
		return "Number of Tasks that have tags (agg: sum)"
	default:
		return "Unknown Metric"
	}
}
