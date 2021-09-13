package fakes

import "github.com/orange-cloudfoundry/bosh-bootloader/storage"

type StateBootstrap struct {
	GetStateCall struct {
		CallCount int
		Returns   struct {
			State storage.State
			Error error
		}
		Receives struct {
			Dir string
		}
	}
}

func (s *StateBootstrap) GetState(dir string) (storage.State, error) {
	s.GetStateCall.CallCount++
	s.GetStateCall.Receives.Dir = dir

	return s.GetStateCall.Returns.State, s.GetStateCall.Returns.Error
}
