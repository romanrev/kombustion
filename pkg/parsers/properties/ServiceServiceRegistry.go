package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ServiceServiceRegistry Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceregistry.html
type ServiceServiceRegistry struct {
	Port        interface{} `yaml:"Port,omitempty"`
	RegistryArn interface{} `yaml:"RegistryArn,omitempty"`
}

// ServiceServiceRegistry validation
func (resource ServiceServiceRegistry) Validate() []error {
	errs := []error{}

	return errs
}
