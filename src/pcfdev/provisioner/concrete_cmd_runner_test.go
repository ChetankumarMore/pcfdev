package provisioner_test

import (
	"pcfdev/provisioner"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("ConcreteCmdRunner", func() {
	Describe("#Run", func() {
		var (
			r      *provisioner.ConcreteCmdRunner
			stdout *gbytes.Buffer
			stderr *gbytes.Buffer
		)

		BeforeEach(func() {
			stdout = gbytes.NewBuffer()
			stderr = gbytes.NewBuffer()

			r = &provisioner.ConcreteCmdRunner{
				Stdout: stdout,
				Stderr: stderr,
			}
		})

		It("should run commands", func() {
			Expect(r.Run("echo", "-n", "some output")).To(Succeed())
			Eventually(stdout).Should(gbytes.Say("some output"))

			Expect(r.Run("bash", "-c", ">&2 echo -n some output")).To(Succeed())
			Eventually(stderr).Should(gbytes.Say("some output"))
		})

		Context("when there is an error", func() {
			It("should return the error and the output", func() {
				Expect(r.Run("/some/bad/binary")).To(MatchError(ContainSubstring("no such file or directory")))
			})
		})
	})
})
