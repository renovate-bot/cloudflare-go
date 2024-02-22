// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AIService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewAIService] method instead.
type AIService struct {
	Options []option.RequestOption
}

// NewAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIService(opts ...option.RequestOption) (r *AIService) {
	r = &AIService{}
	r.Options = opts
	return
}

// This endpoint provides users with the capability to run specific AI models
// on-demand.
//
// By submitting the required input data, users can receive real-time predictions
// or results generated by the chosen AI model. The endpoint supports various AI
// model types, ensuring flexibility and adaptability for diverse use cases.
func (r *AIService) Run(ctx context.Context, accountID string, modelName string, body AIRunParams, opts ...option.RequestOption) (res *AIRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AIRunResponseEnvelope
	path := fmt.Sprintf("accounts/%s/ai/run/%s", accountID, modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIRunResponse = interface{}

type AIRunParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AIRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AIRunResponseEnvelope struct {
	Errors   []AIRunResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                      `json:"messages,required"`
	Result   AIRunResponse                 `json:"result,required"`
	Success  bool                          `json:"success,required"`
	JSON     aiRunResponseEnvelopeJSON     `json:"-"`
}

// aiRunResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIRunResponseEnvelope]
type aiRunResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AIRunResponseEnvelopeErrors struct {
	Message string                          `json:"message,required"`
	JSON    aiRunResponseEnvelopeErrorsJSON `json:"-"`
}

// aiRunResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AIRunResponseEnvelopeErrors]
type aiRunResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
