/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e_test

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"

	ctrl "sigs.k8s.io/controller-runtime"

	"go.openviz.dev/grafana-tools/pkg/operator/cmds/server"
	"go.openviz.dev/grafana-tools/test/e2e/framework"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"gomodules.xyz/logs"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"kmodules.xyz/client-go/tools/clientcmd"
)

const (
	TIMEOUT = 20 * time.Minute
)

var (
	root *framework.Framework
)

func TestE2e(t *testing.T) {
	logs.InitLogs()
	RegisterFailHandler(Fail)
	SetDefaultEventuallyTimeout(TIMEOUT)

	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "e2e Suite", []Reporter{junitReporter})
}

var _ = BeforeSuite(func() {
	By("Using kubeconfig from " + options.kubeConfig)
	clientConfig, err := clientcmd.BuildConfigFromContext(options.kubeConfig, options.kubeContext)
	Expect(err).NotTo(HaveOccurred())
	// raise throttling time. ref: https://github.com/appscode/voyager/issues/640
	clientConfig.Burst = 100
	clientConfig.QPS = 100

	mgr, err := server.GetManager()
	Expect(err).NotTo(HaveOccurred())
	restConfig := mgr.GetConfig()
	restConfig.Burst = 100
	restConfig.QPS = 100

	// Framework
	root = framework.New(restConfig, mgr.GetClient())

	go func() {
		if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
			log.Printf("error from manager: %v", err)
		}
	}()

	By("Creating namespace " + root.Namespace())
	err = root.CreateNamespace()
	Expect(err).NotTo(HaveOccurred())

	By("Creating grafana server")
	err = root.DeployGrafanaServer()
	Expect(err).NotTo(HaveOccurred())

	By("Waiting for grafana server to be ready")
	root.WaitForGrafanaServerToBeReady()

	// temporary crd installation
	crd := &apiextensions.CustomResourceDefinition{}
	by, err := os.ReadFile("../../config/crd/bases/appcatalog.appscode.com_appbindings.yaml")
	Expect(err).NotTo(HaveOccurred())
	Expect(json.Unmarshal(by, crd)).Should(Succeed())
	cc := mgr.GetClient()
	Expect(cc.Create(context.TODO(), crd)).Should(Succeed())

	root.EventuallyCRD().Should(Succeed())
})

var _ = AfterSuite(func() {
	By("Deleting grafana server")
	err := root.DeleteGrafanaServer()
	Expect(err).NotTo(HaveOccurred())

	By("Deleting Namespace")
	err = root.DeleteNamespace()
	Expect(err).NotTo(HaveOccurred())
})
