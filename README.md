# tpass - A simple Go password generater

## Usage 

**Install**
```go
go get github.com/cn-maul-tpass
```

**Cite**
```go
var Default = tpass.Flag{
		Length:  18,
		Digit:   true,
		Upper:   true,
		Lower:   true,
		Special: true,
	}

	pass, err := Default.Generate()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pass)
```