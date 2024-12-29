package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder_Select(t *testing.T) {
	t.Run("should default to 'SELECT *' when no fields are passed", func(t *testing.T) {
		qb := NewQueryBuilder("Users")
		query := qb.Execute()
		expected := "FROM Users SELECT *"
		assert.Equal(t, expected, query)
	})
	t.Run("should parse SELECT query correctly when only one field is passed", func(t *testing.T) {
		qb := NewQueryBuilder("Users").Select("email")
		query := qb.Execute()
		expected := "FROM Users SELECT email"
		assert.Equal(t, expected, query)
	})
	t.Run("should parse SELECT query correctly when multipled fields are passed", func(t *testing.T) {
		qb := NewQueryBuilder("Users").Select("email", "first_name", "last_name")
		query := qb.Execute()
		expected := "FROM Users SELECT email, first_name, last_name"
		assert.Equal(t, expected, query)
	})
}

func TestBuilder_Where(t *testing.T) {
	t.Run("should declare WHERE clause when a condition is passed", func(t *testing.T) {
		qb := NewQueryBuilder("Users").Where("age > 18")
		query := qb.Execute()
		expected := "FROM Users SELECT * WHERE age > 18"
		assert.Equal(t, expected, query)
	})
	t.Run("should concat WHERE clauses correctly when multiple conditions are passed", func(t *testing.T) {
		qb := NewQueryBuilder("Users").Where("age >= 18").Where("role = 'user'")
		query := qb.Execute()
		expected := "FROM Users SELECT * WHERE age >= 18 AND role = 'user'"
		assert.Equal(t, expected, query)
	})
}

func TestBuilder_OrderBy(t *testing.T) {
	t.Run("should declare ORDER BY clause when called", func(t *testing.T) {
		qb := NewQueryBuilder("Users").OrderBy("email", OrderByAsc)
		query := qb.Execute()
		expected := "FROM Users SELECT * ORDER BY email asc"
		assert.Equal(t, expected, query)
	})
	t.Run("should default to desc when no direction is passed", func(t *testing.T) {
		qb := NewQueryBuilder("Users").OrderBy("email")
		query := qb.Execute()
		expected := "FROM Users SELECT * ORDER BY email desc"
		assert.Equal(t, expected, query)
	})
}
func TestBuilder_Limit(t *testing.T) {
	t.Run("should declare LIMIT clause when declared", func(t *testing.T) {
		qb := NewQueryBuilder("Users").Limit(10)
		query := qb.Execute()
		expected := "FROM Users SELECT * LIMIT 10"
		assert.Equal(t, expected, query)
	})

}
func TestBuilder_E2E(t *testing.T) {
	type testCase struct {
		title    string
		input    *QueryBuilder
		expected string
	}

	testCases := []testCase{
		{
			title:    "E2E Test case 1",
			input:    NewQueryBuilder("Users").Select("email", "name").OrderBy("email", OrderByAsc).Where("age > 21"),
			expected: "FROM Users SELECT email, name WHERE age > 21 ORDER BY email asc",
		},
		{
			title:    "E2E Test case 2",
			input:    NewQueryBuilder("Users").Where("age > 21").Where("role = 'user'").OrderBy("name"),
			expected: "FROM Users SELECT * WHERE age > 21 AND role = 'user' ORDER BY name desc",
		},
		{
			title:    "E2E Test case 3",
			input:    NewQueryBuilder("Users").Select("email", "name", "createdAt").Where("age > 21").Where("role = 'admin'").OrderBy("name").Limit(200),
			expected: "FROM Users SELECT email, name, createdAt WHERE age > 21 AND role = 'admin' ORDER BY name desc LIMIT 200",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.title, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.Execute())
		})
	}
}
