//go:build e2e
// +build e2e

/*
Copyright 2020 The Kubernetes Authors.

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
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	ekstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/util/version"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"

	ekscontrolplanev1 "sigs.k8s.io/cluster-api-provider-aws/v2/controlplane/eks/api/v1beta2"
	"sigs.k8s.io/cluster-api/test/framework"
)

type waitForControlPlaneToBeUpgradedInput struct {
	ControlPlane   *ekscontrolplanev1.AWSManagedControlPlane
	AWSSession     *aws.Config
	UpgradeVersion string
}

func waitForControlPlaneToBeUpgraded(ctx context.Context, input waitForControlPlaneToBeUpgradedInput, intervals ...interface{}) {
	Expect(input.ControlPlane).ToNot(BeNil(), "Invalid argument. input.ControlPlane can't be nil")
	Expect(input.AWSSession).ToNot(BeNil(), "Invalid argument. input.AWSSession can't be nil")
	Expect(input.UpgradeVersion).ToNot(BeNil(), "Invalid argument. input.UpgradeVersion can't be nil")

	By(fmt.Sprintf("Ensuring EKS control-plane has been upgraded to kubernetes version %s", input.UpgradeVersion))
	v, err := version.ParseGeneric(input.UpgradeVersion)
	Expect(err).NotTo(HaveOccurred())
	expectedVersion := fmt.Sprintf("%d.%d", v.Major(), v.Minor())

	Eventually(func() (bool, error) {
		cluster, err := getEKSCluster(ctx, input.ControlPlane.Spec.EKSClusterName, input.AWSSession)
		if err != nil {
			return false, err
		}

		switch cluster.Status {
		case ekstypes.ClusterStatusUpdating:
			return false, nil
		case ekstypes.ClusterStatusActive:
			if *cluster.Version == expectedVersion {
				return true, nil
			}
			return false, nil
		default:
			return false, nil
		}
	}, intervals...).Should(BeTrue(), fmt.Sprintf("Eventually failed waiting for EKS control-plane to be upgraded to kubernetes version %q", input.UpgradeVersion))
}

type GetControlPlaneByNameInput struct {
	Getter    framework.Getter
	Name      string
	Namespace string
}

func GetControlPlaneByName(ctx context.Context, input GetControlPlaneByNameInput) *ekscontrolplanev1.AWSManagedControlPlane {
	cp := &ekscontrolplanev1.AWSManagedControlPlane{}
	key := crclient.ObjectKey{
		Name:      input.Name,
		Namespace: input.Namespace,
	}
	Eventually(func() error {
		err := input.Getter.Get(ctx, key, cp)
		if err != nil {
			return err
		}
		return nil
	}, 2*time.Minute, 5*time.Second).Should(Succeed(), fmt.Sprintf("Eventually failed to get AWSManagedControlPlane object '%s/%s'", input.Namespace, input.Name))
	Expect(input.Getter.Get(ctx, key, cp)).To(Succeed(), "Failed to get AWSManagedControlPlane object %s/%s", input.Namespace, input.Name)
	return cp
}
