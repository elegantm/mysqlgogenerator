package model
import "time"

// {{.TableName}}
type {{.StrutName}} struct {
{{- range .Fields}}
	{{.Name}} {{.DataType}}  `db:"{{ .ColumnName}}" gorm:"column:{{ .ColumnName}}" ` //{{.ColumnComment}} 
{{- end}}
}

//select all column {{.TableName}}

sqlstr := "SELECT  {{- range .Fields}} `{{.ColumnName}}`{{.Separator}} {{- end}} FROM {{.TableName}} WHERE "

sqlInsert :="INSERT INTO {{.TableName}}({{- range .Fields}} `{{.ColumnName}}`{{.Separator}} {{- end}}) VALUES  " +
" ({{- range .Fields}} :{{.ColumnName}}{{.Separator}} {{- end}}) "