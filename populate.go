package main

import (
  "github.com/garyburd/redigo/redis"
  "time"
  "io/ioutil"
)

func populate(c redis.Conn) {
  //set
  c.Do("DEL", "bloggo:articles:index")
  articles := make([]Article, 5)

  body, err := ioutil.ReadFile("raw/dhcpcd-timeouts-new-archlinux.md")
  if err != nil {
    panic(err)
  }
  articles[0] = Article{
    Title: "DHCPCD timeouts on new archlinux installs",
    PublishedDate: time.Date(2013, time.August, 17, 8, 0, 0, 0, time.UTC),
    SimpleName: "dhcpcd-timeouts-new-archlinux",
    Published: true,
    Body: body,
  }

  body, err = ioutil.ReadFile("raw/zeus-urxvt-broken.md")
  if err != nil {
    panic(err)
  }
  articles[1] = Article{
    Title: "Archlinux Zeus and URXVT broken with latest glibc",
    PublishedDate: time.Date(2013, time.August, 17, 8, 0, 0, 0, time.UTC),
    SimpleName: "zeus-urxvt-broken",
    Published: true,
    Body: body,
  }

  body, err = ioutil.ReadFile("raw/starting-god-with-init.md")
  if err != nil {
    panic(err)
  }
  articles[2] = Article{
    Title: "Starting God with an init script",
    PublishedDate: time.Date(2013, time.August, 17, 8, 0, 0, 0, time.UTC),
    SimpleName: "starting-god-with-init",
    Published: true,
    Body: body,
  }

  body, err = ioutil.ReadFile("raw/using-papertrail.md")
  if err != nil {
    panic(err)
  }
  articles[3] = Article{
    Title: "Using Papertrail",
    PublishedDate: time.Date(2013, time.August, 17, 8, 0, 0, 0, time.UTC),
    SimpleName: "using-papertrail",
    Published: true,
    Body: body,
  }

  body, err = ioutil.ReadFile("raw/segfaults-ree-archlinux.md")
  if err != nil {
    panic(err)
  }
  articles[4] = Article{
    Title: "Segfaults with REE on ArchLinux",
    PublishedDate: time.Date(2012, time.June, 15, 8, 0, 0, 0, time.UTC),
    SimpleName: "segfaults-ree-archlinux",
    Published: true,
    Body: body,
  }

  for _, article := range articles {
    c.Do("RPUSH", "bloggo:articles:index", article.ToJson())
  }
}
