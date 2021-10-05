package common

import (
	"github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"
)

var _ = ginkgo.Describe(commonTestKey, func() {
	ginkgo.BeforeSuite(func() {
		log.Info("test suite setup")
		//env := config.GetTestEnvironment()
		//env.LoadAndRefresh()
	})
	//
	ginkgo.AfterSuite(func() {
		TeardownNodeDebugSession()
		log.Info("After Suite")
	})
})
