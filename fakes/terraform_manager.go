package fakes

import (
	"github.com/orange-cloudfoundry/bosh-bootloader/storage"
	"github.com/orange-cloudfoundry/bosh-bootloader/terraform"
)

type TerraformManager struct {
	InitCall struct {
		CallCount int
		Receives  struct {
			BBLState storage.State
		}
		Returns struct {
			Error error
		}
	}
	SetupCall struct {
		CallCount int
		Receives  struct {
			BBLState storage.State
		}
		Returns struct {
			Error error
		}
	}
	ApplyCall struct {
		CallCount int
		Receives  struct {
			BBLState storage.State
		}
		Returns struct {
			BBLState storage.State
			Error    error
		}
	}
	DestroyCall struct {
		CallCount int
		Receives  struct {
			BBLState storage.State
		}
		Returns struct {
			BBLState storage.State
			Error    error
		}
	}
	ValidateCall struct {
		CallCount int
		Receives  struct {
			BBLState storage.State
		}
		Returns struct {
			BBLState storage.State
			Error    error
		}
	}
	ImportCall struct {
		CallCount int
		Receives  struct {
			BBLState storage.State
			Outputs  map[string]string
		}
		Returns struct {
			BBLState storage.State
			Error    error
		}
	}
	GetOutputsCall struct {
		CallCount int
		Returns   struct {
			Outputs terraform.Outputs
			Error   error
		}
	}
	VersionCall struct {
		CallCount int
		Returns   struct {
			Version string
			Error   error
		}
	}
	ValidateVersionCall struct {
		CallCount int
		Returns   struct {
			Error error
		}
	}
	IsPavedCall struct {
		CallCount int
		Returns   struct {
			IsPaved bool
			Error   error
		}
	}
}

func (t *TerraformManager) Setup(bblState storage.State) error {
	t.SetupCall.CallCount++
	t.SetupCall.Receives.BBLState = bblState

	return t.SetupCall.Returns.Error
}

func (t *TerraformManager) Init(bblState storage.State) error {
	t.InitCall.CallCount++
	t.InitCall.Receives.BBLState = bblState

	return t.InitCall.Returns.Error
}

func (t *TerraformManager) Apply(bblState storage.State) (storage.State, error) {
	t.ApplyCall.CallCount++
	t.ApplyCall.Receives.BBLState = bblState

	return t.ApplyCall.Returns.BBLState, t.ApplyCall.Returns.Error
}

func (t *TerraformManager) Destroy(bblState storage.State) (storage.State, error) {
	t.DestroyCall.CallCount++
	t.DestroyCall.Receives.BBLState = bblState

	return t.DestroyCall.Returns.BBLState, t.DestroyCall.Returns.Error
}

func (t *TerraformManager) Validate(bblState storage.State) (storage.State, error) {
	t.ValidateCall.CallCount++
	t.ValidateCall.Receives.BBLState = bblState

	return t.ValidateCall.Returns.BBLState, t.ValidateCall.Returns.Error
}

func (t *TerraformManager) Import(bblState storage.State, outputs map[string]string) (storage.State, error) {
	t.ImportCall.CallCount++
	t.ImportCall.Receives.BBLState = bblState
	t.ImportCall.Receives.Outputs = outputs

	return t.ImportCall.Returns.BBLState, t.ImportCall.Returns.Error
}

func (t *TerraformManager) GetOutputs() (terraform.Outputs, error) {
	t.GetOutputsCall.CallCount++
	return t.GetOutputsCall.Returns.Outputs, t.GetOutputsCall.Returns.Error
}

func (t *TerraformManager) Version() (string, error) {
	t.VersionCall.CallCount++
	return t.VersionCall.Returns.Version, t.VersionCall.Returns.Error
}

func (t *TerraformManager) ValidateVersion() error {
	t.ValidateVersionCall.CallCount++
	return t.ValidateVersionCall.Returns.Error
}

func (t *TerraformManager) IsPaved() (bool, error) {
	t.IsPavedCall.CallCount++
	return t.IsPavedCall.Returns.IsPaved, t.IsPavedCall.Returns.Error
}
