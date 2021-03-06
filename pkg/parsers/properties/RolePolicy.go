package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// RolePolicy Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
type RolePolicy struct {
	PolicyDocument interface{} `yaml:"PolicyDocument"`
	PolicyName     interface{} `yaml:"PolicyName"`
}

// RolePolicy validation
func (resource RolePolicy) Validate() []error {
	errs := []error{}

	if resource.PolicyDocument == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyDocument'"))
	}
	if resource.PolicyName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyName'"))
	}
	return errs
}
