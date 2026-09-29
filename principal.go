// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package miruplatform

import (
	"context"
	"net/http"
	"slices"

	"github.com/mirurobotics/go-platform-sdk/internal/apijson"
	"github.com/mirurobotics/go-platform-sdk/internal/requestconfig"
	"github.com/mirurobotics/go-platform-sdk/option"
	"github.com/mirurobotics/go-platform-sdk/packages/respjson"
	"github.com/mirurobotics/go-platform-sdk/shared/constant"
)

// PrincipalService contains methods and other services that help with interacting
// with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrincipalService] method instead.
type PrincipalService struct {
	Options []option.RequestOption
}

// NewPrincipalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrincipalService(opts ...option.RequestOption) (r PrincipalService) {
	r = PrincipalService{}
	r.Options = opts
	return
}

// Retrieve the currently authenticated principal.
func (r *PrincipalService) Self(ctx context.Context, opts ...option.RequestOption) (res *Principal, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "principal"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Principal struct {
	// ID of the principal.
	ID string `json:"id" api:"required"`
	// The display name of the principal.
	Name string `json:"name" api:"required"`
	// The object type, which is always `principal`.
	Object constant.Principal `json:"object" default:"principal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Principal) RawJSON() string { return r.JSON.raw }
func (r *Principal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
