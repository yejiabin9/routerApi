package handler

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/asim/go-micro/v3/logger"
	"github.com/sirupsen/logrus"
	router "github.com/yejiabin9/router/proto/router"
	form "github.com/yejiabin9/routerApi/from"
	routerApi "github.com/yejiabin9/routerApi/proto/routerApi"
	"strconv"
)

type RouterApi struct {
	RouterService router.RouterService
}

// routerApi.FindRouterById 通过API向外暴露为/routerApi/findRouterById，接收http请求
// 即：/routerApi/FindRouterById 请求会调用go.micro.api.routerApi 服务的routerApi.FindRouterById 方法
func (e *RouterApi) FindRouterById(ctx context.Context, req *routerApi.Request, rsp *routerApi.Response) error {
	log.Info("Received routerApi.FindRouterById request")
	if _, ok := req.Get["route_id"]; !ok {
		rsp.StatusCode = 500
		return errors.New("参数异常")
	}
	//获取 route id
	routeIdString := req.Get["route_id"].Values[0]
	routeId, err := strconv.ParseInt(routeIdString, 10, 64)
	if err != nil {
		logrus.Error(err)
		return err
	}
	//获取route信息
	routeInfo, err := e.RouterService.FindRouterByID(ctx, &router.RouterId{
		Id: routeId,
	})
	if err != nil {
		logrus.Error(err)
		return err
	}
	//返回route结果
	rsp.StatusCode = 200
	b, _ := json.Marshal(routeInfo)
	rsp.Body = string(b)
	return nil
}

// routerApi.AddRouter 通过API向外暴露为/routerApi/AddRouter，接收http请求
// 即：/routerApi/AddRouter 请求会调用go.micro.api.routerApi 服务的routerApi.AddRouter 方法
func (e *RouterApi) AddRouter(ctx context.Context, req *routerApi.Request, rsp *routerApi.Response) error {
	log.Info("Received routerApi.AddRouter request")
	addRouteInfo := &router.RouterInfo{}
	routePathName, ok := req.Post["router_path_name"]
	if ok && len(routePathName.Values) > 0 {
		port, err := strconv.ParseInt(req.Post["router_backend_service_port"].Values[0], 10, 32)
		if err != nil {
			logrus.Error(err)
			return err
		}
		//这里如果有多个Path需要处理多个
		routePath := &router.RouterPath{
			RouterPathName:           req.Post["router_path_name"].Values[0],
			RouterBackendService:     req.Post["router_backend_service"].Values[0],
			RouterBackendServicePort: int32(port),
		}
		//合并
		addRouteInfo.RouterPath = append(addRouteInfo.RouterPath, routePath)
	}
	form.FormToSvcStruct(req.Post, addRouteInfo)
	response, err := e.RouterService.AddRouter(ctx, addRouteInfo)
	if err != nil {
		logrus.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	return nil
}

// routerApi.DeleteRouterById 通过API向外暴露为/routerApi/DeleteRouterById，接收http请求
// 即：/routerApi/DeleteRouterById 请求会调用go.micro.api.routerApi 服务的 routerApi.DeleteRouterById 方法
func (e *RouterApi) DeleteRouterById(ctx context.Context, req *routerApi.Request, rsp *routerApi.Response) error {
	log.Info("Received routerApi.DeleteRouterById request")
	if _, ok := req.Get["route_id"]; !ok {
		rsp.StatusCode = 500
		return errors.New("参数异常")
	}
	//获取 route id
	routeIdString := req.Get["route_id"].Values[0]
	routeId, err := strconv.ParseInt(routeIdString, 10, 64)
	if err != nil {
		logrus.Error(err)
		return err
	}
	//调用route 删除服务
	response, err := e.RouterService.DeleteRouter(ctx, &router.RouterId{
		Id: routeId,
	})
	if err != nil {
		logrus.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(response)
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
	allRoute, err := e.RouterService.FindAllRouter(ctx, &router.FindAll{})
	if err != nil {
		logrus.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(allRoute)
	rsp.Body = string(b)
	return nil
}
