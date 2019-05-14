# Getting Started

GO Workspace contains multiple version control repositories:

* `src` contains source files
* `bin` contains executables which are built by go tools

```
bin/
    hello                          # command executable
src/
    github.com/:user/:repo/
      .git/
      hello/
          hello.go                 # command source
```

Packages are resolved by the `$GOPATH` environment variable. Therefore when you see an import statement in a go package like the following it is not importing "by url":

```go
import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/walkerrandolphsmith/go-playground/features/flag"
)
```