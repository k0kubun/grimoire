# grimoire

Utility for naming your softwares or services.  
Grimoire does just outputting a word list, so you should filter the output with
[percol](https://github.com/mooz/percol), [peco](https://github.com/peco/peco), etc.

## Installation

```bash
$ go get github.com/k0kubun/grimoire
```

## Usage

```bash
# Output all words
# It takes much time at first execution
$ grimoire

# Show dictionary list
$ grimoire -h

# Output specific dictionary
$ grimoire greek
```

## TODO
- User-editable dictionary for memoing words you want to use
- Show dictionary name in output
- Add another dictionary

## License

MIT License