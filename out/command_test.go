package out_test

import (
	"errors"
	"path/filepath"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/idahobean/npm-resource"
	"github.com/idahobean/npm-resource/out"
	"github.com/idahobean/npm-resource/npm/fakes"
)

var _ = Describe("Out Command", func() {
	var (
		NPM *fakes.FakeNPM
		request out.Request
		command *out.Command
	)

	BeforeEach(func() {
		NPM = &fakes.FakeNPM{}
		command = out.NewCommand(NPM)

		request = out.Request{
			Source: resource.Source{
				Token: "test",
				PackageName: "foo",
				Registry: "http://my.private.registry/",
			},
			Params: out.Params{
				Path: "bar/baz",
				Version: "0.0.2",
				Tag: "fox",
			},
		}
	})

	Describe("running the command", func() {
		It("publishes package", func() {
			response, err := command.Run(request)
			Ω(err).ShouldNot(HaveOccurred())

			Ω(response.Version.Version).Should(Equal("0.0.1"))
			Ω(response.Metadata[0]).Should(Equal(
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

			By("npm publish")
			Ω(NPM.PublishCallCount()).Should(Equal(1))

			path, tag, registry := NPM.PublishArgsForCall(0)

			basePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				return
			}

			Ω(path).Should(Equal(filepath.Join(basePath, "bar/baz")))
			Ω(tag).Should(Equal("fox"))
			Ω(registry).Should(Equal("http://my.private.registry/"))

		})

// TODO other test patterns

		Describe("handling any errors", func() {
			var expectedError error

			BeforeEach(func() {
				expectedError = errors.New("it all went wrong")
			})

			It("from publish package", func() {
				NPM.PublishReturns(expectedError)

				_, err := command.Run(request)
				Ω(err).Should(MatchError(expectedError))
			})
		})
	})
})
