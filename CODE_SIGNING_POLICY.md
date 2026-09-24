# Code signing policy

Free code signing provided by SignPath.io, certificate by SignPath Foundation.

## Build system

Official Windows release binaries are built only by GitHub Actions from this public repository using GitHub-hosted runners.

## Roles

- Committer: davidev3p0
- Reviewer: davidev3p0
- Approver: davidev3p0

The project currently has one maintainer. If more maintainers are added, these roles will be separated where practical before they participate in release signing.

## Signing rules

- Only artifacts produced by the repository's release workflow may be submitted for signing.
- Release signing requires explicit approval.
- Source code, workflow definitions and release history remain public.
- No packer, UPX, executable encryption or code-obfuscation technique intended to conceal behavior is permitted.
- No antivirus or SmartScreen bypass technique is permitted.
- The project never receives or stores the private code-signing key.
- SHA-256 hashes and build provenance are published for releases.

## Scope

The signing certificate is used only for official Vencord Italiano Installer Windows releases from this repository.
