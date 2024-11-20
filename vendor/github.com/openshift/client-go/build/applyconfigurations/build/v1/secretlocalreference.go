// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// SecretLocalReferenceApplyConfiguration represents an declarative configuration of the SecretLocalReference type for use
// with apply.
type SecretLocalReferenceApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
}

// SecretLocalReferenceApplyConfiguration constructs an declarative configuration of the SecretLocalReference type for use with
// apply.
func SecretLocalReference() *SecretLocalReferenceApplyConfiguration {
	return &SecretLocalReferenceApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SecretLocalReferenceApplyConfiguration) WithName(value string) *SecretLocalReferenceApplyConfiguration {
	b.Name = &value
	return b
}
