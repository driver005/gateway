package registry

import (
	"context"

	"github.com/driver005/gateway/logger"
)

type OptionsModifier func(*options)

type options struct {
	forcedValues map[string]interface{}
	preload      bool
	validate     bool
	migrate      bool
	opts         []OptionsModifier
	config       map[string]interface{} //*config.DefaultProvider
	// The first default refers to determining the NID at startup; the second default referes to the fact that the Contextualizer may dynamically change the NID.
	skipNetworkInit bool
}

func NewOptions() *options {
	return &options{
		validate: true,
		preload:  true,
		migrate:  true,
		opts:     []OptionsModifier{},
	}
}

func WithConfig(config map[string]interface{}) func(o *options) {
	return func(o *options) {
		o.config = config
	}
}

func WithOptions(opts ...OptionsModifier) OptionsModifier {
	return func(o *options) {
		o.opts = append(o.opts, opts...)
	}
}

// DisableValidation validating the config.
//
// This does not affect schema validation!
func DisableValidation() OptionsModifier {
	return func(o *options) {
		o.validate = false
	}
}

// DisablePreloading will not preload the config.
func DisablePreloading() OptionsModifier {
	return func(o *options) {
		o.preload = false
	}
}

// DisablePreloading will not preload the config.
func DisableMigration() OptionsModifier {
	return func(o *options) {
		o.migrate = false
	}
}

func SkipNetworkInit() OptionsModifier {
	return func(o *options) {
		o.skipNetworkInit = true
	}
}

func New(ctx context.Context, opts []OptionsModifier) Registry {
	o := NewOptions()
	for _, f := range opts {
		f(o)
	}

	// l = logger.New("Ory Hydra", config.Version)

	l := logger.New("ORY Hydra", "2.1")

	r, err := NewRegistryFromDSN(ctx, l)
	if err != nil {
		l.WithError(err).Fatal("Unable to create service registry.")
	}

	if err = r.Init(ctx); err != nil {
		l.WithError(err).Fatal("Unable to initialize service registry.")
	}

	if o.preload {
		CallRegistry(ctx, r)
	}

	return r
}
