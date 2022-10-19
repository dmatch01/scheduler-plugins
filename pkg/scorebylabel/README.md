# Overview

This folder holds a tutorial for [KubeCon 2022 Detroit](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/) to write a simple scheduler score plugin, called ScoreByLabel plugin, which prioritize nodes for pods based on user-defined labels.

## Pre-requisites
- A Kubernetes Cluster with 2+ nodes
- [Go](https://golang.org/doc/install) 1.17 installed

## Tutorial to write a ScoreByLabel score plugin
##### 1. Prepare ScoreByLabel Plugin
In this tutorial, we are going to build a Socre Plugin named "ScoreByLabel" that favors nodes with higher scores defined by a specific label.
To start, we create the folder `pkg/scorebylabel` and the file `pkg/scorebylabel/score_by_label.go` in the following structure:

```tree
|- scheduler-plugins
    |- pkg
        |- scorebylabel
            |- score_by_label.go
```
In the `scorebylabel.go` file, we define the `ScoreByLabel` struct. Moreover, the ScorePlugin interface also have the Plugin interface as an embedded field. 
So, we must implement its `Name()` string function
```go
type ScoreByLabel struct {
    handle       framework.Handle
}

const (
    Name = "ScoreByLabel" // Name is the name of the plugin used in Registry and configurations.
)

var _ = framework.ScorePlugin(&ScoreByLabel{})

func (s *ScoreByLabel) Name() string {
    return Name
}
```

##### 2. Prepare `ScoreByLabelArgs` that holds the configurations under `apis/config` folder.
We will then need to declare a new struct called `ScoreByLabelArgs` that will contain the label name that we want to use to prioritize nodes.
We will add the configuration in two places: `apis/config/types.go` and `apis/config/v1beta2/types.go`. 
The `apis/config/types.go` holds the struct we will use in the New function.

Also, the config struct must follow the name pattern <Plugin Name>Args, otherwise, it won't be properly decoded and you will face issues.
In `apis/config/types.go`, we add the following struct.
```Go
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScoreByLabelArgs holds arguments used to configure the ScoreByLabel plugin.
type ScoreByLabelArgs struct {
	metav1.TypeMeta

	// LabelKey is the name of the label to be used for scoring.
	LabelKey string
}
```

In `apis/config/v1beta2/types.go`, we add the following struct.
```go
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScoreByLabelArgs holds arguments used to configure the ScoreByLabel plugin.
type ScoreByLabelArgs struct {
	metav1.TypeMeta `json:",inline"`

	// LabelKey is the name of the label to be used for scoring.
	LabelKey *string `json:"labelKey,omitempty"`
}
```

Furthermore, we will add a new function `SetDefaults_ScoreByLabelArgs` in the `config/v1beta2/defaults.go`. 
The function will set the default values for the ScoreByLabel and users can denote their own node label key if 
it is different from the `DefaultLabelKey`.

```go
const (
	// Defaults for ScoreByLabel plugin
	DefaultLabelKey = "score-by-label"
)

// SetDefaults_ScoreByLabelArgs sets the default parameters for the ScoreByLabel plugin
func SetDefaults_ScoreByLabelArgs(obj *ScoreByLabelArgs) {
	if obj.LabelKey == nil {
		obj.LabelKey = &DefaultLabelKey
	}
}

With the structs added, we need to execute the hack/update-codegen.sh script. 
It will update the generated files with functions as DeepCopy for the added structures.
```bash
$ hack/update-codegen.sh
```

To finish the default values configuration, we need to make sure the function above is registered in the v1beta2 schema. 
Thus, make sure that it is registered in the file `apis/config/v1beta2/register.go` and `apis/config/register.go`.
In the `apis/config/v1beta2/register.go` file, we add the following line:
```go
// addKnownTypes registers known types to the given scheme
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		...
		&ScoreByLabelArgs{},
	)
	return nil
}
```

Similarly, in the `apis/config/register.go` file, we add the following line:
```go
// addKnownTypes registers known types to the given scheme
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		...
		&ScoreByLabelArgs{},
	)
	return nil
}
```

The files that will be updated are shown in the following file structure.
```tree
|- scheduler-plugins
    |- apis
        |- config
            |- types.go
            |- v1beta2
                |- types.go
                |- defaults.go
                |- register.go
                |- zz_generated.deepcopy.go
                |- zz_generated.defaults.go
                |- zz_generated.conversion.go
            |- register.go
            |- zz_generated.deepcopy.go
``` 

With the `ScoreByLabelArgs` struct defined, we can now add the `New` function in the `pkg/scorebylabel/scorebylabel.go` file 
to follow the scheduler framework `PluginFactory` interface.
```go
var LabelKey string

// New initializes a new plugin and returns it.
func New(obj runtime.Object, h framework.Handle) (framework.Plugin, error) {
	var args, ok = obj.(*pluginConfig.ScoreByLabelArgs)
	if !ok {
		return nil, fmt.Errorf("[ScoreByLabelArgs] want args to be of type ScoreByLabelArgs, got %T", obj)
	}

	klog.Infof("[ScoreByLabelArgs] args received. LabelKey: %s", args.LabelKey)
	LabelKey = args.LabelKey

	return &ScoreByLabel{
		handle:     h,
	}, nil
}
```

##### 3. Registering the `ScoreByLabel` plugin
Now we have both `ScoreByLabel` plugin and `ScoreByLabelArgs` struct ready, we then need to register the plugin in the 
`cmd/scheduler/main.go` to include the plugin in the scheduler runtime. In the `main` function, we can register a new
plugin via the `app.NewSchedulerCommand` function with the following line:
```go
	command := app.NewSchedulerCommand(
		// added the scoreByLabel plugin
		app.WithPlugin(scorebylabel.Name, scorebylabel.New),
	)
```
In the `cmd/scheduler/main.go`  file we have the import of sigs.k8s.io/scheduler-plugins/pkg/apis/config/scheme, 
which initializes the scheme with all configurations we have introduced in the pkg/apis/config files.

##### 4. Configure the scheduler to enable only the `ScoreByLabel` plugin
The scheduler built from `scheduler-plugins` can configure scheduler profiles via the `KubeSchedulerConfiguration` struct.
Each profile allows plugins to be enabled, disabled and configured according to the configurations parameters defined by the plugin.
Here is an example configuration of the scheduler to enable `ScoreByLabel` plugin only and disabled all other plugins.
```yaml
apiVersion: kubescheduler.config.k8s.io/v1beta2
kind: KubeSchedulerConfiguration
leaderElection:
  leaderElect: false
profiles:
- schedulerName: scorebylabel
  plugins:
    score:
      enabled:
      - name: ScoreByLabel
      disabled:
        - name: "*" # disable all default plugins
  pluginConfig:
    - name: ScoreByLabel
      args:
        labelKey: "score-by-label"
```

##### 5. Build the scheduler image and run the `scorebylabel` scheduler as a secondary scheduler
1. Build the scheduler image
We can use the provided `Makefile` to build the scheduler image. The `Makefile` will build the scheduler image locally
when we run the following. You can change `LOCAL_REGISTRY=localhost:5000/scheduler-plugins` and `LOCAL_IMAGE=kube-scheduler:latest`
to your own registry and image name.
```bash
make local-image
```

2. Run the scheduler as a secondary scheduler
We can run the scheduler as a secondary scheduler by deploying the built scheduler image in a deployment with all necessary roles.
The deployment yaml file is provided in the `manifests/scorebylabel/install/scorebylabel-scheduler.yaml` as shown below. 
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: scorebylabel-scheduler
  namespace: scorebylabel
  labels:
    app: scorebylabel-scheduler
spec:
  replicas: 1
  selector:
    matchLabels:
      app: scorebylabel-scheduler
  template:
    metadata:
      labels:
        app: scorebylabel-scheduler
    spec:
      volumes:
        - name: etckubernetes
          configMap:
            name: scorebylabel-scheduler-config
      containers:
        - name: kube-scheduler
          image: quay.io/chenw615/scorebylabel:latest
          imagePullPolicy: Always
          args:
            - /bin/kube-scheduler
            - --config=/etc/kubernetes/config.yaml
            - -v=6
          volumeMounts:
            - name: etckubernetes
              mountPath: /etc/kubernetes
```

4. We also mount the
`KubeSchedulerConfiguration` as a configuration file wrapped in a `ConfigMap` to be used in the deployment yaml file. 
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: scorebylabel-scheduler-config
  namespace: scorebylabel
data:
  config.yaml: |
    apiVersion: kubescheduler.config.k8s.io/v1beta2
    kind: KubeSchedulerConfiguration
    leaderElection:
      leaderElect: false
    profiles:
      - schedulerName: scorebylabel
        plugins:
          score:
            enabled:
              - name: ScoreByLabel
            disabled:
              - name: "*" # disable all default plugins
        pluginConfig:
          - name: ScoreByLabel
            args:
              labelKey: "score-by-label"
```

The necessary `ClusterRole` access,  `ServiceAccount` and `ClusterRoleBinding` to grant the scheduler permission can be 
found [here](https://github.com/kubernetes-sigs/scheduler-plugins/blob/master/manifests/install/all-in-one.yaml). We will 
skip the details here.

5. The all-in-one scorebylabel scheduler deployment is given in the `manifests/scorebylabel/install/scorebylabel-scheduler.yaml` file.
We run the following command to deploy the scorebylabel scheduler in the `default` namespace.
```bash
kubectl apply -f manifests/scorebylabel/install/scorebylabel-scheduler.yaml
```

##### 6. Run the testing pods to use the scorebylabel scheduler
1. Let's label the nodes with the label `score-by-label` with the following command.
```bash
kubectl label node <node-name-1> score-by-label=1
kubectl label node <node-name-1> score-by-label=5
kubectl label node <node-name-1> score-by-label=10
```

2. We can create a testing pod to use the `scorebylabel` scheduler by specifying the `schedulerName` in the pod spec as shown below.
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: testpod
spec:
  schedulerName: scorebylabel
  containers:
    - name: testpod
      image: k8s.gcr.io/ubuntu-slim:0.1
      command: ["/bin/sh"]
      args:
        - "-c"
        - "sleep 3600s"
      resources:
        requests:
          cpu: "200m"
          memory: 50Mi
```
We then create the pod and see where the pod is scheduled. We can see that the pod is scheduled to the node with the label.
```bash
kubectl create -f manifests/scorebylabel/tests/testpod.yaml
```

3. We can then check the logs of the scorebylabel scheduler to see the score of each node.
```bash
kubectl logs -f -l app=scorebylabel-scheduler
```
We can see log messages as the following:
```text
I1019 15:39:30.060546       1 schedule_one.go:84] "Attempting to schedule pod" pod="default/testpod"
I1019 15:39:30.061060       1 scorebylabel.go:65] [ScoreByLabel] Label score for node 10.177.222.232 is score-by-label = 10
I1019 15:39:30.061457       1 scorebylabel.go:65] [ScoreByLabel] Label score for node 10.177.222.243 is score-by-label = 5
I1019 15:39:30.061497       1 scorebylabel.go:65] [ScoreByLabel] Label score for node 10.177.222.253 is score-by-label = 1
I1019 15:39:30.061996       1 scorebylabel.go:90] [ScoreByLabel] Nodes final score: [{10.177.222.232 100} {10.177.222.243 50} {10.177.222.253 10}]
```

4. Then when we check the `testpod` events, we will see it is scheduled to a node with highest label score.
```bash
kubectl get events |grep testpod
```

##### 7. Demo video



