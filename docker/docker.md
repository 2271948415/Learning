### Docker简介
docker 是一个开源的应用容器引擎，一个完整的docker有以下几部分组成：dockerClient客户端、dockerDaemon守护进程、dockerImage镜像、dockerContainer容器。  
UnionFS(联合文件系统)：是一种分层、轻量级并且高性能的文件系统，他支持对文件系统的修改作为一次提交来一层层叠加。
### docker命令  
```
docker system df  查看容器/镜像/数据卷所占用的空间  
docker run --name 指定名字
            -d 后台运行并返回容器id  
            -i 已交互模式运行容器  
            -t 为容器分配一个伪终端  
            -p 分配端口    
docker log 查看日志
docker top 查看容器内运行进程
docker inspect 查看容器内部细节
docker exec 进入容器
docker attach 进入容器 直接进入容器启动命令的终端，不会启动新的进程，使用exit退出，会导致容器的停止。
docker cp 容器id：容器内路径  目的路径  拷贝文件
docker export 容器id > abc.tar 将容器导出成tar包
cat abc.tar | docker import - 镜像用户/镜像名：镜像版本号  将tar包导入成容器
``` 
### docker net  
- bridge模式：为每个容器分配设置ip等，将容器连接到docker0虚拟网桥，默认为该模式。docker服务默认会创造一个docker0网桥，该网桥的名称为docker0，他在内核层链通了其他的物理或虚拟网卡，这样将所有容器和本地主机都放到同一个物理网络。docker默认指定了docker0接口的ip和子网掩码，让主机和容器之间可以通过网桥相互通信。  
- host模式：直接使用宿主机ip地址与外界进行通信，不再需要额外的nat转换。  
- none模式：禁用网络功能，不进行任何网络配置  
- container模式：新创建的容器不会创建自己的王卡和配置自己的ip，而是和某个容器共享ip、端口范围等。
