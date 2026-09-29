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

// DeploymentService contains methods and other services that help with interacting
// with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeploymentService] method instead.
type DeploymentService struct {
	Options []option.RequestOption
}

// NewDeploymentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeploymentService(opts ...option.RequestOption) (r DeploymentService) {
	r = DeploymentService{}
	r.Options = opts
	return
}

// Stage or deploy a new deployment.
func (r *DeploymentService) New(ctx context.Context, params DeploymentNewParams, opts ...option.RequestOption) (res *Deployment, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a deployment by its ID.
func (r *DeploymentService) Get(ctx context.Context, deploymentID string, query DeploymentGetParams, opts ...option.RequestOption) (res *Deployment, err error) {
	opts = slices.Concat(r.Options, opts)
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("deployments/%s", deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List all deployments.
func (r *DeploymentService) List(ctx context.Context, query DeploymentListParams, opts ...option.RequestOption) (res *DeploymentList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Archive staged or drifted deployments.
func (r *DeploymentService) Archive(ctx context.Context, deploymentID string, body DeploymentArchiveParams, opts ...option.RequestOption) (res *Deployment, err error) {
	opts = slices.Concat(r.Options, opts)
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("deployments/%s/archive", deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Trigger a deployment to its target device.
func (r *DeploymentService) Deploy(ctx context.Context, deploymentID string, body DeploymentDeployParams, opts ...option.RequestOption) (res *Deployment, err error) {
	opts = slices.Concat(r.Options, opts)
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("deployments/%s/deploy", deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List the deployments which have been deployed since this deployment was staged.
func (r *DeploymentService) ListDrifts(ctx context.Context, deploymentID string, query DeploymentListDriftsParams, opts ...option.RequestOption) (res *DeploymentList, err error) {
	opts = slices.Concat(r.Options, opts)
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("deployments/%s/drifts", deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Deployment struct {
	// ID of the deployment.
	ID string `json:"id" api:"required"`
	// Last known activity state of the deployment.
	//
	// `drifted` means the device's configurations have drifted since this deployment
	// was staged, and the deployment needs to be reviewed before it can be deployed.
	//
	// `staged` means the deployment is ready to be deployed.
	//
	// `queued` means the deployment's config instances are waiting to be received by
	// the device and will be deployed as soon as the device is online.
	//
	// `deployed` means the deployment's config instances are currently available for
	// consumption on the device.
	//
	// `removing` means the deployment's config instances are being removed from the
	// device.
	//
	// `archived` means the deployment is available for historical reference but cannot
	// be deployed and is not active on the device.
	//
	// Any of "drifted", "staged", "queued", "deployed", "removing", "archived".
	ActivityStatus DeploymentActivityStatus `json:"activity_status" api:"required"`
	// Timestamp of when the deployment was archived.
	ArchivedAt time.Time `json:"archived_at" api:"required" format:"date-time"`
	// Timestamp of when the device release was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Timestamp of when the deployment was deployed.
	DeployedAt time.Time `json:"deployed_at" api:"required" format:"date-time"`
	// The description of the deployment.
	Description string `json:"description" api:"required"`
	// ID of the device.
	DeviceID string `json:"device_id" api:"required"`
	// Last known error state of the deployment.
	//
	// `none` means there are no errors.
	//
	// `retrying` means an error has been encountered and the agent is retrying to
	// reach the target status.
	//
	// `failed` means a fatal error has been encountered; the deployment is archived
	// and, if deployed, removed from the device.
	//
	// Any of "none", "failed", "retrying".
	ErrorStatus DeploymentErrorStatus `json:"error_status" api:"required"`
	// The object type, which is always `deployment`.
	Object constant.Deployment `json:"object" default:"deployment"`
	// ID of the deployment that this deployment was patched from. Null if this
	// deployment has no parent.
	ParentID string `json:"parent_id" api:"required"`
	// Timestamp of when the deployment was queued.
	QueuedAt time.Time `json:"queued_at" api:"required" format:"date-time"`
	// ID of the release.
	ReleaseID string `json:"release_id" api:"required"`
	// This status merges the 'activity_status' and 'error_status' fields, with error
	// states taking precedence over activity states when errors are present. For
	// example, if the activity status is 'deployed' but the error status is 'failed',
	// the status is 'failed'. However, if the error status is 'none' and the activity
	// status is 'deployed', the status is 'deployed'.
	//
	// Any of "drifted", "staged", "queued", "deployed", "removing", "archived",
	// "failed", "retrying".
	Status DeploymentStatus `json:"status" api:"required"`
	// Desired state of the deployment.
	//
	// `staged` means the deployment is ready to be deployed.
	//
	// `deployed` means all config instances in the deployment are available for
	// consumption on the device.
	//
	// `archived` means the deployment is available for historical reference but cannot
	// be deployed and is not active on the device.
	//
	// Any of "staged", "deployed", "archived".
	TargetStatus DeploymentTargetStatus `json:"target_status" api:"required"`
	// Timestamp of when the device release was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Expand the config instances using 'expand=config_instances' in the query string.
	ConfigInstances []ConfigInstance `json:"config_instances"`
	// Expand the device using 'expand=device' in the query string.
	Device *Device `json:"device"`
	// The parent deployment that this deployment was patched from. Null if this
	// deployment has no parent. Expand the parent using 'expand=parent' in the query
	// string.
	Parent *Deployment `json:"parent" api:"nullable"`
	// Expand the release using 'expand=release' in the query string.
	Release Release `json:"release"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ActivityStatus  respjson.Field
		ArchivedAt      respjson.Field
		CreatedAt       respjson.Field
		DeployedAt      respjson.Field
		Description     respjson.Field
		DeviceID        respjson.Field
		ErrorStatus     respjson.Field
		Object          respjson.Field
		ParentID        respjson.Field
		QueuedAt        respjson.Field
		ReleaseID       respjson.Field
		Status          respjson.Field
		TargetStatus    respjson.Field
		UpdatedAt       respjson.Field
		ConfigInstances respjson.Field
		Device          respjson.Field
		Parent          respjson.Field
		Release         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Deployment) RawJSON() string { return r.JSON.raw }
func (r *Deployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Last known activity state of the deployment.
//
// `drifted` means the device's configurations have drifted since this deployment
// was staged, and the deployment needs to be reviewed before it can be deployed.
//
// `staged` means the deployment is ready to be deployed.
//
// `queued` means the deployment's config instances are waiting to be received by
// the device and will be deployed as soon as the device is online.
//
// `deployed` means the deployment's config instances are currently available for
// consumption on the device.
//
// `removing` means the deployment's config instances are being removed from the
// device.
//
// `archived` means the deployment is available for historical reference but cannot
// be deployed and is not active on the device.
type DeploymentActivityStatus string

const (
	DeploymentActivityStatusDrifted  DeploymentActivityStatus = "drifted"
	DeploymentActivityStatusStaged   DeploymentActivityStatus = "staged"
	DeploymentActivityStatusQueued   DeploymentActivityStatus = "queued"
	DeploymentActivityStatusDeployed DeploymentActivityStatus = "deployed"
	DeploymentActivityStatusRemoving DeploymentActivityStatus = "removing"
	DeploymentActivityStatusArchived DeploymentActivityStatus = "archived"
)

// Last known error state of the deployment.
//
// `none` means there are no errors.
//
// `retrying` means an error has been encountered and the agent is retrying to
// reach the target status.
//
// `failed` means a fatal error has been encountered; the deployment is archived
// and, if deployed, removed from the device.
type DeploymentErrorStatus string

const (
	DeploymentErrorStatusNone     DeploymentErrorStatus = "none"
	DeploymentErrorStatusFailed   DeploymentErrorStatus = "failed"
	DeploymentErrorStatusRetrying DeploymentErrorStatus = "retrying"
)

// This status merges the 'activity_status' and 'error_status' fields, with error
// states taking precedence over activity states when errors are present. For
// example, if the activity status is 'deployed' but the error status is 'failed',
// the status is 'failed'. However, if the error status is 'none' and the activity
// status is 'deployed', the status is 'deployed'.
type DeploymentStatus string

const (
	DeploymentStatusDrifted  DeploymentStatus = "drifted"
	DeploymentStatusStaged   DeploymentStatus = "staged"
	DeploymentStatusQueued   DeploymentStatus = "queued"
	DeploymentStatusDeployed DeploymentStatus = "deployed"
	DeploymentStatusRemoving DeploymentStatus = "removing"
	DeploymentStatusArchived DeploymentStatus = "archived"
	DeploymentStatusFailed   DeploymentStatus = "failed"
	DeploymentStatusRetrying DeploymentStatus = "retrying"
)

// Desired state of the deployment.
//
// `staged` means the deployment is ready to be deployed.
//
// `deployed` means all config instances in the deployment are available for
// consumption on the device.
//
// `archived` means the deployment is available for historical reference but cannot
// be deployed and is not active on the device.
type DeploymentTargetStatus string

const (
	DeploymentTargetStatusStaged   DeploymentTargetStatus = "staged"
	DeploymentTargetStatusDeployed DeploymentTargetStatus = "deployed"
	DeploymentTargetStatusArchived DeploymentTargetStatus = "archived"
)

type DeploymentList struct {
	// The list of deployments.
	Data []Deployment `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedList
}

// Returns the unmodified JSON received from the API
func (r DeploymentList) RawJSON() string { return r.JSON.raw }
func (r *DeploymentList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentNewParams struct {
	// The IDs of the config instances to deploy. A deployment must include config
	// instances according to the config schemas slots defined in the release.
	ConfigInstanceIDs []string `json:"config_instance_ids,omitzero" api:"required"`
	// The description of the deployment.
	Description string `json:"description" api:"required"`
	// The ID of the device that the deployment is being created for.
	DeviceID string `json:"device_id" api:"required"`
	// The release ID which this deployment adheres to.
	ReleaseID string `json:"release_id" api:"required"`
	// Desired state of the deployment.
	//
	// `staged` means the deployment is ready for deployment. Deployments can only be
	// staged if their release is not the device's current release.
	//
	// `deployed` means the deployment should be deployed to the device. Deployments
	// can only be deployed if their release is the device's current release.
	//
	// Any of "staged", "deployed".
	TargetStatus DeploymentNewParamsTargetStatus `json:"target_status,omitzero" api:"required"`
	// Parent deployment ID used as an optimistic-concurrency token. Tristate:
	//
	//   - omitted: server fills with the device's current `target_deployment_id`; no
	//     concurrency check.
	//   - null: caller asserts the device has no current target; 409 if it does.
	//   - value: caller asserts the device's current target equals this ID; 409 if it
	//     does not match.
	ParentID param.Opt[string] `json:"parent_id,omitzero"`
	// Fields to expand on the deployment resource.
	//
	// Any of "device", "release", "parent", "config_instances".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

func (r DeploymentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DeploymentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeploymentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [DeploymentNewParams]'s query parameters as `url.Values`.
func (r DeploymentNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Desired state of the deployment.
//
// `staged` means the deployment is ready for deployment. Deployments can only be
// staged if their release is not the device's current release.
//
// `deployed` means the deployment should be deployed to the device. Deployments
// can only be deployed if their release is the device's current release.
type DeploymentNewParamsTargetStatus string

const (
	DeploymentNewParamsTargetStatusStaged   DeploymentNewParamsTargetStatus = "staged"
	DeploymentNewParamsTargetStatusDeployed DeploymentNewParamsTargetStatus = "deployed"
)

type DeploymentGetParams struct {
	// Fields to expand on the deployment resource.
	//
	// Any of "device", "release", "parent", "config_instances".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeploymentGetParams]'s query parameters as `url.Values`.
func (r DeploymentGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeploymentListParams struct {
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The deployment IDs to filter by.
	ID []string `query:"id,omitzero" json:"-"`
	// The deployment activity statuses to filter by.
	//
	// Any of "drifted", "staged", "queued", "deployed", "removing", "archived".
	ActivityStatus []string `query:"activity_status,omitzero" json:"-"`
	// The deployment device IDs to filter by.
	DeviceID []string `query:"device_id,omitzero" json:"-"`
	// The deployment error statuses to filter by.
	//
	// Any of "none", "failed", "retrying".
	ErrorStatus []string `query:"error_status,omitzero" json:"-"`
	// Fields to expand on each deployment in the list.
	//
	// Any of "total_count", "device", "release", "parent", "config_instances".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Sort order for the deployment results.
	//
	// Any of "id:asc", "id:desc", "created_at:desc", "created_at:asc".
	OrderBy []string `query:"order_by,omitzero" json:"-"`
	// The deployment release IDs to filter by.
	ReleaseID []string `query:"release_id,omitzero" json:"-"`
	// The deployment target statuses to filter by.
	//
	// Any of "staged", "deployed", "archived".
	TargetStatus []string `query:"target_status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeploymentListParams]'s query parameters as `url.Values`.
func (r DeploymentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeploymentArchiveParams struct {
	// Fields to expand on the deployment resource.
	//
	// Any of "device", "release", "parent", "config_instances".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeploymentArchiveParams]'s query parameters as
// `url.Values`.
func (r DeploymentArchiveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeploymentDeployParams struct {
	// Fields to expand on the deployment resource.
	//
	// Any of "device", "release", "parent", "config_instances".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeploymentDeployParams]'s query parameters as `url.Values`.
func (r DeploymentDeployParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeploymentListDriftsParams struct {
	// Fields to expand on the deployment resource.
	//
	// Any of "device", "release", "parent", "config_instances".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeploymentListDriftsParams]'s query parameters as
// `url.Values`.
func (r DeploymentListDriftsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
