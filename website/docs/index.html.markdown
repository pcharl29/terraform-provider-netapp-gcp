---
layout: "netapp_gcp"
page_title: "Provider: NetApp_GCP"
sidebar_current: "docs-netapp-gcp-index"
description: |-
  The NetApp_GCP provider is used to interact with the NetApp Cloud Volumes Service for Google Cloud resources. The provider needs to be configured with the proper credentials
  before it can be used.
---

# NetApp_GCP Provider

The NetApp_GCP provider is used to interact with the NetApp Cloud Volumes Service for Google Cloud resources.
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

~> **NOTE:** The NetApp_GCP Provider currently represents _initial support_
and therefore may undergo significant changes as the community improves it.

## Example Usage
Please see our [example][tf-example].
[tf-example]: https://github.com/NetApp/terraform-provider-netapp-gcp/tree/master/examples/gcp

```
# Install provider from terraform register, only available for version 0.13 and above.
terraform {
  required_providers {
    netapp-gcp = {
      source = "NetApp/netapp-gcp"
      version = "0.1.1"
    }
  }
}

# Configure the NetApp_GCP Provider
provider "netapp-gcp" {
  project         = "YOUR_PROJECT_NUMBER"
  service_account = "YOUR_SERVICE_ACCOUNT"
}

# Create a volume tied to an account
resource "netapp-gcp_volume" "gcp-volume" {
  provider = netapp-gcp
  name = "deleteme_asapGO_jusitin"
  region = "us-west2"
  protocol_types = ["NFSv3"]
  network = "cvs-terraform-vpc"
  size = 1024
  service_level = "premium"
  volume_path = "deleteme-asapGO"
  snapshot_policy {
    enabled = true
    daily_schedule {
      hour = 10
      minute = 1
    }
  }
  export_policy {
    rule {
      allowed_clients = "0.0.0.0/0"
      access= "ReadWrite"
      nfsv3 {
        checked =  true
      }
      nfsv4 {
        checked = false
      }
    }
    rule {
      allowed_clients= "10.0.0.0"
      access= "ReadWrite"
      nfsv3 {
        checked =  true
      }
      nfsv4 {
        checked = false
      }
    }
  }
}

# Create a snapshot tied to an volume
resource "netapp-gcp_snapshot" "gcp-snapshot" {
  name = "deleteme_asapGO_jusitin-snapshot"
  region = "us-west2"
  volume_name =  "deleteme_asapGO_jusitin"
  # creation_token = "unique-token-number"
}

# Create an Active Directory
resource "netapp-gcp_active_directory" "gcp-active-directory" {
  provider = netapp-gcp
  region = "us-west2"
	username = "test_user"
	password = "netapp"
  domain = "example.com"
  dns_server = "10.0.0.0"
  net_bios = "cvsserver"
}
```

## Argument Reference

The following arguments are used to configure the NetApp_GCP Provider:

* `project` - (Required) This is the project number for NetApp_GCP API operations.
* `service_account` - (Required) This is the path of service_account for NetApp_GCP API operations.

## Required Privileges

These settings were tested with GCP Google Cloud SDK 274.0.0.
For additional information on roles and permissions, please refer to official
NetApp_GCP documentation.

