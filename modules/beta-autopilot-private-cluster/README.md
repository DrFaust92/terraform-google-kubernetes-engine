# Terraform Kubernetes Engine Module

This module handles opinionated Google Cloud Platform Kubernetes Engine cluster creation and configuration with Node Pools, Network Policy, etc. This particular submodule creates a [private cluster](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters)Beta features are enabled in this submodule.
The resources/services/activations/deletions that this module will create/trigger are:

- Create a GKE cluster with the provided addons
- Create GKE Node Pool(s) with provided configuration and attach to cluster
- Replace the default kube-dns configmap if `stub_domains` are provided
- Activate network policy if `network_policy` is true

Sub modules are provided for creating private clusters, beta private clusters, and beta public clusters as well.  Beta sub modules allow for the use of various GKE beta features. See the modules directory for the various sub modules.

## Private Cluster Details

For details on configuring private clusters with this module, check the [troubleshooting guide](https://github.com/terraform-google-modules/terraform-google-kubernetes-engine/blob/master/docs/private_clusters.md).

## Compatibility

This module is meant for use with Terraform 1.3+ and tested using Terraform 1.10+.
If you find incompatibilities using Terraform `>=1.3`, please open an issue.

If you haven't [upgraded to 1.3][terraform-1.3-upgrade] and need a Terraform
0.13.x-compatible version of this module, the last released version
intended for Terraform 0.13.x is [27.0.0].

If you haven't [upgraded to 0.13][terraform-0.13-upgrade] and need a Terraform
0.12.x-compatible version of this module, the last released version
intended for Terraform 0.12.x is [12.3.0].

## Usage

There are multiple examples included in the [examples](https://github.com/terraform-google-modules/terraform-google-kubernetes-engine/tree/master/examples) folder but simple usage is as follows:

```hcl
# google_client_config and kubernetes provider must be explicitly specified like the following.
data "google_client_config" "default" {}

provider "kubernetes" {
  host                   = "https://${module.gke.endpoint}"
  token                  = data.google_client_config.default.access_token
  cluster_ca_certificate = base64decode(module.gke.ca_certificate)
}

module "gke" {
  source                     = "terraform-google-modules/kubernetes-engine/google//modules/beta-autopilot-private-cluster"
  project_id                 = "<PROJECT ID>"
  name                       = "gke-test-1"
  region                     = "us-central1"
  zones                      = ["us-central1-a", "us-central1-b", "us-central1-f"]
  network                    = "vpc-01"
  subnetwork                 = "us-central1-01"
  ip_range_pods              = "us-central1-01-gke-01-pods"
  ip_range_services          = "us-central1-01-gke-01-services"
  horizontal_pod_autoscaling = true
  filestore_csi_driver       = false
  enable_private_endpoint    = true
  enable_private_nodes       = true
  dns_cache                  = false

}
```

<!-- do not understand what this is about -->
Then perform the following commands on the root folder:

- `terraform init` to get the plugins
- `terraform plan` to see the infrastructure plan
- `terraform apply` to apply the infrastructure build
- `terraform destroy` to destroy the built infrastructure

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| add\_cluster\_firewall\_rules | Create additional firewall rules | `bool` | `false` | no |
| add\_master\_webhook\_firewall\_rules | Create master\_webhook firewall rules for ports defined in `firewall_inbound_ports` | `bool` | `false` | no |
| add\_shadow\_firewall\_rules | Create GKE shadow firewall (the same as default firewall rules with firewall logs enabled). | `bool` | `false` | no |
| additional\_ip\_range\_pods | List of _names_ of the additional secondary subnet ip ranges to use for pods | `list(string)` | `[]` | no |
| allow\_net\_admin | (Optional) Enable NET\_ADMIN for the cluster. | `bool` | `null` | no |
| authenticator\_security\_group | The name of the RBAC security group for use with Google security groups in Kubernetes RBAC. Group name must be in format gke-security-groups@yourdomain.com | `string` | `null` | no |
| boot\_disk\_kms\_key | The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool, if not overridden in `node_pools`. This should be of the form projects/[KEY\_PROJECT\_ID]/locations/[LOCATION]/keyRings/[RING\_NAME]/cryptoKeys/[KEY\_NAME]. For more information about protecting resources with Cloud KMS Keys please see: https://cloud.google.com/compute/docs/disks/customer-managed-encryption | `string` | `null` | no |
| cluster\_ipv4\_cidr | The IP address range of the kubernetes pods in this cluster. Default is an automatically assigned CIDR. | `string` | `null` | no |
| cluster\_resource\_labels | The GCE resource labels (a map of key/value pairs) to be applied to the cluster | `map(string)` | `{}` | no |
| create\_service\_account | Defines if service account specified to run nodes should be created. | `bool` | `true` | no |
| database\_encryption | Application-layer Secrets Encryption settings. The object format is {state = string, key\_name = string}. Valid values of state are: "ENCRYPTED"; "DECRYPTED". key\_name is the name of a CloudKMS key. | `list(object({ state = string, key_name = string }))` | <pre>[<br>  {<br>    "key_name": "",<br>    "state": "DECRYPTED"<br>  }<br>]</pre> | no |
| deletion\_protection | Whether or not to allow Terraform to destroy the cluster. | `bool` | `true` | no |
| deploy\_using\_private\_endpoint | A toggle for Terraform and kubectl to connect to the master's internal IP address during deployment. | `bool` | `false` | no |
| description | The description of the cluster | `string` | `""` | no |
| disable\_default\_snat | Whether to disable the default SNAT to support the private use of public IP addresses | `bool` | `false` | no |
| disable\_l4\_lb\_firewall\_reconciliation | Disable L4 Load Balancer firewall reconciliation | `bool` | `null` | no |
| dns\_allow\_external\_traffic | (Optional) Controls whether external traffic is allowed over the dns endpoint. | `bool` | `null` | no |
| dns\_cache | The status of the NodeLocal DNSCache addon. | `bool` | `true` | no |
| enable\_binary\_authorization | Enable BinAuthZ Admission controller | `bool` | `false` | no |
| enable\_cilium\_clusterwide\_network\_policy | Enable Cilium Cluster Wide Network Policies on the cluster | `bool` | `false` | no |
| enable\_confidential\_nodes | An optional flag to enable confidential node config. | `bool` | `false` | no |
| enable\_cost\_allocation | Enables Cost Allocation Feature and the cluster name and namespace of your GKE workloads appear in the labels field of the billing export to BigQuery | `bool` | `false` | no |
| enable\_fqdn\_network\_policy | Enable FQDN Network Policies on the cluster | `bool` | `null` | no |
| enable\_l4\_ilb\_subsetting | Enable L4 ILB Subsetting on the cluster | `bool` | `false` | no |
| enable\_multi\_networking | Whether multi-networking is enabled for this cluster | `bool` | `null` | no |
| enable\_network\_egress\_export | Whether to enable network egress metering for this cluster. If enabled, a daemonset will be created in the cluster to meter network egress traffic. | `bool` | `false` | no |
| enable\_private\_endpoint | Whether the master's internal IP address is used as the cluster endpoint | `bool` | `false` | no |
| enable\_private\_nodes | Whether nodes have internal IP addresses only | `bool` | `true` | no |
| enable\_resource\_consumption\_export | Whether to enable resource consumption metering on this cluster. When enabled, a table will be created in the resource export BigQuery dataset to store resource consumption data. The resulting table can be joined with the resource usage table or with BigQuery billing export. | `bool` | `true` | no |
| enable\_secret\_manager\_addon | Enable the Secret Manager add-on for this cluster | `bool` | `false` | no |
| enable\_tpu | Enable Cloud TPU resources in the cluster. WARNING: changing this after cluster creation is destructive! | `bool` | `false` | no |
| enable\_vertical\_pod\_autoscaling | Vertical Pod Autoscaling automatically adjusts the resources of pods controlled by it | `bool` | `true` | no |
| enterprise\_config | (Optional) Enable or disable GKE enterprise. Valid values are STANDARD and ENTERPRISE. | `string` | `null` | no |
| filestore\_csi\_driver | The status of the Filestore CSI driver addon, which allows the usage of filestore instance as volumes | `bool` | `false` | no |
| firewall\_inbound\_ports | List of TCP ports for admission/webhook controllers. Either flag `add_master_webhook_firewall_rules` or `add_cluster_firewall_rules` (also adds egress rules) must be set to `true` for inbound-ports firewall rules to be applied. | `list(string)` | <pre>[<br>  "8443",<br>  "9443",<br>  "15017"<br>]</pre> | no |
| firewall\_priority | Priority rule for firewall rules | `number` | `1000` | no |
| fleet\_project | (Optional) Register the cluster with the fleet in this project. | `string` | `null` | no |
| fleet\_project\_grant\_service\_agent | (Optional) Grant the fleet project service identity the `roles/gkehub.serviceAgent` and `roles/gkehub.crossProjectServiceAgent` roles. | `bool` | `false` | no |
| gateway\_api\_channel | The gateway api channel of this cluster. Accepted values are `CHANNEL_STANDARD` and `CHANNEL_DISABLED`. | `string` | `null` | no |
| gcp\_public\_cidrs\_access\_enabled | Allow access through Google Cloud public IP addresses | `bool` | `null` | no |
| gke\_backup\_agent\_config | Whether Backup for GKE agent is enabled for this cluster. | `bool` | `false` | no |
| grant\_registry\_access | Grants created cluster-specific service account storage.objectViewer and artifactregistry.reader roles. | `bool` | `false` | no |
| horizontal\_pod\_autoscaling | Enable horizontal pod autoscaling addon | `bool` | `true` | no |
| hpa\_profile | Enable the Horizontal Pod Autoscaling profile for this cluster. Values are "NONE" and "PERFORMANCE". | `string` | `""` | no |
| http\_load\_balancing | Enable httpload balancer addon | `bool` | `true` | no |
| identity\_namespace | The workload pool to attach all Kubernetes service accounts to. (Default value of `enabled` automatically sets project-based pool `[project_id].svc.id.goog`) | `string` | `"enabled"` | no |
| in\_transit\_encryption\_config | Defines the config of in-transit encryption. Valid values are `IN_TRANSIT_ENCRYPTION_DISABLED` and `IN_TRANSIT_ENCRYPTION_INTER_NODE_TRANSPARENT`. | `string` | `null` | no |
| insecure\_kubelet\_readonly\_port\_enabled | Whether or not to set `insecure_kubelet_readonly_port_enabled` for node pool defaults and autopilot clusters. | `bool` | `null` | no |
| ip\_endpoints\_enabled | (Optional) Controls whether to allow direct IP access. Defaults to `true`. | `bool` | `null` | no |
| ip\_range\_pods | The _name_ of the secondary subnet ip range to use for pods | `string` | n/a | yes |
| ip\_range\_services | The _name_ of the secondary subnet range to use for services. If not provided, the default `34.118.224.0/20` range will be used. | `string` | `null` | no |
| issue\_client\_certificate | Issues a client certificate to authenticate to the cluster endpoint. To maximize the security of your cluster, leave this option disabled. Client certificates don't automatically rotate and aren't easily revocable. WARNING: changing this after cluster creation is destructive! | `bool` | `false` | no |
| kubernetes\_version | The Kubernetes version of the masters. If set to 'latest' it will pull latest available version in the selected region. | `string` | `"latest"` | no |
| logging\_enabled\_components | List of services to monitor: SYSTEM\_COMPONENTS, APISERVER, CONTROLLER\_MANAGER, KCP\_CONNECTION, KCP\_SSHD, KCP\_HPA, SCHEDULER, and WORKLOADS. Empty list is default GKE configuration. | `list(string)` | `[]` | no |
| maintenance\_end\_time | Time window specified for recurring maintenance operations in RFC3339 format | `string` | `""` | no |
| maintenance\_exclusions | List of maintenance exclusions. A cluster can have up to three | `list(object({ name = string, start_time = string, end_time = string, exclusion_scope = string }))` | `[]` | no |
| maintenance\_recurrence | Frequency of the recurring maintenance window in RFC5545 format. | `string` | `""` | no |
| maintenance\_start\_time | Time window specified for daily or recurring maintenance operations in RFC3339 format | `string` | `"05:00"` | no |
| master\_authorized\_networks | List of master authorized networks. If none are provided, disallow external access (except the cluster node IPs, which GKE automatically whitelists). | `list(object({ cidr_block = string, display_name = string }))` | `[]` | no |
| master\_global\_access\_enabled | Whether the cluster master is accessible globally (from any region) or only within the same region as the private endpoint. | `bool` | `true` | no |
| master\_ipv4\_cidr\_block | (Optional) The IP range in CIDR notation to use for the hosted master network. | `string` | `null` | no |
| monitoring\_enabled\_components | List of services to monitor: SYSTEM\_COMPONENTS, APISERVER, SCHEDULER, CONTROLLER\_MANAGER, STORAGE, HPA, POD, DAEMONSET, DEPLOYMENT, STATEFULSET, KUBELET, CADVISOR, DCGM, and JOBSET. In beta provider, WORKLOADS is supported on top of those 12 values. (WORKLOADS is deprecated and removed in GKE 1.24.) KUBELET and CADVISOR are only supported in GKE 1.29.3-gke.1093000 and above. JOBSET is only supported in GKE 1.32.1-gke.1357001 and above. Empty list is default GKE configuration. | `list(string)` | `[]` | no |
| monitoring\_metric\_writer\_role | The monitoring metrics writer role to assign to the GKE node service account | `string` | `"roles/monitoring.metricWriter"` | no |
| name | The name of the cluster (required) | `string` | n/a | yes |
| network | The VPC network to host the cluster in (required) | `string` | n/a | yes |
| network\_project\_id | The project ID of the shared VPC's host (for shared vpc support) | `string` | `""` | no |
| network\_tags | (Optional) - List of network tags applied to auto-provisioned node pools. | `list(string)` | `[]` | no |
| node\_pools\_cgroup\_mode | Specifies the Linux cgroup mode for autopilot Kubernetes nodes in the cluster. Accepted values are `CGROUP_MODE_UNSPECIFIED`, `CGROUP_MODE_V1`, and `CGROUP_MODE_V2`, which determine the control group hierarchy used for resource management. | `string` | `null` | no |
| notification\_config\_topic | The desired Pub/Sub topic to which notifications will be sent by GKE. Format is projects/{project}/topics/{topic}. | `string` | `""` | no |
| notification\_filter\_event\_type | Choose what type of notifications you want to receive. If no filters are applied, you'll receive all notification types. Can be used to filter what notifications are sent. Accepted values are UPGRADE\_AVAILABLE\_EVENT, UPGRADE\_EVENT, and SECURITY\_BULLETIN\_EVENT. | `list(string)` | `[]` | no |
| private\_endpoint\_subnetwork | The subnetwork to use for the hosted master network. | `string` | `null` | no |
| project\_id | The project ID to host the cluster in (required) | `string` | n/a | yes |
| ray\_operator\_config | The Ray Operator Addon configuration for this cluster. | <pre>object({<br>    enabled            = bool<br>    logging_enabled    = optional(bool, false)<br>    monitoring_enabled = optional(bool, false)<br>  })</pre> | <pre>{<br>  "enabled": false,<br>  "logging_enabled": false,<br>  "monitoring_enabled": false<br>}</pre> | no |
| region | The region to host the cluster in (optional if zonal cluster / required if regional) | `string` | `null` | no |
| regional | Whether is a regional cluster (zonal cluster if set false. WARNING: changing this after cluster creation is destructive!) | `bool` | `true` | no |
| registry\_project\_ids | Projects holding Google Container Registries. If empty, we use the cluster project. If a service account is created and the `grant_registry_access` variable is set to `true`, the `storage.objectViewer` and `artifactregsitry.reader` roles are assigned on these projects. | `list(string)` | `[]` | no |
| release\_channel | The release channel of this cluster. Accepted values are `UNSPECIFIED`, `RAPID`, `REGULAR` and `STABLE`. Defaults to `REGULAR`. | `string` | `"REGULAR"` | no |
| resource\_usage\_export\_dataset\_id | The ID of a BigQuery Dataset for using BigQuery as the destination of resource usage export. | `string` | `""` | no |
| security\_posture\_mode | Security posture mode. Accepted values are `DISABLED` and `BASIC`. Defaults to `DISABLED`. | `string` | `"DISABLED"` | no |
| security\_posture\_vulnerability\_mode | Security posture vulnerability mode. Accepted values are `VULNERABILITY_DISABLED`, `VULNERABILITY_BASIC`, and `VULNERABILITY_ENTERPRISE`. Defaults to `VULNERABILITY_DISABLED`. | `string` | `"VULNERABILITY_DISABLED"` | no |
| service\_account | The service account to run nodes as if not overridden in `node_pools`. The create\_service\_account variable default value (true) will cause a cluster-specific service account to be created. This service account should already exists and it will be used by the node pools. If you wish to only override the service account name, you can use service\_account\_name variable. | `string` | `""` | no |
| service\_account\_name | The name of the service account that will be created if create\_service\_account is true. If you wish to use an existing service account, use service\_account variable. | `string` | `""` | no |
| service\_external\_ips | Whether external ips specified by a service will be allowed in this cluster | `bool` | `false` | no |
| shadow\_firewall\_rules\_log\_config | The log\_config for shadow firewall rules. You can set this variable to `null` to disable logging. | <pre>object({<br>    metadata = string<br>  })</pre> | <pre>{<br>  "metadata": "INCLUDE_ALL_METADATA"<br>}</pre> | no |
| shadow\_firewall\_rules\_priority | The firewall priority of GKE shadow firewall rules. The priority should be less than default firewall, which is 1000. | `number` | `999` | no |
| stack\_type | The stack type to use for this cluster. Either `IPV4` or `IPV4_IPV6`. Defaults to `IPV4`. | `string` | `"IPV4"` | no |
| stateful\_ha | Whether the Stateful HA Addon is enabled for this cluster. | `bool` | `false` | no |
| subnetwork | The subnetwork to host the cluster in (required) | `string` | n/a | yes |
| timeouts | Timeout for cluster operations. | `map(string)` | `{}` | no |
| workload\_config\_audit\_mode | (beta) Sets which mode of auditing should be used for the cluster's workloads. Accepted values are DISABLED, BASIC. | `string` | `"DISABLED"` | no |
| workload\_vulnerability\_mode | (beta) Sets which mode to use for Protect workload vulnerability scanning feature. Accepted values are DISABLED, BASIC. | `string` | `""` | no |
| zones | The zones to host the cluster in (optional if regional cluster / required if zonal) | `list(string)` | `[]` | no |

## Outputs

| Name | Description |
|------|-------------|
| ca\_certificate | Cluster ca certificate (base64 encoded) |
| cloudrun\_enabled | Whether CloudRun enabled |
| cluster\_id | Cluster ID |
| dns\_cache\_enabled | Whether DNS Cache enabled |
| endpoint | Cluster endpoint |
| endpoint\_dns | Cluster endpoint DNS |
| fleet\_membership | Fleet membership (if registered) |
| gateway\_api\_channel | The gateway api channel of this cluster. |
| horizontal\_pod\_autoscaling\_enabled | Whether horizontal pod autoscaling enabled |
| http\_load\_balancing\_enabled | Whether http load balancing enabled |
| identity\_namespace | Workload Identity pool |
| identity\_service\_enabled | Whether Identity Service is enabled |
| intranode\_visibility\_enabled | Whether intra-node visibility is enabled |
| istio\_enabled | Whether Istio is enabled |
| location | Cluster location (region if regional cluster, zone if zonal cluster) |
| logging\_service | Logging service used |
| master\_authorized\_networks\_config | Networks from which access to master is permitted |
| master\_ipv4\_cidr\_block | The IP range in CIDR notation used for the hosted master network |
| master\_version | Current master kubernetes version |
| min\_master\_version | Minimum master kubernetes version |
| monitoring\_service | Monitoring service used |
| name | Cluster name |
| peering\_name | The name of the peering between this cluster and the Google owned VPC. |
| pod\_security\_policy\_enabled | Whether pod security policy is enabled |
| region | Cluster region |
| release\_channel | The release channel of this cluster |
| secret\_manager\_addon\_enabled | Whether Secret Manager add-on is enabled |
| service\_account | The service account to default running nodes as if not overridden in `node_pools`. |
| tpu\_ipv4\_cidr\_block | The IP range in CIDR notation used for the TPUs |
| type | Cluster type (regional / zonal) |
| vertical\_pod\_autoscaling\_enabled | Whether vertical pod autoscaling enabled |
| zones | List of zones in which the cluster resides |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->


## Requirements

Before this module can be used on a project, you must ensure that the following pre-requisites are fulfilled:

1. Terraform and kubectl are [installed](#software-dependencies) on the machine where Terraform is executed.
2. The Service Account you execute the module with has the right [permissions](#configure-a-service-account).
3. The Compute Engine and Kubernetes Engine APIs are [active](#enable-apis) on the project you will launch the cluster in.
4. If you are using a Shared VPC, the APIs must also be activated on the Shared VPC host project and your service account needs the proper permissions there.

The [project factory](https://github.com/terraform-google-modules/terraform-google-project-factory) can be used to provision projects with the correct APIs active and the necessary Shared VPC connections.

### Software Dependencies

#### Kubectl

- [kubectl](https://github.com/kubernetes/kubernetes/releases) 1.9.x

#### Terraform and Plugins

- [Terraform](https://www.terraform.io/downloads.html) 1.3+
- [Terraform Provider for GCP Beta][terraform-provider-google-beta] v6.41+

#### gcloud

Some submodules use the [terraform-google-gcloud](https://github.com/terraform-google-modules/terraform-google-gcloud) module. By default, this module assumes you already have gcloud installed in your $PATH.
See the [module](https://github.com/terraform-google-modules/terraform-google-gcloud#downloading) documentation for more information.

### Configure a Service Account

In order to execute this module you must have a Service Account with the
following project roles:

- roles/compute.viewer
- roles/compute.securityAdmin (only required if `add_cluster_firewall_rules` is set to `true`)
- roles/container.clusterAdmin
- roles/container.developer
- roles/iam.serviceAccountAdmin
- roles/iam.serviceAccountUser
- roles/resourcemanager.projectIamAdmin (only required if `service_account` is set to `create`)

Additionally, if `service_account` is set to `create` and `grant_registry_access` is requested, the service account requires the following role on the `registry_project_ids` projects:

- roles/resourcemanager.projectIamAdmin

### Enable APIs

In order to operate with the Service Account you must activate the following APIs on the project where the Service Account was created:

- Compute Engine API - compute.googleapis.com
- Kubernetes Engine API - container.googleapis.com

[terraform-provider-google-beta]: <https://github.com/terraform-providers/terraform-provider-google-beta>
[12.3.0]: <https://registry.terraform.io/modules/terraform-google-modules/kubernetes-engine/google/12.3.0>
[terraform-0.13-upgrade]: <https://www.terraform.io/upgrade-guides/0-13.html>
[terraform-1.3-upgrade]: <https://developer.hashicorp.com/terraform/language/v1.3.x/upgrade-guides>
