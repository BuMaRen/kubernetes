package myplugin

import (
	"context"

	v1 "k8s.io/api/core/v1"
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

func (mp *MyPlugin) Score(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) (int64, *framework.Status) {
	return 100, nil
}

func (mp *MyPlugin) NormalizeScore(ctx context.Context, state *framework.CycleState, p *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	return nil
}

func (mp *MyPlugin) ScoreExtensions() framework.ScoreExtensions {
	return mp
}

// // Score is called on each filtered node. It must return success and an integer
// // indicating the rank of the node. All scoring plugins must return success or
// // the pod will be rejected.
// Score(ctx context.Context, state *CycleState, p *v1.Pod, nodeName string) (int64, *Status)

// // ScoreExtensions returns a ScoreExtensions interface if it implements one, or nil if does not.
// ScoreExtensions() ScoreExtensions

// fts是特性门控
func New(_ context.Context, plArgs runtime.Object, h framework.Handle, fts feature.Features) (framework.Plugin, error) {
	return &MyPlugin{}, nil
}
