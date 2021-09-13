package fakes

import (
	"github.com/orange-cloudfoundry/bosh-bootloader/commands"
	"github.com/orange-cloudfoundry/bosh-bootloader/storage"
)

type LBArgsHandler struct {
	GetLBStateCall struct {
		CallCount int
		Returns   struct {
			LB    storage.LB
			Error error
		}
		Receives struct {
			IAAS string
			Args commands.LBArgs
		}
	}
	MergeCall struct {
		CallCount int
		Returns   struct {
			LB storage.LB
		}
		Receives struct {
			New storage.LB
			Old storage.LB
		}
	}
}

func (c *LBArgsHandler) GetLBState(iaas string, args commands.LBArgs) (storage.LB, error) {
	c.GetLBStateCall.CallCount++
	c.GetLBStateCall.Receives.Args = args
	c.GetLBStateCall.Receives.IAAS = iaas
	return c.GetLBStateCall.Returns.LB, c.GetLBStateCall.Returns.Error
}

func (c *LBArgsHandler) Merge(new storage.LB, old storage.LB) storage.LB {
	c.MergeCall.CallCount++
	c.MergeCall.Receives.New = new
	c.MergeCall.Receives.Old = old
	return c.MergeCall.Returns.LB
}
