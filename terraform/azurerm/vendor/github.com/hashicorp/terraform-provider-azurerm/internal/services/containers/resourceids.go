package containers

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Cluster -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ContainerService/managedClusters/cluster1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=NodePool -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ContainerService/managedClusters/cluster1/agentPools/pool1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ContainerGroup -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ContainerInstance/containerGroups/containerGroup1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ContainerRegistryScopeMap -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ContainerRegistry/registries/registry1/scopeMaps/scopeMap1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ContainerRegistryTask -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.ContainerRegistry/registries/registry1/tasks/task1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ContainerRegistryToken -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ContainerRegistry/registries/registry1/tokens/token1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Registry -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.ContainerRegistry/registries/registry1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Webhook -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.ContainerRegistry/registries/registry1/webhooks/webhook1
