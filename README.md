<div align="center">
<img src="assets/ikaros.png?raw=true#gh-dark-mode-only" height="40">
<img src="assets/ikaros-mono.png?raw=true#gh-light-mode-only" height="40">

---
[![Translation Status][weblate-image]][weblate-url]

[weblate-url]: https://hosted.weblate.org/engage/vanilla-os/
[weblate-image]: https://hosted.weblate.org/widgets/vanilla-os/-/apx/svg-badge.svg
  
Ikaros is a drivers backend for Vanilla OS.

This project is meant to be used as a ubuntu-drivers-common replacement.
It's still in development and not ready for production use. 
It's meant to automatically discover and install drivers for your devices.
</div>

## Help

```bash
Usage:
  ikaros [command]

Available Commands:
  auto-install Auto install correct drivers
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  install      Install a driver
  list-devices List devices
  list-drivers List drivers

Flags:
  -h, --help      help for ikaros
  -v, --version   version for ikaros

Use "ikaros [command] --help" for more information about a command.
```

## Translations

Contribute translations for the manpage and help page in [Weblate](https://hosted.weblate.org/projects/vanilla-os/ikaros).

### Generating man pages for translations

Once the translation is complete in Weblate and the changes committed, clone the repository using `git` and perform `go build`, create a directory using the `mkdir man/<language_code>` command, and execute this command `LANG=<language_code> ./ikaros man > man/<language_code>/ikaros.1`. Open a PR for the generated manpage here.
