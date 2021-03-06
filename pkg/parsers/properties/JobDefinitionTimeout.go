package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// JobDefinitionTimeout Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-timeout.html
type JobDefinitionTimeout struct {
	AttemptDurationSeconds interface{} `yaml:"AttemptDurationSeconds,omitempty"`
}

// JobDefinitionTimeout validation
func (resource JobDefinitionTimeout) Validate() []error {
	errs := []error{}

	return errs
}
