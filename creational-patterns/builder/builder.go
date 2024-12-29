package builder

import (
	"fmt"
	"strings"
)

//! Tarea: crear un QueryBuilder para construir consultas SQL
/**
* Debe de tener los siguientes métodos:
* - constructor(table: string)
* - select(fields: string[]): QueryBuilder -- si no se pasa ningún campo, se seleccionan todos con el (*)
* - where(condition: string): QueryBuilder - opcional
* - orderBy(field: string, order: string): QueryBuilder - opcional
* - limit(limit: number): QueryBuilder - opcional
* - execute(): string - retorna la consulta SQL
*
** Ejemplo de uso:
 const usersQuery = new QueryBuilder("users") // users es el nombre de la tabla
   .select("id", "name", "email")
   .where("age > 18")
   .where("country = 'Cri'")
   .orderBy("name", "ASC")
   .limit(10)
   .execute();

 console.log('Consulta: ', usersQuery);
 // Select id, name, email from users where age > 18 and country = 'Cri' order by name ASC limit 10;
*/

type OrderByDirection uint8

const (
	_ OrderByDirection = iota
	OrderByAsc
	OrderByDesc
)

func (d OrderByDirection) String() string {
	switch d {
	case OrderByAsc:
		return "asc"
	default:
		return "desc"
	}
}

type QueryBuilder struct {
	table           string
	fields          []string
	whereConditions []string
	orderBy         string
	limit           *uint
}

func NewQueryBuilder(table string) *QueryBuilder {
	return &QueryBuilder{
		table: table,
	}
}

func (qb *QueryBuilder) Select(fields ...string) *QueryBuilder {
	qb.fields = fields
	return qb
}

func (qb *QueryBuilder) Where(condition string) *QueryBuilder {
	qb.whereConditions = append(qb.whereConditions, condition)
	return qb
}

// If no direction is passed, defaults to descending
func (qb *QueryBuilder) OrderBy(field string, direction ...OrderByDirection) *QueryBuilder {
	if len(direction) == 0 {
		direction = append(direction, OrderByDesc)
	}

	dir := direction[0]

	qb.orderBy = fmt.Sprintf("%s %s", field, dir)
	return qb
}
func (qb *QueryBuilder) Limit(n uint) *QueryBuilder {
	qb.limit = &n
	return qb
}

func (qb *QueryBuilder) Execute() string {
	query := fmt.Sprintf("FROM %s", qb.table)

	if len(qb.fields) == 0 {
		qb.fields = []string{"*"}
	}

	query += fmt.Sprintf(" SELECT %s", strings.Join(qb.fields, ", "))

	if len(qb.whereConditions) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(qb.whereConditions, " AND "))
	}

	if qb.orderBy != "" {
		query += fmt.Sprintf(" ORDER BY %s", qb.orderBy)
	}

	if qb.limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *qb.limit)
	}

	return query
}
