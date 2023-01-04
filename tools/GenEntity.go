package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"shrimp_blog_sever_v2/config"
	_ "shrimp_blog_sever_v2/config"
	"shrimp_blog_sever_v2/constant"
	"shrimp_blog_sever_v2/utils"
)

const structEnd = "}\n\n"

var typeConvert = map[string]string{
	"varchar":          "string",
	"int":              "int",
	"datetime":         "*time.Time",
	"tinyint":          "int",
	"int unsigned":     "uint",
	"tinyint unsigned": "uint",
	"text":             "string",
}

func init() {
	config.InitGoble(".")
}

func main() {
	projectPath, err := filepath.Abs(".")

	if err != nil {
		panic(err)
	}

	genModelFile(projectPath, "shrimp_blog")
}

func genEntity(dbName, tableName string) string {
	content := fmt.Sprintf("package model \n\nimport (\n\t\"time\"\n)  \n\ntype %s struct {\n", utils.ToUpperCamel(tableName))

	columns, err := constant.DBOp.Queryx("select COLUMN_NAME,COLUMN_TYPE from information_schema.columns where  TABLE_SCHEMA = ? and TABLE_NAME = ?", dbName, tableName)
	if err != nil {
		panic(err)
	}
	for columns.Next() {
		column := make(map[string]interface{})
		columns.MapScan(column)
		columnType := regexp.MustCompile(`[^a-zA-Z ]+`).ReplaceAllString(string(column["COLUMN_TYPE"].([]uint8)), "")
		columnName := string(column["COLUMN_NAME"].([]uint8))
		var name = utils.ToUpperCamel(columnName)
		content += fmt.Sprintf("  %s %s   `json:\"%s\" db:\"%s\"` \n", name, typeConvert[columnType], columnName, columnName)
	}
	content += structEnd
	return content
}

func genModelFile(projectPath, dbName string) {
	rows, err := constant.DBOp.Queryx("select table_name from information_schema.TABLES where TABLE_SCHEMA =?", dbName)
	if err != nil {
		panic(err)
	}

	var tableName string
	for rows.Next() {
		result := make(map[string]interface{})
		if err := rows.MapScan(result); err != nil {
			panic(err)
		}
		tableName = string(result["TABLE_NAME"].([]uint8))
		if err != nil {
			panic(err)
		}
		// 生成文件
		fileName := tableName + ".go"

		content := genEntity(dbName, tableName)

		path := filepath.Join(projectPath, "model", fileName)

		_, err = os.Stat(path)
		if err == nil {
			err := os.Remove(path)
			if err != nil {
				fmt.Println(err)
			}
		}

		// 写文件
		file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
		if os.IsNotExist(err) {
			//file, err = os.Create(path)
			//if err != nil {
			//	panic(err)
			//}
			panic(err)
		}
		file.Write([]byte(content))
		fmt.Print(content)
	}
}

func genColumns(dbName, tableName string) string {
	columns, err := constant.DBOp.Queryx("select COLUMN_NAME from information_schema.columns where  TABLE_SCHEMA = ? and TABLE_NAME = ?", dbName, tableName)
	if err != nil {
		panic(err)
	}
	columnsSlice := make([]map[string]interface{}, 0)
	for columns.Next() {
		column := make(map[string]interface{})
		columns.MapScan(column)
		columnsSlice = append(columnsSlice, column)
	}

	content := fmt.Sprintf("type %sTable struct {\n\tTable   BaseTable\n\tColumns %sColumns\n}\n\ntype %sColumns struct {\n", utils.ToUpperCamel(tableName), utils.ToLowerCamel(tableName), utils.ToLowerCamel(tableName))

	for i := range columnsSlice {
		column := columnsSlice[i]

		columnName := string(column["COLUMN_NAME"].([]uint8))
		var name = utils.ToUpperCamel(columnName)
		content += fmt.Sprintf("    %s string\n", name)
	}
	content += structEnd

	content += fmt.Sprintf("var %sColumn = %sColumns{\n", utils.ToLowerCamel(tableName), utils.ToLowerCamel(tableName))
	for i := range columnsSlice {
		column := columnsSlice[i]

		columnName := string(column["COLUMN_NAME"].([]uint8))
		var name = utils.ToUpperCamel(columnName)

		content += fmt.Sprintf("    %s:         \"%s\",\n", name, columnName)
	}
	content += structEnd

	content += fmt.Sprintf("func New%sTable() *%sTable {\n\treturn &%sTable{\n\t\tTable:   BaseTable{Table: \"%s\"},\n\t\tColumns: %sColumn,\n\t}\n}\n\n", utils.ToUpperCamel(tableName),
		utils.ToUpperCamel(tableName), utils.ToUpperCamel(tableName), tableName, utils.ToLowerCamel(tableName))

	return content
}
