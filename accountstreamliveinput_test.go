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

func TestAccountStreamLiveInputGet(t *testing.T) {
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
	_, err := client.Accounts.Streams.LiveInputs.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"66be4bf738797e01e1fca35a7bdecdcd",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountStreamLiveInputUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Streams.LiveInputs.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"66be4bf738797e01e1fca35a7bdecdcd",
		cloudflare.AccountStreamLiveInputUpdateParams{
			DefaultCreator:           cloudflare.F("string"),
			DeleteRecordingAfterDays: cloudflare.F(45.000000),
			Meta: cloudflare.F[any](map[string]interface{}{
				"name": "test stream 1",
			}),
			Recording: cloudflare.F(cloudflare.AccountStreamLiveInputUpdateParamsRecording{
				AllowedOrigins:    cloudflare.F([]string{"example.com"}),
				Mode:              cloudflare.F(cloudflare.AccountStreamLiveInputUpdateParamsRecordingModeOff),
				RequireSignedURLs: cloudflare.F(false),
				TimeoutSeconds:    cloudflare.F(int64(0)),
			}),
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

func TestAccountStreamLiveInputDelete(t *testing.T) {
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
	err := client.Accounts.Streams.LiveInputs.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"66be4bf738797e01e1fca35a7bdecdcd",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountStreamLiveInputStreamLiveInputsNewALiveInputWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Streams.LiveInputs.StreamLiveInputsNewALiveInput(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountStreamLiveInputStreamLiveInputsNewALiveInputParams{
			DefaultCreator:           cloudflare.F("string"),
			DeleteRecordingAfterDays: cloudflare.F(45.000000),
			Meta: cloudflare.F[any](map[string]interface{}{
				"name": "test stream 1",
			}),
			Recording: cloudflare.F(cloudflare.AccountStreamLiveInputStreamLiveInputsNewALiveInputParamsRecording{
				AllowedOrigins:    cloudflare.F([]string{"example.com"}),
				Mode:              cloudflare.F(cloudflare.AccountStreamLiveInputStreamLiveInputsNewALiveInputParamsRecordingModeOff),
				RequireSignedURLs: cloudflare.F(false),
				TimeoutSeconds:    cloudflare.F(int64(0)),
			}),
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

func TestAccountStreamLiveInputStreamLiveInputsListLiveInputsWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Streams.LiveInputs.StreamLiveInputsListLiveInputs(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountStreamLiveInputStreamLiveInputsListLiveInputsParams{
			IncludeCounts: cloudflare.F(true),
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
