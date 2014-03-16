package main

import (
  "github.com/codegangsta/martini"
  "github.com/martini-contrib/render"
  "github.com/garyburd/redigo/redis"
)

type viewContext struct {
  Articles []Article
}

func getArticles(c redis.Conn) (content []Article, err error) {
  encoded, err := redis.Strings(c.Do("LRANGE", "bloggo:articles:index", "0", "-1"))
  content = make([]Article, len(encoded))
  for i, raw := range encoded {
    content[i].FromJson([]byte(raw))
  }
  return
}

func main() {
  m := martini.Classic()
  c, err := redis.Dial("tcp", ":6379")
  if err != nil {
    panic(err)
  }
  defer c.Close()

  populate(c)

  m.Use(render.Renderer(render.Options{
    Layout: "layout",
  }))

  m.Get("/", func(r render.Render) {
    var view viewContext
    view.Articles, err = getArticles(c)

    if err != nil {
      r.HTML(200, "error", err)
    } else {
      r.HTML(200, "home", view)
    }
  })

  m.Run()
}
