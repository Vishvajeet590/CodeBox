package db

// type for paginated result
type PaginatedResult[T any] struct {
	Results []T
	Page    int
	Limit   int
}

// type for updating specific columns
type UpdateColumnsValues = map[string]interface{}
