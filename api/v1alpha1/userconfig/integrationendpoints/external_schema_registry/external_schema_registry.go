// Code generated by user config generator. DO NOT EDIT.
// +kubebuilder:object:generate=true

package externalschemaregistryuserconfig

type ExternalSchemaRegistryUserConfig struct {
	// +kubebuilder:validation:Enum="none";"basic"
	// Authentication method
	Authentication string `groups:"create,update" json:"authentication"`

	// +kubebuilder:validation:MaxLength=256
	// Basic authentication password
	BasicAuthPassword *string `groups:"create,update" json:"basic_auth_password,omitempty"`

	// +kubebuilder:validation:MaxLength=256
	// Basic authentication user name
	BasicAuthUsername *string `groups:"create,update" json:"basic_auth_username,omitempty"`

	// +kubebuilder:validation:MaxLength=2048
	// Schema Registry URL
	Url string `groups:"create,update" json:"url"`
}
