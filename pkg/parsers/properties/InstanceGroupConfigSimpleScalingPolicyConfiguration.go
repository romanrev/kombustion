package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// InstanceGroupConfigSimpleScalingPolicyConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html
type InstanceGroupConfigSimpleScalingPolicyConfiguration struct {
	AdjustmentType    interface{} `yaml:"AdjustmentType,omitempty"`
	CoolDown          interface{} `yaml:"CoolDown,omitempty"`
	ScalingAdjustment interface{} `yaml:"ScalingAdjustment"`
}

// InstanceGroupConfigSimpleScalingPolicyConfiguration validation
func (resource InstanceGroupConfigSimpleScalingPolicyConfiguration) Validate() []error {
	errs := []error{}

	if resource.ScalingAdjustment == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ScalingAdjustment'"))
	}
	return errs
}
