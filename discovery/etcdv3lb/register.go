package etcdv3lb

import (
	"context"
	"fmt"
	"log"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

type etcdRegister struct {
	Address string
}

func NewRegister(target string) *etcdRegister {
	return &etcdRegister{Address: target}
}

func (s *etcdRegister) Register(info ServiceMetadata) error {

	client, err := clientv3.NewFromURL(s.Address)
	if err != nil {
		log.Println("connect etcd fail:", err)
		return err
	}
	// minimum lease TTL is ttl-second
	resp, err := client.Grant(context.TODO(), int64(info.IntervalTime))
	if err != nil {
		log.Println("创建租约失败:", err)
		return err
	}
	// should get first, if not exist, set it
	_, err = client.Get(context.Background(), info.ServiceName)
	serviceValue := fmt.Sprintf("%s:%d", info.Host, info.Port)
	if err != nil {
		if err == rpctypes.ErrKeyNotFound {
			if _, err := client.Put(context.TODO(), info.ServiceName, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
				log.Printf("s: set service '%s' with ttl to etcd3 failed: %s", info.ServiceName, err.Error())
			}
		} else {
			log.Printf("s: service '%s' connect to etcd3 failed: %s", info.ServiceName, err.Error())
			return err
		}
	} else {
		// refresh set to true for not notifying the watcher
		if _, err := client.Put(context.Background(), info.ServiceName, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
			log.Printf("s: refresh service '%s' with ttl to etcd3 failed: %s", info.ServiceName, err.Error())
			return err
		}
	}
	log.Println("register successful")
	return nil
}

func (s *etcdRegister) UnRegister(info ServiceMetadata) error {
	client, err := clientv3.NewFromURL(s.Address)
	if _, err := client.Delete(context.Background(), info.ServiceName); err != nil {
		log.Printf("s: deregister '%s' failed: %s", info.ServiceName, err.Error())
	} else {
		log.Printf("s: deregister '%s' ok.", info.ServiceName)
	}
	return err
}
