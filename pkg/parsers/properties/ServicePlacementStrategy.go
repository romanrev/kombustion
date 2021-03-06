package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ServicePlacementStrategy Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementstrategy.html
type ServicePlacementStrategy struct {
	Field interface{} `yaml:"Field,omitempty"`
	Type  interface{} `yaml:"Type"`
}

// ServicePlacementStrategy validation
func (resource ServicePlacementStrategy) Validate() []error {
	errs := []error{}

	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
