package main_test

import (
	"fmt"
	"runtime"

	cli "github.com/heroku/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("version", func() {
	BeforeEach(func() {
		cli.Start("heroku", "version")
	})

	It("shows the version", func() {
		version := fmt.Sprintf("heroku-cli/%s (%s-%s) %s ?\n", cli.Version, runtime.GOOS, runtime.GOARCH, runtime.Version())
		Expect(stdout()).To(Equal(version))
	})
})
