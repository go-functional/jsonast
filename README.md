# jsonast

This repository is dedicated to parsing a JSON string into an 
abstract syntax tree (AST).

Beware: it is pre-alpha and a major work in progress!

# Example Usage

```go
type Address struct {
    StreetNum int
    StreetName string
    City string
    State string
}

val := jsonast.Parse(`{
    "street_num":1,
    "street_name": "Microsoft Way",
    "city": "Redmond",
    "state": "WA"
}`)

walked := val.
    Walk().
    Field("street_num").
    And().Field("street_name")
    And().Field("city").
    And().Field("state").Validate(func(v Value) bool {
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
    })

if transformer.Err() != nil {
    log.Fatal(transformer.Err())
}

// transformer now has 3 jsonast.String fields in its Strings property
// and 1 jsonast.Number field in its Numbers property.
//
// you can now validate on these fields
address := Address{
    StreetNum: transformer.Numbers[0],
    StreetName: transformer.Strings[0],
    City: transformer.Strings[1],
    State: transformer.Strings[2],
}
