package buildkit

import (
	"testing"

	"github.com/railwayapp/railpack/core/plan"
	"github.com/stretchr/testify/require"
)

func TestConvertPlanToLLB_ServiceLabel(t *testing.T) {
	buildPlan := &plan.BuildPlan{
		Service: "my-test-service",
		Steps:   []plan.Step{},
		Secrets: []string{},
		Deploy: plan.Deploy{
			StartCmd: "echo hello",
		},
	}

	opts := ConvertPlanOptions{
		BuildPlatform: BuildPlatform{
			OS:           "linux",
			Architecture: "amd64",
		},
		SecretsHash: "",
		CacheKey:    "test",
		SessionID:   "test-session",
	}

	_, image, err := ConvertPlanToLLB(buildPlan, opts)
	require.NoError(t, err)
	require.NotNil(t, image)

	// Verify that the service label is set correctly
	require.NotNil(t, image.Config.Labels)
	require.Equal(t, "my-test-service", image.Config.Labels["service"])
}

func TestConvertPlanToLLB_NoServiceLabel(t *testing.T) {
	buildPlan := &plan.BuildPlan{
		Service: "", // Empty service
		Steps:   []plan.Step{},
		Secrets: []string{},
		Deploy: plan.Deploy{
			StartCmd: "echo hello",
		},
	}

	opts := ConvertPlanOptions{
		BuildPlatform: BuildPlatform{
			OS:           "linux",
			Architecture: "amd64",
		},
		SecretsHash: "",
		CacheKey:    "test",
		SessionID:   "test-session",
	}

	_, image, err := ConvertPlanToLLB(buildPlan, opts)
	require.NoError(t, err)
	require.NotNil(t, image)

	// Verify that no service label is set when service is empty
	if image.Config.Labels != nil {
		_, hasServiceLabel := image.Config.Labels["service"]
		require.False(t, hasServiceLabel, "service label should not be present when service is empty")
	}
}
