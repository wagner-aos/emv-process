# EMV Process

Emv Encode and Decode library based on builder pattern

| Date | Author | Comments | Version |
| --- | --- | --- | --- |
| 10/05/2020 | Wagner aka Barão | This is the Beginning of success | 0.0.1 |

### What is EMV?

[Wikipedia](https://en.wikipedia.org/wiki/EMV)

### What you can do with this lib?

With this lib you can Encode and Decode EMV tags for handling in your code.

### How to use in you code:

* import:

```go

    import "github.com/wagner-aos/emv-process"
```

* build and EMV object:

```go

    emv := NewEMV().
            AddTag("5F20", "WAGNER ALVES").
            AddTag("9F0B", "WAGNER ALVES DE OLIVEIRA SILVA").
            Build()
```


* Encode TLV string:

```go
    
    tlv := emv.ToTLV()

```

* More examples you can see in test go files!



