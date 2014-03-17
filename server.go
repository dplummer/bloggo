package main

import (
  "github.com/codegangsta/martini"
  "github.com/martini-contrib/render"
  "github.com/garyburd/redigo/redis"
)

type viewContext struct {
  Articles []Article
  ArticleShow *Article
  PreviousArticle *Article
  NextArticle *Article
}

func getArticles(c redis.Conn) (content []Article, err error) {
  encoded, err := redis.Strings(c.Do("LRANGE", "bloggo:articles:index", "0", "-1"))
  content = make([]Article, len(encoded))
  for i, raw := range encoded {
    content[i].FromJson([]byte(raw))
  }
  return
}

func (v *viewContext) setArticle(name string) {
  for i, article := range v.Articles {
    if article.SimpleName == name {
      v.ArticleShow = &v.Articles[i]
      if i > 0 {
        v.PreviousArticle = &v.Articles[i-1]
      }
      if i < len(v.Articles) - 1 {
        v.NextArticle = &v.Articles[i+1]
      }
    }
  }
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
      r.HTML(500, "error", err)
    } else {
      r.HTML(200, "home", view)
    }
  })

  m.Get("/articles/:name", func(r render.Render, params martini.Params) {
    var view viewContext
    view.Articles, err = getArticles(c)
    view.setArticle(params["name"])

    if err != nil {
      r.HTML(500, "error", err)
    } else {
      r.HTML(200, "articles/show", view)
    }
  })

  m.Run()
}
