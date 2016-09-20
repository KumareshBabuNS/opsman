package uaa_test

import (
	. "github.com/datianshi/opsman/uaa"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Uaa", func() {

	uaaRequest := &UAA{
		URL: "https://opsmgr.haas-22.pez.pivotal.io/uaa",
		Username: "admin",
		Password: "password",
		SkipSsl: true,
	}

	BeforeEach(func(){
		os.Setenv("HTTP_DEBUG", "true")
	})

	Describe("Retrieve Token", func() {
		Context("When requesting token", func() {
			It("err should be nil", func() {
				_, err := uaaRequest.GetToken()
				Ω(err).Should(BeNil())
			})
			It("token should be returned", func() {
				token, _ := uaaRequest.GetToken()
				Ω(token).ShouldNot(BeNil())
			})
		})
	})
})
