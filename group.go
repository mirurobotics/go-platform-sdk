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

// GroupService contains methods and other services that help with interacting with
// the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGroupService] method instead.
type GroupService struct {
	Options []option.RequestOption
}

// NewGroupService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGroupService(opts ...option.RequestOption) (r GroupService) {
	r = GroupService{}
	r.Options = opts
	return
}

// Create a new group.
func (r *GroupService) New(ctx context.Context, params GroupNewParams, opts ...option.RequestOption) (res *Group, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a group by ID.
func (r *GroupService) Get(ctx context.Context, groupID string, query GroupGetParams, opts ...option.RequestOption) (res *Group, err error) {
	opts = slices.Concat(r.Options, opts)
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("groups/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update a group.
func (r *GroupService) Update(ctx context.Context, groupID string, params GroupUpdateParams, opts ...option.RequestOption) (res *Group, err error) {
	opts = slices.Concat(r.Options, opts)
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("groups/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List groups.
func (r *GroupService) List(ctx context.Context, query GroupListParams, opts ...option.RequestOption) (res *GroupList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Move a group to a new parent, or to the top level by setting parent_id to null.
func (r *GroupService) Move(ctx context.Context, groupID string, params GroupMoveParams, opts ...option.RequestOption) (res *Group, err error) {
	opts = slices.Concat(r.Options, opts)
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("groups/%s/move", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type Group struct {
	// ID of the group.
	ID string `json:"id" api:"required"`
	// Timestamp of when the group was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name of the group.
	Name string `json:"name" api:"required"`
	// The object type, which is always `group`.
	Object constant.Group `json:"object" default:"group"`
	// ID of the parent group. Null for top-level groups.
	ParentID string `json:"parent_id" api:"required"`
	// Timestamp of when the group was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The parent group. Null for top-level groups. Expand using `expand=parent` in the
	// query string.
	Parent *Group `json:"parent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		ParentID    respjson.Field
		UpdatedAt   respjson.Field
		Parent      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Group) RawJSON() string { return r.JSON.raw }
func (r *Group) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupList struct {
	// The list of groups.
	Data []Group `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedList
}

// Returns the unmodified JSON received from the API
func (r GroupList) RawJSON() string { return r.JSON.raw }
func (r *GroupList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupNewParams struct {
	// The name of the group.
	Name string `json:"name" api:"required"`
	// ID of the parent group. Omit or set to null to create a top-level group.
	ParentID param.Opt[string] `json:"parent_id,omitzero"`
	// Fields to expand on the group resource.
	//
	// Any of "parent".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

func (r GroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [GroupNewParams]'s query parameters as `url.Values`.
func (r GroupNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GroupGetParams struct {
	// Fields to expand on the group resource.
	//
	// Any of "parent".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GroupGetParams]'s query parameters as `url.Values`.
func (r GroupGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GroupUpdateParams struct {
	// The updated name of the group.
	Name param.Opt[string] `json:"name,omitzero"`
	// Fields to expand on the group resource.
	//
	// Any of "parent".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

func (r GroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow GroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [GroupUpdateParams]'s query parameters as `url.Values`.
func (r GroupUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GroupListParams struct {
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The group IDs to filter by.
	ID []string `query:"id,omitzero" json:"-"`
	// Fields to expand on each group in the list.
	//
	// Any of "total_count", "parent".
	Expand []string `query:"expand,omitzero" json:"-"`
	// The group names to filter by.
	Name []string `query:"name,omitzero" json:"-"`
	// Sort order for the group results.
	//
	// Any of "id:asc", "id:desc", "created_at:asc", "created_at:desc",
	// "updated_at:asc", "updated_at:desc", "name:asc", "name:desc".
	OrderBy []string `query:"order_by,omitzero" json:"-"`
	// The parent group IDs to filter by.
	ParentID []string `query:"parent_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GroupListParams]'s query parameters as `url.Values`.
func (r GroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GroupMoveParams struct {
	// ID of the new parent group. Set to null to make this a top-level group.
	ParentID param.Opt[string] `json:"parent_id,omitzero" api:"required"`
	// Fields to expand on the group resource.
	//
	// Any of "parent".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

func (r GroupMoveParams) MarshalJSON() (data []byte, err error) {
	type shadow GroupMoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroupMoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [GroupMoveParams]'s query parameters as `url.Values`.
func (r GroupMoveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
