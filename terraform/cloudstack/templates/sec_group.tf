variable "cf_external_ports" {
  type = "list"
  default = [
    80,
    443,
    2222,
    4443]
}

variable "bosh_ports" {
  type = "list"
  default = [
    22,
    6868,
    2555,
    4222,
    25250]
}

variable "bosh_external_ports" {
  type = "list"
  default = [
    22,
    6868,
    25555,
    3389,
    8844,
    8443]
}

variable "shared_tcp_ports" {
  type = "list"
  default = [
    9090,
    9091,
    8082,
    8300,
    8301,
    8889,
    8443,
    3000,
    4443,
    8080,
    3457,
    9023,
    9022,
    4222,
    8891,
    4003,
    4103,
    1801]
}

variable "shared_udp_ports" {
  type = "list"
  default = [
    8301,
    8302,
    8600]
}

resource "cloudstack_network_acl" "bosh_subnet_sec_group" {
  count = "${var.secure ? 1 : 0}"
  name = "bosh_subnet"
  vpc_id = "${cloudstack_vpc.vpc.id}"
}


resource "cloudstack_network_acl_rule" "bosh_subnet_sec_group_egress" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.bosh_subnet_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "0.0.0.0/0"]
    protocol = "all"
    ports = []
    traffic_type = "egress"
  }
}

resource "cloudstack_network_acl_rule" "bosh_subnet_sec_group_udp" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.bosh_subnet_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "0.0.0.0/0"]
    protocol = "udp"
    ports = []
    traffic_type = "ingress"
  }
}

resource "cloudstack_network_acl_rule" "bosh_subnet_sec_group_external" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.bosh_subnet_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "0.0.0.0/0"]
    protocol = "tcp"
    ports = "${var.bosh_external_ports}"
    traffic_type = "ingress"
  }
}

resource "cloudstack_network_acl_rule" "bosh_subnet_sec_group_icmp" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.bosh_subnet_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "0.0.0.0/0"]
    protocol = "icmp"
    icmp_type = -1
    icmp_code = -1
    ports = []
    traffic_type = "ingress"
  }
}


resource "cloudstack_network_acl" "control_plane_sec_group" {
  count = "${var.secure ? 1 : 0}"
  name = "control_plane"
  vpc_id = "${cloudstack_vpc.vpc.id}"
}


resource "cloudstack_network_acl_rule" "control_plane_sec_group_egress" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.control_plane_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "0.0.0.0/0"]
    protocol = "all"
    ports = []
    traffic_type = "egress"
  }
}

resource "cloudstack_network_acl_rule" "control_plane_sec_group_bosh" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.control_plane_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "${var.vpc_cidr}"]
    protocol = "tcp"
    ports = "${var.bosh_ports}"
    traffic_type = "ingress"
  }
}

resource "cloudstack_network_acl_rule" "control_plane_sec_group_shared_tcp_ports" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.control_plane_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "${var.vpc_cidr}"]
    protocol = "tcp"
    ports = "${var.shared_tcp_ports}"
    traffic_type = "ingress"
  }
}

resource "cloudstack_network_acl_rule" "control_plane_sec_group_shared_udp_ports" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.control_plane_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "${var.vpc_cidr}"]
    protocol = "udp"
    ports = "${var.shared_udp_ports}"
    traffic_type = "ingress"
  }
}

resource "cloudstack_network_acl_rule" "control_plane_sec_group_icmp" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.control_plane_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "0.0.0.0/0"]
    protocol = "icmp"
    icmp_type = -1
    icmp_code = -1
    ports = []
    traffic_type = "ingress"
  }
}

resource "cloudstack_network_acl" "data_plane_sec_group" {
  count = "${var.secure ? 1 : 0}"
  name = "data_plane"
  vpc_id = "${cloudstack_vpc.vpc.id}"
}


resource "cloudstack_network_acl_rule" "data_plane_sec_group_egress" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.data_plane_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "0.0.0.0/0"]
    protocol = "all"
    ports = []
    traffic_type = "egress"
  }
}

resource "cloudstack_network_acl_rule" "data_plane_sec_group_bosh" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.data_plane_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "${var.vpc_cidr}"]
    protocol = "tcp"
    ports = "${var.bosh_ports}"
    traffic_type = "ingress"
  }
}

resource "cloudstack_network_acl_rule" "data_plane_sec_group_shared_tcp_ports" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.data_plane_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "${var.vpc_cidr}"]
    protocol = "tcp"
    ports = "${var.shared_tcp_ports}"
    traffic_type = "ingress"
  }
}

resource "cloudstack_network_acl_rule" "data_plane_sec_group_shared_udp_ports" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.data_plane_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "${var.vpc_cidr}"]
    protocol = "udp"
    ports = "${var.shared_udp_ports}"
    traffic_type = "ingress"
  }
}

resource "cloudstack_network_acl_rule" "data_plane_sec_group_external" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.data_plane_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "0.0.0.0/0"]
    protocol = "all"
    ports = "${var.cf_external_ports}"
    traffic_type = "ingress"
  }
}

resource "cloudstack_network_acl_rule" "data_plane_sec_group_icmp" {
  count = "${var.secure ? 1 : 0}"
  acl_id = "${cloudstack_network_acl.data_plane_sec_group.id}"

  rule {
    action = "allow"
    cidr_list = [
      "0.0.0.0/0"]
    protocol = "icmp"
    icmp_type = -1
    icmp_code = -1
    ports = []
    traffic_type = "ingress"
  }
}