# Mash Token

Use the TIBCO Cloud Mashery V3 API to get a login token to perform other calls

## Installation

```bash
flogo install github.com/project-flogo/tibco-contrib/activity/mashtoken
```

Link for flogo web:

```
https://github.com/project-flogo/tibco-contrib/activity/mashtoken
```

## Inputs

| Input     | Description    |
|:----------|:---------------|
| username  | The username for which you want to generate a token |
| password  | The password associated with the username |
| scope     | The scope for which you want to generate the token (this is the Area ID) |
| basicauth | The username / password combination used to connect to Mashery (must be in format `user:pass` and is likely not the same as the username and password above) |

## Outputs

| Output       | Description                                             |
|:-------------|:--------------------------------------------------------|
| accesstoken  | The access token generated by the Mashery API           |
| tokentype    | The type of token generated by the Mashery API          |
| expiresin    | The amount of time (seconds) in which the token expires |
| refreshtoken | The token to use to refresh the access token            |
| scope        | The scope for which the access token is valid (Area)    |
