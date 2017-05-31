package out_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"testing"
)

func TestOut(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Out Suite")
}

var binPath string

var _ = BeforeSuite(func() {
	var err error

	binPath, err = Build("github.com/idahobean/npm-resource/out/cmd/out")
	Ω(err).ShouldNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	CleanupBuildArtifacts()
})
