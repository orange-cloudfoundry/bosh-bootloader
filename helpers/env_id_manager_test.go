package helpers_test

import (
	"errors"

	"github.com/orange-cloudfoundry/bosh-bootloader/fakes"
	"github.com/orange-cloudfoundry/bosh-bootloader/helpers"
	"github.com/orange-cloudfoundry/bosh-bootloader/storage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EnvIDManager", func() {
	var (
		envIDGenerator *fakes.EnvIDGenerator
		networkClient  *fakes.NetworkClient
		envIDManager   helpers.EnvIDManager
	)

	BeforeEach(func() {
		envIDGenerator = &fakes.EnvIDGenerator{}
		envIDGenerator.GenerateCall.Returns.EnvID = "some-env-id"
		networkClient = &fakes.NetworkClient{}

		envIDManager = helpers.NewEnvIDManager(envIDGenerator, networkClient)
	})

	Describe("Sync", func() {
		Context("when no previous env id exists", func() {
			It("calls env id generator if name is not passed in", func() {
				state, err := envIDManager.Sync(storage.State{}, "")
				Expect(err).NotTo(HaveOccurred())

				Expect(envIDGenerator.GenerateCall.CallCount).To(Equal(1))
				Expect(state.EnvID).To(Equal("some-env-id"))
			})

			It("uses the name passed in if an environment does not exist", func() {
				state, err := envIDManager.Sync(storage.State{}, "some-other-env-id")
				Expect(err).NotTo(HaveOccurred())

				Expect(envIDGenerator.GenerateCall.CallCount).To(Equal(0))
				Expect(state.EnvID).To(Equal("some-other-env-id"))
			})

			Context("when a pre-existing environment has that id", func() {
				BeforeEach(func() {
					networkClient.CheckExistsCall.Returns.Exists = true
				})

				Context("gcp", func() {
					It("fails", func() {
						_, err := envIDManager.Sync(storage.State{IAAS: "gcp"}, "existing")

						Expect(networkClient.CheckExistsCall.CallCount).To(Equal(1))
						Expect(networkClient.CheckExistsCall.Receives.Name).To(Equal("existing-network"))

						Expect(err).To(MatchError("It looks like a bbl environment already exists with the name 'existing'. Please provide a different name."))
					})
				})

				Context("for aws", func() {
					It("fails", func() {
						_, err := envIDManager.Sync(storage.State{IAAS: "aws"}, "existing-env")

						Expect(networkClient.CheckExistsCall.CallCount).To(Equal(1))
						Expect(networkClient.CheckExistsCall.Receives.Name).To(Equal("existing-env-vpc"))

						Expect(err).To(MatchError("It looks like a bbl environment already exists with the name 'existing-env'. Please provide a different name."))
					})
				})

				Context("for azure", func() {
					It("fails", func() {
						_, err := envIDManager.Sync(storage.State{IAAS: "azure"}, "existing-env")

						Expect(networkClient.CheckExistsCall.CallCount).To(Equal(1))
						Expect(networkClient.CheckExistsCall.Receives.Name).To(Equal("existing-env"))

						Expect(err).To(MatchError("It looks like a bbl environment already exists with the name 'existing-env'. Please provide a different name."))
					})
				})
			})
		})

		Context("for vsphere", func() {
			It("does not call the network client", func() {
				_, err := envIDManager.Sync(storage.State{
					IAAS:  "vsphere",
					EnvID: "existing-environment",
				}, "specified-env-id")
				Expect(err).NotTo(HaveOccurred())

				Expect(networkClient.CheckExistsCall.CallCount).To(Equal(0))
			})
		})

		Context("for openstack", func() {
			It("does not call the network client", func() {
				_, err := envIDManager.Sync(storage.State{IAAS: "openstack"}, "specified-env-id")
				Expect(err).NotTo(HaveOccurred())

				Expect(networkClient.CheckExistsCall.CallCount).To(Equal(0))
			})
		})

		Context("when an env id exists in the state", func() {
			It("returns the existing env id", func() {
				state, err := envIDManager.Sync(storage.State{EnvID: "some-previous-env-id"}, "")
				Expect(err).NotTo(HaveOccurred())

				Expect(envIDGenerator.GenerateCall.CallCount).To(Equal(0))
				Expect(state.EnvID).To(Equal("some-previous-env-id"))
			})
		})

		Context("when the name provided is two characters", func() {
			It("returns the name provided", func() {
				state, err := envIDManager.Sync(storage.State{IAAS: "gcp"}, "ci")
				Expect(err).NotTo(HaveOccurred())
				Expect(state.EnvID).To(Equal("ci"))
			})

			Context("when the name ends with a hyphen", func() {
				It("returns the name", func() {
					_, err := envIDManager.Sync(storage.State{IAAS: "gcp"}, "c-")
					Expect(err).To(MatchError("Names must start with a letter, be all lowercase, and be alphanumeric or hyphenated."))
				})
			})
		})

		Context("failure cases", func() {
			Context("when the NetworkClient cannot check if a network exists", func() {
				BeforeEach(func() {
					networkClient.CheckExistsCall.Returns.Error = errors.New("failed to get network list")
				})

				It("returns an error", func() {
					_, err := envIDManager.Sync(storage.State{IAAS: "gcp"}, "existing")
					Expect(err).To(MatchError("failed to get network list"))
				})
			})

			Context("when an invalid name is provided", func() {
				It("returns an error with a helpful message", func() {
					_, err := envIDManager.Sync(storage.State{}, "some_bad_name")
					Expect(err).To(MatchError("Names must start with a letter, be all lowercase, and be alphanumeric or hyphenated."))
				})
			})

			Context("when the env id generator fails", func() {
				BeforeEach(func() {
					envIDGenerator.GenerateCall.Returns.Error = errors.New("failed to generate")
				})

				It("returns an error", func() {
					_, err := envIDManager.Sync(storage.State{}, "")
					Expect(err).To(MatchError("failed to generate"))
				})
			})

			Context("when regex match string fails", func() {
				BeforeEach(func() {
					helpers.SetMatchString(func(string, string) (bool, error) {
						return false, errors.New("failed to match string")
					})
				})

				It("returns an error", func() {
					_, err := envIDManager.Sync(storage.State{}, "some-name")
					Expect(err).To(MatchError("failed to match string"))

					helpers.ResetMatchString()
				})
			})
		})
	})
})
