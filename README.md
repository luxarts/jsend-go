# jsend-go
Go wrapper for [JSend specification](https://github.com/omniti-labs/jsend)

# Installation
```
go get -u github.com/luxarts/jsend-go
```

# Usage

1. Import package
```
import jsend "github.com/luxarts/jsend-go"
```

2. Use the functions
```
// Create an error response with code number
err := errors.New("error doing something")
errorBody := jsend.NewError("Description about error.", err, 1234)

// Create a fail response for missing parameter
data := struct {
    Parameter string `json:"parameter"`
}{
    Parameter: "Missing parameter",
}
failBody := jsend.NewFail(data)

// Create a success response
data := struct {
    Parameter string `json:"parameter"`
}{
    Parameter: "value",
}
successBody := jsend.NewSuccess(data)

```

