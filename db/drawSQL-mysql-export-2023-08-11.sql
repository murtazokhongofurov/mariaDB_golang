CREATE TABLE `employees`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `branch_id` BIGINT NOT NULL,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(255) NOT NULL,
    `salary` DECIMAL(8, 2) NOT NULL
);
CREATE TABLE `course`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `branch_id` BIGINT NOT NULL
);
CREATE TABLE `studens`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(255) NOT NULL,
    `branch_id` BIGINT NOT NULL,
    `course_id` BIGINT NOT NULL
);
CREATE TABLE `branches`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `addresses` VARCHAR(255) NOT NULL
);
CREATE TABLE `employee_branch`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `employee_id` BIGINT NOT NULL,
    `branch_id` BIGINT NOT NULL
);
ALTER TABLE
    `studens` ADD CONSTRAINT `studens_course_id_foreign` FOREIGN KEY(`course_id`) REFERENCES `course`(`id`);
ALTER TABLE
    `course` ADD CONSTRAINT `course_branch_id_foreign` FOREIGN KEY(`branch_id`) REFERENCES `branches`(`id`);
ALTER TABLE
    `studens` ADD CONSTRAINT `studens_branch_id_foreign` FOREIGN KEY(`branch_id`) REFERENCES `branches`(`id`);
ALTER TABLE
    `employee_branch` ADD CONSTRAINT `employee_branch_branch_id_foreign` FOREIGN KEY(`branch_id`) REFERENCES `branches`(`id`);
ALTER TABLE
    `employee_branch` ADD CONSTRAINT `employee_branch_employee_id_foreign` FOREIGN KEY(`employee_id`) REFERENCES `employees`(`id`);