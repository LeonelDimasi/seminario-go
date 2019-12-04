package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Insertdb(ctx *gin.Context) {

	db, err := sql.Open("mysql", "root:root@/127.0.0.1:3306/agencia_autos")
	if err != nil {
		panic(err.Error())
	}

	queryCreate := "create table if not exists `cars` (" +
		"id integer not null auto_increment," +
		"brand varchar(255)," +
		"model varchar(255)," +
		"primary key (id))"
	stmtCreate, _ := db.Prepare(queryCreate)
	if _, err = stmtCreate.Exec(); err != nil {
		panic(err)
	}

	//defer CloseStatement(stmtCreate)
	defer db.Close()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK ",
	})
}

func GetCar(ctx *gin.Context) {

	idCar := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "id_car " + idCar,
	})
}

func main() {
	conexion()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/auto/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.GET("/auto/:id", GetCar)

	r.GET("/insert", Insertdb)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func conexion() {
	db, err := sql.Open("mysql", "root:root@/agencia_autos")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

/*
db, err := gorm.Open("mysql", "root:root@/sample_database?parseTime=true")
if err != nil {
    panic(err)
}

defer CloseDatabaseConnection(db)
*/
/*
db, err := sql.Open("mysql", "root:root@/sample_database")
if err != nil {
    panic(err.Error())
}

defer db.Close()
*/
/*
err = db.Ping()
if err != nil {
    panic(err.Error())
}

*/

/*


 */

/*
querySelect := "select * from `cars` where id = ?"

stmtSelect, _ := db.Prepare(querySelect)
var aux Car
if err = stmtSelect.QueryRow(1).Scan(&aux.Id, &aux.Brand, &aux.Model); err != nil {
    panic(err)
}

defer CloseStatement(stmtSelect)

*/
