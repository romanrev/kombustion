package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TaskDefinitionKernelCapabilities Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-kernelcapabilities.html
type TaskDefinitionKernelCapabilities struct {
	Add  interface{} `yaml:"Add,omitempty"`
	Drop interface{} `yaml:"Drop,omitempty"`
}

// TaskDefinitionKernelCapabilities validation
func (resource TaskDefinitionKernelCapabilities) Validate() []error {
	errs := []error{}

	return errs
}