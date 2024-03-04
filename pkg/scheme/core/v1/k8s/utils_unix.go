//go:build unix

package k8s

import (
	"os"
	"strings"
	"syscall"

	"k8s.io/klog/v2"
)

// unmountKubeletDirectory unmounts all paths that contain KubeletRunDirectory
func unmountKubeletDirectory(absoluteKubeletRunDirectory string) error {
	raw, err := os.ReadFile("/proc/mounts")
	if err != nil {
		return err
	}
	if !strings.HasSuffix(absoluteKubeletRunDirectory, "/") {
		// trailing "/" is needed to ensure that possibly mounted /var/lib/kubelet is skipped
		absoluteKubeletRunDirectory += "/"
	}

	mounts := strings.Split(string(raw), "\n")
	for _, mount := range mounts {
		m := strings.Split(mount, " ")
		if len(m) < 2 || !strings.HasPrefix(m[1], absoluteKubeletRunDirectory) {
			continue
		}
		if err = syscall.Unmount(m[1], 0); err != nil {
			klog.Warningf("[reset] Failed to unmount mounted directory in %s: %s", absoluteKubeletRunDirectory, m[1])
		}
	}
	return nil
}
