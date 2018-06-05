package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Pipeline_PipelineObject struct {
	Id     interface{} `yaml:"Id"`
	Name   interface{} `yaml:"Name"`
	Fields interface{} `yaml:"Fields"`
}

func (resource Pipeline_PipelineObject) Validate() []error {
	errs := []error{}

	if resource.Id == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Id'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Fields == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Fields'"))
	}
	return errs
}