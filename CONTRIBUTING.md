# Contributing

Grazie per l'interesse in Vencord Italiano Installer.

## Prima di aprire una issue

Per i bug, verifica prima:

- di usare l'ultima release;
- di aver chiuso completamente Discord;
- quale client stai usando: Stable, Canary o PTB;
- la versione di Windows;
- se il problema è riproducibile.

Non pubblicare token Discord, credenziali, dati personali o contenuti privati nei log.

## Sviluppo locale

Richiede Go.

```powershell
go test ./...
go vet ./...
gofmt -w .
go build -trimpath -o Vencord-Italiano-Setup.exe .
```

Prima di una pull request:

1. esegui `gofmt`;
2. esegui `go vet ./...`;
3. esegui `go test ./...`;
4. mantieni il comportamento dell'installer semplice e prevedibile;
5. evita dipendenze esterne quando non necessarie;
6. non introdurre packer, offuscamento, download nascosti o tecniche di bypass di strumenti di sicurezza.

## Pull request

Descrivi:

- cosa cambia;
- perché serve;
- come è stato testato;
- quali client Discord sono coinvolti;
- eventuali rischi di compatibilità o rollback.

## Sicurezza

Per vulnerabilità o problemi che potrebbero esporre utenti o dati, segui [SECURITY.md](SECURITY.md) invece di pubblicare dettagli sensibili in una issue.
