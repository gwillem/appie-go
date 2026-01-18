package appie

import (
	"context"
	"fmt"
	"net/http"
)

const defaultAppVersion = "9.28"

// FeatureFlags represents the feature flags returned by the API.
// Values are rollout percentages (0-100).
type FeatureFlags map[string]int

// GetFeatureFlags retrieves the feature flags for the iOS app.
func (c *Client) GetFeatureFlags(ctx context.Context) (FeatureFlags, error) {
	return c.GetFeatureFlagsForVersion(ctx, defaultAppVersion)
}

// GetFeatureFlagsForVersion retrieves feature flags for a specific app version.
func (c *Client) GetFeatureFlagsForVersion(ctx context.Context, version string) (FeatureFlags, error) {
	path := fmt.Sprintf("/mobile-services/config/v1/features/ios?version=%s", version)

	var result FeatureFlags
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("get feature flags failed: %w", err)
	}

	return result, nil
}

// IsEnabled checks if a feature flag is enabled (rollout >= 100).
func (f FeatureFlags) IsEnabled(flag string) bool {
	return f[flag] >= 100
}

// Rollout returns the rollout percentage for a feature flag.
func (f FeatureFlags) Rollout(flag string) int {
	return f[flag]
}
