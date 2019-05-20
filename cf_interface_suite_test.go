package cf_interface_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCfInterface(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CfInterface Suite")
}
