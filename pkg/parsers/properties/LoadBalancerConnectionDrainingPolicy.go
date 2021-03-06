package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// LoadBalancerConnectionDrainingPolicy Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectiondrainingpolicy.html
type LoadBalancerConnectionDrainingPolicy struct {
	Enabled interface{} `yaml:"Enabled"`
	Timeout interface{} `yaml:"Timeout,omitempty"`
}

// LoadBalancerConnectionDrainingPolicy validation
func (resource LoadBalancerConnectionDrainingPolicy) Validate() []error {
	errs := []error{}

	if resource.Enabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enabled'"))
	}
	return errs
}
