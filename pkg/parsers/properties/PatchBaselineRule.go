package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// PatchBaselineRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html
type PatchBaselineRule struct {
	ApproveAfterDays  interface{}                    `yaml:"ApproveAfterDays,omitempty"`
	ComplianceLevel   interface{}                    `yaml:"ComplianceLevel,omitempty"`
	EnableNonSecurity interface{}                    `yaml:"EnableNonSecurity,omitempty"`
	PatchFilterGroup  *PatchBaselinePatchFilterGroup `yaml:"PatchFilterGroup,omitempty"`
}

// PatchBaselineRule validation
func (resource PatchBaselineRule) Validate() []error {
	errs := []error{}

	return errs
}
