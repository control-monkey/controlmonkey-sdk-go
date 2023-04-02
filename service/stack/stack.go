package stack

import (
	"context"
	"encoding/json"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/uritemplates"
	"io"
	"net/http"
	"time"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
)

// A Product represents the type of an operating system.
type Product int

const (
	// ProductWindows represents the Windows product.
	ProductWindows Product = iota

	// ProductWindowsVPC represents the Windows (Amazon VPC) product.
	ProductWindowsVPC

	// ProductLinuxUnix represents the Linux/Unix product.
	ProductLinuxUnix

	// ProductLinuxUnixVPC represents the Linux/Unix (Amazon VPC) product.
	ProductLinuxUnixVPC

	// ProductSUSELinux represents the SUSE Linux product.
	ProductSUSELinux

	// ProductSUSELinuxVPC represents the SUSE Linux (Amazon VPC) product.
	ProductSUSELinuxVPC
)

var ProductName = map[Product]string{
	ProductWindows:      "Windows",
	ProductWindowsVPC:   "Windows (Amazon VPC)",
	ProductLinuxUnix:    "Linux/UNIX",
	ProductLinuxUnixVPC: "Linux/UNIX (Amazon VPC)",
	ProductSUSELinux:    "SUSE Linux",
	ProductSUSELinuxVPC: "SUSE Linux (Amazon VPC)",
}

func (p Product) String() string {
	return ProductName[p]
}

// region Plan

type Plan struct {
	ID       *string `json:"id,omitempty"`
	Status   *string `json:"status,omitempty"`
	IsActive *bool   `json:"isActive,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// forceSendFields is a read of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	forceSendFields []string

	// nullFields is a read of field names (e.g. "Keys") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this read has a non-empty value.
	// This may be used to include null fields in Patch requests.
	nullFields []string
}

type ListPlansInput struct{}
type ListPlansOutput struct {
	Plans []*Plan `json:"stacks,omitempty"`
}

type CreatePlanInput struct {
	StackId    *string `json:"stackId,omitempty" validate:"required"`
	HeadBranch *string `json:"headBranch,omitempty"`
	HeadSha    *string `json:"headSha,omitempty"`
}

type CreatePlanOutput struct {
	Plan *Plan `json:"plan,omitempty"`
}

type ReadPlanInput struct {
	PlanId *string `json:"planId,omitempty"`
}

type ReadPlanOutput struct {
	Plan *Plan `json:"plan,omitempty"`
}

type UpdatePlanInput struct{}
type UpdatePlanOutput struct{}
type DeletePlanInput struct{}
type DeletePlanOutput struct{}

func planFromJSON(in []byte) (*Plan, error) {
	b := new(Plan)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func plansFromJSON(in []byte) ([]*Plan, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Plan, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := planFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func plansFromHttpResponse(resp *http.Response) ([]*Plan, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return plansFromJSON(body)
}

func (s *ServiceOp) CreatePlan(ctx context.Context, input *CreatePlanInput) (*CreatePlanOutput, error) {
	r := client.NewRequest(http.MethodPost, "/stack/plan")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := plansFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreatePlanOutput)
	if len(gs) > 0 {
		output.Plan = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadPlan(ctx context.Context, input *ReadPlanInput) (*ReadPlanOutput, error) {
	path, err := uritemplates.Expand("/stack/plan/{planId}", uritemplates.Values{
		"planId": controlmonkey.StringValue(input.PlanId),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := plansFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadPlanOutput)
	if len(gs) > 0 {
		output.Plan = gs[0]
	}

	return output, nil
}

func (o Plan) MarshalJSON() ([]byte, error) {
	type noMethod Plan
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Plan) SetID(v *string) *Plan {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

// endregion

//region Deployment

type Deployment struct {
	ID       *string `json:"id,omitempty"`
	Status   *string `json:"status,omitempty"`
	IsActive *bool   `json:"isActive,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// forceSendFields is a read of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	forceSendFields []string

	// nullFields is a read of field names (e.g. "Keys") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this read has a non-empty value.
	// This may be used to include null fields in Patch requests.
	nullFields []string
}

type CreateDeploymentInput struct {
	StackId *string `json:"stackId,omitempty"`
}

type CreateDeploymentOutput struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

type ReadDeploymentInput struct {
	DeploymentId *string `json:"deploymentId,omitempty"`
}

type ReadDeploymentOutput struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

func deploymentFromJSON(in []byte) (*Deployment, error) {
	b := new(Deployment)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func deploymentsFromJSON(in []byte) ([]*Deployment, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Deployment, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := deploymentFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func deploymentsFromHttpResponse(resp *http.Response) ([]*Deployment, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return deploymentsFromJSON(body)
}

func (s *ServiceOp) CreateDeployment(ctx context.Context, input *CreateDeploymentInput) (*CreateDeploymentOutput, error) {
	r := client.NewRequest(http.MethodPost, "/stack/deployment")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := deploymentsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateDeploymentOutput)
	if len(gs) > 0 {
		output.Deployment = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadDeployment(ctx context.Context, input *ReadDeploymentInput) (*ReadDeploymentOutput, error) {
	path, err := uritemplates.Expand("/stack/deployment/{deploymentId}", uritemplates.Values{
		"deploymentId": controlmonkey.StringValue(input.DeploymentId),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := deploymentsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadDeploymentOutput)
	if len(gs) > 0 {
		output.Deployment = gs[0]
	}

	return output, nil
}

func (o Deployment) MarshalJSON() ([]byte, error) {
	type noMethod Deployment
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Deployment) SetID(v *string) *Deployment {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

//endregion
