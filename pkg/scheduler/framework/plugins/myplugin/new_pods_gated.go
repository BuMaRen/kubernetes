package myplugin

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// 【插件】PreEnqueue
// 该扩展点将在Pod被添加到内部活动队列之前被调用，在此队列中Pod被标记为准备好进行调度。
// 只有当所有PreEnqueue插件返回Success时，Pod才允许进入活动队列。
// 否则它将被放置在内部无法调度的Pod列表中，也不会获得Unschedulable状态。
func (mp *MyPlugin) PreEnqueue(ctx context.Context, p *v1.Pod) *framework.Status {
	// 这里可以添加一些逻辑来决定是否允许 Pod 进入活动队列
	// 例如，检查 Pod 的标签、注释或其他属性
	// 如果不允许 Pod 进入活动队列，可以返回一个错误状态

	if _, ok := p.Annotations["testPreEnqueueForError"]; ok {
		// 这里可以添加一些逻辑来决定是否允许 Pod 进入活动队列
		// 例如，检查 Pod 的标签、注释或其他属性
		// 如果不允许 Pod 进入活动队列，可以返回一个错误状态
		return framework.NewStatus(framework.Unschedulable, "Pod is not allowed to enter active queue")
	}
	return framework.NewStatus(framework.Success, "Pod is allowed to enter active queue")
}
