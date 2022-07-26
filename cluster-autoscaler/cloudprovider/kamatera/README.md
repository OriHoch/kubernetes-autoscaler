# Cluster Autoscaler for Kamatera

The cluster autoscaler for Kamatera scales nodes in a Kamatera cluster.

## Kamatera Kubernetes

[Kamatera](https://www.kamatera.com/express/compute/) supports Kubernetes clusters using our Rancher app
or by creating a self-managed cluster directly on Kamatera compute servers, the autoscaler supports 
both methods.

## Cluster Autoscaler Node Groups

An autoscaler node group is composed of multiple Kamatera servers with the same server configuration.
All servers belonging to a node group are identified by Kamatera server tags `k8sca-CLUSTER_NAME`, `k8scang-NODEGROUP_NAME`.
The cluster and node groups must be specified in the autoscaler cloud configuration file.

## Configuration

The cluster autoscaler only considers the cluster and node groups defined in the configuration file.

You can see an example of the cloud config file at [examples/cluster-autoscaler-secret.yaml](examples/cluster-autoscaler-secret.yaml),

**Important Note:** The cluster and node group names must be 15 characters or less.

it is an INI file with the following fields:

| Key | Value | Mandatory | Default |
|-----|-------|-----------|---------|
| global/kamatera-api-client-id | Kamatera API Client ID | yes | none |
| global/kamatera-api-secret | Kamatera API Secret | yes | none |
| global/cluster-name | **max 15 characters**: distinct string used to set the cluster server tag | yes | none |
| global/default-min-size | default minimum size of a node group (must be > 0) | no | 1 |
| global/default-max-size | default maximum size of a node group | no | 254 |
| global/default-<SERVER_CONFIG_KEY> | replace <SERVER_CONFIG_KEY> with the relevant configuration key | see below | see below |
| nodegroup \"name\" | **max 15 characters**: distinct string within the cluster used to set the node group server tag | yes | none |
| nodegroup \"name\"/min-size | minimum size for a specific node group | no | global/defaut-min-size |
| nodegroup \"name\"/max-size | maximum size for a specific node group | no | global/defaut-min-size |
| nodegroup \"name\"/<SERVER_CONFIG_KEY> | replace <SERVER_CONFIG_KEY> with the relevant configuration key | no | global/default-<SERVER_CONFIG_KEY> |

### Server configuration keys

Following are the supported server configuration keys, see the [example config](examples/cluster-autoscaler-secret.yaml) for more details:

| Key | Value | Mandatory | Default |
|-----|-------|-----------|---------|
| name-prefix | Prefix for all created server names | no | none |
| password | Server root password | no | none |
| ssh-key | Public SSH key to add to the server authorized keys | no | none |
| datacenter | Datacenter ID | yes | none |
| image | Image ID or name | yes | none |
| cpu | CPU type and size identifier | yes | none |
| ram | RAM size in MB | yes | none |
| disk | Disk specifications - see below for details | yes | none |
| dailybackup | boolean - set to true to enable daily backups | no | false |
| managed | boolean - set to true to enable managed services | no | false |
| network | Network specifications - see below for details | yes | none |
| billingcycle | \"hourly\" or \"monthly\" | no | \"hourly\" |
| monthlypackage | For monthly billing only - the monthly network package to use | no | none |
| script-base64 | base64 encoded server initialization script, must be provided to connect the server to the cluster, see below for details | no | none |

### Disk specifications

Server disks are specified using an array of strings which are the same as the cloudcli `--disk` argument
as specified in [cloudcli server create](https://github.com/cloudwm/cloudcli/blob/master/docs/cloudcli_server_create.md).
For multiple disks, include the configuration multiple times, example:

```
[global]
; default for all node groups: single 100gb disk
default-disk = "size=100"

[nodegroup "ng1"]
; this node group will use the default

[nodegroup "ng2"]
; override the default and use 2 disks
disk = "size=100"
disk = "size=200"
```

### Network specifications

Networks are specified using an array of strings which are the same as the cloudcli `--network` argument
as specified in [cloudcli server create](https://github.com/cloudwm/cloudcli/blob/master/docs/cloudcli_server_create.md).
For multiple networks, include the configuration multiple times, example:

```
[global]
; default for all node groups: single public network with auto-assigned ip
default-network = "name=wan,ip=auto"

[nodegroup "ng1"]
; this node group will use the default

[nodegroup "ng2"]
; override the default and attach 2 networks - 1 public and 1 private
network = "name=wan,ip=auto"
network = "name=lan-12345-abcde,ip=auto"
```

### Server Initialization Script

This script is required so that the server will connect to the relevant cluster. The specific script depends on
how you create and manage the cluster. The script needs to be provided as a base64 encoded string.

See below for some common configurations, but the exact script may need to be modified depending on your requirements
and server image.

#### Kamatera Rancher Server Initialization Script

Using Kamatera Rancher you need to get the command to join a server to the cluster. This is available from the
following URL: `https://rancher.domain/v3/clusterregistrationtokens`. The relevant command is available under
`data[].nodeCommand`, if you have a single cluster, it will be the first one. If you have multiple cluster you
will have to locate the relevant cluster from the array using `clusterId`. The command will look like this:

```
sudo docker run -d --privileged --restart=unless-stopped --net=host -v /etc/kubernetes:/etc/kubernetes -v /var/run:/var/run  rancher/rancher-agent:v2.6.4 --server https://rancher.domain --token aaa --ca-checksum bbb
```

You can replace this command in the example script at [examples/server-init-rancher.sh](examples/server-init-rancher.sh)

#### Kubeadm Initialization Script

The example script at [examples/server-init-kubeadm.sh](examples/server-init-kubeadm.sh) can be used as a base for
writing your own script to join the server to your cluster.

## Development

Make sure you are inside the `cluster-autoscaler` path of the [autoscaler repository](https://github.com/kubernetes/autoscaler).

Run tests:

```
go test -v k8s.io/autoscaler/cluster-autoscaler/cloudprovider/kamatera
```

Setup a Kamatera cluster, you can use [this guide](https://github.com/Kamatera/rancher-kubernetes/blob/main/README.md)

Get the cluster kubeconfig and set in local file and set in the `KUBECONFIG` environment variable.
Make sure you are connected to the cluster using `kubectl get nodes`.
Create a cloud config file and set it's path in `CLOUD_CONFIG_FILE` env var.

Build the binary and run it:

```
make build &&\
./cluster-autoscaler-amd64 --cloud-config $CLOUD_CONFIG_FILE --cloud-provider kamatera --kubeconfig $KUBECONFIG -v2
```

Open a new terminal and schedule some pods to trigger scale up:

```
echo '
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test
  labels:
    app: test
spec:
  selector:
    matchLabels:
      app: test
  replicas: 4
  template:
    metadata:
      labels:
        app: test
    spec:
      containers:
      - name: test
        image: alpine
        command: [sleep, "86400"]
        resources:
          requests:
            cpu: "800m"
            memory: "100Mi"
' | kubectl apply -f -
```

Check the pods, 

Create the docker image:

```
make container
```

tag the generated docker image and push it to a registry.
