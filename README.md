
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
val := repeater.Value("foobarfoobarfoobar")
z := val.Count()
```

## Http Util Usage

Sum of url response bodies

Cretating a Page struct
```
page := Page{}
```
Defining url string array

```go
urls := []string{"http://golang.org.tr/", "http://www.example.com"}

```

Setting urls to page 

```go
page.SetUrls(urls)
```

Total size of url bodies

```go
totalSize := page.getTotalSize()
```

