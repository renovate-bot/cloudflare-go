// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// BotManagementService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewBotManagementService] method
// instead.
type BotManagementService struct {
	Options []option.RequestOption
}

// NewBotManagementService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBotManagementService(opts ...option.RequestOption) (r *BotManagementService) {
	r = &BotManagementService{}
	r.Options = opts
	return
}

// Retrieve a zone's Bot Management Config
func (r *BotManagementService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *BotManagementGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BotManagementGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/bot_management", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the Bot Management configuration for a zone.
//
// This API is used to update:
//
// - **Bot Fight Mode**
// - **Super Bot Fight Mode**
// - **Bot Management for Enterprise**
//
// See [Bot Plans](https://developers.cloudflare.com/bots/plans/) for more
// information on the different plans
func (r *BotManagementService) Update(ctx context.Context, zoneID string, body BotManagementUpdateParams, opts ...option.RequestOption) (res *BotManagementUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BotManagementUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/bot_management", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [BotManagementGetResponseBotManagementBotFightModeConfig],
// [BotManagementGetResponseBotManagementSbfmDefinitelyConfig],
// [BotManagementGetResponseBotManagementSbfmLikelyConfig] or
// [BotManagementGetResponseBotManagementBmSubscriptionConfig].
type BotManagementGetResponse interface {
	implementsBotManagementGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*BotManagementGetResponse)(nil)).Elem(), "")
}

type BotManagementGetResponseBotManagementBotFightModeConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                        `json:"using_latest_model"`
	JSON             botManagementGetResponseBotManagementBotFightModeConfigJSON `json:"-"`
}

// botManagementGetResponseBotManagementBotFightModeConfigJSON contains the JSON
// metadata for the struct
// [BotManagementGetResponseBotManagementBotFightModeConfig]
type botManagementGetResponseBotManagementBotFightModeConfigJSON struct {
	EnableJs         apijson.Field
	FightMode        apijson.Field
	UsingLatestModel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BotManagementGetResponseBotManagementBotFightModeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementGetResponseBotManagementBotFightModeConfig) implementsBotManagementGetResponse() {
}

type BotManagementGetResponseBotManagementSbfmDefinitelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                          `json:"using_latest_model"`
	JSON             botManagementGetResponseBotManagementSbfmDefinitelyConfigJSON `json:"-"`
}

// botManagementGetResponseBotManagementSbfmDefinitelyConfigJSON contains the JSON
// metadata for the struct
// [BotManagementGetResponseBotManagementSbfmDefinitelyConfig]
type botManagementGetResponseBotManagementSbfmDefinitelyConfigJSON struct {
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *BotManagementGetResponseBotManagementSbfmDefinitelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementGetResponseBotManagementSbfmDefinitelyConfig) implementsBotManagementGetResponse() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedAllow            BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedBlock            BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots string

const (
	BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsAllow BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "allow"
	BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsBlock BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "block"
)

type BotManagementGetResponseBotManagementSbfmLikelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                      `json:"using_latest_model"`
	JSON             botManagementGetResponseBotManagementSbfmLikelyConfigJSON `json:"-"`
}

// botManagementGetResponseBotManagementSbfmLikelyConfigJSON contains the JSON
// metadata for the struct [BotManagementGetResponseBotManagementSbfmLikelyConfig]
type botManagementGetResponseBotManagementSbfmLikelyConfigJSON struct {
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmLikelyAutomated          apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *BotManagementGetResponseBotManagementSbfmLikelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementGetResponseBotManagementSbfmLikelyConfig) implementsBotManagementGetResponse() {}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedAllow            BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedBlock            BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated string

const (
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomatedAllow            BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "allow"
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomatedBlock            BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "block"
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots string

const (
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmVerifiedBotsAllow BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots = "allow"
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmVerifiedBotsBlock BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots = "block"
)

type BotManagementGetResponseBotManagementBmSubscriptionConfig struct {
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel bool `json:"auto_update_model"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore bool `json:"suppress_session_score"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                          `json:"using_latest_model"`
	JSON             botManagementGetResponseBotManagementBmSubscriptionConfigJSON `json:"-"`
}

// botManagementGetResponseBotManagementBmSubscriptionConfigJSON contains the JSON
// metadata for the struct
// [BotManagementGetResponseBotManagementBmSubscriptionConfig]
type botManagementGetResponseBotManagementBmSubscriptionConfigJSON struct {
	AutoUpdateModel      apijson.Field
	EnableJs             apijson.Field
	SuppressSessionScore apijson.Field
	UsingLatestModel     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BotManagementGetResponseBotManagementBmSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementGetResponseBotManagementBmSubscriptionConfig) implementsBotManagementGetResponse() {
}

// Union satisfied by [BotManagementUpdateResponseBotManagementBotFightModeConfig],
// [BotManagementUpdateResponseBotManagementSbfmDefinitelyConfig],
// [BotManagementUpdateResponseBotManagementSbfmLikelyConfig] or
// [BotManagementUpdateResponseBotManagementBmSubscriptionConfig].
type BotManagementUpdateResponse interface {
	implementsBotManagementUpdateResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*BotManagementUpdateResponse)(nil)).Elem(), "")
}

type BotManagementUpdateResponseBotManagementBotFightModeConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                           `json:"using_latest_model"`
	JSON             botManagementUpdateResponseBotManagementBotFightModeConfigJSON `json:"-"`
}

// botManagementUpdateResponseBotManagementBotFightModeConfigJSON contains the JSON
// metadata for the struct
// [BotManagementUpdateResponseBotManagementBotFightModeConfig]
type botManagementUpdateResponseBotManagementBotFightModeConfigJSON struct {
	EnableJs         apijson.Field
	FightMode        apijson.Field
	UsingLatestModel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BotManagementUpdateResponseBotManagementBotFightModeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementUpdateResponseBotManagementBotFightModeConfig) implementsBotManagementUpdateResponse() {
}

type BotManagementUpdateResponseBotManagementSbfmDefinitelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                             `json:"using_latest_model"`
	JSON             botManagementUpdateResponseBotManagementSbfmDefinitelyConfigJSON `json:"-"`
}

// botManagementUpdateResponseBotManagementSbfmDefinitelyConfigJSON contains the
// JSON metadata for the struct
// [BotManagementUpdateResponseBotManagementSbfmDefinitelyConfig]
type botManagementUpdateResponseBotManagementSbfmDefinitelyConfigJSON struct {
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *BotManagementUpdateResponseBotManagementSbfmDefinitelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementUpdateResponseBotManagementSbfmDefinitelyConfig) implementsBotManagementUpdateResponse() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedAllow            BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedBlock            BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots string

const (
	BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsAllow BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "allow"
	BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsBlock BotManagementUpdateResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "block"
)

type BotManagementUpdateResponseBotManagementSbfmLikelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                         `json:"using_latest_model"`
	JSON             botManagementUpdateResponseBotManagementSbfmLikelyConfigJSON `json:"-"`
}

// botManagementUpdateResponseBotManagementSbfmLikelyConfigJSON contains the JSON
// metadata for the struct
// [BotManagementUpdateResponseBotManagementSbfmLikelyConfig]
type botManagementUpdateResponseBotManagementSbfmLikelyConfigJSON struct {
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmLikelyAutomated          apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *BotManagementUpdateResponseBotManagementSbfmLikelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementUpdateResponseBotManagementSbfmLikelyConfig) implementsBotManagementUpdateResponse() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedAllow            BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedBlock            BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated string

const (
	BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomatedAllow            BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "allow"
	BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomatedBlock            BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "block"
	BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots string

const (
	BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmVerifiedBotsAllow BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots = "allow"
	BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmVerifiedBotsBlock BotManagementUpdateResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots = "block"
)

type BotManagementUpdateResponseBotManagementBmSubscriptionConfig struct {
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel bool `json:"auto_update_model"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore bool `json:"suppress_session_score"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                             `json:"using_latest_model"`
	JSON             botManagementUpdateResponseBotManagementBmSubscriptionConfigJSON `json:"-"`
}

// botManagementUpdateResponseBotManagementBmSubscriptionConfigJSON contains the
// JSON metadata for the struct
// [BotManagementUpdateResponseBotManagementBmSubscriptionConfig]
type botManagementUpdateResponseBotManagementBmSubscriptionConfigJSON struct {
	AutoUpdateModel      apijson.Field
	EnableJs             apijson.Field
	SuppressSessionScore apijson.Field
	UsingLatestModel     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BotManagementUpdateResponseBotManagementBmSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementUpdateResponseBotManagementBmSubscriptionConfig) implementsBotManagementUpdateResponse() {
}

type BotManagementGetResponseEnvelope struct {
	Errors   []BotManagementGetResponseEnvelopeErrors   `json:"errors"`
	Messages []BotManagementGetResponseEnvelopeMessages `json:"messages"`
	Result   BotManagementGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success BotManagementGetResponseEnvelopeSuccess `json:"success"`
	JSON    botManagementGetResponseEnvelopeJSON    `json:"-"`
}

// botManagementGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BotManagementGetResponseEnvelope]
type botManagementGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BotManagementGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    botManagementGetResponseEnvelopeErrorsJSON `json:"-"`
}

// botManagementGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BotManagementGetResponseEnvelopeErrors]
type botManagementGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BotManagementGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    botManagementGetResponseEnvelopeMessagesJSON `json:"-"`
}

// botManagementGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [BotManagementGetResponseEnvelopeMessages]
type botManagementGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type BotManagementGetResponseEnvelopeSuccess bool

const (
	BotManagementGetResponseEnvelopeSuccessTrue BotManagementGetResponseEnvelopeSuccess = true
)

// This interface is a union satisfied by one of the following:
// [BotManagementUpdateParamsBotManagementBotFightModeConfig],
// [BotManagementUpdateParamsBotManagementSbfmDefinitelyConfig],
// [BotManagementUpdateParamsBotManagementSbfmLikelyConfig],
// [BotManagementUpdateParamsBotManagementBmSubscriptionConfig].
type BotManagementUpdateParams interface {
	ImplementsBotManagementUpdateParams()
}

type BotManagementUpdateParamsBotManagementBotFightModeConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode param.Field[bool] `json:"fight_mode"`
}

func (r BotManagementUpdateParamsBotManagementBotFightModeConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (BotManagementUpdateParamsBotManagementBotFightModeConfig) ImplementsBotManagementUpdateParams() {

}

type BotManagementUpdateParamsBotManagementSbfmDefinitelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated param.Field[BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots param.Field[BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r BotManagementUpdateParamsBotManagementSbfmDefinitelyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (BotManagementUpdateParamsBotManagementSbfmDefinitelyConfig) ImplementsBotManagementUpdateParams() {

}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedAllow            BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedBlock            BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBots string

const (
	BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsAllow BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "allow"
	BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsBlock BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "block"
)

type BotManagementUpdateParamsBotManagementSbfmLikelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated param.Field[BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated param.Field[BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated] `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots param.Field[BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r BotManagementUpdateParamsBotManagementSbfmLikelyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (BotManagementUpdateParamsBotManagementSbfmLikelyConfig) ImplementsBotManagementUpdateParams() {

}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedAllow            BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedBlock            BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated string

const (
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomatedAllow            BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "allow"
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomatedBlock            BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "block"
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmVerifiedBots string

const (
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmVerifiedBotsAllow BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmVerifiedBots = "allow"
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmVerifiedBotsBlock BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmVerifiedBots = "block"
)

type BotManagementUpdateParamsBotManagementBmSubscriptionConfig struct {
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel param.Field[bool] `json:"auto_update_model"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore param.Field[bool] `json:"suppress_session_score"`
}

func (r BotManagementUpdateParamsBotManagementBmSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (BotManagementUpdateParamsBotManagementBmSubscriptionConfig) ImplementsBotManagementUpdateParams() {

}

type BotManagementUpdateResponseEnvelope struct {
	Errors   []BotManagementUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []BotManagementUpdateResponseEnvelopeMessages `json:"messages"`
	Result   BotManagementUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success BotManagementUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    botManagementUpdateResponseEnvelopeJSON    `json:"-"`
}

// botManagementUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [BotManagementUpdateResponseEnvelope]
type botManagementUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BotManagementUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    botManagementUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// botManagementUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BotManagementUpdateResponseEnvelopeErrors]
type botManagementUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BotManagementUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    botManagementUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// botManagementUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BotManagementUpdateResponseEnvelopeMessages]
type botManagementUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type BotManagementUpdateResponseEnvelopeSuccess bool

const (
	BotManagementUpdateResponseEnvelopeSuccessTrue BotManagementUpdateResponseEnvelopeSuccess = true
)
