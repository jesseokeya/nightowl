# method [![wercker status](https://app.wercker.com/status/e396edd92b8f734e277064f2ffe76adc "wercker status")](https://app.wercker.com/project/bykey/e396edd92b8f734e277064f2ffe76adc)
Martini middleware/handler for handling http method overrides.
This checks for the X-HTTP-Method-Override header and uses it
if the original request method is POST.
GET/HEAD methods shouldn't be overriden, hence they can't be overriden.

This is useful for REST APIs and services making use of many HTTP verbs, and when http clients don't support all of them.

[API Reference](http://godoc.org/github.com/martini-contrib/method)

## Usage

~~~ go
import (
  "github.com/codegangsta/martini"
  "github.com/martini-contrib/method"
)

func main() {
  m := martini.Classic()
  m.Use(method.Override())
  m.Run()
}

~~~

## Authors
* [Vincent Petithory](http://github.com/vincent-petithory)
