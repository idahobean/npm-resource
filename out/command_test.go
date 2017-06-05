package out_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/idahobean/npm-resource"
	"github.com/idahobean/npm-resource/npm"
	"github.com/idahobean/npm-resource/npm/fakes"
	"github.com/idahobean/npm-resource/out"
)

var _ = Describe("Out Command", func() {
	var (
		NPM          *fakes.FakeNPM
		request      out.Request
		command      *out.Command
		returnedInfo *npm.PackageInfo
	)

	BeforeEach(func() {
		NPM = &fakes.FakeNPM{}
		command = out.NewCommand(NPM)

		request = out.Request{
			Source: resource.Source{
				PackageName: "foo-package",
				Registry:    "http://my.private.registry/",
			},
			Params: out.Params{
				UserName: "aaa",
				Password: "bbb",
				Email:    "ccc@ddd.eee",
				Path:     "bar/baz",
				Tag:      "fox",
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

		It("publishes package", func() {
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

			By("npm-cli-login")
			Ω(NPM.LoginCallCount()).Should(Equal(1))

			username, password, email, registry := NPM.LoginArgsForCall(0)

			Ω(username).Should(Equal("aaa"))
			Ω(password).Should(Equal("bbb"))
			Ω(email).Should(Equal("ccc@ddd.eee"))
			Ω(registry).Should(Equal("http://my.private.registry/"))

			By("npm publish")
			Ω(NPM.PublishCallCount()).Should(Equal(1))

			path, tag, registry := NPM.PublishArgsForCall(0)

			Ω(path).Should(Equal("bar/baz"))
			Ω(tag).Should(Equal("fox"))
			Ω(registry).Should(Equal("http://my.private.registry/"))

			By("npm logout")
			Ω(NPM.LogoutCallCount()).Should(Equal(1))

			registry = NPM.LogoutArgsForCall(0)

			Ω(registry).Should(Equal("http://my.private.registry/"))

		})

		Describe("handling any errors", func() {
			var expectedError error

			BeforeEach(func() {
				expectedError = errors.New("it all went wrong")
			})

			It("from login", func() {
				NPM.LoginReturns(expectedError)

				_, err := command.Run(request)
				Ω(err).Should(MatchError(expectedError))
			})

			It("from publish package", func() {
				NPM.PublishReturns(expectedError)

				_, err := command.Run(request)
				Ω(err).Should(MatchError(expectedError))
			})

			It("from logout", func() {
				NPM.LogoutReturns(expectedError)

				_, err := command.Run(request)
				Ω(err).Should(MatchError(expectedError))
			})
		})
	})
})
