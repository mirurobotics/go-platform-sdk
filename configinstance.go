// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package miruplatform

import (
	"context"
	"encoding/json"
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

// ConfigInstanceService contains methods and other services that help with
// interacting with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigInstanceService] method instead.
type ConfigInstanceService struct {
	Options []option.RequestOption
}

// NewConfigInstanceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigInstanceService(opts ...option.RequestOption) (r ConfigInstanceService) {
	r = ConfigInstanceService{}
	r.Options = opts
	return
}

// Create a new config instance.
func (r *ConfigInstanceService) New(ctx context.Context, params ConfigInstanceNewParams, opts ...option.RequestOption) (res *ConfigInstance, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config_instances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a config instance by ID.
func (r *ConfigInstanceService) Get(ctx context.Context, configInstanceID string, query ConfigInstanceGetParams, opts ...option.RequestOption) (res *ConfigInstance, err error) {
	opts = slices.Concat(r.Options, opts)
	if configInstanceID == "" {
		err = errors.New("missing required config_instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("config_instances/%s", configInstanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List config instances.
func (r *ConfigInstanceService) List(ctx context.Context, query ConfigInstanceListParams, opts ...option.RequestOption) (res *ConfigInstanceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config_instances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Download the content of a config instance.
func (r *ConfigInstanceService) Download(ctx context.Context, configInstanceID string, query ConfigInstanceDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if configInstanceID == "" {
		err = errors.New("missing required config_instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("config_instances/%s/content", configInstanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ConfigInstance struct {
	// ID of the config instance.
	ID string `json:"id" api:"required"`
	// ID of the config schema which the config instance must adhere to.
	ConfigSchemaID string `json:"config_schema_id" api:"required"`
	// ID of the config type which the config instance (and its schema) is a part of.
	ConfigTypeID string `json:"config_type_id" api:"required"`
	// The name of the config type.
	ConfigTypeName string `json:"config_type_name" api:"required"`
	// The timestamp of when the config instance was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The absolute file system path where this config instance is written.
	Filepath string `json:"filepath" api:"required"`
	// The object type, which is always `config_instance`.
	Object constant.ConfigInstance `json:"object" default:"config_instance"`
	// The key of the config schema slot this instance is bound to. Must match the
	// `key` of one of the slots declared by this instance's config schema.
	SlotKey string `json:"slot_key" api:"required"`
	// Expand the config schema using 'expand=config_schema' in the query string.
	ConfigSchema ConfigSchema `json:"config_schema"`
	// Expand the config type using 'expand=config_type' in the query string.
	ConfigType ConfigType `json:"config_type"`
	// The configuration values associated with the config instance. Expand the content
	// using 'expand=content' in the query string.
	Content InstanceContent `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ConfigSchemaID respjson.Field
		ConfigTypeID   respjson.Field
		ConfigTypeName respjson.Field
		CreatedAt      respjson.Field
		Filepath       respjson.Field
		Object         respjson.Field
		SlotKey        respjson.Field
		ConfigSchema   respjson.Field
		ConfigType     respjson.Field
		Content        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigInstance) RawJSON() string { return r.JSON.raw }
func (r *ConfigInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceContent struct {
	// The configuration values associated with the config instance.
	Data string `json:"data" api:"required"`
	// The on-disk format used when a config instance is written to the device
	// filesystem.
	//
	// - `json`: standard JSON.
	// - `yaml`: YAML 1.2.
	// - `jsonc`: JSON with comments (JSON plus `//` and `/* */` comment syntax).
	// - `xml`: XML.
	// - `text`: plain, unstructured text with no specific format.
	//
	// Any of "json", "yaml", "jsonc", "xml", "text".
	Format InstanceContentFormat `json:"format" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceContent) RawJSON() string { return r.JSON.raw }
func (r *InstanceContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InstanceContent to a InstanceContentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InstanceContentParam.Overrides()
func (r InstanceContent) ToParam() InstanceContentParam {
	return param.Override[InstanceContentParam](json.RawMessage(r.RawJSON()))
}

// The on-disk format used when a config instance is written to the device
// filesystem.
//
// - `json`: standard JSON.
// - `yaml`: YAML 1.2.
// - `jsonc`: JSON with comments (JSON plus `//` and `/* */` comment syntax).
// - `xml`: XML.
// - `text`: plain, unstructured text with no specific format.
type InstanceContentFormat string

const (
	InstanceContentFormatJson  InstanceContentFormat = "json"
	InstanceContentFormatYaml  InstanceContentFormat = "yaml"
	InstanceContentFormatJsonc InstanceContentFormat = "jsonc"
	InstanceContentFormatXml   InstanceContentFormat = "xml"
	InstanceContentFormatText  InstanceContentFormat = "text"
)

// The properties Data, Format are required.
type InstanceContentParam struct {
	// The configuration values associated with the config instance.
	Data string `json:"data" api:"required"`
	// The on-disk format used when a config instance is written to the device
	// filesystem.
	//
	// - `json`: standard JSON.
	// - `yaml`: YAML 1.2.
	// - `jsonc`: JSON with comments (JSON plus `//` and `/* */` comment syntax).
	// - `xml`: XML.
	// - `text`: plain, unstructured text with no specific format.
	//
	// Any of "json", "yaml", "jsonc", "xml", "text".
	Format InstanceContentFormat `json:"format,omitzero" api:"required"`
	paramObj
}

func (r InstanceContentParam) MarshalJSON() (data []byte, err error) {
	type shadow InstanceContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigInstanceListResponse struct {
	// The list of config instances.
	Data []ConfigInstance `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedList
}

// Returns the unmodified JSON received from the API
func (r ConfigInstanceListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConfigInstanceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigInstanceNewParams struct {
	// The ID of the config schema which this config instance must adhere to.
	ConfigSchemaID string               `json:"config_schema_id" api:"required"`
	Content        InstanceContentParam `json:"content,omitzero" api:"required"`
	// The slot of the config schema this instance is bound to. Optional when the
	// schema has exactly one slot. Must match the `key` of one of the slots declared
	// by the config schema.
	SlotKey param.Opt[string] `json:"slot_key,omitzero"`
	// Fields to expand on the config instance resource.
	//
	// Any of "content", "config_schema", "config_type".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

func (r ConfigInstanceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConfigInstanceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigInstanceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ConfigInstanceNewParams]'s query parameters as
// `url.Values`.
func (r ConfigInstanceNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigInstanceGetParams struct {
	// Fields to expand on the config instance resource.
	//
	// Any of "content", "config_schema", "config_type".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConfigInstanceGetParams]'s query parameters as
// `url.Values`.
func (r ConfigInstanceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigInstanceListParams struct {
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The config instance IDs to filter by.
	ID []string `query:"id,omitzero" json:"-"`
	// The config schema IDs to filter by.
	ConfigSchemaID []string `query:"config_schema_id,omitzero" json:"-"`
	// The config type IDs to filter by.
	ConfigTypeID []string `query:"config_type_id,omitzero" json:"-"`
	// Fields to expand on each config instance in the list.
	//
	// Any of "total_count", "content", "config_schema", "config_type".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Sort order for the config instance results.
	//
	// Any of "id:asc", "id:desc", "created_at:desc", "created_at:asc".
	OrderBy []string `query:"order_by,omitzero" json:"-"`
	// The config schema slot keys to filter by.
	SlotKey []string `query:"slot_key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConfigInstanceListParams]'s query parameters as
// `url.Values`.
func (r ConfigInstanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigInstanceDownloadParams struct {
	// Controls the Content-Disposition behavior. Use `inline` to display in browser or
	// `attachment` to trigger a file download.
	//
	// Any of "inline", "attachment".
	Disposition ConfigInstanceDownloadParamsDisposition `query:"disposition,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConfigInstanceDownloadParams]'s query parameters as
// `url.Values`.
func (r ConfigInstanceDownloadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Controls the Content-Disposition behavior. Use `inline` to display in browser or
// `attachment` to trigger a file download.
type ConfigInstanceDownloadParamsDisposition string

const (
	ConfigInstanceDownloadParamsDispositionInline     ConfigInstanceDownloadParamsDisposition = "inline"
	ConfigInstanceDownloadParamsDispositionAttachment ConfigInstanceDownloadParamsDisposition = "attachment"
)
