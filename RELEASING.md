# Processo di release

Le release ufficiali dell'installer vengono prodotte esclusivamente da GitHub Actions.

## Controlli prima della pubblicazione

Prima di creare una release devono passare:

```text
gofmt
go vet ./...
go test ./...