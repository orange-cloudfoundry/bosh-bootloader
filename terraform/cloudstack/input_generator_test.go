package cloudstack_test

import (
	"github.com/orange-cloudfoundry/bosh-bootloader/fakes"
	"github.com/orange-cloudfoundry/bosh-bootloader/storage"
	"github.com/orange-cloudfoundry/bosh-bootloader/terraform/cloudstack"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InputGenerator", func() {
	var (
		awsClient *fakes.AWSClient

		inputGenerator cloudstack.InputGenerator
	)

	BeforeEach(func() {
		awsClient = &fakes.AWSClient{}
		awsClient.RetrieveAZsCall.Returns.AZs = []string{"z1", "z2", "z3"}

		inputGenerator = cloudstack.NewInputGenerator()
	})

	Describe("Generate", func() {
		Context("when env-id is greater than 18 characters", func() {
			It("creates a short env-id with truncated env_id and sha1sum[0:7]", func() {
				inputs, err := inputGenerator.Generate(storage.State{
					EnvID:      "some-env-id-that-is-pretty-long",
					CloudStack: storage.CloudStack{},
				})
				Expect(err).NotTo(HaveOccurred())

				Expect(inputs["env_id"]).To(Equal("some-env-id-that-is-pretty-long"))
				Expect(inputs["short_env_id"]).To(Equal("some-env-i-1fc794e"))
			})
		})

		It("receives BBL state and returns a map of terraform variables", func() {
			inputs, err := inputGenerator.Generate(storage.State{
				EnvID: "some-env-id",
				CloudStack: storage.CloudStack{
					ApiKey:          "some-access-key-id",
					SecretAccessKey: "some-secret-access-key",
					Zone:            "my-zone",
					Secure:          true,
					IsoSegment:      true,
					Endpoint:        "http://my-cloudstack.com/client/api",
				},
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(inputs).To(Equal(map[string]interface{}{
				"env_id":              "some-env-id",
				"short_env_id":        "some-env-id",
				"cloudstack_zone":     "my-zone",
				"cloudstack_endpoint": "http://my-cloudstack.com/client/api",
				"secure":              "true",
				"iso_segment":         "true",
			}))
		})
	})

	Describe("Credentials", func() {
		It("returns the api key, secret key and the key name", func() {
			state := storage.State{
				CloudStack: storage.CloudStack{
					ApiKey:          "some-access-key-id",
					SecretAccessKey: "some-secret-access-key",
				},
			}
			credentials := inputGenerator.Credentials(state)

			Expect(credentials).To(Equal(map[string]string{
				"cloudstack_api_key":           "some-access-key-id",
				"cloudstack_secret_access_key": "some-secret-access-key",
			}))
		})
	})
})
