package cloudstack

import (
	"github.com/cloudfoundry/bosh-bootloader/storage"
	"strings"
)

type TemplateGenerator struct{}

func NewTemplateGenerator() TemplateGenerator {
	return TemplateGenerator{}
}

func (t TemplateGenerator) Generate(state storage.State) string {
	tf := []string{
		string(MustAsset("templates/base.tf")),
		string(MustAsset("templates/sec_group.tf")),
	}
	return strings.Join(tf, "\n")
}
