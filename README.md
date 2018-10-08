# sp - Symfony console commands Prompt

## Description

Input assist tool for Symfony console commands.

## Installation

```bash
$ go get github.com/kei711/sp
```

## Usage

### Execute command

```bash
$ sp
```

**NOTICE: If you have not registered command, you need to register first.**

### Register command

```bash
$ sp add
```

### Remove registered command

```bash
$ sp remove
```

### Change cache directory

default: `~/.sp`

```bash
$ sp setCacheDir -c ~/.symfony-command-cache
```

### Show version

```bash
$ sp version
```
