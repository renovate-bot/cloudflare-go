// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RumRuleService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRumRuleService] method instead.
type RumRuleService struct {
	Options []option.RequestOption
}

// NewRumRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRumRuleService(opts ...option.RequestOption) (r *RumRuleService) {
	r = &RumRuleService{}
	r.Options = opts
	return
}

// Creates a new rule in a Web Analytics ruleset.
func (r *RumRuleService) New(ctx context.Context, accountIdentifier string, rulesetIdentifier string, body RumRuleNewParams, opts ...option.RequestOption) (res *RumRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RumRuleNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule", accountIdentifier, rulesetIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a rule in a Web Analytics ruleset.
func (r *RumRuleService) Update(ctx context.Context, accountIdentifier string, rulesetIdentifier string, ruleIdentifier string, body RumRuleUpdateParams, opts ...option.RequestOption) (res *RumRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RumRuleUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule/%s", accountIdentifier, rulesetIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all the rules in a Web Analytics ruleset.
func (r *RumRuleService) List(ctx context.Context, accountIdentifier string, rulesetIdentifier string, opts ...option.RequestOption) (res *RumRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RumRuleListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rules", accountIdentifier, rulesetIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing rule from a Web Analytics ruleset.
func (r *RumRuleService) Delete(ctx context.Context, accountIdentifier string, rulesetIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *RumRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RumRuleDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule/%s", accountIdentifier, rulesetIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RumRuleNewResponse struct {
	// The Web Analytics rule identifier.
	ID      string    `json:"id"`
	Created time.Time `json:"created" format:"date-time"`
	// The hostname the rule will be applied to.
	Host string `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive bool `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused bool `json:"is_paused"`
	// The paths the rule will be applied to.
	Paths    []string               `json:"paths"`
	Priority float64                `json:"priority"`
	JSON     rumRuleNewResponseJSON `json:"-"`
}

// rumRuleNewResponseJSON contains the JSON metadata for the struct
// [RumRuleNewResponse]
type rumRuleNewResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Host        apijson.Field
	Inclusive   apijson.Field
	IsPaused    apijson.Field
	Paths       apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleUpdateResponse struct {
	// The Web Analytics rule identifier.
	ID      string    `json:"id"`
	Created time.Time `json:"created" format:"date-time"`
	// The hostname the rule will be applied to.
	Host string `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive bool `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused bool `json:"is_paused"`
	// The paths the rule will be applied to.
	Paths    []string                  `json:"paths"`
	Priority float64                   `json:"priority"`
	JSON     rumRuleUpdateResponseJSON `json:"-"`
}

// rumRuleUpdateResponseJSON contains the JSON metadata for the struct
// [RumRuleUpdateResponse]
type rumRuleUpdateResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Host        apijson.Field
	Inclusive   apijson.Field
	IsPaused    apijson.Field
	Paths       apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleListResponse struct {
	// A list of rules.
	Rules   []RumRuleListResponseRule  `json:"rules"`
	Ruleset RumRuleListResponseRuleset `json:"ruleset"`
	JSON    rumRuleListResponseJSON    `json:"-"`
}

// rumRuleListResponseJSON contains the JSON metadata for the struct
// [RumRuleListResponse]
type rumRuleListResponseJSON struct {
	Rules       apijson.Field
	Ruleset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleListResponseRule struct {
	// The Web Analytics rule identifier.
	ID      string    `json:"id"`
	Created time.Time `json:"created" format:"date-time"`
	// The hostname the rule will be applied to.
	Host string `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive bool `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused bool `json:"is_paused"`
	// The paths the rule will be applied to.
	Paths    []string                    `json:"paths"`
	Priority float64                     `json:"priority"`
	JSON     rumRuleListResponseRuleJSON `json:"-"`
}

// rumRuleListResponseRuleJSON contains the JSON metadata for the struct
// [RumRuleListResponseRule]
type rumRuleListResponseRuleJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Host        apijson.Field
	Inclusive   apijson.Field
	IsPaused    apijson.Field
	Paths       apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleListResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleListResponseRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                         `json:"zone_tag"`
	JSON    rumRuleListResponseRulesetJSON `json:"-"`
}

// rumRuleListResponseRulesetJSON contains the JSON metadata for the struct
// [RumRuleListResponseRuleset]
type rumRuleListResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleListResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleDeleteResponse struct {
	// The Web Analytics rule identifier.
	ID   string                    `json:"id"`
	JSON rumRuleDeleteResponseJSON `json:"-"`
}

// rumRuleDeleteResponseJSON contains the JSON metadata for the struct
// [RumRuleDeleteResponse]
type rumRuleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleNewParams struct {
	Host param.Field[string] `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive param.Field[bool] `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused param.Field[bool]     `json:"is_paused"`
	Paths    param.Field[[]string] `json:"paths"`
}

func (r RumRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RumRuleNewResponseEnvelope struct {
	Errors   []RumRuleNewResponseEnvelopeErrors   `json:"errors"`
	Messages []RumRuleNewResponseEnvelopeMessages `json:"messages"`
	Result   RumRuleNewResponse                   `json:"result"`
	// Whether the API call was successful.
	Success bool                           `json:"success"`
	JSON    rumRuleNewResponseEnvelopeJSON `json:"-"`
}

// rumRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RumRuleNewResponseEnvelope]
type rumRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleNewResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    rumRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// rumRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RumRuleNewResponseEnvelopeErrors]
type rumRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleNewResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    rumRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// rumRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RumRuleNewResponseEnvelopeMessages]
type rumRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleUpdateParams struct {
	Host param.Field[string] `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive param.Field[bool] `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused param.Field[bool]     `json:"is_paused"`
	Paths    param.Field[[]string] `json:"paths"`
}

func (r RumRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RumRuleUpdateResponseEnvelope struct {
	Errors   []RumRuleUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []RumRuleUpdateResponseEnvelopeMessages `json:"messages"`
	Result   RumRuleUpdateResponse                   `json:"result"`
	// Whether the API call was successful.
	Success bool                              `json:"success"`
	JSON    rumRuleUpdateResponseEnvelopeJSON `json:"-"`
}

// rumRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RumRuleUpdateResponseEnvelope]
type rumRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    rumRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// rumRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RumRuleUpdateResponseEnvelopeErrors]
type rumRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    rumRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// rumRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RumRuleUpdateResponseEnvelopeMessages]
type rumRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleListResponseEnvelope struct {
	Errors   []RumRuleListResponseEnvelopeErrors   `json:"errors"`
	Messages []RumRuleListResponseEnvelopeMessages `json:"messages"`
	Result   RumRuleListResponse                   `json:"result"`
	// Whether the API call was successful.
	Success bool                            `json:"success"`
	JSON    rumRuleListResponseEnvelopeJSON `json:"-"`
}

// rumRuleListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RumRuleListResponseEnvelope]
type rumRuleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    rumRuleListResponseEnvelopeErrorsJSON `json:"-"`
}

// rumRuleListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RumRuleListResponseEnvelopeErrors]
type rumRuleListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    rumRuleListResponseEnvelopeMessagesJSON `json:"-"`
}

// rumRuleListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RumRuleListResponseEnvelopeMessages]
type rumRuleListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleDeleteResponseEnvelope struct {
	Errors   []RumRuleDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []RumRuleDeleteResponseEnvelopeMessages `json:"messages"`
	Result   RumRuleDeleteResponse                   `json:"result"`
	// Whether the API call was successful.
	Success bool                              `json:"success"`
	JSON    rumRuleDeleteResponseEnvelopeJSON `json:"-"`
}

// rumRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RumRuleDeleteResponseEnvelope]
type rumRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    rumRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// rumRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RumRuleDeleteResponseEnvelopeErrors]
type rumRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    rumRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// rumRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RumRuleDeleteResponseEnvelopeMessages]
type rumRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
