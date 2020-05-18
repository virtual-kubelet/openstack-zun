package main

import (
	"github.com/virtual-kubelet/openstack-zun"
	"github.com/virtual-kubelet/virtual-kubelet/providers"
)

func registerACI(s *providers.Store) error {
	return s.Register("openstack", func(cfg providers.InitConfig) (providers.Provider, error) {
		return azure.NewZunProvider(cfg.ConfigPath, cfg.ResourceManager, cfg.NodeName, cfg.OperatingSystem, cfg.DaemonPort)
	})
}
