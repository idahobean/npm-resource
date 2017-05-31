package check_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"testing"
)

func TestCheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Check Suite")
}

var binPath string

var _ = BeforeSuite(func() {
	var err error

	binPath, err = Build("github.com/idahobean/npm-resource/check/cmd/check")
	Ω(err).ShouldNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	CleanupBuildArtifacts()
})
