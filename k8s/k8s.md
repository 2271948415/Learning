### kubernetes对象
在 Kubernetes 系统中，Kubernetes 对象是持久化的实体。 Kubernetes 使用这些实体去表示整个集群的状态。 具体而言，它们描述了如下信息：
- 哪些容器化应用正在运行（以及在哪些节点上运行）
- 可以被应用使用的资源
- 关于应用运行时行为的策略，比如重启策略、升级策略以及容错策略   

每个 Kubernetes 对象包含两个嵌套的对象字段，它们负责管理对象的配置： 对象 spec（规约） 和对象 status（状态）。 对于具有 spec 的对象，你必须在创建对象时设置其内容，描述你希望对象所具有的特征： 期望状态（Desired State）。  
### pod
Pod 是可以在 Kubernetes 中创建和管理的、最小的可部署的计算单元。Pod 类似于共享名字空间并共享文件系统卷的一组容器。  
```
apiVersion: v1
kind: Pod
metadata:
  name: nginx \\ pod的名称，唯一
spec:
  containers:
  - name: nginx  \\容器的名称
    image: nginx:1.14.2
    ports:
    - containerPort: 80
    

kubectl apply -f simple-pod.yaml  应用pod
```  
### Pod 怎样管理多个容器   
Pod 被设计成支持形成内聚服务单元的多个协作过程（形式为容器）。 Pod 中的容器被自动安排到集群中的同一物理机或虚拟机上，并可以一起进行调度。 容器之间可以共享资源和依赖、彼此通信、协调何时以及何种方式终止自身  
Pod 天生地为其成员容器提供了两种共享资源：网络和存储。  
### Pod 和控制器  
你可以使用工作负载资源来创建和管理多个 Pod。 资源的控制器能够处理副本的管理、上线，并在 Pod 失效时提供自愈能力。 例如，如果一个节点失败，控制器注意到该节点上的 Pod 已经停止工作， 就可以创建替换性的 Pod。调度器会将替身 Pod 调度到一个健康的节点执行。
- Deployment
    - 创建 Deployment 以将 ReplicaSet 上线。ReplicaSet 在后台创建 Pod。 检查 ReplicaSet 的上线状态，查看其是否成功。
    ```
        apiVersion: apps/v1
        kind: Deployment
        metadata:
            name: nginx-deployment
            labels:
                app: nginx
        spec:
            replicas: 3
            selector:
                matchLabels:
                    app: nginx
            template:
                metadata:
                    labels:
                        app: nginx
                spec:
                    containers:
                    - name: nginx
                    image: nginx:1.14.2
                    ports:
                    - containerPort: 80
        \\ 解释
        apiVersion: apps/v1: 这指定了使用的Kubernetes API版本，apps/v1版本用于Deployment对象。
        kind: Deployment: 这表明这是一个Deployment对象，用于在Kubernetes集群中管理Pod的部署和扩展。
        metadata: 这里包含有关Deployment对象的元数据，如名称和标签。
        name: nginx-deployment: 这是Deployment对象的名称，用于在集群中唯一标识该部署。
        labels: 这里定义了标签，用于标识Deployment对象。
        spec: 这是Deployment对象的规范部分，描述了部署的详细信息。
        replicas: 3: 这指定了要创建的Pod副本数量，这里是3个。
        selector: 这指定了用于选择要控制的Pod的标签。
        matchLabels: 这里指定了用于选择Pod的标签，这些标签与Pod模板中定义的标签相匹配。
        template: 这是用于创建Pod的模板。
        metadata: 这里包含有关Pod模板的元数据，如标签。
        labels: 这里定义了Pod的标签，用于标识Pod。
        spec: 这是Pod的规范部分，描述了Pod的详细信息。
        containers: 这里定义了Pod中包含的容器列表。
        name: nginx: 这是容器的名称。
        image: nginx:1.14.2: 这是要在容器中运行的镜像，这里是nginx:1.14.2。
        ports: 这里定义了容器暴露的端口。
        containerPort: 80: 这指定了容器监听的端口，这里是80端口
    ```
  1. 通过运行以下命令创建 Deployment ：   kubectl apply -f xxx.yaml
  2. 运行 kubectl get deployments 检查 Deployment 是否已创建。在检查集群中的 Deployment 时，所显示的字段有：
     - NAME 列出了名字空间中 Deployment 的名称。
     - READY 显示应用程序的可用的“副本”数。显示的模式是“就绪个数/期望个数”。
     - UP-TO-DATE 显示为了达到期望状态已经更新的副本数。
     - AVAILABLE 显示应用可供用户使用的副本数。
     - AGE 显示应用程序运行的时间
  3. 要查看 Deployment 上线状态，运行 kubectl rollout status deployment/nginx-deployment。
  4. 要查看 Deployment 创建的 ReplicaSet（rs），运行 kubectl get rs。
  5. 要查看每个 Pod 自动生成的标签，运行 kubectl get pods --show-labels。 
  #### 更新deployment
  1. 先来更新 nginx Pod 以使用 nginx:1.16.1 镜像，而不是 nginx:1.14.2 镜像。
    ```
    kubectl set image deployment.v1.apps/nginx-deployment nginx=nginx:1.16.1 or kubectl set image deployment/nginx-deployment nginx=nginx:1.16.1 or kubectl edit deployment/nginx-deployment
    ```
  更新时deployment首先创建一个新的replicaset，并将其扩容为1，等待其就绪后将旧的replicaset缩容到2，以此类推直到旧的为0  
  #### 翻转（多 Deployment 动态更新）
  Deployment 控制器每次注意到新的 Deployment 时，都会创建一个 ReplicaSet 以启动所需的 Pod。 如果更新了 Deployment，则控制标签匹配 .spec.selector 但模板不匹配 .spec.template 的 Pod 的现有 ReplicaSet 被缩容。 最终，新的 ReplicaSet 缩放为 .spec.replicas 个副本， 所有旧 ReplicaSet 缩放为 0 个副本。

   当 Deployment 正在上线时被更新，Deployment 会针对更新创建一个新的 ReplicaSet 并开始对其扩容，之前正在被扩容的 ReplicaSet 会被翻转，添加到旧 ReplicaSet 列表 并开始缩容。

   例如，假定你在创建一个 Deployment 以生成 nginx:1.14.2 的 5 个副本，但接下来 更新 Deployment 以创建 5 个 nginx:1.16.1 的副本，而此时只有 3 个 nginx:1.14.2 副本已创建。在这种情况下，Deployment 会立即开始杀死 3 个 nginx:1.14.2 Pod， 并开始创建 nginx:1.16.1 Pod。它不会等待 nginx:1.14.2 的 5 个副本都创建完成后才开始执行变更动作。  

    #### 更改标签选择算符

    #### 回滚deployment
    当 Deployment 不稳定时（例如进入反复崩溃状态）。 默认情况下，Deployment 的所有上线记录都保留在系统中，以便可以随时回滚 （你可以通过修改修订历史记录限制来更改这一约束）。  
  1. 首先，检查 Deployment 修订历史：
    ```
    kubectl rollout history deployment/nginx-deployment
    ```
  2. 要查看修订历史的详细信息
    ```
    kubectl rollout history deployment/nginx-deployment --revision=2
    ```
  3. 回滚之前的版本,使用 --to-revision 来回滚到特定修订版本： 
    ```
    kubectl rollout undo deployment/nginx-deployment --to-revision=?
    ``` 
    #### 缩放 Deployment
    你可以使用如下指令缩放 Deployment：
    ```
    kubectl scale deployment/nginx-deployment --replicas=10
    也可以自动缩放：kubectl autoscale deployment/nginx-deployment --min=10 --max=15 --cpu-percent=80
    ``` 
    #### 暂停、恢复 Deployment 的上线过程       
    使用如下指令暂停上线：
    ```
    kubectl rollout pause deployment/nginx-deployment
    ```
    恢复上线
    ```
    kubectl rollout resume deployment/nginx-deployment
    ```
    #### Deployment 状态
    ##### progressing
    执行下面的任务期间，Kubernetes 标记 Deployment 为进行中（Progressing）_：
  - Deployment 创建新的 ReplicaSet
  - Deployment 正在为其最新的 ReplicaSet 扩容
  - Deployment 正在为其旧有的 ReplicaSet(s) 缩容
  - 新的 Pod 已经就绪或者可用（就绪至少持续了 MinReadySeconds 秒）  
  当上线过程进入“Progressing”状态时，Deployment 控制器会向 Deployment 的 .status.conditions 中添加包含下面属性的状况条目：
  - type: Progressing
  - status: "True"
  - reason: NewReplicaSetCreated | reason: FoundNewReplicaSet | reason: ReplicaSetUpdated  
    ##### complete
    当上线过程进入“Complete”状态时，Deployment 控制器会向 Deployment 的 .status.conditions 中添加包含下面属性的状况条目：  
  - type: Progressing
  - status: "True"
  - reason: NewReplicaSetAvailable  
    ##### false
    超过截止时间后，Deployment 控制器将添加具有以下属性的 Deployment 状况到 Deployment 的 .status.conditions 中：  
  - type: Progressing
  - status: "False"
  - reason: ProgressDeadlineExceeded
    StatefulSet 是用来管理有状态应用的工作负载 API 对象。StatefulSet 用来管理某 Pod 集合的部署和扩缩， 并为这些 Pod 提供持久存储和持久标识符。  
    和 Deployment 类似， StatefulSet 管理基于相同容器规约的一组 Pod。但和 Deployment 不同的是， StatefulSet 为它们的每个 Pod 维护了一个有粘性的 ID。这些 Pod 是基于相同的规约来创建的， 但是不能相互替换：无论怎么调度，每个 Pod 都有一个永久不变的 ID

  StatefulSet 是用来管理有状态应用的工作负载 API 对象。StatefulSet 用来管理某 Pod 集合的部署和扩缩， 并为这些 Pod 提供持久存储和持久标识符。和 Deployment 类似， StatefulSet 管理基于相同容器规约的一组 Pod。但和 Deployment 不同的是， StatefulSet 为它们的每个 Pod 维护了一个有粘性的 ID。这些 Pod 是基于相同的规约来创建的， 但是不能相互替换：无论怎么调度，每个 Pod 都有一个永久不变的 ID  
- StateFulSet 
  - 给定 Pod 的存储必须由 PersistentVolume Provisioner 基于所请求的 storage class 来制备，或者由管理员预先制备。
  - 删除或者扩缩 StatefulSet 并不会删除它关联的存储卷。 这样做是为了保证数据安全，它通常比自动清除 StatefulSet 所有相关的资源更有价值。
    StatefulSet 当前需要无头服务来负责 Pod 的网络标识。你需要负责创建此服务。
  - 当删除一个 StatefulSet 时，该 StatefulSet 不提供任何终止 Pod 的保证。 为了实现 StatefulSet 中的 Pod 可以有序且体面地终止，可以在删除之前将 StatefulSet 缩容到 0。
  - 在默认 Pod 管理策略(OrderedReady) 时使用滚动更新， 可能进入需要人工干预才能修复的损坏状态。
  examples:
  ```
  apiVersion: v1  //service配置
  kind: Service
  metadata:
    name: nginx
    labels:
      app: nginx
  spec:
    ports:
    - port: 80
      name: web
    clusterIP: None
    selector:
      app: nginx

  apiVersion: apps/v1    //statefulset配置
  kind: StatefulSet   //表示创建一个StatefulSet资源。
  metadata:
    name: web    //字段指定了StatefulSet的名称为web
  spec:
    selector:
      matchLabels:
        app: nginx # 必须匹配 .spec.template.metadata.labels   // 字段定义了StatefulSet控制的Pod的标签选择器，这里选择app为nginx的Pod。
    serviceName: "nginx"
    replicas: 3 # 默认值是 1
    minReadySeconds: 10 # 默认值是 0
    template:
      metadata:
        labels:
          app: nginx # 必须匹配 .spec.selector.matchLabels
      spec:
        terminationGracePeriodSeconds: 10
        containers:   //字段定义了Pod中的容器，这里是一个名为nginx的容器，使用了nginx镜像，监听80端口，并挂载了一个名为www的卷
        - name: nginx
          image: registry.k8s.io/nginx-slim:0.8
          ports:
          - containerPort: 80
            name: web
          volumeMounts:
          - name: www
            mountPath: /usr/share/nginx/html
    volumeClaimTemplates:    //用于创建PersistentVolumeClaim（PVC），这里定义了一个名为www的PVC，使用了ReadWriteOnce的访问模式，存储类为my-storage-class，请求1Gi的存储空间。
    - metadata:
      name: www
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: "my-storage-class"
      resources:
        requests:
          storage: 1Gi
  ```
  #### pod选择器
  你必须设置 StatefulSet 的 .spec.selector 字段，使之匹配其在 .spec.template.metadata.labels 中设置的标签。 未指定匹配的 Pod 选择算符将在创建 StatefulSet 期间导致验证错误。  
  #### 数据卷
  你可以设置 .spec.volumeClaimTemplates， 它可以使用 PersistentVolume 制备程序所准备的 PersistentVolumes 来提供稳定的存储。
  #### 最短就绪秒数
  .spec.minReadySeconds 是一个可选字段。 它指定新创建的 Pod 应该在没有任何容器崩溃的情况下运行并准备就绪，才能被认为是可用的  
  #### pod标识  
  StatefulSet Pod 具有唯一的标识，该标识包括顺序标识、稳定的网络标识和稳定的存储。 该标识和 Pod 是绑定的，与该 Pod 调度到哪个节点上无关。
  #### 序号索引 
  对于具有 N 个副本的 StatefulSet，该 StatefulSet 中的每个 Pod 将被分配一个整数序号， 该序号在此 StatefulSet 中是唯一的。默认情况下，这些 Pod 将被赋予从 0 到 N-1 的序号。 StatefulSet 的控制器也会添加一个包含此索引的 Pod 标签：apps.kubernetes.io/pod-index。
  #### 起始序号
  .spec.ordinals 是一个可选的字段，允许你配置分配给每个 Pod 的整数序号。 该字段默认为 nil 值。你必须启用 StatefulSetStartOrdinal 特性门控才能使用此字段。 一旦启用，你就可以配置以下选项：
   - .spec.ordinals.start：如果 .spec.ordinals.start 字段被设置，则 Pod 将被分配从 .spec.ordinals.start 到 .spec.ordinals.start + .spec.replicas - 1 的序号
  #### 稳定的存储
  对于 StatefulSet 中定义的每个 VolumeClaimTemplate，每个 Pod 接收到一个 PersistentVolumeClaim。 在上面的 nginx 示例中，每个 Pod 将会得到基于 StorageClass my-storage-class 制备的 1 GiB 的 PersistentVolume。如果没有指定 StorageClass，就会使用默认的 StorageClass。 当一个 Pod 被调度（重新调度）到节点上时，它的 volumeMounts 会挂载与其 PersistentVolumeClaims 相关联的 PersistentVolume。 请注意，当 Pod 或者 StatefulSet 被删除时，与 PersistentVolumeClaims 相关联的 PersistentVolume 并不会被删除。要删除它必须通过手动方式来完成。
  #### Pod 名称标签
  当 StatefulSet 控制器创建 Pod 时， 它会添加一个标签 statefulset.kubernetes.io/pod-name，该标签值设置为 Pod 名称。 这个标签允许你给 StatefulSet 中的特定 Pod 绑定一个 Service。
  #### Pod 索引标签 
  当 StatefulSet 控制器创建一个 Pod 时， 新的 Pod 会被打上 apps.kubernetes.io/pod-index 标签。标签的取值为 Pod 的序号索引。 此标签使你能够将流量路由到特定索引值的 Pod、使用 Pod 索引标签来过滤日志或度量值等等。 注意要使用这一特性需要启用特性门控 PodIndexLabel，而该门控默认是被启用的
  #### 扩缩容机制
  - 对于包含 N 个 副本的 StatefulSet，当部署 Pod 时，它们是依次创建的，顺序为 0..N-1，
  - 当删除 Pod 时，它们是逆序终止的，顺序为 N-1..0。
  - 在将扩缩操作应用到 Pod 之前，它前面的所有 Pod 必须是 Running 和 Ready 状态。
  - 在一个 Pod 终止之前，所有的继任者必须完全关闭。
  #### 更新策略
  StatefulSet 的 .spec.updateStrategy 字段让你可以配置和禁用掉自动滚动更新 Pod 的容器、标签、资源请求或限制、以及注解。有两个允许的值
  - OnDelete  
    当 StatefulSet 的 .spec.updateStrategy.type 设置为 OnDelete 时， 它的控制器将不会自动更新 StatefulSet 中的 Pod。 用户必须手动删除 Pod 以便让控制器创建新的 Pod，以此来对 StatefulSet 的 .spec.template 的变动作出反应。  
  - RollingUpdate  (滚动更新)
    RollingUpdate 更新策略对 StatefulSet 中的 Pod 执行自动的滚动更新。这是默认的更新策略。当 StatefulSet 的 .spec.updateStrategy.type 被设置为 RollingUpdate 时， StatefulSet 控制器会删除和重建 StatefulSet 中的每个 Pod。 它将按照与 Pod 终止相同的顺序（从最大序号到最小序号）进行，每次更新一个 Pod。
  #### 最大不可以pod
  你可以通过指定 .spec.updateStrategy.rollingUpdate.maxUnavailable 字段来控制更新期间不可用的 Pod 的最大数量
  #### 强制回滚
  如果更新后 Pod 模板配置进入无法运行或就绪的状态（例如， 由于错误的二进制文件或应用程序级配置错误），StatefulSet 将停止回滚并等待。在这种状态下，仅将 Pod 模板还原为正确的配置是不够的。 由于已知问题，StatefulSet 将继续等待损坏状态的 Pod 准备就绪（永远不会发生），然后再尝试将其恢复为正常工作配置。恢复模板后，还必须删除 StatefulSet 尝试使用错误的配置来运行的 Pod。这样， StatefulSet 才会开始使用被还原的模板来重新创建 Pod。
  #### PersistentVolumeClaim 保留
  - whenDeleted  
    配置删除 StatefulSet 时应用的卷保留行为。
  - whenScaled  
    配置当 StatefulSet 的副本数减少时应用的卷保留行为；例如，缩小集合时。 对于你可以配置的每个策略，你可以将值设置为 Delete 或 Retain。
  - Delete  
    对于受策略影响的每个 Pod，基于 StatefulSet 的 volumeClaimTemplate 字段创建的 PVC 都会被删除。 使用 whenDeleted 策略，所有来自 volumeClaimTemplate 的 PVC 在其 Pod 被删除后都会被删除。 使用 whenScaled 策略，只有与被缩减的 Pod 副本对应的 PVC 在其 Pod 被删除后才会被删除。
  - Retain（默认）  
    来自 volumeClaimTemplate 的 PVC 在 Pod 被删除时不受影响。这是此新功能之前的行为

  DaemonSet 确保全部（或者某些）节点上运行一个 Pod 的副本。 当有节点加入集群时， 也会为他们新增一个 Pod 。 当有节点从集群移除时，这些 Pod 也会被回收。删除 DaemonSet 将会删除它创建的所有 Pod。
- DaemonSet
- 

