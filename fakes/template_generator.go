package fakes

import "github.com/orange-cloudfoundry/bosh-bootloader/storage"

type TemplateGenerator struct {
	GenerateCall struct {
		CallCount int
		Receives  struct {
			State storage.State
		}
		Returns struct {
			Template string
		}
	}
}

func (t *TemplateGenerator) Generate(state storage.State) string {
	t.GenerateCall.CallCount++
	t.GenerateCall.Receives.State = state
	return t.GenerateCall.Returns.Template
}
