package in_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/idahobean/npm-resource"
	"github.com/idahobean/npm-resource/in"
	"github.com/idahobean/npm-resource/fakes"
)

var _ = Describe("In Command", func() {
	var (
		NPM *fakes.FakeNPM
		request in.Request
		command *in.Command
	)

	BeforeEach(func() {
		NPM = &fakes.FakeNPM{}
		command = in.NewCommand(NPM)

		request = in.Request{
			Source: resource.Source{
				Token: "//localhost:4873:_authToken=test",
				PackageName: "foo",
				Registry: "http://my.private.registry/",
			},
			Params: in.Params{},
		}
	})

	Describe("running the command", func() {
		It("pulls package", func() {
			response, err := command.Run(request)
			Ω(err).ShouldNot(HaveOccurred())

			Ω(response.Version.Version).Should(Equal("0.0.1"))
			Ω(reqponse.Metadata[0]).Should(Equal(
				resource.MetadataPair{
					Name: "name",
					Value: "foo-package",
				},
			))
			Ω(response.Metadata[1]).Should(Equal(
				resource.MetadataPair{
					Name: "homepage",
					Value: "foobars page",
				},
			))

			By("npm install")
			Ω(NPM.InstallCallCount()).Should(Equal(1))

			packageName, registry := NPM.InstallArgsForCall(0)
			Ω(packageName).Should(Equal("foo-package"))
			Ω(registry).Should(Equal("http://my.private.registry/"))

		})

		Describe("handling any errors", func() {
			var expectedError error

			BeforeEach(func() {
				expectedError = errors.New("it all went wrong")
			})

			It("from install package", func() {
				NPM.InstallReturns(expectedError)

				_, err := command.Run(request)
				Ω(err).Should(MatchError(expectedError))
			})
		})
	})
})
