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

First we should build our contianers using
```
./ctl.sh build
```

Next we can run our application using
```
./ctl.sh up
```

Interact with the API via <a href="http://localhost:3000/v1/api/flag">http://localhost:3000/v1/api/flag</a> or the database ui via <a href="http://localhost:8081">http://localhost:8081</a>