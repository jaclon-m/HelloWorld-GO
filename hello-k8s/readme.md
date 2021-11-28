作业要求：
一：
编写 Kubernetes 部署脚本将 httpserver 部署到 kubernetes 集群，以下是你可以思考的维度
优雅启动
优雅终止
资源需求和 QoS 保证
探活
日常运维需求，日志等级
配置和代码分离

kubectl apply -f httpserver-deploy.yaml

二：
Service
Ingress
如何确保整个应用高可用
如何通过证书保证httpserver的通讯安全

生成证书:ingress.md
kubectl apply -f secret.yaml kubectl apply -f nginx-ingress-deployment.yaml kubectl apply -f ingress.yaml