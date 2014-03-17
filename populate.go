package main

import (
  "github.com/garyburd/redigo/redis"
  "time"
  "io/ioutil"
)

func populate(c redis.Conn) {
  //set
  c.Do("DEL", "bloggo:articles")
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
    Body: string(body),
    AuthorName: "Donald Plummer",
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
    Body: string(body),
    AuthorName: "Donald Plummer",
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
    Body: string(body),
    AuthorName: "Donald Plummer",
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
    Body: string(body),
    AuthorName: "Donald Plummer",
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
    Body: string(body),
    AuthorName: "Donald Plummer",
  }

  for _, article := range articles {
    c.Do("RPUSH", "bloggo:articles", article.ToJson())
  }

  blog := Blog{
    Title: "Donald Plummer's Blog",
    Subtitle: "I'm a software developer in Seattle, WA.",
    BaseUrl: "http://www.donaldplummer.com",
    AuthorName: "Donald Plummer",
    About: `I enjoy playing board games, drinking stouts, playing video
            games, and learning. At least once a week I'll have something
            cool I want to share with my coworkers, so I figured I should
            write it up for the world to see. Thus this website.`,
  }
  blog.Save(c)
}
