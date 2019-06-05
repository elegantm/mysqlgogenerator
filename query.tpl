package model
import "time"


// {{.TableName}}  only db
type {{.StrutName}} struct {
{{- range .Fields}}
	{{.Name}} {{.DataType}}  `db:"{{ .ColumnName}}"` //{{.ColumnComment}} 
{{- end}}
}

// {{.TableName}}  db and json
type {{.StrutName}} struct {
{{- range .Fields}}
	{{.Name}} {{.DataType}}  `db:"{{ .ColumnName}}" json:"{{ .Name}}"` //{{.ColumnComment}} 
{{- end}}
}


// {{.TableName}}  empty
type {{.StrutName}} struct {
{{- range .Fields}}
	{{.Name}}: 
{{- end}}
}



// {{.TableName}}  db and gorm
type {{.StrutName}} struct {
{{- range .Fields}}
	{{.Name}} {{.DataType}}  `db:"{{ .ColumnName}}" gorm:"column:{{ .ColumnName}}" ` //{{.ColumnComment}} 
{{- end}}
}


//select all column {{.TableName}}


type {{.StrutName}}JSON struct {
{{- range .Fields}}
	{{.Name}} {{.DataType}}  `json:"{{ .Name}}" ` //
{{- end}}
}



sqlstr := "SELECT  {{- range .Fields}} `{{.ColumnName}}`{{.Separator}} {{- end}} FROM {{.TableName}} WHERE "

sqlInsert :="INSERT INTO {{.TableName}}({{- range .Fields}} `{{.ColumnName}}`{{.Separator}} {{- end}}) VALUES  " +
" ({{- range .Fields}} :{{.ColumnName}}{{.Separator}} {{- end}}) "