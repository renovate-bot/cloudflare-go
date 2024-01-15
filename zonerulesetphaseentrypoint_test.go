// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestZoneRulesetPhaseEntrypointTransformRulesListTransformRules(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Rulesets.Phases.Entrypoints.TransformRulesListTransformRules(
		context.TODO(),
		"9f1839b6152d298aca64c4e906b6d074",
		cloudflare.ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPRequestFirewallCustom,
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Rulesets.Phases.Entrypoints.TransformRulesUpdateTransformRules(
		context.TODO(),
		"9f1839b6152d298aca64c4e906b6d074",
		cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPRequestFirewallCustom,
		cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParams{
			ID:          cloudflare.F[any](map[string]interface{}{}),
			Description: cloudflare.F("My ruleset to execute managed rulesets"),
			Kind:        cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsKindRoot),
			Name:        cloudflare.F("My ruleset"),
			Phase:       cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestFirewallCustom),
			Rules: cloudflare.F([]cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRule{cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRule(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRule{
				Action: cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionBlock),
				ActionParameters: cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionParameters{
					Response: cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F[any]("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F[any](map[string]interface{}{}),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRule(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRule{
				Action: cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionBlock),
				ActionParameters: cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionParameters{
					Response: cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F[any]("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F[any](map[string]interface{}{}),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}), cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRule(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRule{
				Action: cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionBlock),
				ActionParameters: cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionParameters{
					Response: cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionParametersResponse{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F[any]("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F[any](map[string]interface{}{}),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				ID:          cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Logging: cloudflare.F(cloudflare.ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleLogging{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			})}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
