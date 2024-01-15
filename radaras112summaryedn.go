// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAs112SummaryEdnService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAs112SummaryEdnService]
// method instead.
type RadarAs112SummaryEdnService struct {
	Options []option.RequestOption
}

// NewRadarAs112SummaryEdnService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAs112SummaryEdnService(opts ...option.RequestOption) (r *RadarAs112SummaryEdnService) {
	r = &RadarAs112SummaryEdnService{}
	r.Options = opts
	return
}

// Percentage distribution of DNS queries, to AS112, by EDNS support.
func (r *RadarAs112SummaryEdnService) List(ctx context.Context, query RadarAs112SummaryEdnListParams, opts ...option.RequestOption) (res *RadarAs112SummaryEdnListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112SummaryEdnListResponse struct {
	Result  RadarAs112SummaryEdnListResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarAs112SummaryEdnListResponseJSON   `json:"-"`
}

// radarAs112SummaryEdnListResponseJSON contains the JSON metadata for the struct
// [RadarAs112SummaryEdnListResponse]
type radarAs112SummaryEdnListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryEdnListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnListResponseResult struct {
	Meta     RadarAs112SummaryEdnListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAs112SummaryEdnListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryEdnListResponseResultJSON     `json:"-"`
}

// radarAs112SummaryEdnListResponseResultJSON contains the JSON metadata for the
// struct [RadarAs112SummaryEdnListResponseResult]
type radarAs112SummaryEdnListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryEdnListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnListResponseResultMeta struct {
	DateRange      []RadarAs112SummaryEdnListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	Normalization  string                                                   `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryEdnListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryEdnListResponseResultMetaJSON           `json:"-"`
}

// radarAs112SummaryEdnListResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarAs112SummaryEdnListResponseResultMeta]
type radarAs112SummaryEdnListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryEdnListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryEdnListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryEdnListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAs112SummaryEdnListResponseResultMetaDateRange]
type radarAs112SummaryEdnListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryEdnListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryEdnListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarAs112SummaryEdnListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryEdnListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAs112SummaryEdnListResponseResultMetaConfidenceInfo]
type radarAs112SummaryEdnListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryEdnListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryEdnListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryEdnListResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAs112SummaryEdnListResponseResultMetaConfidenceInfoAnnotation]
type radarAs112SummaryEdnListResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarAs112SummaryEdnListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnListResponseResultSummary0 struct {
	NotSupported string                                             `json:"NOT_SUPPORTED,required"`
	Supported    string                                             `json:"SUPPORTED,required"`
	JSON         radarAs112SummaryEdnListResponseResultSummary0JSON `json:"-"`
}

// radarAs112SummaryEdnListResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarAs112SummaryEdnListResponseResultSummary0]
type radarAs112SummaryEdnListResponseResultSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112SummaryEdnListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112SummaryEdnListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112SummaryEdnListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112SummaryEdnListParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryEdnListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112SummaryEdnListParamsDateRange string

const (
	RadarAs112SummaryEdnListParamsDateRange1d         RadarAs112SummaryEdnListParamsDateRange = "1d"
	RadarAs112SummaryEdnListParamsDateRange2d         RadarAs112SummaryEdnListParamsDateRange = "2d"
	RadarAs112SummaryEdnListParamsDateRange7d         RadarAs112SummaryEdnListParamsDateRange = "7d"
	RadarAs112SummaryEdnListParamsDateRange14d        RadarAs112SummaryEdnListParamsDateRange = "14d"
	RadarAs112SummaryEdnListParamsDateRange28d        RadarAs112SummaryEdnListParamsDateRange = "28d"
	RadarAs112SummaryEdnListParamsDateRange12w        RadarAs112SummaryEdnListParamsDateRange = "12w"
	RadarAs112SummaryEdnListParamsDateRange24w        RadarAs112SummaryEdnListParamsDateRange = "24w"
	RadarAs112SummaryEdnListParamsDateRange52w        RadarAs112SummaryEdnListParamsDateRange = "52w"
	RadarAs112SummaryEdnListParamsDateRange1dControl  RadarAs112SummaryEdnListParamsDateRange = "1dControl"
	RadarAs112SummaryEdnListParamsDateRange2dControl  RadarAs112SummaryEdnListParamsDateRange = "2dControl"
	RadarAs112SummaryEdnListParamsDateRange7dControl  RadarAs112SummaryEdnListParamsDateRange = "7dControl"
	RadarAs112SummaryEdnListParamsDateRange14dControl RadarAs112SummaryEdnListParamsDateRange = "14dControl"
	RadarAs112SummaryEdnListParamsDateRange28dControl RadarAs112SummaryEdnListParamsDateRange = "28dControl"
	RadarAs112SummaryEdnListParamsDateRange12wControl RadarAs112SummaryEdnListParamsDateRange = "12wControl"
	RadarAs112SummaryEdnListParamsDateRange24wControl RadarAs112SummaryEdnListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112SummaryEdnListParamsFormat string

const (
	RadarAs112SummaryEdnListParamsFormatJson RadarAs112SummaryEdnListParamsFormat = "JSON"
	RadarAs112SummaryEdnListParamsFormatCsv  RadarAs112SummaryEdnListParamsFormat = "CSV"
)
