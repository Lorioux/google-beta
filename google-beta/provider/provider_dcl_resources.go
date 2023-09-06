// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	services "github.com/lorioux/google-beta/google-beta/services"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/assuredworkloads"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigqueryreservation"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudbuild"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudbuildv2"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/clouddeploy"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/compute"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/containeraws"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/containerazure"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataplex"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataproc"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/eventarc"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebaserules"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/gkehub"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/networkconnectivity"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/orgpolicy"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/privateca"
	// "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/recaptchaenterprise"
)

var dclResources = map[string]*schema.Resource{
	"google_apikeys_key":                                        services.apikeys.ResourceApikeysKey(),
	"google_assured_workloads_workload":                         services.assuredworkloads.ResourceAssuredWorkloadsWorkload(),
	"google_bigquery_reservation_assignment":                    services.bigqueryreservation.ResourceBigqueryReservationAssignment(),
	"google_cloudbuild_worker_pool":                             services.cloudbuild.ResourceCloudbuildWorkerPool(),
	"google_cloudbuildv2_connection":                            services.cloudbuildv2.ResourceCloudbuildv2Connection(),
	"google_cloudbuildv2_repository":                            services.cloudbuildv2.ResourceCloudbuildv2Repository(),
	"google_clouddeploy_delivery_pipeline":                      services.clouddeploy.ResourceClouddeployDeliveryPipeline(),
	"google_clouddeploy_target":                                 services.clouddeploy.ResourceClouddeployTarget(),
	"google_compute_firewall_policy":                            services.compute.ResourceComputeFirewallPolicy(),
	"google_compute_firewall_policy_association":                services.compute.ResourceComputeFirewallPolicyAssociation(),
	"google_compute_firewall_policy_rule":                       services.compute.ResourceComputeFirewallPolicyRule(),
	"google_compute_region_network_firewall_policy":             services.compute.ResourceComputeRegionNetworkFirewallPolicy(),
	"google_compute_network_firewall_policy":                    services.compute.ResourceComputeNetworkFirewallPolicy(),
	"google_compute_network_firewall_policy_association":        services.compute.ResourceComputeNetworkFirewallPolicyAssociation(),
	"google_compute_region_network_firewall_policy_association": services.compute.ResourceComputeRegionNetworkFirewallPolicyAssociation(),
	"google_compute_network_firewall_policy_rule":               services.compute.ResourceComputeNetworkFirewallPolicyRule(),
	"google_compute_region_network_firewall_policy_rule":        services.compute.ResourceComputeRegionNetworkFirewallPolicyRule(),
	"google_container_aws_cluster":                              services.containeraws.ResourceContainerAwsCluster(),
	"google_container_aws_node_pool":                            services.containeraws.ResourceContainerAwsNodePool(),
	"google_container_azure_client":                             services.containerazure.ResourceContainerAzureClient(),
	"google_container_azure_cluster":                            services.containerazure.ResourceContainerAzureCluster(),
	"google_container_azure_node_pool":                          services.containerazure.ResourceContainerAzureNodePool(),
	"google_dataplex_asset":                                     services.dataplex.ResourceDataplexAsset(),
	"google_dataplex_lake":                                      services.dataplex.ResourceDataplexLake(),
	"google_dataplex_zone":                                      services.dataplex.ResourceDataplexZone(),
	"google_dataproc_workflow_template":                         services.dataproc.ResourceDataprocWorkflowTemplate(),
	"google_eventarc_channel":                                   services.eventarc.ResourceEventarcChannel(),
	"google_eventarc_google_channel_config":                     services.eventarc.ResourceEventarcGoogleChannelConfig(),
	"google_eventarc_trigger":                                   services.eventarc.ResourceEventarcTrigger(),
	"google_firebaserules_release":                              services.firebaserules.ResourceFirebaserulesRelease(),
	"google_firebaserules_ruleset":                              services.firebaserules.ResourceFirebaserulesRuleset(),
	"google_gke_hub_feature_membership":                         services.gkehub.ResourceGkeHubFeatureMembership(),
	"google_network_connectivity_hub":                           services.networkconnectivity.ResourceNetworkConnectivityHub(),
	"google_network_connectivity_spoke":                         services.networkconnectivity.ResourceNetworkConnectivitySpoke(),
	"google_org_policy_policy":                                  services.orgpolicy.ResourceOrgPolicyPolicy(),
	"google_privateca_certificate_template":                     services.privateca.ResourcePrivatecaCertificateTemplate(),
	"google_recaptcha_enterprise_key":                           services.recaptchaenterprise.ResourceRecaptchaEnterpriseKey(),
}
