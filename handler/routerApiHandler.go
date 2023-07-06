package handler

import (
	"context"
	"encoding/json"
	log "github.com/asim/go-micro/v3/logger"
	routerApi "github.com/coding/routerApi/proto/routerApi"
	router "github.com/yejiabin9/router/proto/router"
)

type RouterApi struct {
	RouterService router.RouterService
}

// routerApi.FindRouterById 通过API向外暴露为/routerApi/findRouterById，接收http请求
// 即：/routerApi/FindRouterById 请求会调用go.micro.api.routerApi 服务的routerApi.FindRouterById 方法
func (e *RouterApi) FindRouterById(ctx context.Context, req *routerApi.Request, rsp *routerApi.Response) error {
	log.Info("Received routerApi.FindRouterById request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/routerApi/FindRouterById'}")
	rsp.Body = string(b)
	return nil
}

// routerApi.AddRouter 通过API向外暴露为/routerApi/AddRouter，接收http请求
// 即：/routerApi/AddRouter 请求会调用go.micro.api.routerApi 服务的routerApi.AddRouter 方法
func (e *RouterApi) AddRouter(ctx context.Context, req *routerApi.Request, rsp *routerApi.Response) error {
	log.Info("Received routerApi.AddRouter request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/routerApi/AddRouter'}")
	rsp.Body = string(b)
	return nil
}

// routerApi.DeleteRouterById 通过API向外暴露为/routerApi/DeleteRouterById，接收http请求
// 即：/routerApi/DeleteRouterById 请求会调用go.micro.api.routerApi 服务的 routerApi.DeleteRouterById 方法
func (e *RouterApi) DeleteRouterById(ctx context.Context, req *routerApi.Request, rsp *routerApi.Response) error {
	log.Info("Received routerApi.DeleteRouterById request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/routerApi/DeleteRouterById'}")
	rsp.Body = string(b)
	return nil
}

// routerApi.UpdateRouter 通过API向外暴露为/routerApi/UpdateRouter，接收http请求
// 即：/routerApi/UpdateRouter 请求会调用go.micro.api.routerApi 服务的routerApi.UpdateRouter 方法
func (e *RouterApi) UpdateRouter(ctx context.Context, req *routerApi.Request, rsp *routerApi.Response) error {
	log.Info("Received routerApi.UpdateRouter request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/routerApi/UpdateRouter'}")
	rsp.Body = string(b)
	return nil
}

// 默认的方法routerApi.Call 通过API向外暴露为/routerApi/call，接收http请求
// 即：/routerApi/call或/routerApi/ 请求会调用go.micro.api.routerApi 服务的routerApi.FindRouterById 方法
func (e *RouterApi) Call(ctx context.Context, req *routerApi.Request, rsp *routerApi.Response) error {
	log.Info("Received routerApi.Call request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问：Call'}")
	rsp.Body = string(b)
	return nil
}
