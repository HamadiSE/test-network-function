package ocpclient

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/test-network-function/test-network-function/test-network-function/common"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

type OcpClient struct {
	Coreclient *corev1client.CoreV1Client
	ready      bool
}

var ocpClient = OcpClient{}

func GetOcpClient() *OcpClient {
	NewOcpClient()
	return &ocpClient
}

func getDefaultPath() (path string) {
	home := os.Getenv("HOME")
	if home != "" {
		path = filepath.Join(home, ".kube", "config")
	}
	return
}
func getKubeConfig() (config string) {
	config = os.Getenv("KUBECONFIG")
	return
}

func NewOcpClient(filenames ...string) {

	if ocpClient.ready {
		return
	}
	ocpClient.ready = true

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()

	precedence := []string{}
	if len(filenames) > 0 {
		precedence = append(precedence, filenames...)
	}
	if f := getKubeConfig(); f != "" {
		precedence = append(precedence, f)
	}
	if f := getDefaultPath(); f != "" {
		precedence = append(precedence, f)
	}
	// Follows the logic presented in README.md,
	// we start by loading configuration from files supplied using -k command
	// then we look for files using KUBECONFIG
	// then we look for file $HOME/.kube/config
	loadingRules.Precedence = precedence

	configOverrides := &clientcmd.ConfigOverrides{}

	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		configOverrides,
	)
	// Get a rest.Config from the kubeconfig file.  This will be passed into all
	// the client objects we create.
	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}
	restconfig.Timeout = common.DefaultTimeout
	ocpClient.Coreclient, err = corev1client.NewForConfig(restconfig)
	if err != nil {
		logrus.Panic("can't instantiate corev1client", err)
	}
}
