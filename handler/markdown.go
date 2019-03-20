package markdown

import (
    "os"
    "fmt"
    "net/http"

    "github.com/labstack/echo"
    "gopkg.in/russross/blackfriday.v2"
)

type MarkdownData struct {
    MD string `json:"md" form:"md" query:"md"`
}

func Default(c echo.Context) (err error) {
    md := new(MarkdownData)
    if err = c.Bind(md); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }

    extFlags := blackfriday.CommonExtensions
    extFlags |= blackfriday.FencedCode
    extFlags |= blackfriday.Footnotes
    extFlags |= blackfriday.HeadingIDs
    extFlags |= blackfriday.Titleblock

    html := blackfriday.Run(([]byte)(md.MD), blackfriday.WithExtensions(extFlags))

    return c.HTML(http.StatusOK, fmt.Sprintf("%s", html))
}
