// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package appengine

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/lorioux/google-beta/google-beta/tpgresource"
	transport_tpg "github.com/lorioux/google-beta/google-beta/transport"
	"github.com/lorioux/google-beta/google-beta/verify"
)

func ResourceAppEngineServiceNetworkSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppEngineServiceNetworkSettingsCreate,
		Read:   resourceAppEngineServiceNetworkSettingsRead,
		Update: resourceAppEngineServiceNetworkSettingsUpdate,
		Delete: resourceAppEngineServiceNetworkSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAppEngineServiceNetworkSettingsImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"network_settings": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Ingress settings for this service. Will apply to all versions.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ingress_traffic_allowed": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED", "INGRESS_TRAFFIC_ALLOWED_ALL", "INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY", "INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB", ""}),
							Description:  `The ingress settings for version or service. Default value: "INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED" Possible values: ["INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED", "INGRESS_TRAFFIC_ALLOWED_ALL", "INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY", "INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB"]`,
							Default:      "INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED",
						},
					},
				},
			},
			"service": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The name of the service these settings apply to.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAppEngineServiceNetworkSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	idProp, err := expandAppEngineServiceNetworkSettingsService(d.Get("service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service"); !tpgresource.IsEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	networkSettingsProp, err := expandAppEngineServiceNetworkSettingsNetworkSettings(d.Get("network_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkSettingsProp)) && (ok || !reflect.DeepEqual(v, networkSettingsProp)) {
		obj["networkSettings"] = networkSettingsProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/services/{{service}}?updateMask=networkSettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServiceNetworkSettings: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceNetworkSettings: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating ServiceNetworkSettings: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}/services/{{service}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = AppEngineOperationWaitTime(
		config, res, project, "Creating ServiceNetworkSettings", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ServiceNetworkSettings: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ServiceNetworkSettings %q: %#v", d.Id(), res)

	return resourceAppEngineServiceNetworkSettingsRead(d, meta)
}

func resourceAppEngineServiceNetworkSettingsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/services/{{service}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceNetworkSettings: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AppEngineServiceNetworkSettings %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ServiceNetworkSettings: %s", err)
	}

	if err := d.Set("service", flattenAppEngineServiceNetworkSettingsService(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceNetworkSettings: %s", err)
	}
	if err := d.Set("network_settings", flattenAppEngineServiceNetworkSettingsNetworkSettings(res["networkSettings"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceNetworkSettings: %s", err)
	}

	return nil
}

func resourceAppEngineServiceNetworkSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceNetworkSettings: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	idProp, err := expandAppEngineServiceNetworkSettingsService(d.Get("service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	networkSettingsProp, err := expandAppEngineServiceNetworkSettingsNetworkSettings(d.Get("network_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, networkSettingsProp)) {
		obj["networkSettings"] = networkSettingsProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/services/{{service}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ServiceNetworkSettings %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("service") {
		updateMask = append(updateMask, "id")
	}

	if d.HasChange("network_settings") {
		updateMask = append(updateMask, "networkSettings")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating ServiceNetworkSettings %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ServiceNetworkSettings %q: %#v", d.Id(), res)
	}

	err = AppEngineOperationWaitTime(
		config, res, project, "Updating ServiceNetworkSettings", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceAppEngineServiceNetworkSettingsRead(d, meta)
}

func resourceAppEngineServiceNetworkSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] AppEngine ServiceNetworkSettings resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceAppEngineServiceNetworkSettingsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"apps/(?P<project>[^/]+)/services/(?P<service>[^/]+)",
		"(?P<project>[^/]+)/(?P<service>[^/]+)",
		"(?P<service>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}/services/{{service}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAppEngineServiceNetworkSettingsService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAppEngineServiceNetworkSettingsNetworkSettings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["ingress_traffic_allowed"] =
		flattenAppEngineServiceNetworkSettingsNetworkSettingsIngressTrafficAllowed(original["ingressTrafficAllowed"], d, config)
	return []interface{}{transformed}
}
func flattenAppEngineServiceNetworkSettingsNetworkSettingsIngressTrafficAllowed(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandAppEngineServiceNetworkSettingsService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineServiceNetworkSettingsNetworkSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIngressTrafficAllowed, err := expandAppEngineServiceNetworkSettingsNetworkSettingsIngressTrafficAllowed(original["ingress_traffic_allowed"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIngressTrafficAllowed); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ingressTrafficAllowed"] = transformedIngressTrafficAllowed
	}

	return transformed, nil
}

func expandAppEngineServiceNetworkSettingsNetworkSettingsIngressTrafficAllowed(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
