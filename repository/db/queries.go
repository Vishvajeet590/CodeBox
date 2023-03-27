package db

import "fmt"

func GetOne[T any](model T) (T, error) {
	var models T
	err := DB.Where(model).First(&models).Error
	return models, err
}

// Create a new record for the model
func CreateOne[T any](class T) (interface{}, error) {
	err := DB.Create(&class).Error
	return class, err
}

// Create multiple records for the model
func CreateBulk[T any](data []T) ([]T, error) {
	err := DB.Create(&data).Error
	return data, err
}

// Update a record for the model
func UpdateOne[T any](class T) (interface{}, error) {
	err := DB.Save(&class).Error
	return class, err
}

// Delete a record for the model
func DeleteOne[T any](class T) (interface{}, error) {
	err := DB.Delete(&class).Error
	return class, err
}

// Get all the records of the model with pagination
func GetAll[T any](model T, page int, limit int, order_by string, where_query string) (PaginatedResult[T], error) {
	offset := (page - 1) * limit
	var results []T
	// Get all the records of the model
	if order_by == "" && where_query == "" {
		err := DB.Where(model).Offset(offset).Limit(limit).Find(&results).Error
		result := PaginatedResult[T]{
			Results: results,
			Page:    page,
			Limit:   limit,
		}
		return result, err
	}

	if where_query == "" {
		err := DB.Where(model).Offset(offset).Limit(limit).Order(order_by).Find(&results).Error
		result := PaginatedResult[T]{
			Results: results,
			Page:    page,
			Limit:   limit,
		}
		return result, err
	}

	if order_by == "" {
		if limit == 0 {
			err := DB.Where(where_query).Find(&results).Error
			result := PaginatedResult[T]{
				Results: results,
				Page:    page,
				Limit:   limit,
			}
			return result, err
		}
	}

	err := DB.Where(where_query).Order(order_by).Offset(offset).Limit(limit).Find(&results).Error
	result := PaginatedResult[T]{
		Results: results,
		Page:    page,
		Limit:   limit,
	}
	return result, err

}

func JoinGetAll[T any](tableName string, selection string, joins []string, page int, limit int, query string, model T) (PaginatedResult[T], error) {
	if tableName == "" || selection == "" || len(joins) == 0 || page < 1 || query == "" {
		panic("Missing Parameters")
	}
	switch len(joins) {
	default:
		return JoinTwoGetAll(tableName, selection, joins[0], page, limit, query, "", model)
	}
}

func JoinTwoGetAll[T any](tableName string, selection string, join1 string, page int, limit int, query string, order_by string, model T) (PaginatedResult[T], error) {
	offset := (page - 1) * limit
	var results []T
	// Get all the records of the model
	var err error
	if query == "" {
		err = DB.Table(tableName).Select(selection).Joins(join1).Offset(offset).Limit(limit).Order(order_by).Where(model).Find(&results).Error
	} else {
		err = DB.Table(tableName).Select(selection).Joins(join1).Where(query).Offset(offset).Limit(limit).Find(&results).Error
	}
	if err != nil {
		return PaginatedResult[T]{}, err
	}
	result := PaginatedResult[T]{
		Results: results,
		Page:    page,
		Limit:   limit,
	}
	return result, err
}

func JoinTwoGetOne[T any](tableName string, selection string, join1 string, model T) (T, error) {
	//var result T
	// Get all the records of the model
	var models T
	err := DB.Table(tableName).Select(selection).Joins(join1).Where(model).First(&models).Error
	return models, err
}

// Update specific columns for a model
func UpdateColumns[T any](class T, updateColumns UpdateColumnsValues) (interface{}, error) {
	return nil, DB.Model(class).Updates(updateColumns).Error
}

// find or Create if not exists
func FirstOrCreate[T any](class T) (int64, error) {
	result := DB.Where(class).FirstOrCreate(&class)
	return result.RowsAffected, result.Error
}

func CallPostgresGetAllFunction[T any](class T, query string, page int, limit int) (PaginatedResult[T], error) {
	var result []T
	err := DB.Raw(query).Scan(&result).Error
	results := PaginatedResult[T]{
		Results: result,
		Page:    page,
		Limit:   limit,
	}
	fmt.Println(results, "result")
	return results, err
}
