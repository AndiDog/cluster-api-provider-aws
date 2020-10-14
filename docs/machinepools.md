# MachinePools
- **Feature status:** Experimental
- **Feature gate:** MachinePool=true

MachinePool allows users to manage many machines as a single entity. Infrastructure providers 
implement a separate CRD that handles infrastructure side of the feature. 


## AWSMachinePool
Cluster API Provider AWS (CAPA) has experimental support for `MachinePool` though the infrastructure
type `AWSMachinePool`. An `AWSMachinePool` corresponds to an [AWS AutoScaling Groups](https://docs.aws.amazon.com/autoscaling/ec2/userguide/AutoScalingGroup.html), which provides the cloud provider specific resource for orchestrating a group of EC2 machines.

The AWSMachinePool controller creates and manages an AWS AutoScaling Group using launch templates 
so users don't have to manage individual machines. You can use Autoscaling health checks for replacing 
instances and it will maintain the number of instances specified.


### Using `clusterctl` to deploy
To deploy a MachinePool / AWSMachinePool via `clusterctl config` there's a [flavor](https://cluster-api.sigs.k8s.io/clusterctl/commands/config-cluster.html#flavors)
for that.

Make sure to set up your AWS environment as described [here](https://cluster-api.sigs.k8s.io/user/quick-start.html).

```shell
export EXP_MACHINE_POOL=true
clusterctl init --infrastructure aws
clusterctl config cluster my-cluster --kubernetes-version v1.16.8 --flavor machinepool > my-cluster.yaml
```

The template used for this [flavor](https://cluster-api.sigs.k8s.io/clusterctl/commands/config-cluster.html#flavors)
is located [here](../../templates/cluster-template-machinepool.yaml). For example you can add it to `~/.cluster-api/overrides/infrastructure-aws/0.0.0/cluster-template-machinepool.yaml`

### Example MachinePool, AWSMachinePool and KubeadmConfig Resources
Below is an example of the resources needed to create a pool of EC2 machines orchestrated with
an AWS Auto Scaling Group.
```yaml
---
apiVersion: exp.cluster.x-k8s.io/v1alpha3
kind: MachinePool
metadata:
  name: capa-mp-0
spec:
  clusterName: capa
  replicas: 2
  template:
    spec:
      bootstrap:
        configRef:
          apiVersion: bootstrap.cluster.x-k8s.io/v1alpha3
          kind: KubeadmConfig
          name: capa-mp-0
      clusterName: capa
      infrastructureRef:
        apiVersion: infrastructure.cluster.x-k8s.io/v1alpha3
        kind: AWSMachinePool
        name: capa-mp-0
      version: v1.16.8
---
apiVersion: infrastructure.cluster.x-k8s.io/v1alpha3
kind: AWSMachinePool
metadata:
  name: capa-mp-0
spec:
  minSize: 1
  maxSize: 10
  availabilityZones:
    - us-east-1
  awsLaunchTemplate:
    instanceType: "${AWS_CONTROL_PLANE_MACHINE_TYPE}"
    sshKeyName: "${AWS_SSH_KEY_NAME}"
  subnets:
  - ${AWS_SUBNET}
---
apiVersion: bootstrap.cluster.x-k8s.io/v1alpha3
kind: KubeadmConfig
metadata:
  name: capa-mp-0
  namespace: default
spec:
  joinConfiguration:
    nodeRegistration:
      name: '{{ ds.meta_data.local_hostname }}'
      kubeletExtraArgs:
        cloud-provider: aws
```