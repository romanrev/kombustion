package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// PipelineArtifactStore Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html
type PipelineArtifactStore struct {
	Location      interface{}            `yaml:"Location"`
	Type          interface{}            `yaml:"Type"`
	EncryptionKey *PipelineEncryptionKey `yaml:"EncryptionKey,omitempty"`
}

// PipelineArtifactStore validation
func (resource PipelineArtifactStore) Validate() []error {
	errs := []error{}

	if resource.Location == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Location'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
