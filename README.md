# grimoire

Utility for naming your softwares or services.  
Grimoire does just outputting a sorted unique word list, so you should filter the output with
[percol](https://github.com/mooz/percol), [peco](https://github.com/peco/peco), etc.

![](https://gist.github.com/k0kubun/626746935c742894786e/raw/d18056bead3a799e0987ded8502548adbe29fde5/ttt.gif)

## Installation

```bash
$ go get github.com/k0kubun/grimoire
```

## Usage

```bash
# Output all words
# It takes TOO MUCH time at first execution
$ grimoire

# Show dictionary list
$ grimoire -h

# Output specific dictionary
$ grimoire greek
```

## Supported Dictionaries

- English
- Greek Mythology
- Norse Mythology
- Person name in British, French, Italy, Spain, Greek, Finland and Russia

## TODO
- User-editable dictionary for memoing words you want to use
- Show dictionary name in output
- Add another dictionary (Italian, French ...)
- Cache update option `-u`
- Ignore case output sort
- Alphabetize non-alphabet chars

## License

MIT License
