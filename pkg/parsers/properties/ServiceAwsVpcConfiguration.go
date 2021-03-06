package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ServiceAwsVpcConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-awsvpcconfiguration.html
type ServiceAwsVpcConfiguration struct {
	AssignPublicIp interface{} `yaml:"AssignPublicIp,omitempty"`
	SecurityGroups interface{} `yaml:"SecurityGroups,omitempty"`
	Subnets        interface{} `yaml:"Subnets"`
}

// ServiceAwsVpcConfiguration validation
func (resource ServiceAwsVpcConfiguration) Validate() []error {
	errs := []error{}

	if resource.Subnets == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Subnets'"))
	}
	return errs
}
