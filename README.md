# GoPlay

## A wrapper around `playerctl`

I had problems with waybar scripts not working with media displays, so I made a wrapper around `playerctl` to make it easier to manage.

## Installation

```bash
git clone https://github.com/Courtcircuits/goplay.git
cd goplay/
just build
mv goplay ~/.local/bin/ # or anywhere in your PATH
```

## Usage

### Waybar

```json
{
  [...]
  "module-center": "custom/media",
  [...]
  "custom/media": {
    "exec": "goplay",
    "on-click": "playerctl play-pause"
  },
  [...]
}
```

### Shell

```bash
goplay
```
