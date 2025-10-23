// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use thisfile except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package v1alpha1

import "fmt"

var (
	// ProvisionedThroughputManagedByAnnotation is the annotation key used to set the management style for the
	// provisioned throughput of a table. This annotation can only be set on a
	// table custom resource.
	//
	// The value of this annotation must be one of the following:
	//
	// - 'external-autoscaler': The provisioned throughput is managed by an external entity, causing the
	//                          controller to ignore the `spec.provisionedThroughput` field.
	//
	// - 'ack-dynamodb-controller': The provisioned throughput is managed by the ACK controller.
	//
	// By default, the provisioned throughput is managed by the controller.
	ProvisionedThroughputManagedByAnnotation = fmt.Sprintf("%s/provisioned-throughput-managed-by", GroupVersion.Group)
)

const (
	// ProvisionedThroughputManagedByExternalAutoscaler is the value of the ProvisionedThroughputManagedByAnnotation
	// annotation that indicates that the provisioned throughput of a table is managed by an external autoscaler.
	ProvisionedThroughputManagedByExternalAutoscaler = "external-autoscaler"

	// ProvisionedThroughputManagedByACKController is the value of the ProvisionedThroughputManagedByAnnotation
	// annotation that indicates that the provisioned throughput of a table is managed by the ACK
	// controller.
	ProvisionedThroughputManagedByACKController = "ack-eks-controller"
)