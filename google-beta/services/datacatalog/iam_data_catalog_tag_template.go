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

package datacatalog

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/lorioux/google-beta/google-beta/tpgiamresource"
	"github.com/lorioux/google-beta/google-beta/tpgresource"
	transport_tpg "github.com/lorioux/google-beta/google-beta/transport"
)

var DataCatalogTagTemplateIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"region": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"tag_template": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type DataCatalogTagTemplateIamUpdater struct {
	project     string
	region      string
	tagTemplate string
	d           tpgresource.TerraformResourceData
	Config      *transport_tpg.Config
}

func DataCatalogTagTemplateIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	region, _ := tpgresource.GetRegion(d, config)
	if region != "" {
		if err := d.Set("region", region); err != nil {
			return nil, fmt.Errorf("Error setting region: %s", err)
		}
	}
	values["region"] = region
	if v, ok := d.GetOk("tag_template"); ok {
		values["tag_template"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/tagTemplates/(?P<tag_template>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<tag_template>[^/]+)", "(?P<region>[^/]+)/(?P<tag_template>[^/]+)", "(?P<tag_template>[^/]+)"}, d, config, d.Get("tag_template").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataCatalogTagTemplateIamUpdater{
		project:     values["project"],
		region:      values["region"],
		tagTemplate: values["tag_template"],
		d:           d,
		Config:      config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("region", u.region); err != nil {
		return nil, fmt.Errorf("Error setting region: %s", err)
	}
	if err := d.Set("tag_template", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting tag_template: %s", err)
	}

	return u, nil
}

func DataCatalogTagTemplateIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		values["project"] = project
	}

	region, _ := tpgresource.GetRegion(d, config)
	if region != "" {
		values["region"] = region
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/tagTemplates/(?P<tag_template>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<tag_template>[^/]+)", "(?P<region>[^/]+)/(?P<tag_template>[^/]+)", "(?P<tag_template>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataCatalogTagTemplateIamUpdater{
		project:     values["project"],
		region:      values["region"],
		tagTemplate: values["tag_template"],
		d:           d,
		Config:      config,
	}
	if err := d.Set("tag_template", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting tag_template: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *DataCatalogTagTemplateIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyTagTemplateUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := tpgresource.GetProject(u.d, u.Config)
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
		Method:    "POST",
		Project:   project,
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

func (u *DataCatalogTagTemplateIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyTagTemplateUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := tpgresource.GetProject(u.d, u.Config)
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
		Project:   project,
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

func (u *DataCatalogTagTemplateIamUpdater) qualifyTagTemplateUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{DataCatalogBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", u.project, u.region, u.tagTemplate), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *DataCatalogTagTemplateIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", u.project, u.region, u.tagTemplate)
}

func (u *DataCatalogTagTemplateIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-datacatalog-tagtemplate-%s", u.GetResourceId())
}

func (u *DataCatalogTagTemplateIamUpdater) DescribeResource() string {
	return fmt.Sprintf("datacatalog tagtemplate %q", u.GetResourceId())
}
