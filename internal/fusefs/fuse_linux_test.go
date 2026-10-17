//go:build linux

package fusefs

import (
	"testing"

	"github.com/jacobsa/fuse"
)

func TestPlatformMountConfigAllowOther(t *testing.T) {
	t.Run("disabled by default", func(t *testing.T) {
		t.Setenv("ARTIFACT_FS_ALLOW_OTHER", "")
		cfg := &fuse.MountConfig{}
		platformMountConfig(cfg)
		if cfg.Options != nil {
			if _, ok := cfg.Options["allow_other"]; ok {
				t.Fatalf("allow_other should not be set by default")
			}
		}
	})

	t.Run("enabled via env 1", func(t *testing.T) {
		t.Setenv("ARTIFACT_FS_ALLOW_OTHER", "1")
		cfg := &fuse.MountConfig{}
		platformMountConfig(cfg)
		if cfg.Options == nil || cfg.Options["allow_other"] != "" {
			t.Fatalf("expected allow_other in cfg.Options, got %v", cfg.Options)
		}
	})

	t.Run("enabled via env true", func(t *testing.T) {
		t.Setenv("ARTIFACT_FS_ALLOW_OTHER", "true")
		cfg := &fuse.MountConfig{}
		platformMountConfig(cfg)
		if cfg.Options == nil || cfg.Options["allow_other"] != "" {
			t.Fatalf("expected allow_other in cfg.Options, got %v", cfg.Options)
		}
	})
}
