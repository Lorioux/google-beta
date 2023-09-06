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

package healthcare

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/lorioux/google-beta/google-beta/tpgiamresource"
	"github.com/lorioux/google-beta/google-beta/tpgresource"
	transport_tpg "github.com/lorioux/google-beta/google-beta/transport"
)

var HealthcareConsentStoreIamSchema = map[string]*schema.Schema{
	"dataset": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"consent_store_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type HealthcareConsentStoreIamUpdater struct {
	dataset        string
	consentStoreId string
	d              tpgresource.TerraformResourceData
	Config         *transport_tpg.Config
}

func HealthcareConsentStoreIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	if v, ok := d.GetOk("dataset"); ok {
		values["dataset"] = v.(string)
	}

	if v, ok := d.GetOk("consent_store_id"); ok {
		values["consent_store_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"(?P<dataset>.+)/consentStores/(?P<consent_store_id>[^/]+)", "(?P<consent_store_id>[^/]+)"}, d, config, d.Get("consent_store_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &HealthcareConsentStoreIamUpdater{
		dataset:        values["dataset"],
		consentStoreId: values["consent_store_id"],
		d:              d,
		Config:         config,
	}

	if err := d.Set("dataset", u.dataset); err != nil {
		return nil, fmt.Errorf("Error setting dataset: %s", err)
	}
	if err := d.Set("consent_store_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting consent_store_id: %s", err)
	}

	return u, nil
}

func HealthcareConsentStoreIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	m, err := tpgresource.GetImportIdQualifiers([]string{"(?P<dataset>.+)/consentStores/(?P<consent_store_id>[^/]+)", "(?P<consent_store_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &HealthcareConsentStoreIamUpdater{
		dataset:        values["dataset"],
		consentStoreId: values["consent_store_id"],
		d:              d,
		Config:         config,
	}
	if err := d.Set("consent_store_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting consent_store_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *HealthcareConsentStoreIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyConsentStoreUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "GET",
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
	})
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = tpgresource.Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *HealthcareConsentStoreIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyConsentStoreUrl("setIamPolicy")
	if err != nil {
		return err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   u.d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *HealthcareConsentStoreIamUpdater) qualifyConsentStoreUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{HealthcareBasePath}}%s:%s", fmt.Sprintf("%s/consentStores/%s", u.dataset, u.consentStoreId), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *HealthcareConsentStoreIamUpdater) GetResourceId() string {
	return fmt.Sprintf("%s/consentStores/%s", u.dataset, u.consentStoreId)
}

func (u *HealthcareConsentStoreIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-healthcare-consentstore-%s", u.GetResourceId())
}

func (u *HealthcareConsentStoreIamUpdater) DescribeResource() string {
	return fmt.Sprintf("healthcare consentstore %q", u.GetResourceId())
}
