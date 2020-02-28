package reader

import (
	"log"

	"gopkg.in/yaml.v2"
)

type AnsibleMetaData struct {
	AnsibleDependencies []interface{} `yaml:"dependencies"`
	AnsibleGalaxyInfo   interface{}   `yaml:"galaxy_info"`
}

func ReadMetaString(data string) AnsibleMetaData {
	m := AnsibleMetaData{}

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Printf("[WARN] %s", err)
	}

	return m
}

type AnsibleHandler struct {
	Name    string
	Service interface{}
}

func ReadHandlersString(data string) []AnsibleHandler {
	var m []AnsibleHandler

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Printf("[WARN] %s", err)
	}

	return m
}

type AnsibleTask struct {
	Name                   string                 `yaml:"name"`
	SetFact                map[string]interface{} `yaml:"set_fact"`
	WhenClause             interface{}            `yaml:"when"`
	LoopClause             interface{}            `yaml:"loop"`
	WithListClause         interface{}            `yaml:"with_list"`
	WithItemsClause        interface{}            `yaml:"with_items"`
	WithIndexedItemsClause interface{}            `yaml:"with_indexed_items"`
	WithFlattenedClause    interface{}            `yaml:"with_flattened"`
	WithTogetherClause     interface{}            `yaml:"with_together"`
	WithDictClause         interface{}            `yaml:"with_dict"`
	WithSequenceClause     interface{}            `yaml:"with_sequence"`
	WithSubelementsClause  interface{}            `yaml:"with_subelements"`
	WithNestedClause       []interface{}          `yaml:"with_nested"`
	WithCartesianClause    []interface{}          `yaml:"with_cartesian"`
	Block                  []AnsibleTask          `yaml:"block"`
	RescueBlock            []AnsibleTask          `yaml:"rescue"`
	AlwaysBlock            []AnsibleTask          `yaml:"always"`
	Tags                   interface{}            `yaml:"tags"`
	Assert                 interface{}            `yaml:"assert"`
}

func ReadTasksString(data string) []AnsibleTask {
	var m []AnsibleTask

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Printf("[WARN] %s", err)
	}

	return m
}

type AnsiblePlay struct {
	Tasks     []AnsibleTask
	PreTasks  []AnsibleTask
	Handlers  []AnsibleHandler
	PostTasks []AnsibleTask
	Roles     []interface{}
}

func ReadPlaybookString(data string) []AnsiblePlay {
	var m []AnsiblePlay

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Printf("[WARN] %s", err)
	}

	return m
}
