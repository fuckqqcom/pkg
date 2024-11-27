package main

import (
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

func main() {
	// Build a SQL to select a user from database.
	//sb := sqlbuilder.NewSelectBuilder().Select("name", "level").From("users")
	//sb.Where(
	//	sb.Equal("id", 1234),
	//)
	//fmt.Println(sb)

	ub := sqlbuilder.Update("users")
	ub.Set(
		ub.Incr("level"),
	)

	// Set the WHERE clause of UPDATE to the WHERE clause of SELECT.
	//ub.WhereClause = sb.WhereClause
	sql, args := ub.Build()
	fmt.Println(sql, args)

	// Output:
	// SELECT name, level FROM users WHERE id = ?
	// UPDATE users SET level = level + ? WHERE id = ?
}
