# RadioTray

![Logo](./lib/icons/radio.red.128.png)

Simple radio player in your tray bar for macOS

## Installation

```bash
# clone repo
brew install go
make darwin-app
# move build/RadioTray.app to /Applications
```

## Configuration

Radios configurations are stored in `$HOME/.radiotray.yaml`.

After each change, restart the app.

Example configuration :

```yaml
radios:
- name: Maxxima
  format: mp3
  source: https://maxxima.mine.nu/maxxima.mp3
- name: Swiss Pop
  format: mp3
  source: http://stream.srg-ssr.ch/m/rsp/mp3_128
- name: Swiss Jazz
  format: mp3
  source: http://stream.srg-ssr.ch/m/rsj/mp3_128
```

## Credits

- https://github.com/angelodlfrtr/radiotray