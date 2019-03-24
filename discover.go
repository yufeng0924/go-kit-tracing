package tracing

import (
	"context"

	log "github.com/go-kit/kit/log"
	sdetcd "github.com/go-kit/kit/sd/etcd"
)

// RegisterService register server.
func RegisterService(logger log.Logger, etcdServer string, prefix string, instance string) (*sdetcd.Registrar, error) {
	var (
		key = prefix + instance
	)
	client, err := sdetcd.NewClient(context.Background(), []string{etcdServer}, sdetcd.ClientOptions{})
	if err != nil {
		return nil, err
	}
	registrar := sdetcd.NewRegistrar(client, sdetcd.Service{
		Key:   key,
		Value: instance,
	}, logger)
	registrar.Register()
	return registrar, nil
}

// DiscoverServer get the server address via etcd
func DiscoverServer(logger log.Logger, etcdServer string, prefix string) (string, error) {

	client, err := sdetcd.NewClient(context.Background(), []string{etcdServer}, sdetcd.ClientOptions{})
	if err != nil {
		logger.Log("cannot connect to etcd", err.Error())
		return "", err
	}
	entries, err := client.GetEntries(prefix)
	if err != nil || len(entries) == 0 {
		return "", err
	}
	return entries[0], nil
}
