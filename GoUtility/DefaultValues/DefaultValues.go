package DefaultValues

import (
	"reflect"
	"strconv"
)

var type_info_by_type_name map[string]*typeInfo = make(map[string]*typeInfo)

/*
This method will return the type information stored
in the static global for this module. If the type
information does not exist this method will create it

NOTICE: If you are trying to follow the logic in this
module this is where the "LAZY" loading occurs.

Process: When you call New for the first time New
will call this function. This function will determine
that there is no type information for the provided type.

This will trigger this function to actively load the
default values for this new type and create the type_information

This will allow future calls in the same process to resuse the
default values
*/
func get_type_information_from_global(reflect_type reflect.Type) *typeInfo {
	var type_name = reflect_type.Name()
	type_information, ok := type_info_by_type_name[type_name]
	if !ok {
		//create a new type object and populate it with
		//the field information
		type_information = new(typeInfo)
		type_information.ReflectType = reflect_type
		type_info_by_type_name[type_name] = type_information
		populate_type_info(type_information)
	}
	return type_information
}

/*
This method will use reflection to set
the default value we have stored to the
value in the struct that we are actively creating
*/
func set_default_value_to_concrete_object(type_information *typeInfo, index int, reflected_element_value reflect.Value) {
	//Get the default value we stored in "static" memory
	var default_value = type_information.DefaultValues[index]
	//Don't set nil, just let the default default do that
	if default_value != nil {
		//The names get really confusing here
		var default_value_realized = reflect.Indirect(reflect.ValueOf(default_value))
		var field = reflected_element_value.Field(index)
		//if canset fails this will panic
		//and no one wants to panic!
		if field.CanSet() {
			field.Set(default_value_realized)
		}
	}
}

/*
Will create a new instance of the struct
provided and call default values on the struct

NOTICE: This method is where to look to discover the lazy
loading of the default values

New calls

	get_type_information_from_global

which calls

	populate_type_info
*/
func New[T any]() *T {
	var new_object = new(T)
	var type_information = get_type_information_from_global(reflect.TypeOf(new_object))
	//assign values to the new object
	var reflected_element_value = reflect.Indirect(reflect.ValueOf(new_object).Elem())
	for i := 0; i < type_information.ReflectType.Elem().NumField(); i++ {
		set_default_value_to_concrete_object(type_information, i, reflected_element_value)
	}
	return new_object
}

/*
This method will loop through the types in the
struct and look for annotations or tags that have
default values specified and popuate the fieldInfo
list
*/
func populate_type_info(type_information *typeInfo) {
	type_information.DefaultValues = []any{}
	var reflected_element_type = type_information.ReflectType.Elem()
	//var common_error error
	//Loop through every field by index
	for i := 0; i < reflected_element_type.NumField(); i++ {
		var current_field = reflected_element_type.Field(i)
		process_field_and_get_default_value(i, &current_field, type_information)

	} //end for
} //end populate_type_info

/*
This function will process the field value
and do it's best to ascertain, determine, discover, interpret, or get
the value provided inside the tag value.

SORRY, this is a massive function with a lot of resused code. This
should be relatively easy for a learner to follow
*/
func process_field_and_get_default_value(index int, current_field *reflect.StructField, type_information *typeInfo) {
	//The value for tag value is the
	//value provided in the annotation
	//in the struct. For example
	//Foo        string `default:"bar"`
	// fmt.PrintLn(Foo)
	// bar
	var tag_value = current_field.Tag.Get("default")
	//nil is used to tell the New method to ignore this information
	//if you want a nil value don't use default, duh
	type_information.DefaultValues = append(type_information.DefaultValues, nil)
	//Ditto with the empty string. The empty string is the default value for
	//a string value so it is silly to go any further if the tag value
	//is set to the empty string ""
	if tag_value != "" {
		//Search the "built-in" types first
		switch current_field.Type.Kind() {
		case reflect.Uint:
			value, common_error := strconv.ParseUint(tag_value, 10, 0)
			if common_error == nil {
				type_information.DefaultValues[index] = uint(value)
			}
		case reflect.Uint64:
			value, common_error := strconv.ParseUint(tag_value, 10, 64)
			if common_error == nil {
				type_information.DefaultValues[index] = uint64(value)
			}
		case reflect.Uint32:
			value, common_error := strconv.ParseUint(tag_value, 10, 32)
			if common_error == nil {
				type_information.DefaultValues[index] = uint32(value)
			}
		case reflect.Uint16:
			value, common_error := strconv.ParseUint(tag_value, 10, 16)
			if common_error == nil {
				type_information.DefaultValues[index] = uint16(value)
			}
		case reflect.Uint8:
			value, common_error := strconv.ParseUint(tag_value, 10, 8)
			if common_error == nil {
				type_information.DefaultValues[index] = uint8(value)
			}
		case reflect.Int:
			value, common_error := strconv.ParseInt(tag_value, 10, 0)
			if common_error == nil {
				type_information.DefaultValues[index] = int(value)
			}
		case reflect.Int64:
			value, common_error := strconv.ParseInt(tag_value, 10, 64)
			if common_error == nil {
				type_information.DefaultValues[index] = int64(value)
			}
		case reflect.Int32:
			value, common_error := strconv.ParseInt(tag_value, 10, 32)
			if common_error == nil {
				type_information.DefaultValues[index] = int32(value)
			}
		case reflect.Int16:
			value, common_error := strconv.ParseInt(tag_value, 10, 16)
			if common_error == nil {
				type_information.DefaultValues[index] = int16(value)
			}
		case reflect.Int8:
			value, common_error := strconv.ParseInt(tag_value, 10, 8)
			if common_error == nil {
				type_information.DefaultValues[index] = int8(value)
			}
		case reflect.Complex64:
			value, common_error := strconv.ParseComplex(tag_value, 64)
			if common_error == nil {
				type_information.DefaultValues[index] = complex64(value)
			}
		case reflect.Complex128:
			value, common_error := strconv.ParseComplex(tag_value, 128)
			if common_error == nil {
				type_information.DefaultValues[index] = complex128(value)
			}
		case reflect.Bool:
			value, common_error := strconv.ParseBool(tag_value)
			if common_error == nil {
				type_information.DefaultValues[index] = bool(value)
			}
		case reflect.String:
			type_information.DefaultValues[index] = tag_value
		default:
			type_information.DefaultValues[index] = nil
		} //end switch
	} //end if not empty string
} //end process_field_and_get_default_value

type typeInfo struct {
	ReflectType   reflect.Type
	DefaultValues []any
}
