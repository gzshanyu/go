package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"time"
)

type Connect struct {
	config clientv3.Config
	client *clientv3.Client
}

func NewConnect(addr []string) (*Connect, error) {
	var (
		err  error
		conf clientv3.Config
		conn *Connect
		cli  *clientv3.Client
	)
	// 配置客户端
	conf = clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}
	// 建立连接
	if cli, err = clientv3.New(conf); err != nil {
		return nil, err
	}

	conn = &Connect{
		config: conf,
		client: cli,
	}

	return conn, nil
}
