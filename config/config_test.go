package config_test

import (
	. "github.com/datianshi/sshdownload/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Describe("Test Config parse", func() {
		Context("With correct config", func() {
			It("can be loaded from YML", func() {
				filePath := "config-test.yml"
				cf := ParseConfig(filePath)
				Expect(cf.Username).To(Equal("username"))
				Expect(cf.Password).To(Equal("password"))
				Expect(cf.Host).To(Equal("192.168.1.33"))
				Expect(cf.Port).To(Equal("22"))
				Expect(cf.Cmd).To(Equal("abc && cde && def"))
				Expect(cf.File).To(Equal("/tmp/abc"))
			})
		})

	})

})
