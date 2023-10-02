# jwt-viewer
Basic program to view base64 encoded jwts

`jwt-viewer` reads from stdin. Just pipe your jwt into the program.

Example usage:

```sh
echo $JWT_IN_VARIABLE | jwt-viewer
```

Is your jwt stored in a file? Redirect it into the program

```sh
jwt-viewer < file-with-jwt
```
