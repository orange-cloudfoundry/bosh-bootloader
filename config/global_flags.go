package config

type GlobalFlags struct {
	Help        bool   `short:"h" long:"help"`
	Debug       bool   `short:"d" long:"debug"        env:"BBL_DEBUG"`
	Version     bool   `short:"v" long:"version"`
	NoConfirm   bool   `short:"n" long:"no-confirm"`
	StateDir    string `short:"s" long:"state-dir"    env:"BBL_STATE_DIRECTORY"`
	StateBucket string `          long:"state-bucket" env:"BBL_STATE_BUCKET"`
	EnvID       string `          long:"name"`
	IAAS        string `          long:"iaas"         env:"BBL_IAAS"`

	AWSAccessKeyID     string `long:"aws-access-key-id"       env:"BBL_AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `long:"aws-secret-access-key"   env:"BBL_AWS_SECRET_ACCESS_KEY"`
	AWSRegion          string `long:"aws-region"              env:"BBL_AWS_REGION"`

	AzureClientID       string `long:"azure-client-id"        env:"BBL_AZURE_CLIENT_ID"`
	AzureClientSecret   string `long:"azure-client-secret"    env:"BBL_AZURE_CLIENT_SECRET"`
	AzureRegion         string `long:"azure-region"           env:"BBL_AZURE_REGION"`
	AzureSubscriptionID string `long:"azure-subscription-id"  env:"BBL_AZURE_SUBSCRIPTION_ID"`
	AzureTenantID       string `long:"azure-tenant-id"        env:"BBL_AZURE_TENANT_ID"`

	GCPServiceAccountKey string `long:"gcp-service-account-key" env:"BBL_GCP_SERVICE_ACCOUNT_KEY"`
	GCPRegion            string `long:"gcp-region"              env:"BBL_GCP_REGION"`

	VSphereNetwork          string `long:"vsphere-network"            env:"BBL_VSPHERE_NETWORK"`
	VSphereSubnetCIDR       string `long:"vsphere-subnet-cidr"        env:"BBL_VSPHERE_SUBNET_CIDR"`
	VSphereVCenterCluster   string `long:"vsphere-vcenter-cluster"    env:"BBL_VSPHERE_VCENTER_CLUSTER"`
	VSphereVCenterDC        string `long:"vsphere-vcenter-dc"         env:"BBL_VSPHERE_VCENTER_DC"`
	VSphereVCenterDS        string `long:"vsphere-vcenter-ds"         env:"BBL_VSPHERE_VCENTER_DS"`
	VSphereVCenterIP        string `long:"vsphere-vcenter-ip"         env:"BBL_VSPHERE_VCENTER_IP"`
	VSphereVCenterPassword  string `long:"vsphere-vcenter-password"   env:"BBL_VSPHERE_VCENTER_PASSWORD"`
	VSphereVCenterRP        string `long:"vsphere-vcenter-rp"         env:"BBL_VSPHERE_VCENTER_RP"`
	VSphereVCenterUser      string `long:"vsphere-vcenter-user"       env:"BBL_VSPHERE_VCENTER_USER"`
	VSphereVCenterDisks     string `long:"vsphere-vcenter-disks"      env:"BBL_VSPHERE_VCENTER_DISKS"`
	VSphereVCenterVMs       string `long:"vsphere-vcenter-vms"        env:"BBL_VSPHERE_VCENTER_VMS"`
	VSphereVCenterTemplates string `long:"vsphere-vcenter-templates"  env:"BBL_VSPHERE_VCENTER_TEMPLATES"`

	OpenStackInternalCidr         string `long:"openstack-internal-cidr"          env:"BBL_OPENSTACK_INTERNAL_CIDR"`
	OpenStackExternalIP           string `long:"openstack-external-ip"            env:"BBL_OPENSTACK_EXTERNAL_IP"`
	OpenStackAuthURL              string `long:"openstack-auth-url"               env:"BBL_OPENSTACK_AUTH_URL"`
	OpenStackAZ                   string `long:"openstack-az"                     env:"BBL_OPENSTACK_AZ"`
	OpenStackDefaultKeyName       string `long:"openstack-default-key-name"       env:"BBL_OPENSTACK_DEFAULT_KEY_NAME"`
	OpenStackDefaultSecurityGroup string `long:"openstack-default-security-group" env:"BBL_OPENSTACK_DEFAULT_SECURITY_GROUP"`
	OpenStackNetworkID            string `long:"openstack-network-id"             env:"BBL_OPENSTACK_NETWORK_ID"`
	OpenStackPassword             string `long:"openstack-password"               env:"BBL_OPENSTACK_PASSWORD"`
	OpenStackUsername             string `long:"openstack-username"               env:"BBL_OPENSTACK_USERNAME"`
	OpenStackProject              string `long:"openstack-project"                env:"BBL_OPENSTACK_PROJECT"`
	OpenStackDomain               string `long:"openstack-domain"                 env:"BBL_OPENSTACK_DOMAIN"`
	OpenStackRegion               string `long:"openstack-region"                 env:"BBL_OPENSTACK_REGION"`
	OpenStackPrivateKey           string `long:"openstack-private-key"            env:"BBL_OPENSTACK_PRIVATE_KEY"`

	CloudStackEndpoint           string `long:"cloudstack-endpoint"               env:"BBL_CLOUDSTACK_ENDPOINT"`
	CloudStackApiKey             string `long:"cloudstack-api-key"                env:"BBL_CLOUDSTACK_API_KEY"`
	CloudStackSecretAccessKey    string `long:"cloudstack-secret-access-key"      env:"BBL_CLOUDSTACK_SECRET_ACCESS_KEY"`
	CloudStackZone               string `long:"cloudstack-zone"                   env:"BBL_CLOUDSTACK_ZONE"`
	CloudStackNetworkVpcOffering string `long:"cloudstack-network-vpc-offering"   env:"BBL_CLOUDSTACK_NETWORK_VPC_OFFERING"`
	CloudStackComputeOffering    string `long:"cloudstack-compute-offering"       env:"BBL_CLOUDSTACK_COMPUTE_OFFERING"`
	CloudStackSecure             bool   `long:"cloudstack-secure"                 env:"BBL_CLOUDSTACK_SECURE"`
	CloudStackIsoSegment         bool   `long:"cloudstack-iso-segment"            env:"BBL_CLOUDSTACK_ISO_SEGMENT"`
}
