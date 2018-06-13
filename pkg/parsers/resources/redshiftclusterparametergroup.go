package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// RedshiftClusterParameterGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html
type RedshiftClusterParameterGroup struct {
	Type       string                                  `yaml:"Type"`
	Properties RedshiftClusterParameterGroupProperties `yaml:"Properties"`
	Condition  interface{}                             `yaml:"Condition,omitempty"`
	Metadata   interface{}                             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                             `yaml:"DependsOn,omitempty"`
}

// RedshiftClusterParameterGroup Properties
type RedshiftClusterParameterGroupProperties struct {
	Description          interface{} `yaml:"Description"`
	ParameterGroupFamily interface{} `yaml:"ParameterGroupFamily"`
	Parameters           interface{} `yaml:"Parameters,omitempty"`
	Tags                 interface{} `yaml:"Tags,omitempty"`
}

// NewRedshiftClusterParameterGroup constructor creates a new RedshiftClusterParameterGroup
func NewRedshiftClusterParameterGroup(properties RedshiftClusterParameterGroupProperties, deps ...interface{}) RedshiftClusterParameterGroup {
	return RedshiftClusterParameterGroup{
		Type:       "AWS::Redshift::ClusterParameterGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseRedshiftClusterParameterGroup parses RedshiftClusterParameterGroup
func ParseRedshiftClusterParameterGroup(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource RedshiftClusterParameterGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RedshiftClusterParameterGroup - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseRedshiftClusterParameterGroup validator
func (resource RedshiftClusterParameterGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseRedshiftClusterParameterGroupProperties validator
func (resource RedshiftClusterParameterGroupProperties) Validate() []error {
	errs := []error{}
	if resource.Description == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Description'"))
	}
	if resource.ParameterGroupFamily == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ParameterGroupFamily'"))
	}
	return errs
}