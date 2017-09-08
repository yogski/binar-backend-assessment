package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	_ "log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Order []struct {
	OrderID    string    `json:"orderId"`
	Timestamp  time.Time `json:"timestamp"`
	CustomerID string    `json:"customerId"`
	SellerID   string    `json:"sellerId"`
	Payment    string    `json:"payment"`
	Ordered    []struct {
		MenuID   string `json:"menuId"`
		MenuName string `json:"menuName,omitempty"`
		Price    int    `json:"price,omitempty"`
		Qty    int    `json:"votes,omitempty"`
	} `json:"ordered"`
}

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/binar_go")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	router := gin.Default()

	// GET a order detail
	router.GET("/order/:id", func(c *gin.Context) {
		var (
			order Order
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select * from order where orderid = ?;", id)
		err = row.Scan(&order.OrderID, &order.Timestamp, &order.CustomerID, &order.SellerID, &order.Payment)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": order,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// GET all orders
	router.GET("/orders", func(c *gin.Context) {
		var (
			order  Order
			orders []Order
		)
		rows, err := db.Query("select * from order;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&order.OrderID, &order.Ordered)
			orders = append(orders, order)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": orders,
			"count":  len(orders),
		})
	})

	// POST new order details
	router.POST("/order", func(c *gin.Context) {
		var buffer bytes.Buffer
		name := c.PostForm("name")
		age := c.PostForm("age")
		weight := c.PostForm("weight")
		verified := c.PostForm("verified")
		stmt, err := db.Prepare("insert into order (name, age, weight, verified) values(?,?,?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(name, age, weight, verified)

		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(name)
		defer stmt.Close()
		disp_name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", disp_name),
		})
	})

	// PUT - update a order details
	router.PUT("/order", func(c *gin.Context) {
		var buffer bytes.Buffer
		id := c.Query("id")
		first_name := c.PostForm("first_name")
		last_name := c.PostForm("last_name")
		stmt, err := db.Prepare("update order set first_name= ?, last_name= ? where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(first_name, last_name, id)
		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(first_name)
		buffer.WriteString(" ")
		buffer.WriteString(last_name)
		defer stmt.Close()
		name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", name),
		})
	})

	// Delete resources
	router.DELETE("/order", func(c *gin.Context) {
		id := c.Query("id")
		stmt, err := db.Prepare("delete from order where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted order: %s", id),
		})
	})
	router.Run(":8989")
}