# GO AST Sandbox
Stay tuned for Quantum Dempsey blog post on using Go to create a Go code sanbox safely by using the AST library. 

## Settup 
```
docker build -t go-server . --no-cache
docker run -p 8080:8080 go-server
```
Navigate to localhost:8080


## Usage
Can be used and placed as an iframe in an existing web app. 


## TODO:
- Need to handle if the user already has imports
- handle all types of compilation errors and 

## Possible sandbox escapes:
- entering into a defer func for code recovery to continue execution. 
- os.Exit(1)