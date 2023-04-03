<div align="center">
  <h1 align="center">Ikaros</h1>
  
[![Translation Status][weblate-image]][weblate-url]
[![build result][build-image]][build-url]

[weblate-url]: https://hosted.weblate.org/engage/vanilla-os/
[weblate-image]: https://hosted.weblate.org/widgets/vanilla-os/-/apx/svg-badge.svg
[weblate-status-image]: https://hosted.weblate.org/widgets/vanilla-os/-/ikaros/multi-auto.svg
[build-image]:https://build.opensuse.org/projects/home:fabricators:orchid/packages/ikaros/badge.svg?type=default
[build-url]: https://build.opensuse.org/package/show/home:fabricators:orchid/ikaros
  
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

[![Translation Status][weblate-status-image]][weblate-url]

### Generating man pages for translations

Once the translation is complete in Weblate and the changes committed, clone the repository using `git` and perform `go build`, create a directory using the `mkdir man/<language_code>` command, and execute this command `LANG=<language_code> ./ikaros man > man/<language_code>/ikaros.1`. Open a PR for the generated manpage here.
