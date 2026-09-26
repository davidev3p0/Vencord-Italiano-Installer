# Code signing policy

## Current status

Official Windows binaries are built publicly through GitHub Actions. Current releases may be unsigned with Authenticode.

The project intends to use a publicly trusted code-signing provider when it meets that provider's eligibility and identity requirements. No self-signed certificate is represented as a publicly trusted signature.

## Build system

Official Windows release binaries are built only by GitHub Actions from this public repository using GitHub-hosted runners.

## Roles

- Authors / committers: [davidev3p0](https://github.com/davidev3p0)
- Reviewers: [davidev3p0](https://github.com/davidev3p0)
- Approvers: [davidev3p0](https://github.com/davidev3p0)

The project currently has one maintainer. If additional maintainers join, these roles will be separated where practical before they participate in release signing.

## Signing rules

- Only artifacts produced by the repository's release workflow may be submitted for signing.
- Release signing requires explicit approval.
- Source code, workflow definitions and release history remain public.
- No packer, UPX, executable encryption or code-obfuscation technique intended to conceal behavior is permitted.
- No antivirus or SmartScreen bypass technique is permitted.
- The project must never receive or store an exportable private signing key when a managed signing provider is used.
- SHA-256 hashes and build provenance are published for releases.
- Signed artifacts must correspond to a public source revision and a successful CI run.

## Scope

A future publicly trusted signing certificate may be used only for official Vencord Italiano Installer Windows releases from this repository.
