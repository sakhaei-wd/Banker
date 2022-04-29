The go-github library does not directly handle authentication. Instead, when
creating a new client, pass an `http.Client` that can handle authentication for
you. The easiest and recommended way to do this is using the [oauth2][]
library, but you can always use any other library that provides an
`http.Client`. If you have an OAuth2 access token (for example, a [personal
API token][]), you can use it with the oauth2 library using:


```go
import "golang.org/x/oauth2"

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
}
```
For complete usage of go-github, see the full [package docs][].

[GitHub API v3]: https://docs.github.com/en/rest
[oauth2]: https://github.com/golang/oauth2
[oauth2 docs]: https://godoc.org/golang.org/x/oauth2
[personal API token]: https://github.com/blog/1509-personal-api-tokens
[package docs]: https://pkg.go.dev/github.com/google/go-github/v43/github
[GraphQL API v4]: https://developer.github.com/v4/
[shurcooL/githubv4]: https://github.com/shurcooL/githubv4
