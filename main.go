/*
 * Vencord Italiano Installer
 * Copyright (c) 2026 contributors
 * SPDX-License-Identifier: GPL-3.0-only
 */

package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/davidev3p0/Vencord-Italiano-Installer/internal/discord"
	"github.com/davidev3p0/Vencord-Italiano-Installer/internal/release"
)

var (
	version = "dev"
	commit  = "unknown"
)

type operation string

const (
	opInstall   operation = "install"
	opRepair    operation = "repair"
	opUninstall operation = "uninstall"
)

func main() {
	installFlag := flag.Bool("install", false, "Installa o aggiorna Vencord Italiano")
	repairFlag := flag.Bool("repair", false, "Ripara Vencord Italiano")
	uninstallFlag := flag.Bool("uninstall", false, "Disinstalla Vencord da Discord")
	branchFlag := flag.String("branch", "all", "Client Discord: stable, canary, ptb o all")
	showVersion := flag.Bool("version", false, "Mostra versione installer")
	flag.Parse()

	if *showVersion {
		fmt.Printf("Vencord Italiano Installer %s (%s)\n", version, commit)
		return
	}

	directCount := boolInt(*installFlag) + boolInt(*repairFlag) + boolInt(*uninstallFlag)
	interactive := directCount == 0

	fmt.Println("============================================================")
	fmt.Printf(" Vencord Italiano Installer  %s (%s)\n", version, commit)
	fmt.Println("============================================================")
	fmt.Println("Installer nativo Windows · nessun PowerShell · nessun packer")
	fmt.Println()

	code := run(interactive, *installFlag, *repairFlag, *uninstallFlag, *branchFlag)
	if interactive {
		fmt.Println()
		fmt.Print("Premi INVIO per chiudere...")
		_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	}
	os.Exit(code)
}

func run(interactive, installFlag, repairFlag, uninstallFlag bool, branch string) int {
	directCount := boolInt(installFlag) + boolInt(repairFlag) + boolInt(uninstallFlag)
	if directCount > 1 {
		return fail(errors.New("usa una sola operazione tra --install, --repair e --uninstall"))
	}

	clients, err := discord.Detect()
	if err != nil {
		return fail(err)
	}
	if len(clients) == 0 {
		return fail(errors.New("nessuna installazione Discord Stable, Canary o PTB trovata"))
	}

	fmt.Println("Installazioni rilevate:")
	for i := range clients {
		state := "Vencord non rilevato"
		if clients[i].Patched {
			state = "Vencord presente"
		}
		fmt.Printf("  %d) %-15s - %s\n", i+1, clients[i].Name, state)
	}
	fmt.Println()

	var selected []int
	var op operation

	if interactive {
		selected, err = promptClients(clients)
		if err != nil {
			return fail(err)
		}
		op, err = promptOperation()
		if err != nil {
			return fail(err)
		}
	} else {
		selected, err = selectByBranch(clients, branch)
		if err != nil {
			return fail(err)
		}
		switch {
		case installFlag:
			op = opInstall
		case repairFlag:
			op = opRepair
		case uninstallFlag:
			op = opUninstall
		}
	}

	fmt.Println()
	fmt.Println("Chiudi completamente Discord prima di continuare.")
	fmt.Println()

	if op == opUninstall {
		for _, idx := range selected {
			fmt.Printf("Ripristino %s...\n", clients[idx].Name)
			if err := clients[idx].Unpatch(); err != nil {
				return fail(fmt.Errorf("%s: %w", clients[idx].Name, err))
			}
			fmt.Printf("  OK: %s ripristinato\n", clients[idx].Name)
		}
		fmt.Println()
		fmt.Println("Disinstallazione completata. Le impostazioni personali Vencord non sono state eliminate.")
		return 0
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return fail(fmt.Errorf("cartella configurazione Windows: %w", err))
	}
	distDir := filepath.Join(configDir, "Vencord", "dist")

	fmt.Println("Download build Vencord Italiano da GitHub...")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	rel, err := release.InstallLatest(ctx, distDir)
	if err != nil {
		return fail(err)
	}
	fmt.Printf("  Release verificata: %s\n", rel.TagName)

	patcherPath := filepath.Join(distDir, "patcher.js")
	for _, idx := range selected {
		fmt.Printf("Patch %s...\n", clients[idx].Name)
		if err := clients[idx].Patch(patcherPath); err != nil {
			return fail(fmt.Errorf("%s: %w", clients[idx].Name, err))
		}
		fmt.Printf("  OK: %s\n", clients[idx].Name)
	}

	fmt.Println()
	if op == opRepair {
		fmt.Println("Riparazione completata con successo.")
	} else {
		fmt.Println("Installazione / aggiornamento completato con successo.")
	}
	fmt.Println("Apri Discord: Vencord Italiano userà il proprio aggiornamento automatico.")
	return 0
}

func promptClients(clients []discord.Client) ([]int, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		if len(clients) > 1 {
			fmt.Printf("Seleziona Discord [1-%d] oppure A per tutti: ", len(clients))
		} else {
			fmt.Print("Seleziona Discord [1]: ")
		}
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.TrimSpace(strings.ToLower(line))
		if len(clients) > 1 && (line == "a" || line == "all" || line == "tutti") {
			result := make([]int, len(clients))
			for i := range clients {
				result[i] = i
			}
			return result, nil
		}
		n, err := strconv.Atoi(line)
		if err == nil && n >= 1 && n <= len(clients) {
			return []int{n - 1}, nil
		}
		fmt.Println("Scelta non valida.")
	}
}

func promptOperation() (operation, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		fmt.Println("Operazione:")
		fmt.Println("  1) Installa / Aggiorna Vencord Italiano")
		fmt.Println("  2) Ripara / Reinstalla")
		fmt.Println("  3) Disinstalla Vencord")
		fmt.Println("  4) Esci")
		fmt.Print("Scelta: ")

		line, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		switch strings.TrimSpace(line) {
		case "1":
			return opInstall, nil
		case "2":
			return opRepair, nil
		case "3":
			return opUninstall, nil
		case "4":
			return "", errors.New("operazione annullata")
		default:
			fmt.Println("Scelta non valida.")
		}
	}
}

func selectByBranch(clients []discord.Client, branch string) ([]int, error) {
	branch = strings.ToLower(strings.TrimSpace(branch))
	if branch == "all" || branch == "tutti" {
		result := make([]int, len(clients))
		for i := range clients {
			result[i] = i
		}
		return result, nil
	}
	if branch != "stable" && branch != "canary" && branch != "ptb" {
		return nil, fmt.Errorf("branch non valida: %s", branch)
	}
	for i := range clients {
		if clients[i].Branch == branch {
			return []int{i}, nil
		}
	}
	return nil, fmt.Errorf("Discord %s non trovato", branch)
}

func fail(err error) int {
	fmt.Println("ERRORE:", err)
	return 1
}

func boolInt(v bool) int {
	if v {
		return 1
	}
	return 0
}
