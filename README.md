# encrypt-secret-value-for-gh-actions
A Golang example of encrypting secrets in order to programmatically write them to the GitHub Actions Secrets API. This is not to be used an extension or program, but as an example that you can place in your own code.

# Prerequisites
- [Golang 1.19](https://go.dev/doc/install)
- [GitHub CLI](https://cli.github.com/)

# Usage
You'll need your repositories public key in order to encrypt a value. See the examaple below once you've retreived the public key.

## 1. Get Repo Public Key
1. Install [GH CLI](https://cli.github.com/)
2. Authenticate with your user who owns the repo you want the public key for (see [authentication instructions](https://cli.github.com/manual/gh_auth_login))
3. Execute in terminal (replace `<owner-of-repo>` and `<repo-name>` respectively): 

    ```sh
    gh api -q ".key" /repos/<owner-of-repo>/<repo-name>/actions/secrets/public-key
    ```

## 2. Example

1. Clone this repository
2. Execute in terminal:

    ```sh
    go get
    # replace <secret-value> and <repo-public-key> with real values
    go run encrypt.go "<secret-value>" "<repo-public-key>"
    ```