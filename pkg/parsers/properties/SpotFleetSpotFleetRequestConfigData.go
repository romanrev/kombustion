package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// SpotFleetSpotFleetRequestConfigData Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html
type SpotFleetSpotFleetRequestConfigData struct {
	AllocationStrategy               interface{} `yaml:"AllocationStrategy,omitempty"`
	ExcessCapacityTerminationPolicy  interface{} `yaml:"ExcessCapacityTerminationPolicy,omitempty"`
	IamFleetRole                     interface{} `yaml:"IamFleetRole"`
	ReplaceUnhealthyInstances        interface{} `yaml:"ReplaceUnhealthyInstances,omitempty"`
	SpotPrice                        interface{} `yaml:"SpotPrice,omitempty"`
	TargetCapacity                   interface{} `yaml:"TargetCapacity"`
	TerminateInstancesWithExpiration interface{} `yaml:"TerminateInstancesWithExpiration,omitempty"`
	Type                             interface{} `yaml:"Type,omitempty"`
	ValidFrom                        interface{} `yaml:"ValidFrom,omitempty"`
	ValidUntil                       interface{} `yaml:"ValidUntil,omitempty"`
	LaunchSpecifications             interface{} `yaml:"LaunchSpecifications,omitempty"`
}

// SpotFleetSpotFleetRequestConfigData validation
func (resource SpotFleetSpotFleetRequestConfigData) Validate() []error {
	errs := []error{}

	if resource.IamFleetRole == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IamFleetRole'"))
	}
	if resource.TargetCapacity == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetCapacity'"))
	}
	return errs
}
