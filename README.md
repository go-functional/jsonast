# jsonast

This repository is dedicated to parsing a JSON string into an 
abstract syntax tree (AST).

Based on the following Scala libraries:

- [lift-json](https://github.com/lift/lift/tree/master/framework/lift-base/lift-json)
- [lift-json-scalaz](https://github.com/lift/framework/tree/master/core/json-scalaz)

Beware: it is pre-alpha and a major work in progress!

# Example Usage

```go
type Address struct {
    StreetNum int
    StreetName string
    City string
    State string
    GateCode int
}

val := jsonast.Parse(`{
    "street_num":1,
    "street_name": "Microsoft Way",
    "city": "Redmond",
    "state": "WA",
    "other_info": {
        "needs_signature": false,
        "gate_codes": ["abcd", "efgh"]
    }
}`)

walked := jsonast.NewWalker(val).
    Field("street_num").
    Field("street_name").
    Field("city").
    Field("state").Validate(func(v Value) bool {
        str, ok := v.(jsonast.String)
        if !ok {
            return false
        }
        // only accept the JSON if the state has a coast on the pacific
        // ocean
        return str.String() == "AK" || 
            str.String() == "WA" ||
            str.String() == "OR" ||
            str.String() == "CA"
    }).
    Path(
        jsonast.NewStringPathElt("other_info"),
        jsonast.NewStringPathElt("gate_codes"),
        jsonast.NewIntPathElt(0),
    )

if transformer.Err() != nil {
    log.Fatal(transformer.Err())
}

// walked now has the following fields in the following properties:
//
// - .Strings has 3 jsonast.String values (street_name, city, state)
// - .Numbers has 2 jsonast.Number values (street_num, other_info.gate_codes[0])
//
// you can now create an Address from these fields
address := Address{
    StreetNum: walked.Numbers[0],
    StreetName: walked.Strings[0],
    City: walked.Strings[1],
    State: walked.Strings[2],
    GateCode: walked.Numbers[1],
}
```
