// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package miruplatform

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/mirurobotics/go-platform-sdk/internal/apijson"
	"github.com/mirurobotics/go-platform-sdk/internal/requestconfig"
	"github.com/mirurobotics/go-platform-sdk/option"
	"github.com/mirurobotics/go-platform-sdk/packages/respjson"
)

// ProvisioningTokenService contains methods and other services that help with
// interacting with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProvisioningTokenService] method instead.
type ProvisioningTokenService struct {
	Options []option.RequestOption
}

// NewProvisioningTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProvisioningTokenService(opts ...option.RequestOption) (r ProvisioningTokenService) {
	r = ProvisioningTokenService{}
	r.Options = opts
	return
}

// Create a new provisioning token.
func (r *ProvisioningTokenService) New(ctx context.Context, opts ...option.RequestOption) (res *ProvisioningToken, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "provisioning_tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ProvisioningToken struct {
	// The provisioning token. This value is only returned when the token is created.
	Token string `json:"token" api:"required"`
	// The expiration date and time of the token.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExpiresAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProvisioningToken) RawJSON() string { return r.JSON.raw }
func (r *ProvisioningToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
