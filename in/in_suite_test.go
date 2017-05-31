package in_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"testing"
)

func TestIn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "In Suite")
}

var binPath string

var _ = BeforeSuite(func() {
	var err error

	binPath, err = Build("github.com/idahobean/npm-resource/in/cmd/in")
	Ω(err).ShouldNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	CleanupBuildArtifacts()
})
