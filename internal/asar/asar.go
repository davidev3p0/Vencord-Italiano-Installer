/*
 * Vencord Italiano Installer
 * Copyright (c) 2026 contributors
 * SPDX-License-Identifier: GPL-3.0-only
 */

package asar

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type entry struct {
	Size   int    `json:"size"`
	Offset string `json:"offset"`
}

const packageJSON = "{\n\t\"name\": \"discord\",\n\t\"main\": \"index.js\"\n}"

// Write creates the minimal ASAR bootstrap used to load Vencord's patcher.
func Write(outFile, patcherPath string) error {
	pathJSON, err := json.Marshal(patcherPath)
	if err != nil {
		return fmt.Errorf("encode patcher path: %w", err)
	}

	indexJS := "require(" + string(pathJSON) + ")"
	files := map[string]entry{
		"index.js": {
			Size:   len([]byte(indexJS)),
			Offset: "0",
		},
		"package.json": {
			Size:   len([]byte(packageJSON)),
			Offset: strconv.Itoa(len([]byte(indexJS))),
		},
	}
	header := map[string]any{"files": files}
	headerBytes, err := json.Marshal(header)
	if err != nil {
		return fmt.Errorf("encode asar header: %w", err)
	}

	const dataSize uint32 = 4
	headerStringSize := uint32(len(headerBytes))
	alignedSize := (headerStringSize + dataSize - 1) & ^(dataSize - 1)
	headerSize := alignedSize + 8
	headerObjectSize := alignedSize + dataSize

	f, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("create %s: %w", outFile, err)
	}
	defer f.Close()

	for _, n := range []uint32{dataSize, headerSize, headerObjectSize, headerStringSize} {
		if err := binary.Write(f, binary.LittleEndian, int32(n)); err != nil {
			return fmt.Errorf("write asar header: %w", err)
		}
	}

	if _, err := f.Write(headerBytes); err != nil {
		return fmt.Errorf("write asar json: %w", err)
	}
	if padding := int(alignedSize - headerStringSize); padding > 0 {
		if _, err := f.Write(make([]byte, padding)); err != nil {
			return fmt.Errorf("write asar padding: %w", err)
		}
	}
	if _, err := f.WriteString(indexJS + packageJSON); err != nil {
		return fmt.Errorf("write asar payload: %w", err)
	}
	return f.Sync()
}
