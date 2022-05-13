package common

import (
	"github.com/containers/podman/v4/cmd/podman/registry"
	"github.com/containers/podman/v4/libpod/define"
	"github.com/containers/podman/v4/pkg/domain/entities"
)

func ulimits() []string {
	if !registry.IsRemote() {
		return containerConfig.Ulimits()
	}
	return nil
}

func cgroupConfig() string {
	if !registry.IsRemote() {
		return containerConfig.Cgroups()
	}
	return ""
}

func devices() []string {
	if !registry.IsRemote() {
		return containerConfig.Devices()
	}
	return nil
}

func Env() []string {
	if !registry.IsRemote() {
		return containerConfig.Env()
	}
	return nil
}

func initPath() string {
	if !registry.IsRemote() {
		return containerConfig.InitPath()
	}
	return ""
}

func pidsLimit() int64 {
	if !registry.IsRemote() {
		return containerConfig.PidsLimit()
	}
	return -1
}

func policy() string {
	if !registry.IsRemote() {
		return containerConfig.Engine.PullPolicy
	}
	return ""
}

func shmSize() string {
	if !registry.IsRemote() {
		return containerConfig.ShmSize()
	}
	return ""
}

func volumes() []string {
	if !registry.IsRemote() {
		return containerConfig.Volumes()
	}
	return nil
}

func LogDriver() string {
	if !registry.IsRemote() {
		return containerConfig.Containers.LogDriver
	}
	return ""
}

// DefineCreateDefault is used to initialize ctr create options before flag initialization
func DefineCreateDefaults(opts *entities.ContainerCreateOptions) {
	opts.LogDriver = LogDriver()
	opts.CgroupsMode = cgroupConfig()
	opts.MemorySwappiness = -1
	opts.ImageVolume = containerConfig.Engine.ImageVolumeMode
	opts.Pull = policy()
	opts.ReadOnlyTmpFS = true
	opts.SdNotifyMode = define.SdNotifyModeContainer
	opts.StopTimeout = containerConfig.Engine.StopTimeout
	opts.Systemd = "true"
	opts.Timezone = containerConfig.TZ()
	opts.Umask = containerConfig.Umask()
	opts.Ulimit = ulimits()
	opts.SeccompPolicy = "default"
	opts.Volume = volumes()
}
