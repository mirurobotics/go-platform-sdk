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

// ConfigSchemaService contains methods and other services that help with
// interacting with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigSchemaService] method instead.
type ConfigSchemaService struct {
	Options []option.RequestOption
}

// NewConfigSchemaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigSchemaService(opts ...option.RequestOption) (r ConfigSchemaService) {
	r = ConfigSchemaService{}
	r.Options = opts
	return
}

// Create a new config schema.
func (r *ConfigSchemaService) New(ctx context.Context, params ConfigSchemaNewParams, opts ...option.RequestOption) (res *ConfigSchema, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config_schemas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a config schema by ID.
func (r *ConfigSchemaService) Get(ctx context.Context, configSchemaID string, query ConfigSchemaGetParams, opts ...option.RequestOption) (res *ConfigSchema, err error) {
	opts = slices.Concat(r.Options, opts)
	if configSchemaID == "" {
		err = errors.New("missing required config_schema_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("config_schemas/%s", configSchemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List config schemas.
func (r *ConfigSchemaService) List(ctx context.Context, query ConfigSchemaListParams, opts ...option.RequestOption) (res *ConfigSchemaList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config_schemas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ConfigSchema struct {
	// ID of the config schema.
	ID string `json:"id" api:"required"`
	// ID of the config type.
	ConfigTypeID string `json:"config_type_id" api:"required"`
	// The name of the config type.
	ConfigTypeName string `json:"config_type_name" api:"required"`
	// Timestamp of when the config schema was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The digest of the config schema.
	Digest string `json:"digest" api:"required"`
	// Any of "json", "yaml", "cue".
	Format ConfigSchemaFormat `json:"format" api:"required"`
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
	InstanceFormat ConfigSchemaInstanceFormat `json:"instance_format" api:"required"`
	// The file system destinations this config schema writes to. Every config schema
	// has at least one slot. Slots share the schema's validation and differ only in
	// where the file is written; a file needing different validation is a different
	// config type, not a slot. Slot keys and filepaths must each be unique within the
	// schema, and slot filepaths must be unique across every config schema in a
	// release.
	InstanceSlots []InstanceSlot `json:"instance_slots" api:"required"`
	// Any of "jsonschema", "cue", "opaque".
	Language SchemaLanguage `json:"language" api:"required"`
	// The object type, which is always `config_schema`.
	Object constant.ConfigSchema `json:"object" default:"config_schema"`
	// An operating system family, from the agent's build-time OS vocabulary. Shared
	// across resources (devices, config schemas, file rules, releases). This is the
	// machine-readable OS kind, distinct from a human-readable version string.
	//
	// - `linux`
	// - `windows`
	//
	// Any of "linux", "windows".
	Os ConfigSchemaOs `json:"os" api:"required"`
	// Timestamp of when the config schema was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Expand the config type using 'expand=config_type' in the query string.
	ConfigType ConfigType `json:"config_type"`
	// Expand the config schema documents using `expand=documents` in the query string.
	Documents []SchemaDocument `json:"documents"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ConfigTypeID   respjson.Field
		ConfigTypeName respjson.Field
		CreatedAt      respjson.Field
		Digest         respjson.Field
		Format         respjson.Field
		InstanceFormat respjson.Field
		InstanceSlots  respjson.Field
		Language       respjson.Field
		Object         respjson.Field
		Os             respjson.Field
		UpdatedAt      respjson.Field
		ConfigType     respjson.Field
		Documents      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigSchema) RawJSON() string { return r.JSON.raw }
func (r *ConfigSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigSchemaFormat string

const (
	ConfigSchemaFormatJson ConfigSchemaFormat = "json"
	ConfigSchemaFormatYaml ConfigSchemaFormat = "yaml"
	ConfigSchemaFormatCue  ConfigSchemaFormat = "cue"
)

// The on-disk format used when a config instance is written to the device
// filesystem.
//
// - `json`: standard JSON.
// - `yaml`: YAML 1.2.
// - `jsonc`: JSON with comments (JSON plus `//` and `/* */` comment syntax).
// - `xml`: XML.
// - `text`: plain, unstructured text with no specific format.
type ConfigSchemaInstanceFormat string

const (
	ConfigSchemaInstanceFormatJson  ConfigSchemaInstanceFormat = "json"
	ConfigSchemaInstanceFormatYaml  ConfigSchemaInstanceFormat = "yaml"
	ConfigSchemaInstanceFormatJsonc ConfigSchemaInstanceFormat = "jsonc"
	ConfigSchemaInstanceFormatXml   ConfigSchemaInstanceFormat = "xml"
	ConfigSchemaInstanceFormatText  ConfigSchemaInstanceFormat = "text"
)

// An operating system family, from the agent's build-time OS vocabulary. Shared
// across resources (devices, config schemas, file rules, releases). This is the
// machine-readable OS kind, distinct from a human-readable version string.
//
// - `linux`
// - `windows`
type ConfigSchemaOs string

const (
	ConfigSchemaOsLinux   ConfigSchemaOs = "linux"
	ConfigSchemaOsWindows ConfigSchemaOs = "windows"
)

type ConfigSchemaList struct {
	// The list of config schemas.
	Data []ConfigSchema `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedList
}

// Returns the unmodified JSON received from the API
func (r ConfigSchemaList) RawJSON() string { return r.JSON.raw }
func (r *ConfigSchemaList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceSlot struct {
	// The absolute file system path where instances bound to this slot are written.
	// Must be unique within the schema, and unique across every config schema in a
	// release.
	Filepath string `json:"filepath" api:"required"`
	// An immutable, code-friendly identifier for this slot, unique within the schema.
	// Lowercase alphanumerics, underscores, and hyphens; must start with an
	// alphanumeric; at most 128 characters.
	Key string `json:"key" api:"required"`
	// The human-readable name of this slot.
	Name string `json:"name" api:"required"`
	// Whether every deployment of a release containing this schema must include an
	// instance for this slot. A deployment includes at most one config instance per
	// slot.
	Required bool `json:"required" api:"required"`
	// An optional human-readable description of this slot.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filepath    respjson.Field
		Key         respjson.Field
		Name        respjson.Field
		Required    respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceSlot) RawJSON() string { return r.JSON.raw }
func (r *InstanceSlot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InstanceSlot to a InstanceSlotParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InstanceSlotParam.Overrides()
func (r InstanceSlot) ToParam() InstanceSlotParam {
	return param.Override[InstanceSlotParam](json.RawMessage(r.RawJSON()))
}

// The properties Filepath, Key, Name, Required are required.
type InstanceSlotParam struct {
	// The absolute file system path where instances bound to this slot are written.
	// Must be unique within the schema, and unique across every config schema in a
	// release.
	Filepath string `json:"filepath" api:"required"`
	// An immutable, code-friendly identifier for this slot, unique within the schema.
	// Lowercase alphanumerics, underscores, and hyphens; must start with an
	// alphanumeric; at most 128 characters.
	Key string `json:"key" api:"required"`
	// The human-readable name of this slot.
	Name string `json:"name" api:"required"`
	// Whether every deployment of a release containing this schema must include an
	// instance for this slot. A deployment includes at most one config instance per
	// slot.
	Required bool `json:"required" api:"required"`
	// An optional human-readable description of this slot.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r InstanceSlotParam) MarshalJSON() (data []byte, err error) {
	type shadow InstanceSlotParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceSlotParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SchemaDocument struct {
	// The unique identifier for this document.
	ID string `json:"id" api:"required"`
	// The raw document content.
	Data string `json:"data" api:"required"`
	// The document filename.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Data        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchemaDocument) RawJSON() string { return r.JSON.raw }
func (r *SchemaDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SchemaLanguage string

const (
	SchemaLanguageJsonschema SchemaLanguage = "jsonschema"
	SchemaLanguageCue        SchemaLanguage = "cue"
	SchemaLanguageOpaque     SchemaLanguage = "opaque"
)

type ConfigSchemaNewParams struct {
	// A reference to a config type. At least one of `id` or `slug` must be provided.
	// When both are provided, `id` takes precedence and `slug` is ignored.
	ConfigTypeRef ConfigSchemaNewParamsConfigTypeRef `json:"config_type_ref,omitzero" api:"required"`
	// The schema document files.
	Documents []ConfigSchemaNewParamsDocument `json:"documents,omitzero" api:"required"`
	// Any of "json", "yaml", "cue".
	Format ConfigSchemaNewParamsFormat `json:"format,omitzero" api:"required"`
	// Any of "jsonschema", "cue", "opaque".
	Language SchemaLanguage `json:"language,omitzero" api:"required"`
	// Fields to expand on the config schema resource.
	//
	// Any of "documents", "config_type".
	Expand []string `query:"expand,omitzero" json:"-"`
	// The git commit to link to this config schema.
	GitCommit ConfigSchemaNewParamsGitCommit `json:"git_commit,omitzero"`
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
	InstanceFormat ConfigSchemaNewParamsInstanceFormat `json:"instance_format,omitzero"`
	// The file system destinations this config schema writes to. Slot keys and
	// filepaths must each be unique within the schema, and slot filepaths must be
	// unique across every config schema in a release. Optional: when omitted, the
	// server creates a single required slot with key `default`, name `Default`, and a
	// filepath of `/srv/miru/configs/<config_type_slug>.<json|yaml>` derived from the
	// config type slug and the schema format. When provided, the array must contain at
	// least one slot.
	InstanceSlots []InstanceSlotParam `json:"instance_slots,omitzero"`
	paramObj
}

func (r ConfigSchemaNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConfigSchemaNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigSchemaNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ConfigSchemaNewParams]'s query parameters as `url.Values`.
func (r ConfigSchemaNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A reference to a config type. At least one of `id` or `slug` must be provided.
// When both are provided, `id` takes precedence and `slug` is ignored.
type ConfigSchemaNewParamsConfigTypeRef struct {
	// ID of the config type. Takes precedence over `slug` when both are provided.
	ID param.Opt[string] `json:"id,omitzero"`
	// An immutable, code-friendly name for the config type. Used only when `id` is not
	// provided.
	Slug param.Opt[string] `json:"slug,omitzero"`
	paramObj
}

func (r ConfigSchemaNewParamsConfigTypeRef) MarshalJSON() (data []byte, err error) {
	type shadow ConfigSchemaNewParamsConfigTypeRef
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigSchemaNewParamsConfigTypeRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, Name are required.
type ConfigSchemaNewParamsDocument struct {
	// The raw document content.
	Data string `json:"data" api:"required"`
	// The name of this document.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r ConfigSchemaNewParamsDocument) MarshalJSON() (data []byte, err error) {
	type shadow ConfigSchemaNewParamsDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigSchemaNewParamsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigSchemaNewParamsFormat string

const (
	ConfigSchemaNewParamsFormatJson ConfigSchemaNewParamsFormat = "json"
	ConfigSchemaNewParamsFormatYaml ConfigSchemaNewParamsFormat = "yaml"
	ConfigSchemaNewParamsFormatCue  ConfigSchemaNewParamsFormat = "cue"
)

// The git commit to link to this config schema.
//
// The properties CommitRef, SchemaFilepaths are required.
type ConfigSchemaNewParamsGitCommit struct {
	// A reference to a git commit. At least one of `id` or `sha` must be provided.
	// When both are provided, `id` takes precedence and `sha` is ignored.
	CommitRef GitCommitRefParam `json:"commit_ref,omitzero" api:"required"`
	// The file paths of the config schema files relative to the git repository root.
	// JSON Schema only supports a single file for a schema. CUE support multiple files
	// via [packages](https://cuelang.org/docs/references/spec/#packages).
	SchemaFilepaths []string `json:"schema_filepaths,omitzero" api:"required"`
	paramObj
}

func (r ConfigSchemaNewParamsGitCommit) MarshalJSON() (data []byte, err error) {
	type shadow ConfigSchemaNewParamsGitCommit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigSchemaNewParamsGitCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The on-disk format used when a config instance is written to the device
// filesystem.
//
// - `json`: standard JSON.
// - `yaml`: YAML 1.2.
// - `jsonc`: JSON with comments (JSON plus `//` and `/* */` comment syntax).
// - `xml`: XML.
// - `text`: plain, unstructured text with no specific format.
type ConfigSchemaNewParamsInstanceFormat string

const (
	ConfigSchemaNewParamsInstanceFormatJson  ConfigSchemaNewParamsInstanceFormat = "json"
	ConfigSchemaNewParamsInstanceFormatYaml  ConfigSchemaNewParamsInstanceFormat = "yaml"
	ConfigSchemaNewParamsInstanceFormatJsonc ConfigSchemaNewParamsInstanceFormat = "jsonc"
	ConfigSchemaNewParamsInstanceFormatXml   ConfigSchemaNewParamsInstanceFormat = "xml"
	ConfigSchemaNewParamsInstanceFormatText  ConfigSchemaNewParamsInstanceFormat = "text"
)

type ConfigSchemaGetParams struct {
	// Fields to expand on the config schema resource.
	//
	// Any of "documents", "config_type".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConfigSchemaGetParams]'s query parameters as `url.Values`.
func (r ConfigSchemaGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigSchemaListParams struct {
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The config schema IDs to filter by.
	ID []string `query:"id,omitzero" json:"-"`
	// The config type IDs to filter by.
	ConfigTypeID []string `query:"config_type_id,omitzero" json:"-"`
	// Fields to expand on each config schema in the list.
	//
	// Any of "total_count", "documents", "config_type".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Sort order for the config schema results.
	//
	// Any of "id:asc", "id:desc", "created_at:desc", "created_at:asc".
	OrderBy []string `query:"order_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConfigSchemaListParams]'s query parameters as `url.Values`.
func (r ConfigSchemaListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
