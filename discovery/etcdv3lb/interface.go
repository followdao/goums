package etcdv3lb

import "time"

// gRPC service description 服务描述信息
type ServiceMetadata struct {
	// service name 服务名称
	ServiceName string
	// host ip地址
	Host string
	// port 端口
	Port int
	// heart beat 心跳间隔 秒
	IntervalTime time.Duration
}

// service register 服务注册和下线的接口
type RegisterI interface {
	Register(serviceInfo ServiceMetadata) error
	UnRegister(serviceInfo ServiceMetadata) error
}
