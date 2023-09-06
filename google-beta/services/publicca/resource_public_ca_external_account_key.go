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

package publicca

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/lorioux/google-beta/google-beta/tpgresource"
	transport_tpg "github.com/lorioux/google-beta/google-beta/transport"
)

func ResourcePublicCAExternalAccountKey() *schema.Resource {
	return &schema.Resource{
		Create: resourcePublicCAExternalAccountKeyCreate,
		Read:   resourcePublicCAExternalAccountKeyRead,
		Delete: resourcePublicCAExternalAccountKeyDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Location for the externalAccountKey. Currently only 'global' is supported.`,
				Default:     "global",
			},
			"b64_mac_key": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Base64-URL-encoded HS256 key. It is generated by the PublicCertificateAuthorityService
when the ExternalAccountKey is created.`,
				Sensitive: true,
			},
			"key_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `It is generated by the PublicCertificateAuthorityService when the ExternalAccountKey is created.`,
				Sensitive:   true,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Resource name. projects/{project}/locations/{location}/externalAccountKeys/{keyId}.`,
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

func resourcePublicCAExternalAccountKeyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})

	url, err := tpgresource.ReplaceVars(d, config, "{{PublicCABasePath}}projects/{{project}}/locations/{{location}}/externalAccountKeys")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ExternalAccountKey: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ExternalAccountKey: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating ExternalAccountKey: %s", err)
	}
	if err := d.Set("name", flattenPublicCAExternalAccountKeyName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}
	if err := d.Set("key_id", flattenPublicCAExternalAccountKeyKeyId(res["keyId"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "key_id": %s`, err)
	}
	if err := d.Set("b64_mac_key", flattenPublicCAExternalAccountKeyB64MacKey(res["b64MacKey"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "b64_mac_key": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ExternalAccountKey %q: %#v", d.Id(), res)

	return resourcePublicCAExternalAccountKeyRead(d, meta)
}

func resourcePublicCAExternalAccountKeyRead(d *schema.ResourceData, meta interface{}) error {
	// This resource could not be read from the API.
	return nil
}

func resourcePublicCAExternalAccountKeyDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] PublicCA ExternalAccountKey resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func flattenPublicCAExternalAccountKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenPublicCAExternalAccountKeyKeyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenPublicCAExternalAccountKeyB64MacKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
