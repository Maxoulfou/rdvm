<a href="#">
    <img src="https://cdn.brochier.xyz/github.com/rdvm/rdvm_logo.png" alt="rdvm logo" title="rdvm" align="right" height="60" />
</a>

# Remote Dyjix VPS Management CLI Tool

> **Disclaimer:** this tool is absolutely not affiliated to Dyjix for the moment, the code is open source, and I made it on my own for me, and I share it with you today

## Basic

To start, you will have to put the executable in an isolated folder, then create a file named "config.json".
Fill in the information as follows:

```json
{
  "key": "key-from-solusvm",
  "hash": "hash-from-solusvm",
  "master-url": "https://vps.dyjix.eu/api/client/command.php"
}
```

## Detail

```bash
rdvm is a simple CLI to manage your Dyjix VPS
easily without the need to log into the management
panel (vps.dyjix.eu)

Usage:
  rdvm [command]

Available Commands:
  boot        Turn your VPS in "unrompiche" mode
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  info        Get actual status, HN and IP about your VPS
  reboot      Reboot your VPS
  shutdown    Turn your VPS in "rompiche" mode
  status      Get actual status of your VPS

Flags:
  -h, --help     help for rdvm
  -t, --toggle   Help message for toggle

Use "rdvm [command] --help" for more information about a command.

```

## Compilation

`go build main.go -o bin/`

## Other

- Deb package comming soon
- exe installer comming soon
- executable comming soon
