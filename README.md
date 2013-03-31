# Rota

Package Rota provides simple string matcher which uses it's own mark up language. 

## Example

```go
package main

import "github.com/oguzbilgic/rota"

func main() {
    r := rota.New("/articles/{id}")

    params, err := r.Match("/articles/3245")
    if err != nil {
        print("Rota did not match")
    }

    print("Rota matched and article id is", params["id"])
}
```

## Markup

Here are the most basic examples:

```go
// matches only /
rota.New("/") 

// matches /articles and /articles/
rota.New("/articles") 
```

### Variables

`{name:type}` is used for defining variables with in paths. the first argument is
name of the parameter and the second one is the type of it. Here we define an
integer variable:

```go
// matches /articles/1923 and /articles/1923/ but does not match /articles/
rota.New("/articles/{id:int}")

// also matches /articles/1923 and /articles/1923/
rota.New("/articles/{id}") 

// matches /users/oguzbilgic
// str captures all alphanumeric characters
rota.New("/users/{username:str}")
```

### Or statements

`{option1|option2|etc}` is used for creating an array of accepted constants. 
For example:

```go
// matches /aboutus.html and /about_us.html
rota.New("/{aboutus|about_us}.html")

// value of a or statement can be captured as a named parameter
rota.New("/{dir:css|js|img}/{file:str}")

// or statement can also be used for optional parts
// matches /articles and /articles/
rota.New("/articles{/}")	
```

### Catch-all

Star is a catch-all character

```go
// matches everything
rota.New("/*")

// maches all paths that ends with .json 
rota.New("/*.json")
```
