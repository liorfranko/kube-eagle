package options

// Options are configuration options that can be set by Environment Variables
type Options struct {
	// General
	Version string `envconfig:"VERSION" required:"false"`

	// Kubernetes
	// IsInCluster - Whether to use in cluster communication (if deployed inside of Kubernetes) or to look for a kubeconfig in home directory
	IsInCluster bool `envconfig:"IS_IN_CLUSTER" default:"false"`

	// Prometheus
	// Host - Host to bind socket on for the prometheus exporter
	// Port - Port to listen on for the prometheus exporter
	// Namespace - Prefix of exposed prometheus metrics
	Host         string `envconfig:"TELEMETRY_HOST" default:"0.0.0.0"`
	Port         int    `envconfig:"TELEMETRY_PORT" default:"8080"`
	Namespace    string `envconfig:"METRICS_NAMESPACE" default:"eagle"`
	EnableLabels bool   `envconfig:"ENABLE_LABELS" default:"false"`

	// Logger
	// LogLevel - Logger's log granularity (debug, info, warn, error, fatal, panic)
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
}

// NewOptions provides Application Options
func NewOptions() *Options {
	return &Options{}
}
