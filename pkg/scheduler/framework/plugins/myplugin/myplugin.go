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

//【插件】PreEnqueue
// 该扩展点将在Pod被添加到内部活动队列之前被调用，在此队列中Pod被标记为准备好进行调度。
// 只有当所有PreEnqueue插件返回Success时，Pod才允许进入活动队列。
// 否则它将被放置在内部无法调度的Pod列表中，也不会获得Unschedulable状态。

//【接口】EnqueueExtension
// 插件可以在此接口上根据集群中的变化来控制是否重新尝试调度被插件拒绝的Pod。
// 实现了PreEnqueue、PreFilter、Filter、Reserve或Permit的插件应实现此接口

//【回调函数】QueueingHint
// 用于决定是否将Pod重新排队到活跃队列或回退队列。
// 每当集群中发生某种事件或变化时，此函数就会被执行。
// 当QueueingHint发现事件可能使Pod可调度时，Pod将被放入活跃队列或回退队列，以便调度器可以重新尝试调度Pod。

//【插件】QueueSort
// 这些插件用于对调度队列中的Pod进行排序。
// 队列排序插件本质上提供Less(Pod1, Pod2)函数。
// 一次只能启动一个队列插件。

//【插件】PreFilter——检查环境和自身
// 这些插件用于预处理Pod的相关信息，或者检查集群或Pod必须满足的某些条件。
// 如果PreFilter插件返回错误，则调度周期将终止。

//【插件】Filter
// 这些插件用于过滤出不能运行该Pod的节点。
// 对于每个节点，调度器将按照其配置顺序调用这些过滤插件。
// 如果任何过滤插件将节点标记为不可行，则不会为该节点调用剩下的过滤插件。节点可以被同时进行评估。

//【插件】PostFilter
// 在Filter阶段后调用，但仅在该Pod没有可行的节点时调用。插件按其配置的顺序调用。
// 如果任何PostFilter插件标记节点为“Schedulable”，则其余的插件不会调用。
// 典型的PostFilter实现是抢占，试图通过抢占其他Pod的资源使该Pod可以调度。

//【插件】PreScore
// 这些插件用于执行“前置评分（pre-scoring）”工作，即生成一个可共享状态供Score插件使用。
// 如果PreScore插件返回错误，则调度周期将终止。

//【插件】Score
// 这些插件用于对通过Filter阶段的节点进行排序。
// 调度器将为每个节点调用每个Score插件。将有一个定义明确的整数范围，代表最小和最大分数。
// 在NormalizeScore阶段之后，调度器将根据配置的插件权重合并所有插件的节点分数。

//【插件】NormalizeSocre
// 这些插件用于在调度器计算Node排名之前修改分数。
// 在此扩展点注册的插件被调用时会使用同一插件的Score结果，每个插件在每个调度周期调用一次。
// 如果任何NormalizeScore返回错误，则调度阶段将终止。

//【插件】Reserve
// 实现了Reserve接口的插件，拥有两个方法，即Reserve和Unreserve，他们分别支持两个名为Reserve和Unreserve的信息传递性质的调度阶段。
// 维护运行时状态的插件（又称"有状态插件"）应该使用这两个阶段，以便在节点上的资源被保留和解除保留给特定的Pod时，得到调度器的通知。
// Reserve阶段发生在调度器实际将一个Pod绑定到其指定节点之前。它的存在是为了防止在调度器等待绑定成功时发生竞争情况。
// 每个Reserve插件的Reserve方法可能成功，也可能失败；如果一个Reserve方法调用失败，后面的插件就不会被执行Reserve阶段被认为失败。
// 如果所有插件的Reserve方法都成功了，Reserve阶段就被认为是成功的，剩下的调度周期和绑定周期就会被执行。
// 如果Reserve阶段或后续阶段失败了，则触发Unreserve阶段。发生这种情况时，所有Reserve插件的Unreserve方法将按照Reserve方法调用的相反顺序执行。这个阶段的存在是为了清理与保留的Pod相关的状态。
// Reserve插件中Unreserve方法的实现必须是幂等的，并且不能失败。
// 这个是调度周期的最后一步。一旦Pod处于保留状态，它将在绑定周期结束时触发Unreserve插件（失败时）或PostBind插件（成功时）。

//【插件】Permit
// Permit插件在每个Pod调度周期的最后调用，用于防止或延迟Pod的绑定。
// 一个允许插件可以做以下三件事之一
// 1. 批准：一旦所有Permit插件批准Pod后，该Pod将被发送以进行绑定。
// 2. 拒绝：如果任何Permit拒绝Pod，则该Pod将被返回到调度队列。这将触发Reserve插件中的Unreserve阶段。
// 3. 等待（带有超时的）：如果一个Permit插件返回“等待”结果，则Pod将保持在一个内部的“等待中”的Pod列表，同时该Pod的绑定周期启动时即直接阻塞直到得到批准。如果超时发生，等待变成拒绝，并且Pod将返回调度队列，从而触发Reserve插件中的Unreserve阶段。
// 一旦Pod被批准了，它将发送到PreBind阶段。

//【插件】PreBind
// 这些插件用于执行Pod绑定前所需的所有工作。
// 例如，一个PreBind插件可能需要制备网络卷并且在允许Pod运行在该节点之前将其挂载到目标节点上。
// 如果任何PreBind插件返回错误，则Pod将被拒绝并且回退到调度队列中。

//【插件】Bind
// Bind插件用于将Pod绑定到节点上，直到所有的PreBind插件都完成，Bind插件才会被调用。
// 各Bind插件按照配置顺序被调用。Bind插件可以选择是否处理指定的Pod。如果某Bind插件选择处理某Pod，剩余的Bind插件将被跳过。

//【插件】PostBind
// 这是个信息传递性质的接口。
// PostBind插件在Pod成功绑定后被调用。这是绑定周期的结尾，可用于清理相关的资源。

// fts是特性门控
func New(_ context.Context, plArgs runtime.Object, h framework.Handle, fts feature.Features) (framework.Plugin, error) {
	return &MyPlugin{}, nil
}
