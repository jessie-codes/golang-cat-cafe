# REST API

<img src="rest.jpg"
     style="height:500px;" />

For this workshop, we will be using the popular REST framework [Gin](https://github.com/gin-gonic/gin).

We'll begin by examing the code in `tutorial/main.go`. Here's a few things to keep in mind when reading through this code:

- Every executable must be `package main` and have a `main` function.
- We can import 3rd party packages with a github path, such as `github.com/gin-gonic/gin`
- We can import our own subpackages using a fully qualified github path, such as `github.com/jessie-codes/golang-cat-cafe/tutorial/cat`.

If you look at the `main` function, it's doing the following things:

1. Creating an instance of `gin`
2. Adding an endpoint for `GET /cat`
3. Starting the server and running on port `8080`.

We can start our server by using the command `go run main.go` within the `tutorial` folder. If you navigate to [http://localhost:8080/cat](http://localhost:8080/cat) you should see a json response listing all of our cats.

## Task Two - Better XML Support

Try hitting out endpoint using the header `Accepts: application/xml`. You'll notice that all of the XML nodes are capitialized. This is because by default, Go will marshal attributes using their exact name. Let's change these to be lowercase variables.

Go ahead and open up `tutorial/cat/cat.go` and examine our structs.

```go
type Cat struct {
	Name        string `json:"name"`
	Breed       string `json:"breed"`
	Personality string `json:"personality"`
	Available   bool   `json:"available"`
}

type Cats struct {
	List []*Cat `json:"cats"`
}
```

See the strings with back ticks? Those are annotations. Annotations can be used for a variety of functionality, and you can write your own custom annotations as well. One of the main use cases is to let Go's marshal functions know how to serialize your data. Let's update these to also specify that we want lowercase XML nodes.

```go
Name  string  `json:"name" xml:"name"`
```

If you restart your server and issue the request again, you should see that the XML nodes are now lowercase.

## Task Three - Adding an Endpoint

In the previous section, we created a function to return a filtered list of cats. Let's hook that into our API.

We'll create an endpoint for `GET /cat/:personality.`. In the `Gin` framework, you can specify variable parts of the URL by using the `:` prefix.

```go
r.GET("/cat/:personality", func(c *gin.Context) {
		accepts := c.Request.Header.Get("Accepts")
		result := cats.GetByPersonality(c.Param("personality"))
		...
	})
```

We can also access the value of that parameter using `c.Param("personality")`. Build out the endpoint to return the filtered result. You can test your endpoint by going to [http://localhost:8080/cat/diva](http://localhost:8080/cat/diva). If your code is working, you should recieve a response with just two cats.

## Task Four - Handling Empty Responses

What if we pass in a personality that doesn't match any of our cats? Right now, we would return a `Cats` struct with an empty `List`. A better experience might be to return a `404` to indicate we could not find any matching cats.

In go, you can check the length of a slice using the `len` function.

```go
if len(mySlice) > 0 {

}
```

You'll notice that we're checking to see if the value of `len(mySlice)` is greater than zero. This is because Go requires if statements to result in `boolean` values.

Once we do our check, if the length is less than zero, we'll want to create a custom response. We can do this by passing `http.StatusNotFound` as the first argument to our response. Additionally, we can create a custom message using `gin.H`, which is a map of interfaces.

```go
c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
```

**Note**: Gin will continue execution after `c.JSON` or `c.XML` is called. In order to avoid errors, you must have a `return` statement at the end of each `if` in order to not send multiple status codes back to the client.

[Part Three - Concurrency](CONCURRENCY.md)