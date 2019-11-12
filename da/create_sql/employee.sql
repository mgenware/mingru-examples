CREATE TABLE `employees` (
	`emp_no` INT NOT NULL,
	`first_name` VARCHAR(50) NOT NULL,
	`last_name` VARCHAR(50) NOT NULL,
	`gender` VARCHAR(10) NOT NULL,
	`birth_date` DATE NOT NULL,
	`hire_date` DATE NOT NULL,
	PRIMARY KEY (`emp_no`)
)
CHARACTER SET=utf8mb4
COLLATE=utf8mb4_unicode_ci
;