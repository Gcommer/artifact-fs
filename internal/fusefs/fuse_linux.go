package fusefs

import (
	"os"

	"github.com/jacobsa/fuse"
)

func platformMountConfig(cfg *fuse.MountConfig) {
	// FUSEImplFuseT is the zero value and the correct choice for Linux.
	cfg.FuseImpl = fuse.FUSEImplFuseT
	if os.Getenv("ARTIFACT_FS_ALLOW_OTHER") == "1" || os.Getenv("ARTIFACT_FS_ALLOW_OTHER") == "true" {
		if cfg.Options == nil {
			cfg.Options = make(map[string]string)
		}
		cfg.Options["allow_other"] = ""
	}
}

