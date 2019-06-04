package app

import (
	"github.com/iancoleman/strcase"

)

func ConvertToCamel(name string) string {
	return strcase.ToCamel(name)
}


func ConvertMysqlType(dt string)string  {
	precision := 0
	var  nilVal string

	var typ string

switchDT:
	switch dt {
	case "bit":
		nilVal = "0"
		if precision == 1 {
			nilVal = "false"
			typ = "bool"
			break switchDT
		} else if precision <= 8 {
			typ = "uint8"
		} else if precision <= 16 {
			typ = "uint16"
		} else if precision <= 32 {
			typ = "uint32"
		} else {
			typ = "uint64"
		}

	case "bool", "boolean":
		nilVal = "false"
		typ = "bool"


	case "char", "varchar", "tinytext", "text", "mediumtext", "longtext":
		nilVal = `""`
		typ = "string"


	case "tinyint":
		//people using tinyint(1) really want a bool
		if precision == 1 {
			nilVal = "false"
			typ = "bool"

			break
		}
		nilVal = "0"
		typ = "int8"


	case "smallint":
		nilVal = "0"
		typ = "int16"


	case "mediumint", "int", "integer":
		nilVal = "0"
		typ = "int"


	case "bigint":
		nilVal = "0"
		typ = "int64"


	case "float":
		nilVal = "0.0"
		typ = "float32"


	case "decimal", "double":
		nilVal = "0.0"
		typ = "float64"


	case "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
		typ = "[]byte"

	case "timestamp", "datetime", "date":
		nilVal = "time.Time{}"
		typ = "time.Time"


	case "time":
		// time is not supported by the MySQL driver. Can parse the string to time.Time in the user code.
		typ = "string"

	default:
		typ = "unknown"
	}

	_ = nilVal

	return  typ
}