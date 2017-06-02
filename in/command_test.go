package in_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/idahobean/npm-resource"
	"github.com/idahobean/npm-resource/in"
	"github.com/idahobean/npm-resource/npm"
	"github.com/idahobean/npm-resource/npm/fakes"
)

var _ = Describe("In Command", func() {
	var (
		NPM          *fakes.FakeNPM
		request      in.Request
		command      *in.Command
		returnedInfo *npm.PackageInfo
	)

	BeforeEach(func() {
		NPM = &fakes.FakeNPM{}
		command = in.NewCommand(NPM)

		request = in.Request{
			Source: resource.Source{
				PackageName: "foo-package",
				Registry:    "http://my.private.registry/",
			},
		}

		returnedInfo = &npm.PackageInfo{}
	})

	JustBeforeEach(func() {
		NPM.ViewReturns(returnedInfo, nil)
	})

	Describe("running the command", func() {
		BeforeEach(func() {
			returnedInfo = &npm.PackageInfo{
				Name:     "foo-package",
				Version:  "0.0.1",
				Homepage: "http://foobar.com",
			}
		})

		It("pulls package", func() {
			response, err := command.Run(request)
			Ω(err).ShouldNot(HaveOccurred())

			Ω(response.Version.Version).Should(Equal("0.0.1"))
			Ω(response.Metadata[0]).Should(Equal(
				resource.MetadataPair{
					Name:  "name",
					Value: "foo-package",
				},
			))
			Ω(response.Metadata[1]).Should(Equal(
				resource.MetadataPair{
					Name:  "homepage",
					Value: "http://foobar.com",
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
