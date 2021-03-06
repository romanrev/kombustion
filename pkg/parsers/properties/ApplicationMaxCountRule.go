package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ApplicationMaxCountRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html
type ApplicationMaxCountRule struct {
	DeleteSourceFromS3 interface{} `yaml:"DeleteSourceFromS3,omitempty"`
	Enabled            interface{} `yaml:"Enabled,omitempty"`
	MaxCount           interface{} `yaml:"MaxCount,omitempty"`
}

// ApplicationMaxCountRule validation
func (resource ApplicationMaxCountRule) Validate() []error {
	errs := []error{}

	return errs
}
