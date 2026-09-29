// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/mirurobotics/go-platform-sdk/internal/apijson"
	"github.com/mirurobotics/go-platform-sdk/packages/param"
	"github.com/mirurobotics/go-platform-sdk/packages/respjson"
	"github.com/mirurobotics/go-platform-sdk/shared/constant"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type PaginatedList struct {
	// True if there are more items in the list to return. False if there are no more
	// items to return.
	HasMore bool `json:"has_more" api:"required"`
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit int64 `json:"limit" api:"required"`
	// The object type, which is always `list`.
	Object constant.List `json:"object" default:"list"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset int64 `json:"offset" api:"required"`
	// The total number of items in the list. By default the total count is not
	// returned. The total count must be expanded (using expand=total_count) to get the
	// total number of items in the list.
	TotalCount int64 `json:"total_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Limit       respjson.Field
		Object      respjson.Field
		Offset      respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginatedList) RawJSON() string { return r.JSON.raw }
func (r *PaginatedList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
