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


## Useful links

[EMV Lab](https://emvlab.org/emvtags/all/)

## Adding new EMV Tags to the project

1. in the file **constants.go**

Put a new constant
```javascript
    TagCardholderName = "5F20"
```

Insert in **loadTags()** methods
```javascript
    //load Emv tags available
    func loadTags() {
        tagMap[TagCardholderName] = tag{name: TagCardholderName, minSize: 0, maxSize: 0}

    }    
```

2. in the file **format.go** insert the conditional

```javascript
    if t.name == TagCardholderName {
		return format(TagCardholderName, t.value)
	}
```

## Useful commands:

###  Creating a version of this lib:

* 1- Test all!
```sh
    go test ./... -v
```

* 2- Create a git tag with the version number
```sh
    git tag -a v0.0.3 -m "my version 0.0.3"
```

* 3- Push the tag to git repo
```sh
    git push origin v0.0.3
```

* 4- Now the new version can be dowloaded to the other projects
```sh
    go get -v "github.com/wagner-aos/emv-process"
```


