# EMV Process

Emv Encode and Decode library based on builder pattern

| Date | Author | Comments | Version |
| --- | --- | --- | --- |
| 10/05/2020 | Wagner aka Barão | This is the Beginning of success | 0.0.3 |

### What is EMV?

[Wikipedia](https://en.wikipedia.org/wiki/EMV)

### What you can do with this lib?

With this lib you can Encode and Decode EMV tags for handling in your code.

### How to use in you code:

* Get from repository

```sh
    go get -v "github.com/wagner-aos/emv-process"
```


* import in your go file:

```go

    import "github.com/wagner-aos/emv-process"
```

* build and EMV object:

```go

    emvStruct := emv.NewEMV().
            AddTag("5F20", "WAGNER ALVES").
            AddTag("9F0B", "WAGNER ALVES DE OLIVEIRA SILVA").
            Build()
```


* EMV Struct to TLV string:

```go
    
    tlv := emvStruct.ToTLV()

```

* TLV string to EMV Struct:

```go

    payload := "9F260854BD15AE210A51179F2701809F10180210A50002020000000000000000000000FF"

    emvStruct := emv.NewEMV().Build()
	emvStruct.Parse(payload)
 

```

* Get a tag values by tag name:

```go

    tagName, tagValue, tagSize, err := emvStruct.GetTag("5F20")

``` 

* More examples you can see in test go files!


## TODO List

* Validate tag field size in builder
* Validate tag field size in tlv

