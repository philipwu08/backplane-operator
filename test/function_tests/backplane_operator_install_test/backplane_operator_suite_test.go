// Copyright (c) 2021 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package backplane_install_test

import (
	"flag"
	"testing"

	configv1 "github.com/openshift/api/config/v1"
	operatorv1 "github.com/openshift/api/operator/v1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	backplane "github.com/stolostron/backplane-operator/api/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	reportFile string
)

var (
	scheme             = runtime.NewScheme()
	BackplaneNamespace = flag.String("namespace", "backplane-operator-system", "The namespace to run tests")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(backplane.AddToScheme(scheme))

	utilruntime.Must(corev1.AddToScheme(scheme))

	utilruntime.Must(configv1.AddToScheme(scheme))

	utilruntime.Must(operatorv1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

func init() {
	flag.StringVar(&reportFile, "report-file", "../results/install-results.xml", "Provide the path to where the junit results will be printed.")

}

func TestBackplaneOperatorInstall(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BackplaneOperator Install Suite")
}

var _ = BeforeSuite(func() {
	ctrl.SetLogger(
		zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)),
	)

	By("bootstrapping test environment")
	c, err := config.GetConfig()
	Expect(err).ToNot(HaveOccurred())
	Expect(c).ToNot(BeNil())

	k8sClient, err = client.New(c, client.Options{Scheme: scheme})
	Expect(err).ToNot(HaveOccurred())
	Expect(k8sClient).ToNot(BeNil())

})
