package out_test

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	"github.com/idahobean/npm-resource"
	"github.com/idahobean/npm-resource/out"
)

var _ = Describe("Out", func() {
	var (
		cmd *exec.Cmd
		request out.Request
	)

	BeforeEach(func() {
		request = out.Request{
			Source: resource.Source{
				Token: "test-token",
				PackageName: "foobar-pack",
				Registry: "http://my.private.registry/",
			},
			Params: out.Params{
				Path: "baz/fox",
				Version: "0.1.2",
				Tag: "taag",
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
				Ω(session.Err).Should(gbytes.Say("npm publish baz/fox --tag taag --registry http://my.private.registry/"))
				Ω(session.Err).Should(gbytes.Say("npm view foobar-pack --registry http://my.private.registry/"))

			})
		})
	})

	Context("when required option is empty", func() {
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
