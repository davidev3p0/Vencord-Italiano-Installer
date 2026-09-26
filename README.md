# Vencord Italiano Installer

[![Build](https://github.com/davidev3p0/Vencord-Italiano-Installer/actions/workflows/build.yml/badge.svg)](https://github.com/davidev3p0/Vencord-Italiano-Installer/actions/workflows/build.yml)
[![Release](https://github.com/davidev3p0/Vencord-Italiano-Installer/actions/workflows/release.yml/badge.svg)](https://github.com/davidev3p0/Vencord-Italiano-Installer/actions/workflows/release.yml)
[![Latest release](https://img.shields.io/github/v/release/davidev3p0/Vencord-Italiano-Installer)](https://github.com/davidev3p0/Vencord-Italiano-Installer/releases/latest)
[![Downloads](https://img.shields.io/github/downloads/davidev3p0/Vencord-Italiano-Installer/total)](https://github.com/davidev3p0/Vencord-Italiano-Installer/releases)
[![Stars](https://img.shields.io/github/stars/davidev3p0/Vencord-Italiano-Installer)](https://github.com/davidev3p0/Vencord-Italiano-Installer/stargazers)
[![License](https://img.shields.io/github/license/davidev3p0/Vencord-Italiano-Installer)](LICENSE)

Installer Windows open source per **[Vencord Italiano](https://github.com/davidev3p0/Vencord-Italiano)**.

> Per gli utenti finali questo repository è solo il punto di download dell'installer. Traduzioni, plugin, RealVoiceTTS e aggiornamenti runtime vengono gestiti dal repository principale Vencord Italiano.

## Download

### [⬇️ Scarica Vencord-Italiano-Setup.exe](https://github.com/davidev3p0/Vencord-Italiano-Installer/releases/latest/download/Vencord-Italiano-Setup.exe)

Non sono necessari Git, Node.js, pnpm, PowerShell o file ZIP.

## Come funziona

L'installer ha un compito volutamente limitato:

1. rileva Discord Stable, Canary e PTB;
2. scarica i runtime della release più recente di Vencord Italiano;
3. verifica i digest SHA-256 pubblicati da GitHub;
4. crea un backup prima di modificare `app.asar`;
5. installa, ripara o rimuove la patch;
6. lascia a Vencord Italiano la gestione degli aggiornamenti successivi.

Dopo la prima installazione **non è necessario scaricare un nuovo installer per ogni aggiornamento di Vencord Italiano**. Gli aggiornamenti runtime vengono gestiti direttamente dal progetto principale.

| Client Discord | Supporto |
| --- | --- |
| Stable | ✅ |
| Canary | ✅ |
| PTB | ✅ |

## Installazione

1. Chiudi completamente Discord.
2. Scarica `Vencord-Italiano-Setup.exe` dalla release più recente.
3. Avvia l'installer.
4. Seleziona il client Discord rilevato.
5. Scegli **Installa / Aggiorna Vencord Italiano**.
6. Riapri Discord.

## Verifica del download

Ogni release pubblica anche `Vencord-Italiano-Setup.exe.sha256`.

```powershell
Get-FileHash .\Vencord-Italiano-Setup.exe -Algorithm SHA256
Get-Content .\Vencord-Italiano-Setup.exe.sha256
```

I due SHA-256 devono coincidere. Le build ufficiali sono inoltre prodotte da GitHub Actions con GitHub Artifact Attestation / build provenance.

## Sicurezza e trasparenza

Il progetto usa la libreria standard Go e non utilizza:

- packer o UPX;
- cifratura dell'eseguibile;
- offuscamento volto a nascondere il comportamento;
- script PowerShell eseguiti sul PC dell'utente;
- tecniche di bypass antivirus o SmartScreen;
- telemetria o analytics specifici del progetto.

### Firma Authenticode

Le build pubbliche possono essere **non firmate Authenticode**. In questo caso SmartScreen può mostrare **Editore sconosciuto** anche quando hash e provenienza GitHub risultano corretti.

Il progetto è predisposto per integrare in futuro un provider di code signing pubblicamente attendibile quando ne soddisferà i requisiti. Non viene usato un certificato self-signed per simulare una firma attendibile.

Approfondimenti:

- [Architettura e modello di fiducia](ARCHITECTURE.md)
- [Code signing policy](CODE_SIGNING_POLICY.md)
- [Privacy policy](PRIVACY.md)
- [Security policy](SECURITY.md)
- [Processo di release](RELEASING.md)
- [Roadmap](ROADMAP.md)
- [Changelog](CHANGELOG.md)

## Supporto e community

- [Discussions](https://github.com/davidev3p0/Vencord-Italiano-Installer/discussions) per domande e feedback.
- [Issues](https://github.com/davidev3p0/Vencord-Italiano-Installer/issues) per bug riproducibili.
- [Vencord Italiano](https://github.com/davidev3p0/Vencord-Italiano) per traduzione, plugin, RealVoiceTTS e updater.

Se il progetto ti è utile, una ⭐ su GitHub aiuta altre persone a trovarlo.

## Contribuire

Leggi [CONTRIBUTING.md](CONTRIBUTING.md) prima di aprire una pull request. Per assistenza consulta [SUPPORT.md](SUPPORT.md).

## Build locale

Richiede Go.

```powershell
go test ./...
go vet ./...
go build -trimpath -o Vencord-Italiano-Setup.exe .
```

## Licenza e marchi

GPL-3.0. Questo progetto non è affiliato, sponsorizzato o approvato da Discord Inc. Vencord è un progetto separato; i relativi nomi e marchi appartengono ai rispettivi titolari.
