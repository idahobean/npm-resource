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
