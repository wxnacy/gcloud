package main

import (

    "github.com/kataras/iris"
    // "github.com/kataras/iris/middleware/logger"
)

var app *iris.Application

func main() {
    app = iris.Default()
    app.Logger().SetLevel("debug")
    // app.Use(logger.New())
    app.RegisterView(iris.HTML("views", ".html"))
    app.StaticWeb("static", "./static")


    app.Get("/", func(c iris.Context) {
        c.ViewData("name", "wxnacy")
        c.View("index.html")
    })

    app.Get("/hello", func(c iris.Context) {
        c.WriteString("Hello World")
    })

    // listen and serve on http://0.0.0.0:8080.
    app.Run(iris.Addr(":8080"))

}
