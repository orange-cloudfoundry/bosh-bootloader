variable "cloudstack_zone" {
  type = string
}

variable "cloudstack_endpoint" {
  type = string
}

variable "cloudstack_api_key" {
  type = string
}

variable "cloudstack_secret_access_key" {
  type = string
}

variable "network_vpc_offering" {
  type    = string
  default = "DefaultIsolatedNetworkOfferingForVpcNetworks"
}

variable "vpc_offering" {
  type    = string
  default = "Default VPC offering"
}

variable "cloudstack_compute_offering" {
  type    = string
  default = "shared.large"
}

variable "env_id" {
  type = string
}

variable "short_env_id" {
  type = string
}

variable "vpc_cidr" {
  type    = string
  default = "10.0.0.0/16"
}

variable "dns" {
  type = list(string)
  default = [
    "8.8.8.8",
  ]
}

variable "secure" {
  default = false
}

variable "iso_segment" {
  default = false
}

locals {
  director_name        = "bosh-${var.env_id}"
  internal_cidr        = cidrsubnet(var.vpc_cidr, 8, 0)
  internal_gw          = cidrhost(local.internal_cidr, 1)
  jumpbox_internal_ip  = cidrhost(local.internal_cidr, 5)
  director_internal_ip = cidrhost(local.internal_cidr, 6)
}

provider "cloudstack" {
  api_url    = var.cloudstack_endpoint
  api_key    = var.cloudstack_api_key
  secret_key = var.cloudstack_secret_access_key
  timeout    = "900"
}

resource "tls_private_key" "bosh_vms" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "cloudstack_ssh_keypair" "bosh_vms" {
  name       = "${var.env_id}_bosh_vms"
  public_key = tls_private_key.bosh_vms.public_key_openssh
}

resource "cloudstack_vpc" "vpc" {
  name         = "${var.env_id}-vpc"
  cidr         = var.vpc_cidr
  vpc_offering = var.vpc_offering
  zone         = var.cloudstack_zone
}

resource "cloudstack_network_acl" "allow_all" {
  name   = "allow_all"
  vpc_id = cloudstack_vpc.vpc.id
}

resource "cloudstack_network_acl_rule" "ingress_allow_all" {
  acl_id = cloudstack_network_acl.allow_all.id

  rule {
    action = "allow"
    cidr_list = [
      "0.0.0.0/0",
    ]
    protocol     = "all"
    ports        = []
    traffic_type = "ingress"
  }
}

resource "cloudstack_network_acl_rule" "egress_allow_all" {
  acl_id = cloudstack_network_acl.allow_all.id

  rule {
    action = "allow"
    cidr_list = [
      "0.0.0.0/0",
    ]
    protocol     = "all"
    ports        = []
    traffic_type = "egress"
  }
}

resource "cloudstack_network" "bosh_subnet" {
  cidr             = cidrsubnet(var.vpc_cidr, 8, 0)
  name             = "${var.short_env_id}-bosh-subnet"
  vpc_id           = cloudstack_vpc.vpc.id
  display_text     = "bosh-subnet"
  network_offering = var.network_vpc_offering
  zone             = var.cloudstack_zone
  acl_id = var.secure ? element(
    concat(cloudstack_network_acl.bosh_subnet_sec_group.*.id, [""]),
    0,
  ) : cloudstack_network_acl.allow_all.id
}

resource "cloudstack_network" "compilation_subnet" {
  cidr             = cidrsubnet(var.vpc_cidr, 8, 1)
  name             = "${var.short_env_id}-compilation-subnet"
  vpc_id           = cloudstack_vpc.vpc.id
  display_text     = "compilation-subnet"
  network_offering = var.network_vpc_offering
  zone             = var.cloudstack_zone
  acl_id = var.secure ? element(
    concat(cloudstack_network_acl.bosh_subnet_sec_group.*.id, [""]),
    0,
  ) : cloudstack_network_acl.allow_all.id
}

resource "cloudstack_network" "control_plane" {
  cidr             = cidrsubnet(var.vpc_cidr, 6, 1)
  name             = "${var.short_env_id}-control-plane"
  vpc_id           = cloudstack_vpc.vpc.id
  display_text     = "control-plane"
  network_offering = var.network_vpc_offering
  zone             = var.cloudstack_zone
  acl_id = var.secure ? element(
    concat(cloudstack_network_acl.control_plane_sec_group.*.id, [""]),
    0,
  ) : cloudstack_network_acl.allow_all.id
}

resource "cloudstack_network" "data_plane" {
  cidr             = cidrsubnet(var.vpc_cidr, 6, 2)
  name             = "${var.short_env_id}-data-plane"
  vpc_id           = cloudstack_vpc.vpc.id
  display_text     = "data-plane"
  network_offering = var.network_vpc_offering
  zone             = var.cloudstack_zone
  acl_id = var.secure ? element(
    concat(cloudstack_network_acl.data_plane_sec_group.*.id, [""]),
    0,
  ) : cloudstack_network_acl.allow_all.id
}

resource "cloudstack_network" "data_plane_public" {
  count            = var.iso_segment ? 1 : 0
  cidr             = cidrsubnet(var.vpc_cidr, 6, 3)
  name             = "${var.short_env_id}-data-plane-public"
  vpc_id           = cloudstack_vpc.vpc.id
  display_text     = "data-plane-public"
  network_offering = var.network_vpc_offering
  zone             = var.cloudstack_zone
  acl_id = var.secure ? element(
    concat(cloudstack_network_acl.data_plane_sec_group.*.id, [""]),
    0,
  ) : cloudstack_network_acl.allow_all.id
}

resource "cloudstack_ipaddress" "jumpbox_eip" {
  vpc_id = cloudstack_vpc.vpc.id
  zone   = var.cloudstack_zone
}

output "cloudstack_default_key_name" {
  value = "${var.env_id}_bosh_vms"
}

output "private_key" {
  value     = tls_private_key.bosh_vms.private_key_pem
  sensitive = true
}

output "external_ip" {
  value = cloudstack_ipaddress.jumpbox_eip.ip_address
}

output "jumpbox_url" {
  value = "${cloudstack_ipaddress.jumpbox_eip.ip_address}:22"
}

output "director_address" {
  value = "https://${cloudstack_ipaddress.jumpbox_eip.ip_address}:25555"
}

output "network_name" {
  value = "${var.short_env_id}-bosh-subnet"
}

output "director_name" {
  value = local.director_name
}

output "internal_cidr" {
  value = local.internal_cidr
}

output "internal_gw" {
  value = local.internal_gw
}

output "jumpbox__internal_ip" {
  value = local.jumpbox_internal_ip
}

output "director__internal_ip" {
  value = local.director_internal_ip
}

output "cloudstack_compute_offering" {
  value = var.cloudstack_compute_offering
}

output "cloudstack_zone" {
  value = var.cloudstack_zone
}

output "cloudstack_endpoint" {
  value = var.cloudstack_endpoint
}

output "dns" {
  value = [
    var.dns,
  ]
}

output "internal_subnet_cidr_mapping" {
  value = {
    cloudstack_network.data_plane.name                                    = cloudstack_network.data_plane.cidr
    cloudstack_network.control_plane.name                                 = cloudstack_network.control_plane.cidr
    cloudstack_network.bosh_subnet.name                                   = cloudstack_network.bosh_subnet.cidr
    element(concat(cloudstack_network.data_plane_public.*.name, [""]), 0) = element(concat(cloudstack_network.data_plane_public.*.cidr, [""]), 0)
    cloudstack_network.compilation_subnet.name                            = cloudstack_network.compilation_subnet.cidr
  }
}

output "internal_subnet_gw_mapping" {
  value = {
    cloudstack_network.data_plane.name    = cloudstack_network.data_plane.gateway
    cloudstack_network.control_plane.name = cloudstack_network.control_plane.gateway
    cloudstack_network.bosh_subnet.name   = cloudstack_network.bosh_subnet.gateway
    element(concat(cloudstack_network.data_plane_public.*.name, [""]), 0) = element(
      concat(cloudstack_network.data_plane_public.*.gateway, [""]),
      0,
    )
    cloudstack_network.compilation_subnet.name = cloudstack_network.compilation_subnet.gateway
  }
}

