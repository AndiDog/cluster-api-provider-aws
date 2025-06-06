//go:build e2e
// +build e2e

/*
Copyright 2021 The Kubernetes Authors.

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

package managed

import (
	"context"
	"fmt"

	awsv2 "github.com/aws/aws-sdk-go-v2/aws"
	ekstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"

	ekscontrolplanev1 "sigs.k8s.io/cluster-api-provider-aws/v2/controlplane/eks/api/v1beta2"
	"sigs.k8s.io/cluster-api/test/framework"
	"sigs.k8s.io/cluster-api/test/framework/clusterctl"
)

type CheckAddonExistsSpecInput struct {
	E2EConfig             *clusterctl.E2EConfig
	BootstrapClusterProxy framework.ClusterProxy
	AWSSession            *awsv2.Config
	Namespace             *corev1.Namespace
	ClusterName           string
	AddonName             string
	AddonVersion          string
	AddonConfiguration    string
}

// CheckAddonExistsSpec implements a test for a cluster having an addon.
func CheckAddonExistsSpec(ctx context.Context, inputGetter func() CheckAddonExistsSpecInput) {
	input := inputGetter()
	Expect(input.E2EConfig).ToNot(BeNil(), "Invalid argument. input.E2EConfig can't be nil")
	Expect(input.BootstrapClusterProxy).ToNot(BeNil(), "Invalid argument. input.BootstrapClusterProxy can't be nil")
	Expect(input.AWSSession).ToNot(BeNil(), "Invalid argument. input.AWSSession can't be nil")
	Expect(input.Namespace).NotTo(BeNil(), "Invalid argument. input.Namespace can't be nil")
	Expect(input.ClusterName).ShouldNot(BeEmpty(), "Invalid argument. input.ClusterName can't be empty")
	Expect(input.AddonName).ShouldNot(BeEmpty(), "Invalid argument. input.AddonName can't be empty")
	Expect(input.AddonVersion).ShouldNot(BeEmpty(), "Invalid argument. input.AddonVersion can't be empty")

	mgmtClient := input.BootstrapClusterProxy.GetClient()
	controlPlaneName := getControlPlaneName(input.ClusterName)

	By(fmt.Sprintf("Getting control plane: %s", controlPlaneName))
	controlPlane := &ekscontrolplanev1.AWSManagedControlPlane{}
	Eventually(func() error {
		return mgmtClient.Get(ctx, crclient.ObjectKey{Namespace: input.Namespace.Name, Name: controlPlaneName}, controlPlane)
	}, input.E2EConfig.GetIntervals("", "wait-client-request")...).Should(Succeed(), "eventually failed trying to get the AWSManagedControlPlane")

	By(fmt.Sprintf("Checking EKS addon %s is installed on cluster %s and is active", input.AddonName, input.ClusterName))
	waitForEKSAddonToHaveStatus(ctx, waitForEKSAddonToHaveStatusInput{
		ControlPlane: controlPlane,
		AWSSession:   input.AWSSession,
		AddonName:    input.AddonName,
		AddonVersion: input.AddonVersion,
		AddonStatus:  []string{string(ekstypes.AddonStatusActive), string(ekstypes.AddonStatusDegraded)},
	}, input.E2EConfig.GetIntervals("", "wait-addon-status")...)

	if input.AddonConfiguration != "" {
		By(fmt.Sprintf("Checking EKS addon %s has the correct configuration", input.AddonName))
		checkEKSAddonConfiguration(ctx, checkEKSAddonConfigurationInput{
			ControlPlane:       controlPlane,
			AWSSession:         input.AWSSession,
			AddonName:          input.AddonName,
			AddonVersion:       input.AddonVersion,
			AddonConfiguration: input.AddonConfiguration,
		}, input.E2EConfig.GetIntervals("", "wait-addon-status")...)
	}
}
