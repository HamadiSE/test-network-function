package liveness

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/test-network-function/test-network-function/internal/ocpclient"
	"github.com/test-network-function/test-network-function/pkg/tnf"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Liveness struct {
	timeout       time.Duration
	podName       string
	containerName string
	namespace     string
	liveness      bool // true if container has liveness section defined
	Result        int
}

func NewLiveness(ns, pod, container string, timeout time.Duration) Liveness {
	return Liveness{
		timeout:       timeout,
		podName:       pod,
		containerName: container,
		namespace:     ns,
		liveness:      false,
		Result:        tnf.ERROR,
	}
}

func (l *Liveness) RunTest() {
	cl := ocpclient.GetOcpClient()
	ctx := context.TODO()
	options := v1.GetOptions{}

	log.Debugf("Check liveness ns=%s podname=%s container=%s", l.namespace, l.podName, l.containerName)
	pod, err := cl.Coreclient.Pods(l.namespace).Get(ctx, l.podName, options)
	if err != nil {
		log.Errorf("Getting pod information using go-client %s", err.Error())
		l.Result = tnf.ERROR
		return
	}
	for i, _ := range pod.Spec.Containers {
		cut := pod.Spec.Containers[i]
		fmt.Println("Debug ", cut.LivenessProbe)
		if cut.Name == l.containerName {
			if cut.LivenessProbe != nil {
				l.Result = tnf.SUCCESS
			} else {
				l.Result = tnf.FAILURE
			}
			return
		}
	}
}
