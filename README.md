# Go packages on custom domain

This is a recipe for how you put your Github hosted packages on a custom domain.  We are going to use this repository as an example.  It contains a package called `customdomain` and the `go.mod` file looks like this:

```text
module borud.no/customdomain

go 1.17
```

In order to redirect the Go tools to this github repository we need to create an HTML file that looks like this:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>customdomain</title>
    <meta name="go-import" content="borud.no/customdomain git https://github.com/borud/customdomain">
    <meta name="go-source" content="borud.no/customdomain https://github.com/borud/customdomain https://github.com/borud/customdomain/tree/master{/dir} https://github.com/borud/customdomain/blob/master{/dir}/{file}#L{line}">
  </head>
  <body>
    <h1>customdomain</h1>
    Description here.
  </body>
</html>
```

Substitute for whatever your domain name, package name and github repository should be and you now have a custom domain name in your package name.

## Using this package

```go
package main

import "borud.no/customdomain"

func main() {
  dummy := customdomain.Sometype{}
  dummy.Hello()
}
```

## Private repositories

If you are using private repositories you want to specify the `git+ssh` URL rather than the `https` URL for the actual repository.
