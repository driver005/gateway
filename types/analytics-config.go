package types

type CreateAnalyticsConfig struct {
	OptOut    bool `json:"opt_out"`
	Anonymize bool `json:"anonymize"`
}

type UpdateAnalyticsConfig struct {
	OptOut    bool `json:"opt_out,omitempty" validate:"omitempty"`
	Anonymize bool `json:"anonymize,omitempty" validate:"omitempty"`
}
