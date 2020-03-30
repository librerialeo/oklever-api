CREATE TABLE `roles` (
	`rol_id` INT AUTO_INCREMENT,
	`rol_name` VARCHAR(16) UNIQUE NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`rol_id`)
);

CREATE TABLE `countries` (
	`country_id` INT AUTO_INCREMENT,
	`country_code` VARCHAR(3) UNIQUE NOT NULL,
	`country_name` VARCHAR(32) NOT NULL,
	`country_fullname` VARCHAR(128) NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`country_id`)
);

CREATE TABLE `degrees` (
	`degree_id` INT AUTO_INCREMENT,
	`degree_name` VARCHAR(64) NOT NULL,
	`degree_description` TEXT,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`degree_id`)
);

CREATE TABLE `insignias` (
	`insignia_id` INT AUTO_INCREMENT,
	`insignia_name` VARCHAR(64) NOT NULL,
	`insignia_description` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`insignia_id`)
);

CREATE TABLE `chats` (
	`chat_id` INT AUTO_INCREMENT,
	`log_data` LONGTEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`chat_id`)
);

CREATE TABLE `users` (
	`user_id` INT AUTO_INCREMENT,
	`user_email` VARCHAR(64) UNIQUE NOT NULL,
	`user_password` VARCHAR(128) NOT NULL,
	`user_firstname` VARCHAR(32) NOT NULL,
	`user_lastname` VARCHAR(32) NOT NULL,
	`user_gender` VARCHAR(16),
	`user_image` VARCHAR(64),
	`user_birthdate` DATE,
	`user_phone` VARCHAR(24),
	`country_id` INT,
	`rol_id` INT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`user_id`),
	FOREIGN KEY (`country_id`) REFERENCES `countries`(`country_id`),
	FOREIGN KEY (`rol_id`) REFERENCES `roles`(`rol_id`)
);

CREATE TABLE `users_degrees` (
	`user_degree_id` INT AUTO_INCREMENT,
	`user_id` INT NOT NULL,
	`degree_id` INT NOT NULL,
	`user_degree_name` VARCHAR(64) NOT NULL,
	`user_degree_institution` VARCHAR(128) NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`user_degree_id`),
	FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`),
	FOREIGN KEY (`degree_id`) REFERENCES `degrees`(`degree_id`)
);

CREATE TABLE `users_insignias` (
	`user_insignia_id` INT AUTO_INCREMENT,
	`user_id` INT NOT NULL,
	`insignia_id` INT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`user_insignia_id`),
	FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`),
	FOREIGN KEY (`insignia_id`) REFERENCES `insignias`(`insignia_id`)
);

CREATE TABLE `users_chats` (
	`chat_user_id` INT AUTO_INCREMENT,
	`chat_id` INT NOT NULL,
	`user_id` INT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`chat_user_id`),
	FOREIGN KEY (`chat_id`) REFERENCES `chats`(`chat_id`),
	FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`)
);

CREATE TABLE `languages` (
	`language_id` INT AUTO_INCREMENT,
	`language_name` VARCHAR(32) NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`language_id`)
);

CREATE TABLE `teachers` (
	`teacher_id` INT AUTO_INCREMENT,
	`user_id` INT NOT NULL,
	`teacher_professional_license` VARCHAR(32) UNIQUE NOT NULL,
	`teacher_rfc` VARCHAR(16) UNIQUE NOT NULL,
	`teacher_biography` TEXT NOT NULL,
	`teacher_accepted` BOOLEAN NOT NULL DEFAULT 0,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`teacher_id`),
	FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`)
);

CREATE TABLE `teachers_languages` (
	`teacher_language_id` INT AUTO_INCREMENT,
	`teacher_id` INT NOT NULL,
	`language_id` INT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`teacher_language_id`),
	FOREIGN KEY (`teacher_id`) REFERENCES `teachers`(`teacher_id`),
	FOREIGN KEY (`language_id`) REFERENCES `languages`(`language_id`)
);

CREATE TABLE `teachers_research` (
	`teacher_research_id` INT AUTO_INCREMENT,
	`teacher_id` INT NOT NULL,
	`teacher_research_reference` TEXT NOT NULL,
	`teacher_research_year` YEAR NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`teacher_research_id`),
	FOREIGN KEY (`teacher_id`) REFERENCES `teachers`(`teacher_id`)
);

CREATE TABLE `teachers_managements` (
	`teacher_management_id` INT AUTO_INCREMENT,
	`teacher_id` INT NOT NULL,
	`teacher_management_job` VARCHAR(128) NOT NULL,
	`teacher_management_institution` VARCHAR(128) NOT NULL,
	`teacher_management_months` SMALLINT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`teacher_management_id`),
	FOREIGN KEY (`teacher_id`) REFERENCES `teachers`(`teacher_id`)
);

CREATE TABLE `teachers_expertises` (
	`teacher_expertise_id` INT AUTO_INCREMENT,
	`teacher_id` INT NOT NULL,
	`teacher_expertise_name` VARCHAR(128) NOT NULL,
	`teacher_expertise_months` SMALLINT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`teacher_expertise_id`),
	FOREIGN KEY (`teacher_id`) REFERENCES `teachers`(`teacher_id`)
);

CREATE TABLE `teachers_teaching` (
	`teacher_teaching_id` INT AUTO_INCREMENT,
	`teacher_id` INT NOT NULL,
	`degree_id` INT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`teacher_teaching_id`),
	FOREIGN KEY (`teacher_id`) REFERENCES `teachers`(`teacher_id`),
	FOREIGN KEY (`degree_id`) REFERENCES `degrees`(`degree_id`)
);

CREATE TABLE `teachers_teaching_institutions` (
	`teacher_teaching_institution_id` INT AUTO_INCREMENT,
	`teacher_teaching_id` INT NOT NULL,
	`teacher_teaching_institution_name` VARCHAR(128) NOT NULL,
	`teacher_degree_institution_months` SMALLINT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`teacher_teaching_institution_id`),
	FOREIGN KEY (`teacher_teaching_id`) REFERENCES `teachers_teaching`(`teacher_teaching_id`)
);

CREATE TABLE `teachers_teaching_signatures` (
	`teacher_teaching_signature_id` INT AUTO_INCREMENT,
	`teacher_teaching_id` INT NOT NULL,
	`teacher_teaching_signature_name` VARCHAR(128) NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`teacher_teaching_signature_id`),
	FOREIGN KEY (`teacher_teaching_id`) REFERENCES `teachers_teaching`(`teacher_teaching_id`)
);

CREATE TABLE `courses_levels` (
	`course_level_id` INT AUTO_INCREMENT,
	`course_level_name` VARCHAR(32) NOT NULL,
	`course_level_description` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`course_level_id`)
);

CREATE TABLE `resources_types` (
	`resource_type_id` INT AUTO_INCREMENT,
	`resource_type_name` VARCHAR(32) UNIQUE NOT NULL,
	`resource_type_description` TEXT,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`resource_type_id`)
);

CREATE TABLE `quizzes` (
	`quiz_id` INT AUTO_INCREMENT,
	`quiz_attemps` TINYINT NOT NULL,
	`quiz_approval` TINYINT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`quiz_id`)
);

CREATE TABLE `questions_types` (
	`question_type_id` INT AUTO_INCREMENT,
	`question_type_name` VARCHAR(32) NOT NULL,
	`question_type_description` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`question_type_id`)
);

CREATE TABLE `questions` (
	`question_id` INT AUTO_INCREMENT,
	`quiz_id` INT NOT NULL,
	`question_type_id` INT  NOT NULL,
	`question` TEXT NOT NULL,
	`question_resource` TEXT,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`question_id`),
	FOREIGN KEY (`quiz_id`) REFERENCES `quizzes`(`quiz_id`)
);

CREATE TABLE `questions_options` (
	`question_option_id` INT AUTO_INCREMENT,
	`question_id` INT NOT NULL,
	`question_option` TEXT NOT NULL,
	`question_option_resource` TEXT,
	`question_option_correct` BOOLEAN NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`question_option_id`),
	FOREIGN KEY (`question_id`) REFERENCES `questions`(`question_id`)
);

CREATE TABLE `students` (
	`student_id` INT AUTO_INCREMENT,
	`user_id` INT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`student_id`),
	FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`)
);

CREATE TABLE `courses` (
	`course_id` INT AUTO_INCREMENT,
	`course_level_id` INT NOT NULL,
	`teacher_id` INT NOT NULL,
	`course_name` VARCHAR(128) NOT NULL,
	`course_description` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`course_id`),
	FOREIGN KEY (`course_level_id`) REFERENCES `courses_levels`(`course_level_id`),
	FOREIGN KEY (`teacher_id`) REFERENCES `teachers`(`teacher_id`)
);

CREATE TABLE `synchronous_classes` (
	`synchronous_class_id` INT AUTO_INCREMENT,
	`course_id` INT NOT NULL,
	`synchronous_class_name` VARCHAR(64) NOT NULL,
	`synchronous_class_description` TEXT NOT NULL,
	`synchronous_class_scheduled` DATETIME NOT NULL,
	`synchronous_class_duration` INT NOT NULL,
	`synchronous_class_finished` BOOLEAN NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`synchronous_class_id`),
	FOREIGN KEY (`course_id`) REFERENCES `courses`(`course_id`)
);

CREATE TABLE `synchronous_classes_resources` (
	`synchronous_class_resource_id` INT AUTO_INCREMENT,
	`synchronous_class_id` INT NOT NULL,
	`resource_type_id` INT NOT NULL,
	`synchronous_class_resource_data` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`synchronous_class_resource_id`),
	FOREIGN KEY (`synchronous_class_id`) REFERENCES `synchronous_classes`(`synchronous_class_id`),
	FOREIGN KEY (`resource_type_id`) REFERENCES `resources_types`(`resource_type_id`)
);

CREATE TABLE `courses_projects` (
	`course_project_id` INT AUTO_INCREMENT,
	`course_id` INT NOT NULL,
	`course_project_name` VARCHAR(128) NOT NULL,
	`course_project_description` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`course_project_id`),
	FOREIGN KEY (`course_id`) REFERENCES `courses`(`course_id`)
);

CREATE TABLE `courses_reviews` (
	`course_review_id` INT AUTO_INCREMENT,
	`course_id` INT NOT NULL,
	`student_id` INT NOT NULL,
	`course_review_rating` TINYINT NOT NULL,
	`course_review_message` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`course_review_id`),
	FOREIGN KEY (`course_id`) REFERENCES `courses`(`course_id`),
	FOREIGN KEY (`student_id`) REFERENCES `students`(`student_id`)
);

CREATE TABLE `courses_datasheets` (
	`course_datasheet_id` INT AUTO_INCREMENT,
	`course_id` INT NOT NULL,
	`academic_id` INT NOT NULL,
	`course_datasheet_sentence` VARCHAR(128) NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`course_datasheet_id`),
	FOREIGN KEY (`course_id`) REFERENCES `courses`(`course_id`),
	FOREIGN KEY (`academic_id`) REFERENCES `users`(`user_id`)
);

CREATE TABLE `courses_forums` (
	`course_forum_id` INT AUTO_INCREMENT,
	`course_id` INT NOT NULL,
	`course_forum_topic` VARCHAR(128) NOT NULL,
	`course_forum_description` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`course_forum_id`),
	FOREIGN KEY (`course_id`) REFERENCES `courses`(`course_id`)
);

CREATE TABLE `courses_forums_comments` (
	`course_forum_comment_id` INT AUTO_INCREMENT,
	`course_forum_id` INT NOT NULL,
	`user_id` INT NOT NULL,
	`course_forum_comment_message` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`course_forum_comment_id`),
	FOREIGN KEY (`course_forum_id`) REFERENCES `courses_forums`(`course_forum_id`),
	FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`)
);

CREATE TABLE `modules` (
	`module_id` INT AUTO_INCREMENT,
	`course_id` INT NOT NULL,
	`module_sort` TINYINT NOT NULL,
	`module_name` VARCHAR(64) NOT NULL,
	`module_description` TEXT NOT NULL,
	`quiz_id` INT NOT NULL,
	`exercise_id` INT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`module_id`),
	FOREIGN KEY (`course_id`) REFERENCES `courses`(`course_id`),
	FOREIGN KEY (`quiz_id`) REFERENCES `quizzes`(`quiz_id`),
	FOREIGN KEY (`exercise_id`) REFERENCES `quizzes`(`quiz_id`)
);

CREATE TABLE `modules_feedback` (
	`class_feedback_id` INT AUTO_INCREMENT,
	`module_id` INT NOT NULL,
	`academic_id` INT NOT NULL,
	`class_feedback_message` TEXT NOT NULL,
	`class_feedback_corrected` BOOLEAN NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`class_feedback_id`),
	FOREIGN KEY (`module_id`) REFERENCES `modules`(`module_id`),
	FOREIGN KEY (`academic_id`) REFERENCES `users`(`user_id`)
);

CREATE TABLE `modules_resources` (
	`module_resource_id` INT AUTO_INCREMENT,
	`module_id` INT NOT NULL,
	`resource_type_id` INT NOT NULL,
	`module_resource_data` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`module_resource_id`),
	FOREIGN KEY (`module_id`) REFERENCES `modules`(`module_id`),
	FOREIGN KEY (`resource_type_id`) REFERENCES `resources_types`(`resource_type_id`)
);

CREATE TABLE `topics` (
	`topic_id` INT AUTO_INCREMENT,
	`module_id` INT NOT NULL,
	`topic_sort` TINYINT NOT NULL,
	`topic_name` VARCHAR(128) NOT NULL,
	`topic_description` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`topic_id`),
	FOREIGN KEY (`module_id`) REFERENCES `modules`(`module_id`)
);

CREATE TABLE `topics_resources` (
	`topic_resource_id` INT AUTO_INCREMENT,
	`topic_id` INT NOT NULL,
	`resource_type_id` INT NOT NULL,
	`topic_resource_data` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`topic_resource_id`),
	FOREIGN KEY (`topic_id`) REFERENCES `topics`(`topic_id`),
	FOREIGN KEY (`resource_type_id`) REFERENCES `resources_types`(`resource_type_id`)
);

CREATE TABLE `classes` (
	`class_id` INT AUTO_INCREMENT,
	`topic_id` INT NOT NULL,
	`class_sort` TINYINT NOT NULL,
	`class_name` VARCHAR(64) NOT NULL,
	`class_description` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`class_id`),
	FOREIGN KEY (`topic_id`) REFERENCES `topics`(`topic_id`)
);

CREATE TABLE `classes_resources` (
	`class_resource_id` INT AUTO_INCREMENT,
	`class_id` INT NOT NULL,
	`resource_type_id` INT NOT NULL,
	`class_resource_data` TEXT NOT NULL,
	`class_resource_sort` TINYINT,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`class_resource_id`),
	FOREIGN KEY (`class_id`) REFERENCES `classes`(`class_id`),
	FOREIGN KEY (`resource_type_id`) REFERENCES `resources_types`(`resource_type_id`)
);

CREATE TABLE `classes_comments` (
	`class_comment_id` INT AUTO_INCREMENT,
	`class_id` INT NOT NULL,
	`user_id` INT NOT NULL,
	`class_comment_message` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`class_comment_id`),
	FOREIGN KEY (`class_id`) REFERENCES `classes`(`class_id`),
	FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`)
);

CREATE TABLE `test_classes` (
	`test_class_id` INT AUTO_INCREMENT,
	`teacher_id` INT NOT NULL,
	`test_class_type` INT NOT NULL,
	`test_class_name` VARCHAR(64) NOT NULL,
	`test_class_description` VARCHAR(64) NOT NULL,
	`test_class_data` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`test_class_id`),
	FOREIGN KEY (`teacher_id`) REFERENCES `teachers`(`teacher_id`)
);

CREATE TABLE `test_classes_feedback` (
	`test_class_feedback_id` INT AUTO_INCREMENT,
	`test_class_id` INT NOT NULL,
	`academic_id` INT NOT NULL,
	`test_class_feedback_message` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`test_class_feedback_id`),
	FOREIGN KEY (`test_class_id`) REFERENCES `test_classes`(`test_class_id`),
	FOREIGN KEY (`academic_id`) REFERENCES `users`(`user_id`)
);

CREATE TABLE `subscriptions` (
	`subscription_id` INT AUTO_INCREMENT,
	`subscription_name` VARCHAR(64) UNIQUE NOT NULL,
	`subscription_price` DECIMAL(8,2) NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`subscription_id`)
);

CREATE TABLE `purchases` (
	`purchase_id` INT AUTO_INCREMENT,
	`student_id` INT NOT NULL,
	`purchase_total` DECIMAL(8,2) NOT NULL,
	`purchase_discount` DECIMAL(8,2) NOT NULL,
	`purchase_currency` DECIMAL(8,2) NOT NULL,
	`purchase_currency_value` DECIMAL(8,2) NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`purchase_id`),
	FOREIGN KEY (`student_id`) REFERENCES `students`(`student_id`)
);

CREATE TABLE `purchases_products` (
	`purchase_product_id` INT AUTO_INCREMENT,
	`purchase_id` INT NOT NULL,
	`purchase_product_name` VARCHAR(64),
	`purchase_product_price` DECIMAL(8,2) NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`purchase_product_id`),
	FOREIGN KEY (`purchase_id`) REFERENCES `purchases`(`purchase_id`)
);

CREATE TABLE `students_courses` (
	`student_course_id` INT AUTO_INCREMENT,
	`purchase_id` INT NOT NULL,
	`course_id` INT NOT NULL,
	`student_id` INT NOT NULL,
	`student_course_approved` BOOLEAN,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`student_course_id`),
	FOREIGN KEY (`purchase_id`) REFERENCES `purchases`(`purchase_id`),
	FOREIGN KEY (`course_id`) REFERENCES `courses`(`course_id`),
	FOREIGN KEY (`student_id`) REFERENCES `students`(`student_id`)
);

CREATE TABLE `students_quizzes` (
	`student_quiz_id` INT AUTO_INCREMENT,
	`student_id` INT NOT NULL,
	`quiz_id` INT NOT NULL,
	`student_quiz_aproved` BOOLEAN NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`student_quiz_id`),
	FOREIGN KEY (`student_id`) REFERENCES `students`(`student_id`),
	FOREIGN KEY (`quiz_id`) REFERENCES `quizzes`(`quiz_id`)
);

CREATE TABLE `students_answers` (
	`student_answer_id` INT AUTO_INCREMENT,
	`student_quiz_id` INT NOT NULL,
	`student_id` INT NOT NULL,
	`question_option_id` INT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`student_answer_id`),
	FOREIGN KEY (`student_quiz_id`) REFERENCES `students_quizzes`(`student_quiz_id`),
	FOREIGN KEY (`student_id`) REFERENCES `students`(`student_id`),
	FOREIGN KEY (`question_option_id`) REFERENCES `questions_options`(`question_option_id`)
);

CREATE TABLE `students_projects` (
	`student_project_id` INT AUTO_INCREMENT,
	`course_project_id` INT NOT NULL,
	`student_id` INT NOT NULL,
	`student_project_type` VARCHAR(64) NOT NULL,
	`student_project_data` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`student_project_id`),
	FOREIGN KEY (`course_project_id`) REFERENCES `courses_projects`(`course_project_id`),
	FOREIGN KEY (`student_id`) REFERENCES `students`(`student_id`)
);

CREATE TABLE `students_subscriptions` (
	`students_subscriptions_id` INT AUTO_INCREMENT,
	`purchase_id` INT NOT NULL,
	`subscription_id` INT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`students_subscriptions_id`),
	FOREIGN KEY (`purchase_id`) REFERENCES `purchases`(`purchase_id`),
	FOREIGN KEY (`subscription_id`) REFERENCES `subscriptions`(`subscription_id`)
);

CREATE TABLE `students_projects_history` (
	`student_project_history_id` INT AUTO_INCREMENT,
	`student_project_id` INT NOT NULL,
	`student_project_type` VARCHAR(64) NOT NULL,
	`student_project_data` TEXT NOT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME,
	PRIMARY KEY (`student_project_history_id`),
	FOREIGN KEY (`student_project_id`) REFERENCES `students_projects`(`student_project_id`)
);