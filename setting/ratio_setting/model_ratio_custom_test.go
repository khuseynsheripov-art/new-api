package ratio_setting

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOfficialGPT5PointReleasePricingRatios(t *testing.T) {
	InitRatioSettings()

	tests := []struct {
		model           string
		modelRatio      float64
		completionRatio float64
		cacheRatio      float64
	}{
		{model: "gpt-5.5", modelRatio: 2.5, completionRatio: 6, cacheRatio: 0.1},
		{model: "gpt-5.4", modelRatio: 1.25, completionRatio: 6, cacheRatio: 0.1},
		{model: "gpt-5.4-mini", modelRatio: 0.375, completionRatio: 6, cacheRatio: 0.1},
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			modelRatio, usePrice, ok := GetModelRatioOrPrice(tt.model)
			require.True(t, ok)
			require.False(t, usePrice)
			assert.Equal(t, tt.modelRatio, modelRatio)
			assert.Equal(t, tt.completionRatio, GetCompletionRatio(tt.model))

			cacheRatio, ok := GetCacheRatio(tt.model)
			require.True(t, ok)
			assert.Equal(t, tt.cacheRatio, cacheRatio)
		})
	}
}

func TestCustomCodexModelsUseReducedCacheRatio(t *testing.T) {
	InitRatioSettings()

	for _, model := range []string{"gpt-5.3-codex-spark", "codex-auto-review"} {
		t.Run(model, func(t *testing.T) {
			cacheRatio, ok := GetCacheRatio(model)
			require.True(t, ok)
			assert.Equal(t, 0.1, cacheRatio)
		})
	}
}
