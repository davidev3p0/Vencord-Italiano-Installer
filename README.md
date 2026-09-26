# Vencord Italiano Installer

[![Build](https://github.com/davidev3p0/Vencord-Italiano-Installer/actions/workflows/build.yml/badge.svg)](https://github.com/davidev3p0/Vencord-Italiano-Installer/actions/workflows/build.yml)
[![Release](https://github.com/davidev3p0/Vencord-Italiano-Installer/actions/workflows/release.yml/badge.svg)](https://github.com/davidev3p0/Vencord-Italiano-Installer/actions/workflows/release.yml)
[![Latest release](https://img.shields.io/github/v/release/davidev3p0/Vencord-Italiano-Installer)](https://github.com/davidev3p0/Vencord-Italiano-Installer/releases/latest)
[![License](https://img.shields.io/github/license/davidev3p0/Vencord-Italiano-Installer)](LICENSE)

Installer Windows open source per [Vencord Italiano](https://github.com/davidev3p0/Vencord-Italiano).

**English:** open-source Windows installer for Vencord Italiano, with transparent GitHub Actions builds, SHA-256 verification and public source code.

## Download

Scarica sempre l'ultima release ufficiale:

**[Scarica Vencord-Italiano-Setup.exe](https://github.com/davidev3p0/Vencord-Italiano-Installer/releases/latest/download/Vencord-Italiano-Setup.exe)**

Non sono necessari Git, Node.js, pnpm, PowerShell o archivi ZIP.

## Cosa fa

- rileva Discord Stable, Canary e PTB;
- installa o aggiorna Vencord Italiano;
- ripara una installazione esistente;
- disinstalla Vencord ripristinando l'originale `app.asar`;
- scarica i runtime esclusivamente dalle release pubbliche di `davidev3p0/Vencord-Italiano`;
- verifica i digest SHA-256 pubblicati da GitHub;
- usa backup e rollback durante le operazioni sensibili;
- non elimina le impostazioni personali di Vencord.

| Client Discord | Supporto |
| --- | --- |
| Stable | ✅ |
| Canary | ✅ |
| PTB | ✅ |

## Installazione

1. Chiudi completamente Discord.
2. Scarica `Vencord-Italiano-Setup.exe` dalla pagina Releases.
3. Avvia l'installer.
4. Seleziona il client Discord rilevato.
5. Scegli **Installa / Aggiorna Vencord Italiano**.
6. Al termine riapri Discord.

Vencord Italiano dispone poi del proprio sistema di aggiornamento automatico collegato alle release del progetto principale.

## Verifica del download

Ogni release pubblica anche `Vencord-Italiano-Setup.exe.sha256`.

Su PowerShell:

```powershell
Get-FileHash .\Vencord-Italiano-Setup.exe -Algorithm SHA256
Get-Content .\Vencord-Italiano-Setup.exe.sha256
```

I due SHA-256 devono coincidere.

Le build ufficiali vengono inoltre prodotte da GitHub Actions e accompagnate da GitHub Artifact Attestation / build provenance.

## Sicurezza e trasparenza

Il progetto usa solo la libreria standard Go e non utilizza:

- packer o UPX;
- cifratura dell'eseguibile;
- offuscamento volto a nascondere il comportamento;
- script PowerShell eseguiti sul PC dell'utente;
- tecniche di bypass antivirus o SmartScreen;
- telemetria o analytics specifici del progetto.

### Stato della firma Authenticode

Le build pubbliche possono essere **non firmate Authenticode**. Per questo Windows SmartScreen può mostrare **Editore sconosciuto** anche quando hash e provenienza GitHub risultano corretti.

Il progetto ha una policy pubblica di code signing ed è predisposto per integrare un provider di firma attendibile quando ne soddisferà i requisiti. Non viene usato alcun certificato self-signed per fingere una firma pubblicamente attendibile.

Documentazione:

- [Code signing policy](CODE_SIGNING_POLICY.md)
- [Privacy policy](PRIVACY.md)
- [Security policy](SECURITY.md)
- [Roadmap](ROADMAP.md)
- [Changelog](CHANGELOG.md)

## Contribuire

Bug report, test su Stable/Canary/PTB e contributi sono benvenuti.

Leggi [CONTRIBUTING.md](CONTRIBUTING.md) prima di aprire una pull request. Per assistenza consulta [SUPPORT.md](SUPPORT.md).

## Build locale

Richiede Go.

```powershell
go test ./...
go vet ./...
go build -trimpath -o Vencord-Italiano-Setup.exe .
```

## Progetto correlato

- [Vencord Italiano](https://github.com/davidev3p0/Vencord-Italiano)

## Licenza e marchi

GPL-3.0. Questo progetto non è affiliato, sponsorizzato o approvato da Discord Inc. Vencord è un progetto separato; i relativi nomi e marchi appartengono ai rispettivi titolari.
