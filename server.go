package main

//import "martini"
import (
  "fmt"
  "net/http"
  "io/ioutil"
  "github.com/go-martini/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())
  m.Get("/", func(r render.Render) {
    r.HTML(200, "node", "Drupal")
  })

  // HTTP Client to connect to our Drupal site
  client := &http.Client{
  }
  response, err  := client.Get("http://drupal8.dev")
  if err != nil {
    fmt.Printf("%s", err)
  } else {
    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
      fmt.Printf("%s", err)
    }
    fmt.Printf("%s\n", string(contents))
  }
  m.Run()
}
