package jsonast

// Value is a generic node in a JSON AST. This implements fmt.Stringer
type Value interface {
	// IsString returns true if the data at this node is a string
	IsString() bool
	// IsNumber returns true if the data at this node is numeric
	IsNumber() bool
	// IsBool returns true if the data at this node is a bool
	IsBool() bool
	// IsObject returns true if the data at this node is an object
	IsObject() bool
	// IsArray returns true if the data at this node is an array
	IsArray() bool
	// IsNull returns true if the data at this node is null
	IsNull() bool
}

type valueImpl struct {
	isString bool
	isNumber bool
	isBool   bool
	isObject bool
	isArray  bool
	isNull   bool
}

func (v valueImpl) IsString() bool { return v.isString }
func (v valueImpl) IsNumber() bool { return v.isNumber }
func (v valueImpl) IsBool() bool   { return v.isBool }
func (v valueImpl) IsObject() bool { return v.isObject }
func (v valueImpl) IsArray() bool  { return v.isArray }
func (v valueImpl) IsNull() bool   { return v.isNull }
