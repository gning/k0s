package v1beta1

import "fmt"

// ImageSpec container image settings
type ImageSpec struct {
	Image   string `yaml:"image"`
	Version string `yaml:"version"`
}

// URI build image uri
func (is ImageSpec) URI() string {
	return fmt.Sprintf("%s:%s", is.Image, is.Version)
}

// ClusterImages sets docker images for addon components
type ClusterImages struct {
	Konnectivity  ImageSpec
	MetricsServer ImageSpec
	KubeProxy     ImageSpec
	CoreDNS       ImageSpec

	Calico CalicoImageSpec
}

// CalicoImageSpec config group for calico related image settings
type CalicoImageSpec struct {
	CNI             ImageSpec
	FlexVolume      ImageSpec
	Node            ImageSpec
	KubeControllers ImageSpec
}

// DefaultClusterImages default image settings
func DefaultClusterImages() *ClusterImages {
	return &ClusterImages{
		Konnectivity: ImageSpec{
			Image:   "us.gcr.io/k8s-artifacts-prod/kas-network-proxy/proxy-agent",
			Version: "v0.0.12",
		},
		MetricsServer: ImageSpec{
			Image:   "gcr.io/k8s-staging-metrics-server/metrics-server",
			Version: "v0.3.7",
		},
		KubeProxy: ImageSpec{
			Image:   "k8s.gcr.io/kube-proxy",
			Version: "v1.19.0",
		},
		CoreDNS: ImageSpec{
			Image:   "docker.io/coredns/coredns",
			Version: "1.7.0",
		},
		Calico: CalicoImageSpec{
			CNI: ImageSpec{
				Image:   "calico/cni",
				Version: "v3.16.2",
			},
			FlexVolume: ImageSpec{
				Image:   "calico/pod2daemon-flexvol",
				Version: "v3.16.2",
			},
			Node: ImageSpec{
				Image:   "calico/node",
				Version: "v3.16.2",
			},
			KubeControllers: ImageSpec{
				Image:   "calico/kube-controllers",
				Version: "v3.16.2",
			},
		},
	}
}