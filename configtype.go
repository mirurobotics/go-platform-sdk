// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package miruplatform

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/mirurobotics/go-platform-sdk/internal/apijson"
	"github.com/mirurobotics/go-platform-sdk/internal/apiquery"
	"github.com/mirurobotics/go-platform-sdk/internal/requestconfig"
	"github.com/mirurobotics/go-platform-sdk/option"
	"github.com/mirurobotics/go-platform-sdk/packages/param"
	"github.com/mirurobotics/go-platform-sdk/packages/respjson"
	"github.com/mirurobotics/go-platform-sdk/shared"
	"github.com/mirurobotics/go-platform-sdk/shared/constant"
)

// ConfigTypeService contains methods and other services that help with interacting
// with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigTypeService] method instead.
type ConfigTypeService struct {
	Options []option.RequestOption
}

// NewConfigTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigTypeService(opts ...option.RequestOption) (r ConfigTypeService) {
	r = ConfigTypeService{}
	r.Options = opts
	return
}

// Create a new config type.
func (r *ConfigTypeService) New(ctx context.Context, body ConfigTypeNewParams, opts ...option.RequestOption) (res *ConfigType, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config_types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a config type by ID.
func (r *ConfigTypeService) Get(ctx context.Context, configTypeID string, opts ...option.RequestOption) (res *ConfigType, err error) {
	opts = slices.Concat(r.Options, opts)
	if configTypeID == "" {
		err = errors.New("missing required config_type_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("config_types/%s", configTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a config type by ID.
func (r *ConfigTypeService) Update(ctx context.Context, configTypeID string, body ConfigTypeUpdateParams, opts ...option.RequestOption) (res *ConfigType, err error) {
	opts = slices.Concat(r.Options, opts)
	if configTypeID == "" {
		err = errors.New("missing required config_type_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("config_types/%s", configTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List config types.
func (r *ConfigTypeService) List(ctx context.Context, query ConfigTypeListParams, opts ...option.RequestOption) (res *ConfigTypeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config_types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ConfigType struct {
	// ID of the config type.
	ID string `json:"id" api:"required"`
	// Timestamp of when the config type was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name of the config type.
	Name string `json:"name" api:"required"`
	// The object type, which is always `config_type`.
	Object constant.ConfigType `json:"object" default:"config_type"`
	// An immutable, code-friendly name for the config type.
	Slug string `json:"slug" api:"required"`
	// Timestamp of when the config type was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Slug        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigType) RawJSON() string { return r.JSON.raw }
func (r *ConfigType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigTypeListResponse struct {
	// The list of config types.
	Data []ConfigType `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedList
}

// Returns the unmodified JSON received from the API
func (r ConfigTypeListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConfigTypeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigTypeNewParams struct {
	// The name of the config type.
	Name string `json:"name" api:"required"`
	// An immutable, code-friendly name for the config type.
	Slug string `json:"slug" api:"required"`
	paramObj
}

func (r ConfigTypeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConfigTypeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigTypeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigTypeUpdateParams struct {
	// The updated name of the config type.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ConfigTypeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConfigTypeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigTypeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigTypeListParams struct {
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The config type IDs to filter by.
	ID []string `query:"id,omitzero" json:"-"`
	// Fields to expand on each config type in the list.
	//
	// Any of "total_count".
	Expand []string `query:"expand,omitzero" json:"-"`
	// The config type names to filter by.
	Name []string `query:"name,omitzero" json:"-"`
	// Sort order for the config type results.
	//
	// Any of "id:asc", "id:desc", "created_at:desc", "created_at:asc".
	OrderBy []string `query:"order_by,omitzero" json:"-"`
	// The config type slugs to filter by.
	Slug []string `query:"slug,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConfigTypeListParams]'s query parameters as `url.Values`.
func (r ConfigTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
