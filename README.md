# Vencord Italiano Installer

Installer Windows open source per [Vencord Italiano](https://github.com/davidev3p0/Vencord-Italiano).

## Obiettivo

Un solo `Vencord-Italiano-Setup.exe` per installare o riparare Vencord Italiano senza Git, Node.js, pnpm, PowerShell o file ZIP.

Supporta:

- Discord Stable
- Discord Canary
- Discord PTB
- installazione e aggiornamento
- riparazione
- disinstallazione
- download esclusivamente dalle release pubbliche di Vencord Italiano
- verifica SHA-256 degli asset quando GitHub espone il digest della release
- rollback del patching in caso di errore

Le impostazioni personali di Vencord non vengono eliminate.

## Sicurezza

Il progetto usa solo la libreria standard Go e non contiene packer, UPX, cifratura dell'eseguibile, offuscamento, script PowerShell eseguiti sul PC dell'utente o tecniche di bypass antivirus.

## Code signing policy

Free code signing provided by SignPath.io, certificate by SignPath Foundation.

- [Code signing policy completa](CODE_SIGNING_POLICY.md)
- [Privacy policy](PRIVACY.md)
- [Security policy](SECURITY.md)

## Build locale

```powershell
go test ./...
go build -trimpath -o Vencord-Italiano-Setup.exe .
```

## Licenza

GPL-3.0. Vencord e Discord sono progetti/marchi separati; questo repository non implica affiliazione con Discord Inc.
