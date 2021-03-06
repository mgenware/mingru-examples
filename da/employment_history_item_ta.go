/******************************************************************************************
 * This file was automatically generated by mingru (https://github.com/mgenware/mingru)
 * Do not edit this file manually, your changes will be overwritten.
 ******************************************************************************************/

package da

import (
	"fmt"
	"time"

	"github.com/mgenware/mingru-go-lib"
)

// TableTypeEmploymentHistoryItem ...
type TableTypeEmploymentHistoryItem struct {
}

// EmploymentHistoryItem ...
var EmploymentHistoryItem = &TableTypeEmploymentHistoryItem{}

// ------------ Actions ------------

// EmploymentHistoryItemTableSelectAllHistoryResult ...
type EmploymentHistoryItemTableSelectAllHistoryResult struct {
	DepartmentName    string
	EmployeeFirstName string
	EmployeeLastName  string
	FromDate          time.Time
	ToDate            time.Time
}

// SelectAllHistory ...
func (da *TableTypeEmploymentHistoryItem) SelectAllHistory(queryable mingru.Queryable, page int, pageSize int) ([]*EmploymentHistoryItemTableSelectAllHistoryResult, bool, error) {
	if page <= 0 {
		err := fmt.Errorf("Invalid page %v", page)
		return nil, false, err
	}
	if pageSize <= 0 {
		err := fmt.Errorf("Invalid page size %v", pageSize)
		return nil, false, err
	}
	limit := pageSize + 1
	offset := (page - 1) * pageSize
	max := pageSize
	rows, err := queryable.Query("SELECT `dept_emp`.`from_date` AS `from_date`, `dept_emp`.`to_date` AS `to_date`, `join_1`.`first_name` AS `employee_first_name`, `join_1`.`last_name` AS `employee_last_name`, `join_2`.`dept_name` AS `department_name` FROM `dept_emp` AS `dept_emp` INNER JOIN `employees` AS `join_1` ON `join_1`.`emp_no` = `dept_emp`.`emp_no` INNER JOIN `departments` AS `join_2` ON `join_2`.`dept_no` = `dept_emp`.`dept_no` ORDER BY `dept_emp`.`emp_no`, `dept_emp`.`dept_no` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, false, err
	}
	result := make([]*EmploymentHistoryItemTableSelectAllHistoryResult, 0, limit)
	itemCounter := 0
	defer rows.Close()
	for rows.Next() {
		itemCounter++
		if itemCounter <= max {
			item := &EmploymentHistoryItemTableSelectAllHistoryResult{}
			err = rows.Scan(&item.FromDate, &item.ToDate, &item.EmployeeFirstName, &item.EmployeeLastName, &item.DepartmentName)
			if err != nil {
				return nil, false, err
			}
			result = append(result, item)
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, false, err
	}
	return result, itemCounter > len(result), nil
}
