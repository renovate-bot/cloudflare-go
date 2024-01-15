// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserLoadBalancerMonitorReferenceService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewUserLoadBalancerMonitorReferenceService] method instead.
type UserLoadBalancerMonitorReferenceService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerMonitorReferenceService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserLoadBalancerMonitorReferenceService(opts ...option.RequestOption) (r *UserLoadBalancerMonitorReferenceService) {
	r = &UserLoadBalancerMonitorReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided monitor.
func (r *UserLoadBalancerMonitorReferenceService) LoadBalancerMonitorsListMonitorReferences(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/monitors/%s/references", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponse struct {
	Errors   []UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseError   `json:"errors"`
	Messages []UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseMessage `json:"messages"`
	// List of resources that reference a given monitor.
	Result     []UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResult   `json:"result"`
	ResultInfo UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseSuccess `json:"success"`
	JSON    userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseJSON    `json:"-"`
}

// userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponse]
type userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseError struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseErrorJSON `json:"-"`
}

// userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseErrorJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseError]
type userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseMessage struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseMessageJSON `json:"-"`
}

// userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseMessageJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseMessage]
type userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResult struct {
	ReferenceType UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType `json:"reference_type"`
	ResourceID    string                                                                                               `json:"resource_id"`
	ResourceName  string                                                                                               `json:"resource_name"`
	ResourceType  string                                                                                               `json:"resource_type"`
	JSON          userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultJSON          `json:"-"`
}

// userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResult]
type userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType string

const (
	UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceTypeStar     UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType = "*"
	UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceTypeReferral UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType = "referral"
	UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceTypeReferrer UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType = "referrer"
)

type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                         `json:"total_count"`
	JSON       userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultInfoJSON `json:"-"`
}

// userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultInfo]
type userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseSuccess bool

const (
	UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseSuccessTrue UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseSuccess = true
)
