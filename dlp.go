// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DlpService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewDlpService] method instead.
type DlpService struct {
	Options  []option.RequestOption
	Datasets *DlpDatasetService
}

// NewDlpService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDlpService(opts ...option.RequestOption) (r *DlpService) {
	r = &DlpService{}
	r.Options = opts
	r.Datasets = NewDlpDatasetService(opts...)
	return
}
