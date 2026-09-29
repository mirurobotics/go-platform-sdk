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

// DeviceService contains methods and other services that help with interacting
// with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceService] method instead.
type DeviceService struct {
	Options []option.RequestOption
}

// NewDeviceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDeviceService(opts ...option.RequestOption) (r DeviceService) {
	r = DeviceService{}
	r.Options = opts
	return
}

// Create a new device.
func (r *DeviceService) New(ctx context.Context, params DeviceNewParams, opts ...option.RequestOption) (res *Device, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a device by ID.
func (r *DeviceService) Get(ctx context.Context, deviceID string, query DeviceGetParams, opts ...option.RequestOption) (res *Device, err error) {
	opts = slices.Concat(r.Options, opts)
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s", deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update a device by ID.
func (r *DeviceService) Update(ctx context.Context, deviceID string, params DeviceUpdateParams, opts ...option.RequestOption) (res *Device, err error) {
	opts = slices.Concat(r.Options, opts)
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s", deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List devices.
func (r *DeviceService) List(ctx context.Context, query DeviceListParams, opts ...option.RequestOption) (res *DeviceList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Ping a device to check connectivity.
func (r *DeviceService) Ping(ctx context.Context, deviceID string, body DevicePingParams, opts ...option.RequestOption) (res *DevicePingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/ping", deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type Device struct {
	// ID of the device.
	ID string `json:"id" api:"required"`
	// The version of the agent the device is running.
	AgentVersion string `json:"agent_version" api:"required"`
	// A CPU architecture, from the agent's build-time architecture vocabulary. Shared
	// across resources (devices, config schemas, file rules). These are the canonical
	// wire values.
	//
	// - `x86_64`
	// - `aarch64`
	//
	// Any of "x86_64", "aarch64".
	Arch DeviceArch `json:"arch" api:"required"`
	// Timestamp of when the device was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A free-form description of the device.
	Description string `json:"description" api:"required"`
	// ID of the group this device belongs to. Null if the device is unassigned.
	GroupID string `json:"group_id" api:"required"`
	// The hostname of the device as reported by the agent.
	Hostname string `json:"hostname" api:"required"`
	// The operating system kernel version string reported by the agent.
	KernelVersion string `json:"kernel_version" api:"required"`
	// Timestamp of when the device was last made an initial connection (this is not
	// the same as the last time the device was seen).
	LastConnectedAt time.Time `json:"last_connected_at" api:"required" format:"date-time"`
	// Timestamp of when the device was last disconnected (this is not the same as the
	// last time the device was seen).
	LastDisconnectedAt time.Time `json:"last_disconnected_at" api:"required" format:"date-time"`
	// Name of the device.
	Name string `json:"name" api:"required"`
	// The object type, which is always `device`.
	Object constant.Device `json:"object" default:"device"`
	// An operating system family, from the agent's build-time OS vocabulary. Shared
	// across resources (devices, config schemas, file rules, releases). This is the
	// machine-readable OS kind, distinct from a human-readable version string.
	//
	// - `linux`
	// - `windows`
	//
	// Any of "linux", "windows".
	Os DeviceOs `json:"os" api:"required"`
	// The human-readable long operating system version string reported by the agent
	// (for example `Ubuntu 22.04` or `Windows 11 Pro`). Distinct from the
	// machine-readable `os` family.
	OsVersion string `json:"os_version" api:"required"`
	// The status of the device.
	//
	//   - Inactive: The miru agent has not yet been installed / authenticated
	//   - Activating: The miru agent is currently being installed / authenticated
	//     (should only last for a few seconds)
	//   - Online: The miru agent has successfully pinged the server within the last 60
	//     seconds.
	//   - Offline: The miru agent has not successfully pinged the server within the last
	//     60 seconds (e.g. network issues, device is powered off, etc.)
	//   - Archived: The device has been archived. The agent's session has been revoked
	//     and the device no longer participates in workspace operations. See
	//     `archived_at` for when this occurred.
	//
	// Any of "inactive", "activating", "online", "offline", "archived".
	Status DeviceStatus `json:"status" api:"required"`
	// Timestamp of when the device was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The current deployment for the device.
	CurrentDeployment *Deployment `json:"current_deployment" api:"nullable"`
	// The current release for the device.
	CurrentRelease Release `json:"current_release" api:"nullable"`
	// The group this device belongs to. Null if the device is unassigned.
	Group Group `json:"group" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AgentVersion       respjson.Field
		Arch               respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		GroupID            respjson.Field
		Hostname           respjson.Field
		KernelVersion      respjson.Field
		LastConnectedAt    respjson.Field
		LastDisconnectedAt respjson.Field
		Name               respjson.Field
		Object             respjson.Field
		Os                 respjson.Field
		OsVersion          respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		CurrentDeployment  respjson.Field
		CurrentRelease     respjson.Field
		Group              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Device) RawJSON() string { return r.JSON.raw }
func (r *Device) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A CPU architecture, from the agent's build-time architecture vocabulary. Shared
// across resources (devices, config schemas, file rules). These are the canonical
// wire values.
//
// - `x86_64`
// - `aarch64`
type DeviceArch string

const (
	DeviceArchX86_64  DeviceArch = "x86_64"
	DeviceArchAarch64 DeviceArch = "aarch64"
)

// An operating system family, from the agent's build-time OS vocabulary. Shared
// across resources (devices, config schemas, file rules, releases). This is the
// machine-readable OS kind, distinct from a human-readable version string.
//
// - `linux`
// - `windows`
type DeviceOs string

const (
	DeviceOsLinux   DeviceOs = "linux"
	DeviceOsWindows DeviceOs = "windows"
)

// The status of the device.
//
//   - Inactive: The miru agent has not yet been installed / authenticated
//   - Activating: The miru agent is currently being installed / authenticated
//     (should only last for a few seconds)
//   - Online: The miru agent has successfully pinged the server within the last 60
//     seconds.
//   - Offline: The miru agent has not successfully pinged the server within the last
//     60 seconds (e.g. network issues, device is powered off, etc.)
//   - Archived: The device has been archived. The agent's session has been revoked
//     and the device no longer participates in workspace operations. See
//     `archived_at` for when this occurred.
type DeviceStatus string

const (
	DeviceStatusInactive   DeviceStatus = "inactive"
	DeviceStatusActivating DeviceStatus = "activating"
	DeviceStatusOnline     DeviceStatus = "online"
	DeviceStatusOffline    DeviceStatus = "offline"
	DeviceStatusArchived   DeviceStatus = "archived"
)

type DeviceList struct {
	// The list of devices.
	Data []Device `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedList
}

// Returns the unmodified JSON received from the API
func (r DeviceList) RawJSON() string { return r.JSON.raw }
func (r *DeviceList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePingResponse struct {
	// The result of the ping operation.
	//
	// Any of "success", "timeout".
	Code DevicePingResponseCode `json:"code" api:"required"`
	// The round trip time in nanoseconds to send the ping to the device and receive
	// the response. Only present when `code` is `success`.
	RoundTripNanos int64 `json:"round_trip_nanos"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code           respjson.Field
		RoundTripNanos respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DevicePingResponse) RawJSON() string { return r.JSON.raw }
func (r *DevicePingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result of the ping operation.
type DevicePingResponseCode string

const (
	DevicePingResponseCodeSuccess DevicePingResponseCode = "success"
	DevicePingResponseCodeTimeout DevicePingResponseCode = "timeout"
)

type DeviceNewParams struct {
	// The name of the device.
	Name string `json:"name" api:"required"`
	// Fields to expand on the device resource.
	//
	// Any of "current_deployment", "current_release", "group".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

func (r DeviceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [DeviceNewParams]'s query parameters as `url.Values`.
func (r DeviceNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceGetParams struct {
	// Fields to expand on the device resource.
	//
	// Any of "current_deployment", "current_release", "group".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceGetParams]'s query parameters as `url.Values`.
func (r DeviceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceUpdateParams struct {
	// The new name of the device. If excluded from the request, the device name is not
	// updated.
	Name param.Opt[string] `json:"name,omitzero"`
	// Fields to expand on the device resource.
	//
	// Any of "current_deployment", "current_release", "group".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

func (r DeviceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [DeviceUpdateParams]'s query parameters as `url.Values`.
func (r DeviceUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceListParams struct {
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The device IDs to filter by.
	ID []string `query:"id,omitzero" json:"-"`
	// The agent versions to filter by.
	AgentVersion []string `query:"agent_version,omitzero" json:"-"`
	// The release IDs to filter devices by their current release.
	CurrentReleaseID []string `query:"current_release_id,omitzero" json:"-"`
	// Fields to expand on each device in the list.
	//
	// Any of "total_count", "current_deployment", "current_release", "group".
	Expand []string `query:"expand,omitzero" json:"-"`
	// The group IDs to filter devices by.
	GroupID []string `query:"group_id,omitzero" json:"-"`
	// The device names to filter by.
	Name []string `query:"name,omitzero" json:"-"`
	// Sort order for the device results.
	//
	// Any of "id:asc", "id:desc", "created_at:desc", "created_at:asc", "status:asc",
	// "status:desc".
	OrderBy []string `query:"order_by,omitzero" json:"-"`
	// The device statuses to filter by.
	//
	// Any of "inactive", "activating", "online", "offline", "archived".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceListParams]'s query parameters as `url.Values`.
func (r DeviceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DevicePingParams struct {
	// The timeout in nanoseconds to wait for the ping operation to complete. The
	// maximum timeout is 10 seconds.
	TimeoutNanos int64 `json:"timeout_nanos" api:"required"`
	paramObj
}

func (r DevicePingParams) MarshalJSON() (data []byte, err error) {
	type shadow DevicePingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DevicePingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
