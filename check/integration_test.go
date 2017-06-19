package check_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	"github.com/idahobean/npm-resource"
	"github.com/idahobean/npm-resource/check"
)

var _ = Describe("Check", func() {
	var (
		tmpDir  string
		cmd     *exec.Cmd
		request check.Request
	)

	loginArgs := []string{"-u", "abc", "-p", "def", "-e", "ghi@jkl.mno", "-r", "http://localhost:8080"}

	BeforeEach(func() {
		var err error

		tmpDir, err = ioutil.TempDir("", "npm_resource_check")
		Ω(err).ShouldNot(HaveOccurred())

		packagePath, err := filepath.Abs("../sample-node")
		Ω(err).ShouldNot(HaveOccurred())

		request = check.Request{
			Source: resource.Source{
				PackageName: "sample-node",
				Registry:    "http://localhost:8080",
			},
		}

		err = exec.Command("npm-cli-login", loginArgs...).Run()
		Ω(err).ShouldNot(HaveOccurred())

		args := []string{"publish", packagePath, "--registry", "http://localhost:8080"}
		err = exec.Command("npm", args...).Run()
		Ω(err).ShouldNot(HaveOccurred())

	})

	JustBeforeEach(func() {
		stdin := &bytes.Buffer{}

		err := json.NewEncoder(stdin).Encode(request)
		Ω(err).ShouldNot(HaveOccurred())

		cmd = exec.Command(binPath, tmpDir) // builded from test suite
		cmd.Stdin = stdin
		cmd.Dir = tmpDir
	})

	AfterEach(func() {
		err := os.RemoveAll(tmpDir)
		Ω(err).ShouldNot(HaveOccurred())

		args := []string{"unpublish", "sample-node", "--registry", "http://localhost:8080", "--force"}
		err = exec.Command("npm", args...).Run()
		Ω(err).ShouldNot(HaveOccurred())
	})

	Context("when command terminates correctly", func() {
		Context("packagename is fullfilled", func() {
			It("returns npm version", func() {
				session, err := gexec.Start(
					cmd,
					GinkgoWriter,
					GinkgoWriter,
				)
				Ω(err).ShouldNot(HaveOccurred())

				Eventually(session, "15s").Should(gexec.Exit(0))

				var response check.Response
				err = json.Unmarshal(session.Out.Contents(), &response)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(response).Should(Equal(check.Response{
					Version: resource.Version{
						Version: "0.0.1",
					},
					Metadata: []resource.MetadataPair{
						{
							Name:  "name",
							Value: "sample-node",
						},
						{
							Name:  "homepage",
							Value: "https://github.com/idahobean/sample-node#readme"},
					},
				}))
			})
		})
	})

	Context("when required option is empty", func() {
		Context("packagename is empty", func() {
			BeforeEach(func() {
				request.Source.PackageName = ""
			})

			It("returns an error", func() {
				session, err := gexec.Start(
					cmd,
					GinkgoWriter,
					GinkgoWriter,
				)
				Ω(err).ShouldNot(HaveOccurred())

				Eventually(session).Should(gexec.Exit(1))

				errMsg := fmt.Sprintf("error parameter required: package_name")
				Ω(session.Err).Should(gbytes.Say(errMsg))
			})
		})
	})
})
