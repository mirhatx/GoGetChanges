# GoGetChanges
Check if particular webpage's HTML content has changed

## Installation:

```
go install github.com/mirhatx/gogetchanges@latest
```

## Usage:

```
gogetchanges example.com 60
```

This will check if the HTML content of a page has changed every 60 seconds

If it detects a change it prints it and saves it to a file.
