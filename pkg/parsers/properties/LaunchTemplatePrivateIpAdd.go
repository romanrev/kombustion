package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// LaunchTemplatePrivateIpAdd Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-privateipadd.html
type LaunchTemplatePrivateIpAdd struct {
	Primary          interface{} `yaml:"Primary,omitempty"`
	PrivateIpAddress interface{} `yaml:"PrivateIpAddress,omitempty"`
}

// LaunchTemplatePrivateIpAdd validation
func (resource LaunchTemplatePrivateIpAdd) Validate() []error {
	errs := []error{}

	return errs
}
