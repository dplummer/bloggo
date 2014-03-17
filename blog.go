package main

import (
  "github.com/garyburd/redigo/redis"
  "time"
  "encoding/json"
)

type Blog struct {
  Title         string    `json:"title"`
  Subtitle      string    `json:"subtitle"`
  BaseUrl       string    `json:"baseUrl"`
  LastUpdatedAt time.Time `json:"lastUpdatedAt"`
  AuthorName    string    `json:"authorName"`
  About         string    `json:"about"`
}

func (blog *Blog) FromJson(raw []byte) {
  json.Unmarshal(raw, &blog)
}

func getBlog(c redis.Conn) (blog Blog) {
  raw, _ := redis.String(c.Do("GET", "bloggo:blog"))
  blog.FromJson([]byte(raw))
  return
}

func (blog *Blog) ToJson() []byte {
  encoded, err := json.Marshal(blog)
  if err != nil {
    panic(err)
  }
  return encoded
}

func (blog *Blog) Save(c redis.Conn) {
  c.Do("SET", "bloggo:blog", blog.ToJson())
}
