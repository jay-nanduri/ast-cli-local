[![CircleCI](https://circleci.com/gh/CheckmarxDev/ast-cli/tree/master.svg?style=svg&circle-token=32eeef7505db60c11294e63db64e70882bde83b0)](https://circleci.com/gh/CheckmarxDev/ast-cli/tree/master)
## Building from source code
### Windows 
``` powershell
setx GOOS=windows 
setx GOARCH=am
go build -o ./bin/cx.exe ./cmd
```

### Linux

``` bash
export GOARCH=amd64
export GOOS=linux
go build -o ./bin/cx ./cmd
```

### Macintosh

``` bash
export GOOS=darwin 
export GOARCH=amd64
go build -o ./bin/cx-mac ./cmd
```

** **

## Basic CLI Operation

### Windows
``` powershell
cx.exe [commands]
```

### Linux/Mac
``` bash
./cx [commands]
```

The parameters accepted by the CLI vary based on the commands issued and they will be described thoroughly throughout this document. The following global parameters affect all actions:

- (--base-uri), the URL of the AST server.
- (--base-auth-uri), optionally provides alternative KeyCloak endpoint to (--base-uri).
- (--client-id), the client ID used for authentication (see Authentication documentation).
- (--secret), the secret that corrosponds to the client-id  (see Authentication documentation).
- (--token), the token to authenticate with (see Authentication documentation).
- (--proxy), optional proxy server to use (see Proxy Support documentation).
- (--insecure), indicates CLI should ignore TLS certificate validations.
- (--profile), specifies the CLI profile to store options in (see Environment and Configuration documentation).

Many CLI variables can be provided using environment variables, configuration variables or CLI parameters. The follow precidence is used when the same value is found in settings:

1. CLI parameters, these always overide configuration and environment variables.
2. Configuration variables always overide environment variables.
3. Environment variables are the first order precidence. 

## CLI Configurations

The CLI allows you to permanently store some CLI options in configuration files. The configuration files are kept in the users home directory under a subdirectory named  ($HOME/.checkmarx). 

``` bash
./cx configure set cx_base_uri "http://<your-server>[:<port>]"
./cx configure set cx_ast_access_key_id <your-key>
./cx configure set cx_ast_access_key_secret <your-secret>
./cx configure set cx_http_proxy <your-proxy>
./cx configure set cx_token <your-token>
```

The (--profile) option provides a powerful tool to quickly switch through differnent sets of configurations. You can add a profile name to any CLI command and it utilize the corsponding profile settings. The following example setups up altnerative profile named "test" and calls a CLI command utilizing it.

``` bash
# Create the alternative profile
./cx configure set cx_base_uri "http://your-test-server" --profile test
# Use the cx_base_uri from the "test" profile.
./cx scan list --profile test
# This uses the default profile (if it exists)
./cx scan list
```

These values can be stored in CLI configurations:

- cx_token: the token to authenticate with (see Authentication documentation).
- cx_base-uri: the URL of the AST server.
- cx_http_proxy: optional proxy server to use (see Proxy Support documentation). 
- cx_ast_access_key_id: the client ID used for authentication (see Authentication documentation).
- cx_ast_access_key_secret: the secret that corrosponds to the client-id  (see Authentication documentation).

## Authentication

The CLI supports token and key/secret based authentication.

Token based authentication is the easiest method to use in CI/CD environments. Tokens are generated through KeyCloak and can be created with a predictable lifetime. Once you have a token you can use it from the CLI like this:

``` bash
./cx --token <your-token> scan list 
```

You can optionally configure the token into your stored CLI configuration values like this:

``` bash
./cx configure set cx_token <your-token>
# The following command will automatically use the stored token
./cx scan list
```

You can also store the token in the environment like this:

``` bash
export CX_TOKEN=<your-token>
./cx scan list
```

Key/secret authentication requires you to first use an AST username and password to create the key and secret for authentication. The following example shows how to create a key/secret and then use it:

``` bash
./cx auth register -u <username> -p <password>
CX_AST_ACCESS_KEY_ID=<generated-key>
CX_AST_ACCESS_KEY_SECRET=<generated-secret>
```

Once you generated your key and secret they can be used like this:

``` bash
./cx --client-id <your-key> --secret <your-secret> scan list 
```

You can optionally configure the key/secret into your stored CLI configuration values like this:

``` bash
./cx configure set cx_ast_access_key_id <your-key>
./cx configure set cx_ast_access_key_secret <your-secret>
# The following command will automatically use the stored key/secret
./cx scan list
```

You can also store the key/secret in the environment like this:

``` bash
export CX_AST_ACCESS_KEY_ID=<your-key>
export CX_AST_ACCESS_KEY_SECRET=<your-secret>
./cx scan list
```



## Triggering Scans

You have many options when it comes to creating scans, the most important think you need to figure out is where you're going to create the scan from. You have the following 3 possible scan sources:

1. A zip file with your source code.
2. A directory with your source code.
3. A host git repo.

**NOTE**: for simplicity the following examples assume you have stored your authentication and base-uri information in either environment variables or CLI configuration parameters.

If you decide to scan a local directory you can provide filters that determine which resources are sent to the AST server. The filters are based on an inclusion and exclusion model. The following example shows how to scan a folder:

``` bash
./cx scan create -d <path-to-your-folder> -f "s*.go"
```

Todo. more examples

You can create a scan like this:

``` bash
./cx
```



## Managing Projects



## Retreiving Results

...todo, we're still waiting to work out result retreival

## Proxy Support

The CLI full supports proxy servers and optional proxy server authentication. When the proxy server variable is found all CLI operations will be routed through the target server. Proxy server URLs should like this: "http[s]://your-server.com:[port]"

Proxy support is enabled by creating an environment variable named CX_HTTP_PROXY. You can also specify the proxy by storing the proxy URL in a CLI configuration variable (cx_proxy_http), or directly through with the CLI with the parameter  (--proxy).

The following example demonstraights the use of a proxy server:

``` bash
./cx scan list --proxy "http://<your-proxy>:8081"
```



## Environment Variables

| Environment Variable         | Description                                                  |
| ---------------------------- | ------------------------------------------------------------ |
| **CX_AST_ACCESS_KEY_ID**     | Key portion of key/secret authentication pair.               |
| **CX_AST_ACCESS_KEY_SECRET** | Secret portion of key/secret authentication pair.            |
| **CX_TOKEN**                 | Token for token based authentication.                        |
| **CX_BASE_URI**              | The URI of the AST server.                                   |
| **CX_BASE_IAM_URI**          | The URI of KeyCloak instance. This optional and only required when you're not using AST's built in KeyCloak instance. |
| **CX_HTTP_PROXY**            | When provided this variable will trigger the CLI to use the proxy server pointed to (see proxy support documentation). |

