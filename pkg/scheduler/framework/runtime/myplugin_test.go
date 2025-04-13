package runtime

import (
	"context"
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2/ktesting"
	"k8s.io/kubernetes/pkg/scheduler/apis/config"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type myTestPlugin struct{}

func (mtp *myTestPlugin) Name() string {
	return "MyPlugin"
}

func myTestPluginNew(_ context.Context, _ runtime.Object, f framework.Handle) (framework.Plugin, error) {
	return &myTestPlugin{}, nil
}

func TestMypluginNewFramework(t *testing.T) {
	tests := []struct {
		name    string
		plugins *config.Plugins
		wantErr error
	}{
		{
			name: "my plugin register",
			plugins: &config.Plugins{
				Score: config.PluginSet{
					Enabled: []config.Plugin{
						{Name: "MyPlugin"},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, ctx := ktesting.NewTestContext(t)
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()
			reg := make(Registry)
			err := reg.Register("MyPlugin", myTestPluginNew)
			if err != nil {
				t.Fatal("Unexpected error")
			}
			fwk, err := NewFramework(ctx, reg, &config.KubeSchedulerProfile{Plugins: tc.plugins})
			if err != nil {
				t.Fatal("Unexpected error")
			}
			defer fwk.Close()
			if err != nil {
				t.Error("Unexpected error")
			}

		})
	}
}
