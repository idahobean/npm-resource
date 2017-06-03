package out_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	"github.com/idahobean/npm-resource"
	"github.com/idahobean/npm-resource/out"
)

var _ = Describe("Out", func() {
	var (
		cmd     *exec.Cmd
		request out.Request
	)

	BeforeEach(func() {
		request = out.Request{
			Source: resource.Source{
				PackageName: "sample-node",
				Registry:    "http://localhost:8080/",
			},
			Params: out.Params{
				UserName: "abc",
				Password: "def",
				Email:    "ghi@jkl.mno",
				Path:     "/sample-node",
				Tag:      "stable",
			},
		}
	})

	JustBeforeEach(func() {
		stdin := &bytes.Buffer{}

		err := json.NewEncoder(stdin).Encode(request)
		Ω(err).ShouldNot(HaveOccurred())

		cmd = exec.Command(binPath) // builded from test suite
		cmd.Stdin = stdin
	})

	Context("when command terminates correctly", func() {
		Context("option is fullfilled", func() {
			It("publishes npm package", func() {
				session, err := gexec.Start(
					cmd,
					GinkgoWriter,
					GinkgoWriter,
				)
				Ω(err).ShouldNot(HaveOccurred())

				Eventually(session).Should(gexec.Exit(0))

				var response out.Response
				err = json.Unmarshal(session.Out.Contents(), &response)
				Ω(err).ShouldNot(HaveOccurred())

				// shim outputs arguments
				Ω(session.Err).Should(gbytes.Say("npm-cli-login -u abc -p def -e ghi@jkl.mno -r http://localhost:8080/"))
				Ω(session.Err).Should(gbytes.Say("npm publish /sample-node --tag stable --registry http://localhost:8080/"))
				Ω(session.Err).Should(gbytes.Say("npm view sample-node --registry http://localhost:8080/"))
				Ω(session.Err).Should(gbytes.Say("npm logout --registry http://localhost:8080/"))
			})
		})
	})

	Context("when required option is empty", func() {
		Context("username is empty", func() {
			BeforeEach(func() {
				request.Params.UserName = ""
			})

			It("returns an error", func() {
				session, err := gexec.Start(
					cmd,
					GinkgoWriter,
					GinkgoWriter,
				)
				Ω(err).ShouldNot(HaveOccurred())

				Eventually(session).Should(gexec.Exit(1))

				errMsg := fmt.Sprintf("error parameter required: username")
				Ω(session.Err).Should(gbytes.Say(errMsg))
			})
		})

		Context("password is empty", func() {
			BeforeEach(func() {
				request.Params.Password = ""
			})

			It("returns an error", func() {
				session, err := gexec.Start(
					cmd,
					GinkgoWriter,
					GinkgoWriter,
				)
				Ω(err).ShouldNot(HaveOccurred())

				Eventually(session).Should(gexec.Exit(1))

				errMsg := fmt.Sprintf("error parameter required: password")
				Ω(session.Err).Should(gbytes.Say(errMsg))
			})
		})

		Context("email is empty", func() {
			BeforeEach(func() {
				request.Params.Email = ""
			})

			It("returns an error", func() {
				session, err := gexec.Start(
					cmd,
					GinkgoWriter,
					GinkgoWriter,
				)
				Ω(err).ShouldNot(HaveOccurred())

				Eventually(session).Should(gexec.Exit(1))

				errMsg := fmt.Sprintf("error parameter required: email")
				Ω(session.Err).Should(gbytes.Say(errMsg))
			})
		})

		Context("path is empty", func() {
			BeforeEach(func() {
				request.Params.Path = ""
			})

			It("returns an error", func() {
				session, err := gexec.Start(
					cmd,
					GinkgoWriter,
					GinkgoWriter,
				)
				Ω(err).ShouldNot(HaveOccurred())

				Eventually(session).Should(gexec.Exit(1))

				errMsg := fmt.Sprintf("error parameter required: path")
				Ω(session.Err).Should(gbytes.Say(errMsg))
			})
		})
	})
})
