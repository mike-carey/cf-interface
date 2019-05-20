package cf_interface_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-community/go-cfclient"
	. "github.com/mike-carey/cf-interface"
)

func ShouldBeAbleToCompile(cf CFClient) error {
	fmt.Println("This compiled properly and therefore the cfclient.Client identifies as the CFClient interface")
	return nil
}

var _ = Describe("CfInterface", func() {

	It("CFClient should identify as the interface", func() {
		client := &cfclient.Client{Config: cfclient.Config{}}

		Expect(ShouldBeAbleToCompile(client)).To(BeNil())
	})

})
