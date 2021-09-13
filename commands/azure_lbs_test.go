package commands_test

import (
	"errors"

	"github.com/orange-cloudfoundry/bosh-bootloader/commands"
	"github.com/orange-cloudfoundry/bosh-bootloader/fakes"
	"github.com/orange-cloudfoundry/bosh-bootloader/storage"
	"github.com/orange-cloudfoundry/bosh-bootloader/terraform"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Azure LBs", func() {
	var (
		command commands.AzureLBs

		terraformManager *fakes.TerraformManager
		logger           *fakes.Logger

		incomingState storage.State
	)

	BeforeEach(func() {
		terraformManager = &fakes.TerraformManager{}
		logger = &fakes.Logger{}

		command = commands.NewAzureLBs(terraformManager, logger)
	})

	Describe("Execute", func() {
		Context("when the lb type is cf", func() {
			BeforeEach(func() {
				incomingState = storage.State{
					IAAS: "azure",
					LB: storage.LB{
						Type: "cf",
					},
				}
				terraformManager.GetOutputsCall.Returns.Outputs = terraform.Outputs{Map: map[string]interface{}{
					"cf_app_gateway_name": "some-app-gateway-name",
				}}
			})

			It("prints LB name", func() {
				err := command.Execute([]string{}, incomingState)
				Expect(err).NotTo(HaveOccurred())

				Expect(terraformManager.GetOutputsCall.CallCount).To(Equal(1))
				Expect(logger.PrintfCall.Messages).To(ConsistOf([]string{
					"CF LB: some-app-gateway-name\n",
				}))
			})
		})

		Context("when the lb type is concourse", func() {
			BeforeEach(func() {
				incomingState = storage.State{
					IAAS: "azure",
					LB: storage.LB{
						Type: "concourse",
					},
				}
				terraformManager.GetOutputsCall.Returns.Outputs = terraform.Outputs{Map: map[string]interface{}{
					"concourse_lb_name": "some-load-balancer-name",
					"concourse_lb_ip":   "5.6.7.8",
				}}
			})

			It("prints LB name", func() {
				err := command.Execute([]string{}, incomingState)
				Expect(err).NotTo(HaveOccurred())

				Expect(terraformManager.GetOutputsCall.CallCount).To(Equal(1))
				Expect(logger.PrintfCall.Messages).To(ConsistOf([]string{
					"Concourse LB: some-load-balancer-name (5.6.7.8)\n",
				}))
			})
		})

		Context("when lb type is not cf or concourse", func() {
			BeforeEach(func() {
				incomingState = storage.State{
					IAAS: "azure",
					LB: storage.LB{
						Type: "other",
					},
				}
			})

			It("returns error", func() {
				err := command.Execute([]string{}, incomingState)
				Expect(err).To(MatchError("no lbs found"))
			})
		})

		Context("failure cases", func() {
			BeforeEach(func() {
				incomingState = storage.State{}
			})

			Context("when terraform manager fails", func() {
				It("returns an error", func() {
					terraformManager.GetOutputsCall.Returns.Error = errors.New("terraform manager failed")
					err := command.Execute([]string{}, incomingState)

					Expect(err).To(MatchError("terraform manager failed"))
				})
			})
		})
	})
})
