package servicegraphprocessor

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/servicegraphprocessor/internal/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/service/servicetest"
)

func TestLoadConfig(t *testing.T) {
	// Prepare
	factories, err := componenttest.NopFactories()
	require.NoError(t, err)

	factories.Processors[typeStr] = NewFactory()

	// Test
	cfg, err := servicetest.LoadConfigAndValidate(filepath.Join("testdata", "service-graph-config.yaml"), factories)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Equal(t,
		&Config{
			ProcessorSettings:       config.NewProcessorSettings(config.NewComponentID(typeStr)),
			MetricsExporter:         "metrics",
			LatencyHistogramBuckets: []time.Duration{1, 2, 3, 4, 5},
			Dimensions:              []string{"dimension-1", "dimension-2"},
			Store: store.Config{
				TTL:      time.Second,
				MaxItems: 10,
			},
		},
		cfg.Processors[config.NewComponentID(typeStr)],
	)
}
