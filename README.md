# go-akerun
Akerun API Client for Golang

## Installation
```sh
$ go get -u github.com/yshimada0330/go-akerun
```

## Usage

### Authenticate
Create a new client for the Akerun API.
```sh
$ export CLIENT_ID=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
$ export CLIENT_SECRET=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
$ export REDIRECT_URL=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

```go
clientID := os.Getenv("CLIENT_ID")
clientSecret := os.Getenv("CLIENT_SECRET")
redirectURL := os.Getenv("REDIRECT_URL")
conf := akerun.NewConfig(clientID, clientSecret, redirectURL)
client := akerun.NewClient(conf)
```

Getting an Authentication Code
```go
url := client.AuthCodeURL("state")
fmt.Println(url)
```

Creating an Access Token
```go
ctx := context.Background()
token, err := client.Exchange(ctx, authCode)
if err != nil {
    log.Fatal(err)
}
fmt.Println(token.AccessToken)
fmt.Println(token.RefreshToken)
```

Refresh an Access Token
```go
ctx := context.Background()
token := &oauth2.Token{
    AccessToken:  accessToken,
    RefreshToken: refreshToken,
    Expiry:       time.Now().Add(-time.Hour), # expires
}
reshTokentoken, err := client.RefreshToken(ctx, token)
if err != nil {
    log.Fatal(err)
}
fmt.Println(reshTokentoken.AccessToken)
fmt.Println(reshTokentoken.RefreshToken)
```

Revoke an Access Token
```go
ctx := context.Background()
err := client.Revoke(ctx, token)
if err != nil {
    log.Fatal(err)
}
```



### Example

Get a list of Organization IDs to which the token owner belongs.
```go
token := &oauth2.Token{
    AccessToken:  token.AccessToken,
    RefreshToken: token.RefreshToken,
}
params := &akerun.OrganizationsParams{
    Limit: 100,
}
result, err := client.GetOrganizations(ctx, token, *params)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%#v\n", result)
```
