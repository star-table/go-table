package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(mysql.Open("root:JoLNCfUX0-7Xgre_W@(172.19.166.130:33309)/lesscode-go?charset=utf8mb4&parseTime=True&loc=Local"))
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/data/po/gen/output/",
	})

	g.UseDB(db)

	// generate all table from database
	g.ApplyBasic(g.GenerateModel("lc_app_table"))

	g.Execute()
}
