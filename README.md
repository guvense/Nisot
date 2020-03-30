
## NISOT :shirt:

[![Build Status](https://travis-ci.com/guvense/Nisot.svg?branch=master)](https://travis-ci.com/guvense/Nisot)

Nisot is a go library 

Usage

Read file and parse it 

```
p := mapper.Conf(':', "a.txt")
```

Create a string string map

```
myMap := p.ParseToMap()
```

Change key value in file empty string for same file

```
p.Replace("")
```
