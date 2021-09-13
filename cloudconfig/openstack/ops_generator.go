package openstack

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"github.com/orange-cloudfoundry/bosh-bootloader/storage"
	"github.com/orange-cloudfoundry/bosh-bootloader/terraform"
)

type OpsGenerator struct {
	terraformManager terraformManager
}

type terraformManager interface {
	GetOutputs() (terraform.Outputs, error)
}

func NewOpsGenerator(terraformManager terraformManager) OpsGenerator {
	return OpsGenerator{
		terraformManager: terraformManager,
	}
}

func (o OpsGenerator) Generate(state storage.State) (string, error) {
	return BaseOps, nil
}

func (o OpsGenerator) GenerateVars(state storage.State) (string, error) {
	terraformOutputs, err := o.terraformManager.GetOutputs()
	if err != nil {
		return "", fmt.Errorf("Get terraform outputs: %s", err)
	}

	varsBytes, err := yaml.Marshal(terraformOutputs.Map)
	if err != nil {
		return "", fmt.Errorf("Unmarshalling terraform outputs: %s", err)
	}

	return string(varsBytes), nil
}
