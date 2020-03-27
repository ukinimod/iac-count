# Domain Model for iac-count

## Domain Objects

- SubjectCreator
  - Responsibilities: Creates Subject from path
  - Collaborators: Subject
- Subject
  - Collaborators: SubjectCreator, Node, Metric, MetricCalculator
  - Examples: Directory, File, Yamlfile
- MetricCalculator
  - Responsibilities: Calculates Metric from content
  - Collaborators: Metric, Subject
- Node
  - Collaborators: Metric, Subject
- Metric
  - Collaborators: MetricCalculator
  - Examples: Loc

### Extension Points

- Dependency
