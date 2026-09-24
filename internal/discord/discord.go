/*
 * Vencord Italiano Installer
 * Copyright (c) 2026 contributors
 * SPDX-License-Identifier: GPL-3.0-only
 */

package discord

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/davidev3p0/Vencord-Italiano-Installer/internal/asar"
)

type Client struct {
	Name         string
	Branch       string
	Root         string
	AppDir       string
	ResourcesDir string
	Patched      bool
}

func Detect() ([]Client, error) {
	local := os.Getenv("LOCALAPPDATA")
	if local == "" {
		return nil, errors.New("LOCALAPPDATA non disponibile")
	}

	defs := []struct {
		name   string
		branch string
		dir    string
	}{
		{"Discord Stable", "stable", "Discord"},
		{"Discord Canary", "canary", "DiscordCanary"},
		{"Discord PTB", "ptb", "DiscordPTB"},
	}

	var clients []Client
	for _, d := range defs {
		root := filepath.Join(local, d.dir)
		appDir, resources, ok := latestApp(root)
		if !ok {
			continue
		}
		_, backupErr := os.Stat(filepath.Join(resources, "_app.asar"))
		clients = append(clients, Client{
			Name:         d.name,
			Branch:       d.branch,
			Root:         root,
			AppDir:       appDir,
			ResourcesDir: resources,
			Patched:      backupErr == nil,
		})
	}
	return clients, nil
}

func latestApp(root string) (string, string, bool) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return "", "", false
	}

	var apps []string
	for _, e := range entries {
		if e.IsDir() && strings.HasPrefix(e.Name(), "app-") {
			apps = append(apps, e.Name())
		}
	}
	sort.Slice(apps, func(i, j int) bool {
		return versionGreater(strings.TrimPrefix(apps[i], "app-"), strings.TrimPrefix(apps[j], "app-"))
	})

	for _, name := range apps {
		appDir := filepath.Join(root, name)
		resources := filepath.Join(appDir, "resources")
		appAsar := filepath.Join(resources, "app.asar")
		backup := filepath.Join(resources, "_app.asar")
		if fileExists(appAsar) || fileExists(backup) {
			return appDir, resources, true
		}
	}
	return "", "", false
}

func versionGreater(a, b string) bool {
	pa := strings.Split(a, ".")
	pb := strings.Split(b, ".")
	n := len(pa)
	if len(pb) > n {
		n = len(pb)
	}
	for i := 0; i < n; i++ {
		ai, bi := 0, 0
		if i < len(pa) {
			ai, _ = strconv.Atoi(pa[i])
		}
		if i < len(pb) {
			bi, _ = strconv.Atoi(pb[i])
		}
		if ai != bi {
			return ai > bi
		}
	}
	return false
}

func (c *Client) Patch(patcherPath string) error {
	appAsar := filepath.Join(c.ResourcesDir, "app.asar")
	backup := filepath.Join(c.ResourcesDir, "_app.asar")
	tmp := filepath.Join(c.ResourcesDir, "app.asar.vencord-italiano.tmp")
	_ = os.Remove(tmp)

	if fileExists(backup) {
		if fileExists(appAsar) {
			if err := os.Rename(appAsar, tmp); err != nil {
				return fmt.Errorf("impossibile sostituire app.asar; chiudi Discord: %w", err)
			}
		}
		if err := asar.Write(appAsar, patcherPath); err != nil {
			if fileExists(tmp) {
				_ = os.Rename(tmp, appAsar)
			}
			return err
		}
		_ = os.Remove(tmp)
		c.Patched = true
		return nil
	}

	if !fileExists(appAsar) {
		return errors.New("app.asar originale non trovato")
	}
	if err := os.Rename(appAsar, backup); err != nil {
		return fmt.Errorf("backup app.asar fallito; chiudi Discord: %w", err)
	}
	if err := asar.Write(appAsar, patcherPath); err != nil {
		_ = os.Rename(backup, appAsar)
		return err
	}
	c.Patched = true
	return nil
}

func (c *Client) Unpatch() error {
	appAsar := filepath.Join(c.ResourcesDir, "app.asar")
	backup := filepath.Join(c.ResourcesDir, "_app.asar")
	tmp := filepath.Join(c.ResourcesDir, "app.asar.vencord-italiano.tmp")
	_ = os.Remove(tmp)

	if !fileExists(backup) {
		return errors.New("Vencord non risulta installato su questo client")
	}
	if fileExists(appAsar) {
		if err := os.Rename(appAsar, tmp); err != nil {
			return fmt.Errorf("chiudi Discord prima di disinstallare: %w", err)
		}
	}
	if err := os.Rename(backup, appAsar); err != nil {
		if fileExists(tmp) {
			_ = os.Rename(tmp, appAsar)
		}
		return fmt.Errorf("ripristino app.asar originale fallito: %w", err)
	}
	_ = os.Remove(tmp)
	c.Patched = false
	return nil
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
