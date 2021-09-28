package nodedebugshell

import (
	"path"

	"github.com/onsi/gomega"
	"github.com/test-network-function/test-network-function/pkg/tnf"
	"github.com/test-network-function/test-network-function/pkg/tnf/handlers/generic"
	"github.com/test-network-function/test-network-function/pkg/tnf/reel"
	"github.com/test-network-function/test-network-function/test-network-function/common"
)

var (
	nodeDebugShellPath = path.Join("pkg", "tnf", "handlers", "nodedebugshell", "nodedebugshell.json")

	// relativeShutdownTestPath is the relative path to the shutdown.json test case.
	relativeNodeDebugTestPath = path.Join(common.PathRelativeToRoot, nodeDebugShellPath)
)

func NewNodeDebugShell(node string) (*tnf.Tester, []reel.Handler) {
	test, handlers, result, err := generic.NewGenericFromMap(relativeNodeDebugTestPath, common.RelativeSchemaPath, nil)
	gomega.Expect(err).To(gomega.BeNil())
	gomega.Expect(result).ToNot(gomega.BeNil())
	gomega.Expect(result.Valid()).To(gomega.BeTrue())
	gomega.Expect(handlers).ToNot(gomega.BeNil())
	gomega.Expect(handlers).ToNot(gomega.BeNil())
	gomega.Expect(test).ToNot(gomega.BeNil())
	return test, handlers
}
