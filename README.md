
# NISOT :shirt:

[![Build Status](https://travis-ci.com/guvense/Nisot.svg?branch=master)](https://travis-ci.com/guvense/Nisot)
![GitHub](https://img.shields.io/github/license/guvense/Nisot)

Nisot is a go library 

## Parser Usage

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
p.Replace("foo.txt")
```

## Repeat Usage

Generate Value

```
s := repeater.Value("val")
```
Repeat n times

```
repeated := s.Repeat(4)
```

Calculate Repeated Word

```
count := repeater.Value("foobarfoobarfoobar")
```



