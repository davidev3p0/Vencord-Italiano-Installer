# Architettura e modello di fiducia

Vencord Italiano Installer è intenzionalmente piccolo e separato dal fork principale.

## Responsabilità

Questo repository gestisce soltanto:

- rilevamento delle installazioni Discord;
- download dei runtime pubblicati da Vencord Italiano;
- verifica dei digest SHA-256;
- patch e ripristino di `app.asar`;
- backup e rollback;
- installazione, riparazione e disinstallazione.

Non contiene traduzioni, plugin o logica dell'updater interno di Vencord.

## Flusso

```text
GitHub Release Vencord Italiano
        ↓
metadata + digest SHA-256
        ↓
Vencord Italiano Installer
        ↓
%APPDATA%\Vencord\dist
        ↓
patch di Discord
        ↓
Vencord Italiano avviato
        ↓
aggiornamenti runtime gestiti dal fork principale