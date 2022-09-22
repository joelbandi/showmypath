# showmypath
A cli util to pretty print you `$PATH` variable in any linux based env - written in go


## Usage

1. make sure you have go installed and $GOPATH setup
2. run `go get github.com/joelbandi/showmypath@v1.0.0`
3. use `$ showmypath` to get a indented list view of the path


If you run into problems like it cant find the comman, make sure youre `$GOPATH/bin` is in your `$PATH`. This command `export $PATH=$PATH:$GOPATH/bin` should do the trick. 
