package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/container"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func ContainsString(strings []string, matchString string) bool {
	for _, stringValue := range strings {
		if stringValue == matchString {
			return true
		}
	}
	return false
}

func GetMOName(dn string) string {
	splittedDn := strings.Split(dn, "/")
	if len(splittedDn) > 1 {
		return strings.Join(strings.Split(splittedDn[len(splittedDn)-1], "-")[1:], "-")
	}
	return splittedDn[0]
}

func CheckDn(ctx context.Context, client *client.Client, dn string, diags *diag.Diagnostics) {
	tflog.Debug(ctx, fmt.Sprintf("validate relation dn: %s", dn))
	_, err := client.Get(dn)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("failed validate relation dn: %s", dn))
		diags.AddError(
			"Relation target dn validation failed",
			fmt.Sprintf("The relation target dn is not found: %s", dn),
		)
	}
}

func DoRestRequestEscapeHtml(ctx context.Context, diags *diag.Diagnostics, client *client.Client, path, method string, payload *container.Container, escapeHtml bool) *container.Container {

	// Ensure path starts with a slash to assure signature is created correctly
	if !strings.HasPrefix("/", path) {
		path = fmt.Sprintf("/%s", path)
	}
	var restRequest *http.Request
	var err error
	if escapeHtml {
		restRequest, err = client.MakeRestRequest(method, path, payload, true)
	} else {
		restRequest, err = client.MakeRestRequestRaw(method, path, payload.EncodeJSON(), true)
	}
	if err != nil {
		diags.AddError(
			"Creation of rest request failed",
			fmt.Sprintf("err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	cont, restResponse, err := client.Do(restRequest)

	if restResponse != nil && cont.Data() != nil && restResponse.StatusCode != 200 {
		errCode := models.StripQuotes(models.StripSquareBrackets(cont.Search("imdata", "error", "attributes", "code").String()))
		if errCode != "1" && errCode != "103" && errCode != "107" && errCode != "120" {
			diags.AddError(
				fmt.Sprintf("The %s rest request failed", strings.ToLower(method)),
				fmt.Sprintf("Code: %d Response: %s, err: %s. Please report this issue to the provider developers.", restResponse.StatusCode, cont.Data().(map[string]interface{})["imdata"], err),
			)
			return nil
		}
		tflog.Debug(ctx, models.StripQuotes(models.StripSquareBrackets(cont.Search("imdata", "error", "attributes", "text").String())))
	} else if err != nil {
		diags.AddError(
			fmt.Sprintf("The %s rest request failed", strings.ToLower(method)),
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	return cont
}

func DoRestRequest(ctx context.Context, diags *diag.Diagnostics, client *client.Client, path, method string, payload *container.Container) *container.Container {
	return DoRestRequestEscapeHtml(ctx, diags, client, path, method, payload, true)
}

func GetDeleteJsonPayload(ctx context.Context, diags *diag.Diagnostics, className, dn string) *container.Container {

	jsonString := fmt.Sprintf(`{"%s":{"attributes":{"dn": "%s","status": "deleted"}}}`, className, dn)
	jsonPayload, err := container.ParseJSON([]byte(jsonString))
	if err != nil {
		diags.AddError(
			"Construction of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}
	return jsonPayload
}

type setToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate struct{}

func SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate() planmodifier.String {
	return setToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate{}
}

func (m setToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate) Description(_ context.Context) string {
	return "During the update phase, set the value of this attribute to StringNull when the state value is null and the plan value is unknown."
}

func (m setToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate) MarkdownDescription(_ context.Context) string {
	return "During the update phase, set the value of this attribute to StringNull when the state value is null and the plan value is unknown."
}

// Custom plan modifier to set the plan value to null under certain conditions
func (m setToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Set the plan value to StringType null when state value is null and plan value is unknown during an Update
	if !req.State.Raw.IsNull() && req.StateValue.IsNull() && req.PlanValue.IsUnknown() {
		resp.PlanValue = types.StringNull()
	}
	return
}

// type setToUnknownWhenStateIsNullPlanIsUnknownDuringUpdate struct{}

// func SetToUnknownWhenStateIsNullPlanIsUnknownDuringUpdate() planmodifier.String {
// 	return setToUnknownWhenStateIsNullPlanIsUnknownDuringUpdate{}
// }

// func (m setToUnknownWhenStateIsNullPlanIsUnknownDuringUpdate) Description(_ context.Context) string {
// 	return "During the update phase, set the value of this attribute to StringNull when the state value is null and the plan value is unknown."
// }

// func (m setToUnknownWhenStateIsNullPlanIsUnknownDuringUpdate) MarkdownDescription(_ context.Context) string {
// 	return "During the update phase, set the value of this attribute to StringNull when the state value is null and the plan value is unknown."
// }

// // Custom plan modifier to set the plan value to null under certain conditions
// func (m setToUnknownWhenStateIsNullPlanIsUnknownDuringUpdate) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
// 	// Set the plan value to StringType null when state value is null and plan value is unknown during an Update
// 	if !req.State.Raw.IsNull() && req.StateValue.IsNull() && req.PlanValue.IsNull() {
// 		resp.PlanValue = types.StringUnknown()
// 	}
// 	return
// }

type setToSetNullWhenStateIsNullPlanIsUnknownDuringUpdate struct{}

func SetToSetNullWhenStateIsNullPlanIsUnknownDuringUpdate() planmodifier.Set {
	return setToSetNullWhenStateIsNullPlanIsUnknownDuringUpdate{}
}

func (m setToSetNullWhenStateIsNullPlanIsUnknownDuringUpdate) Description(_ context.Context) string {
	return "During the update phase, set the value of this attribute to SetNull when the state value is null and the plan value is unknown."
}

func (m setToSetNullWhenStateIsNullPlanIsUnknownDuringUpdate) MarkdownDescription(_ context.Context) string {
	return "During the update phase, set the value of this attribute to SetNull when the state value is null and the plan value is unknown."
}

// Custom plan modifier to set the plan value to null under certain conditions
func (m setToSetNullWhenStateIsNullPlanIsUnknownDuringUpdate) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	// Set the plan value to SetType null when state value is null and plan value is unknown during an Update
	log.Printf("hereStateval %v", req.StateValue)
	log.Printf("hereStatevalIsUnknown %v", req.StateValue.IsUnknown())
	log.Printf("hereStatevalIsNull %v", req.StateValue.IsNull())
	log.Printf("herePlanval %v", req.PlanValue)
	log.Printf("herePlanvalIsUnknown %v", req.PlanValue.IsUnknown())
	log.Printf("herePlanvalIsNull %v", req.PlanValue.IsNull())

	// 	log.Printf("hereType %v %v", idx, element.Type(ctx))
	// }
	// if !req.State.Raw.IsNull() && req.PlanValue.Equal(req.StateValue) && req.ConfigValue.IsNull(){
	// 	resp.PlanValue = types.SetNull(req.PlanValue.ElementType(ctx))
	// }

	// if !req.State.Raw.IsNull() && req.StateValue.IsNull() && req.PlanValue.IsUnknown() {
	// 	//log.Printf("hereStateval %v;hereStateIsUnknown %v; hereStateIsNull %v;herePlanVal %v, herePlanIsUnknown %v, herePlanIsNull %v, hereConfigVal %v, hereConfisIsUnkown %v, hereConfigIsNull %v", req.StateValue, req.StateValue.IsUnknown(), req.StateValue.IsNull(), req.PlanValue, req.PlanValue.IsUnknown(), req.PlanValue.IsNull(), req.ConfigValue, req.ConfigValue.IsUnknown(), req.ConfigValue.IsNull())
	// 	resp.PlanValue = types.SetNull(req.PlanValue.ElementType(ctx))
	// }
	return
}

func areAllStringFieldsEmpty(v interface{}) bool {
	val := reflect.ValueOf(v)

	// Check if the passed interface is a struct
	if val.Kind() != reflect.Struct {
		return false
	}

	// Iterate over all fields in the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := field.Type()

		// Check if the field is of type basestring
		if fieldType != reflect.TypeOf(basetypes.NewStringValue("")) {
			return false
		}

		// Check if the string is not empty
		if !field.IsZero() {
			return false
		}
	}

	// All string fields are empty
	return true
}
