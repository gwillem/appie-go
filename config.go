package appie

import (
	"context"
	"fmt"
	"net/http"
)

const defaultAppVersion = "9.28"

// FeatureFlags represents the feature flags returned by the API.
// Values are rollout percentages (0-100), where 100 means fully enabled
// and 0 means disabled. Intermediate values indicate partial rollout.
//
// Example usage:
//
//	flags, _ := client.GetFeatureFlags(ctx)
//	if flags.IsEnabled("dark-mode") {
//	    // Dark mode is available
//	}
type FeatureFlags map[string]int

// GetFeatureFlags retrieves the current feature flags for the iOS app.
// These flags control feature availability and A/B testing.
func (c *Client) GetFeatureFlags(ctx context.Context) (FeatureFlags, error) {
	return c.GetFeatureFlagsForVersion(ctx, defaultAppVersion)
}

// GetFeatureFlagsForVersion retrieves feature flags for a specific app version.
// Different versions may have different features enabled.
func (c *Client) GetFeatureFlagsForVersion(ctx context.Context, version string) (FeatureFlags, error) {
	path := fmt.Sprintf("/mobile-services/config/v1/features/ios?version=%s", version)

	var result FeatureFlags
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("get feature flags failed: %w", err)
	}

	return result, nil
}

// IsEnabled returns true if a feature flag is fully enabled (rollout >= 100).
func (f FeatureFlags) IsEnabled(flag string) bool {
	return f[flag] >= 100
}

// Rollout returns the rollout percentage (0-100) for a feature flag.
// Returns 0 if the flag doesn't exist.
func (f FeatureFlags) Rollout(flag string) int {
	return f[flag]
}
