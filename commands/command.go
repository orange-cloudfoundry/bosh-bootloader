package commands

import "github.com/orange-cloudfoundry/bosh-bootloader/storage"

type Command interface {
	CheckFastFails(subcommandFlags []string, state storage.State) error
	Execute(subcommandFlags []string, state storage.State) error
	Usage() string
}
