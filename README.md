# Configuration
## Modules
Every submodule should have its own configuration. The command
to run is:
```
 go mod init [module] 
```

For example:
```
go mod init github.com/ocampor/go-api/src
```

The file `go.mod` is created with the module configuration. Then
to install the dependencies run

```
go get
```

All the dependencies are installed, and saved in the `go.mod`
for future references. The installed packages are stored in `$GOPATH/pkg`.

## Run the API

To run the API for the first time, install all the dependencies
with 

```
go run app.go
```