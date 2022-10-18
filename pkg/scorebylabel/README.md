# Overview

This folder holds a tutorial for [KubeCon 2022 Detroit](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/) to write a simple scheduler score plugin, called ScoreByLabel plugin, which prioritize nodes for pods based on user-defined labels.

## Pre-requisites
- A Kubernetes Cluster with 2+ nodes
- [Go](https://golang.org/doc/install) 1.17 installed

## Tutorial to write a ScoreByLabel score plugin
1. Prepare ScoreByLabel Plugin

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

2. Prepare `ScoreByLabelArgs` that holds the configurations under `apis/config` folder.

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
