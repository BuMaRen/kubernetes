package myplugin

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/feature"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/names"
)

const Name = names.MyPlugin

type MyPlugin struct {
}

func (mp *MyPlugin) Name() string {
	return Name
}

// fts是特性门控
func New(_ context.Context, plArgs runtime.Object, h framework.Handle, fts feature.Features) (framework.Plugin, error) {
	return &MyPlugin{}, nil
}
