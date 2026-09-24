/*
 * Vencord Italiano Installer
 * Copyright (c) 2026 contributors
 * SPDX-License-Identifier: GPL-3.0-only
 */

package release

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const apiURL = "https://api.github.com/repos/davidev3p0/Vencord-Italiano/releases/latest"

var requiredFiles = []string{"patcher.js", "preload.js", "renderer.js", "renderer.css"}

type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
	Size               int64  `json:"size"`
	Digest             string `json:"digest"`
}

type Release struct {
	Name    string  `json:"name"`
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

var httpClient = &http.Client{Timeout: 45 * time.Second}

func FetchLatest(ctx context.Context) (Release, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return Release{}, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "Vencord-Italiano-Installer")

	res, err := httpClient.Do(req)
	if err != nil {
		return Release{}, fmt.Errorf("contatto GitHub: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Release{}, fmt.Errorf("GitHub ha risposto %s", res.Status)
	}

	var r Release
	if err := json.NewDecoder(io.LimitReader(res.Body, 2<<20)).Decode(&r); err != nil {
		return Release{}, fmt.Errorf("leggo metadata release: %w", err)
	}
	if r.TagName == "" {
		return Release{}, errors.New("release GitHub senza tag")
	}
	return r, nil
}

func InstallLatest(ctx context.Context, destDir string) (Release, error) {
	r, err := FetchLatest(ctx)
	if err != nil {
		return Release{}, err
	}

	assets := make(map[string]Asset, len(r.Assets))
	for _, a := range r.Assets {
		assets[a.Name] = a
	}
	for _, name := range requiredFiles {
		a, ok := assets[name]
		if !ok {
			return Release{}, fmt.Errorf("asset mancante nella release: %s", name)
		}
		if !strings.HasPrefix(strings.ToLower(a.Digest), "sha256:") {
			return Release{}, fmt.Errorf("asset %s senza digest SHA-256 pubblicato da GitHub", name)
		}
	}

	parent := filepath.Dir(destDir)
	if err := os.MkdirAll(parent, 0o755); err != nil {
		return Release{}, fmt.Errorf("creo cartella Vencord: %w", err)
	}
	tmpDir, err := os.MkdirTemp(parent, ".dist-new-")
	if err != nil {
		return Release{}, fmt.Errorf("creo cartella temporanea: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	for _, name := range requiredFiles {
		if err := downloadAsset(ctx, assets[name], filepath.Join(tmpDir, name)); err != nil {
			return Release{}, err
		}
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte("{}\n"), 0o644); err != nil {
		return Release{}, fmt.Errorf("scrivo package.json: %w", err)
	}

	backupDir := destDir + ".installer-backup"
	_ = os.RemoveAll(backupDir)

	if _, err := os.Stat(destDir); err == nil {
		if err := os.Rename(destDir, backupDir); err != nil {
			return Release{}, fmt.Errorf("backup della build corrente fallito; chiudi Discord e riprova: %w", err)
		}
	}
	if err := os.Rename(tmpDir, destDir); err != nil {
		if _, statErr := os.Stat(backupDir); statErr == nil {
			_ = os.Rename(backupDir, destDir)
		}
		return Release{}, fmt.Errorf("installazione nuova build fallita: %w", err)
	}
	_ = os.RemoveAll(backupDir)

	return r, nil
}

func downloadAsset(ctx context.Context, asset Asset, outPath string) error {
	if asset.Size <= 0 || asset.Size > 64<<20 {
		return fmt.Errorf("dimensione non valida per %s", asset.Name)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, asset.BrowserDownloadURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Vencord-Italiano-Installer")

	res, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("download %s: %w", asset.Name, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("download %s: HTTP %s", asset.Name, res.Status)
	}

	f, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	h := sha256.New()
	written, copyErr := io.Copy(io.MultiWriter(f, h), io.LimitReader(res.Body, 64<<20))
	closeErr := f.Close()
	if copyErr != nil {
		return fmt.Errorf("download %s incompleto: %w", asset.Name, copyErr)
	}
	if closeErr != nil {
		return closeErr
	}
	if written != asset.Size {
		return fmt.Errorf("download %s: attesi %d byte, ricevuti %d", asset.Name, asset.Size, written)
	}

	expected := strings.TrimPrefix(strings.ToLower(asset.Digest), "sha256:")
	actual := hex.EncodeToString(h.Sum(nil))
	if actual != expected {
		return fmt.Errorf("SHA-256 non valido per %s", asset.Name)
	}
	return nil
}
