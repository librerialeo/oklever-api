CREATE OR REPLACE FUNCTION update_modified_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.modified_at = NOW(); 
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE roles (
	rol_id SERIAL PRIMARY KEY,
	rol_name VARCHAR(16) UNIQUE NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_roles_modified_at BEFORE UPDATE
ON roles FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE countries (
	country_id SERIAL PRIMARY KEY,
	country_code VARCHAR(3) UNIQUE NOT NULL,
	country_name VARCHAR(32) NOT NULL,
	country_fullname VARCHAR(128) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_countries_modified_at BEFORE UPDATE
ON countries FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE degrees (
	degree_id SERIAL PRIMARY KEY,
	degree_name VARCHAR(64) UNIQUE NOT NULL,
	degree_description TEXT,
	degree_class VARCHAR(20),
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_degrees_modified_at BEFORE UPDATE
ON degrees FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE insignias (
	insignia_id SERIAL PRIMARY KEY,
	insignia_name VARCHAR(64) UNIQUE NOT NULL,
	insignia_description TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_insignias_modified_at BEFORE UPDATE
ON insignias FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE chats (
	chat_id SERIAL PRIMARY KEY,
	log_data TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_chats_modified_at BEFORE UPDATE
ON chats FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users (
	user_id SERIAL PRIMARY KEY,
	user_email VARCHAR(64) UNIQUE NOT NULL,
	user_password VARCHAR(128) NOT NULL,
	user_firstname VARCHAR(32) NOT NULL,
	user_lastname VARCHAR(32) NOT NULL,
	user_gender VARCHAR(16),
	user_image VARCHAR(64),
	user_birthdate DATE,
	user_phone VARCHAR(24),
	user_license VARCHAR(32) UNIQUE,
	user_rfc VARCHAR(16) UNIQUE,
	user_biography TEXT,
	user_months SMALLINT,
	user_status VARCHAR(32) NOT NULL DEFAULT "new",
	country_id INT REFERENCES countries(country_id),
	rol_id INT NOT NULL REFERENCES roles(rol_id),
	user_lastaction TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_modified_at BEFORE UPDATE
ON users FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_degrees (
	user_degree_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(user_id),
	degree_id INT NOT NULL REFERENCES degrees(degree_id),
	user_degree_name VARCHAR(64) NOT NULL,
	user_degree_institution VARCHAR(128) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_degrees_modified_at BEFORE UPDATE
ON users_degrees FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_insignias (
	user_insignia_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(user_id),
	insignia_id INT NOT NULL REFERENCES insignias(insignia_id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_insignias_modified_at BEFORE UPDATE
ON users_insignias FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_chats (
	chat_user_id SERIAL PRIMARY KEY,
	chat_id INT NOT NULL REFERENCES chats(chat_id),
	user_id INT NOT NULL REFERENCES users(user_id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_chats_modified_at BEFORE UPDATE
ON users_chats FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE languages (
	language_id SERIAL PRIMARY KEY,
	language_name VARCHAR(32) NOT NULL,
	language_iso VARCHAR(3) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_languages_modified_at BEFORE UPDATE
ON languages FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_languages (
	user_language_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(user_id),
	language_id INT NOT NULL REFERENCES languages(language_id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_languages_modified_at BEFORE UPDATE
ON users_languages FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_research (
	user_research_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(user_id),
	user_research_type VARCHAR(32) NOT NULL,
	user_research_reference TEXT NOT NULL,
	user_research_year INT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_research_modified_at BEFORE UPDATE
ON users_research FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_managements (
	user_management_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(user_id),
	user_management_job VARCHAR(128) NOT NULL,
	user_management_institution VARCHAR(128) NOT NULL,
	user_management_months SMALLINT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_managements_modified_at BEFORE UPDATE
ON users_managements FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_expertises (
	user_expertise_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(user_id),
	user_expertise_name VARCHAR(128) NOT NULL,
	user_expertise_months SMALLINT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_expertises_modified_at BEFORE UPDATE
ON users_expertises FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_institutions (
	user_institution_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(user_id),
	user_institution_name VARCHAR(128) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_institutions_modified_at BEFORE UPDATE
ON users_institutions FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_signatures (
	user_signature_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(user_id),
	degree_id INT NOT NULL REFERENCES degrees(degree_id),
	user_signature_name VARCHAR(128) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_signatures_modified_at BEFORE UPDATE
ON users_signatures FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE courses_levels (
	course_level_id SERIAL PRIMARY KEY,
	course_level_name VARCHAR(32) NOT NULL,
	course_level_description TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_courses_levels_modified_at BEFORE UPDATE
ON courses_levels FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE resources_types (
	resource_type_id SERIAL PRIMARY KEY,
	resource_type_name VARCHAR(32) UNIQUE NOT NULL,
	resource_type_description TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_resources_types_modified_at BEFORE UPDATE
ON resources_types FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE quizzes (
	quiz_id SERIAL PRIMARY KEY,
	quiz_attemps SMALLINT NOT NULL,
	quiz_approval SMALLINT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_quizzes_modified_at BEFORE UPDATE
ON quizzes FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE questions_types (
	question_type_id SERIAL PRIMARY KEY,
	question_type_name VARCHAR(32) NOT NULL,
	question_type_description TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_questions_types_modified_at BEFORE UPDATE
ON questions_types FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE questions (
	question_id SERIAL PRIMARY KEY,
	quiz_id INT NOT NULL REFERENCES quizzes(quiz_id),
	question_type_id INT  NOT NULL REFERENCES questions_types(question_type_id),
	question TEXT NOT NULL,
	question_resource TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_questions_modified_at BEFORE UPDATE
ON questions FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE questions_options (
	question_option_id SERIAL PRIMARY KEY,
	question_id INT NOT NULL REFERENCES questions(question_id),
	question_option TEXT NOT NULL,
	question_option_resource TEXT,
	question_option_correct BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_questions_options_modified_at BEFORE UPDATE
ON questions_options FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE courses (
	course_id SERIAL PRIMARY KEY,
	course_level_id INT NOT NULL REFERENCES courses_levels(course_level_id),
	teacher_id INT NOT NULL REFERENCES users(user_id),
	course_name VARCHAR(128) NOT NULL,
	course_description TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_courses_modified_at BEFORE UPDATE
ON courses FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE synchronous_classes (
	synchronous_class_id SERIAL PRIMARY KEY,
	course_id INT NOT NULL REFERENCES courses(course_id),
	synchronous_class_name VARCHAR(64) NOT NULL,
	synchronous_class_description TEXT NOT NULL,
	synchronous_class_scheduled TIMESTAMPTZ NOT NULL,
	synchronous_class_duration INT NOT NULL,
	synchronous_class_finished BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_synchronous_classes_modified_at BEFORE UPDATE
ON synchronous_classes FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE synchronous_classes_resources (
	synchronous_class_resource_id SERIAL PRIMARY KEY,
	synchronous_class_id INT NOT NULL REFERENCES synchronous_classes(synchronous_class_id),
	resource_type_id INT NOT NULL REFERENCES resources_types(resource_type_id),
	synchronous_class_resource_data TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_synchronous_classes_resources_modified_at BEFORE UPDATE
ON synchronous_classes_resources FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE courses_projects (
	course_project_id SERIAL PRIMARY KEY,
	course_id INT NOT NULL REFERENCES courses(course_id),
	course_project_name VARCHAR(128) NOT NULL,
	course_project_description TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_courses_projects_modified_at BEFORE UPDATE
ON courses_projects FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE courses_reviews (
	course_review_id SERIAL PRIMARY KEY,
	course_id INT NOT NULL REFERENCES courses(course_id),
	user_id INT NOT NULL REFERENCES users(user_id),
	rol_id INT NOT NULL REFERENCES roles(rol_id),
	course_review_rating SMALLINT NOT NULL,
	course_review_message TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_courses_reviews_modified_at BEFORE UPDATE
ON courses_reviews FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE courses_datasheets (
	course_datasheet_id SERIAL PRIMARY KEY,
	course_id INT NOT NULL REFERENCES courses(course_id),
	academic_id INT NOT NULL REFERENCES users(user_id),
	course_datasheet_sentence VARCHAR(128) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_courses_datasheets_modified_at BEFORE UPDATE
ON courses_datasheets FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE courses_forums (
	course_forum_id SERIAL PRIMARY KEY,
	course_id INT NOT NULL REFERENCES courses(course_id),
	course_forum_topic VARCHAR(128) NOT NULL,
	course_forum_description TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_courses_forums_modified_at BEFORE UPDATE
ON courses_forums FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE courses_forums_comments (
	course_forum_comment_id SERIAL PRIMARY KEY,
	course_forum_id INT NOT NULL REFERENCES courses_forums(course_forum_id),
	user_id INT NOT NULL REFERENCES users(user_id),
	rol_id INT NOT NULL REFERENCES roles(rol_id),
	course_forum_comment_message TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_courses_forums_comments_modified_at BEFORE UPDATE
ON courses_forums_comments FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE modules (
	module_id SERIAL PRIMARY KEY,
	course_id INT NOT NULL REFERENCES courses(course_id),
	module_sort SMALLINT NOT NULL,
	module_name VARCHAR(64) NOT NULL,
	module_description TEXT NOT NULL,
	quiz_id INT NOT NULL REFERENCES quizzes(quiz_id),
	exercise_id INT NOT NULL REFERENCES quizzes(quiz_id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_modules_modified_at BEFORE UPDATE
ON modules FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE modules_feedback (
	class_feedback_id SERIAL PRIMARY KEY,
	module_id INT NOT NULL REFERENCES modules(module_id),
	academic_id INT NOT NULL REFERENCES users(user_id),
	class_feedback_message TEXT NOT NULL,
	class_feedback_corrected BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_modules_feedback_modified_at BEFORE UPDATE
ON modules_feedback FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE modules_resources (
	module_resource_id SERIAL PRIMARY KEY,
	module_id INT NOT NULL REFERENCES modules(module_id),
	resource_type_id INT NOT NULL REFERENCES resources_types(resource_type_id),
	module_resource_data TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_modules_resources_modified_at BEFORE UPDATE
ON modules_resources FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE topics (
	topic_id SERIAL PRIMARY KEY,
	module_id INT NOT NULL REFERENCES modules(module_id),
	topic_sort SMALLINT NOT NULL,
	topic_name VARCHAR(128) NOT NULL,
	topic_description TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_topics_modified_at BEFORE UPDATE
ON topics FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE topics_resources (
	topic_resource_id SERIAL PRIMARY KEY,
	topic_id INT NOT NULL REFERENCES topics(topic_id),
	resource_type_id INT NOT NULL REFERENCES resources_types(resource_type_id),
	topic_resource_data TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_topics_resources_modified_at BEFORE UPDATE
ON topics_resources FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE classes (
	class_id SERIAL PRIMARY KEY,
	topic_id INT NOT NULL REFERENCES topics(topic_id),
	class_sort SMALLINT NOT NULL,
	class_name VARCHAR(64) NOT NULL,
	class_description TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_classes_modified_at BEFORE UPDATE
ON classes FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE classes_resources (
	class_resource_id SERIAL PRIMARY KEY,
	class_id INT NOT NULL REFERENCES classes(class_id),
	resource_type_id INT NOT NULL REFERENCES resources_types(resource_type_id),
	class_resource_data TEXT NOT NULL,
	class_resource_sort SMALLINT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_classes_resources_modified_at BEFORE UPDATE
ON classes_resources FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE classes_comments (
	class_comment_id SERIAL PRIMARY KEY,
	class_id INT NOT NULL REFERENCES classes(class_id),
	user_id INT NOT NULL REFERENCES users(user_id),
	rol_id INT NOT NULL REFERENCES roles(rol_id),
	class_comment_message TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_classes_comments_modified_at BEFORE UPDATE
ON classes_comments FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE test_classes (
	test_class_id SERIAL PRIMARY KEY,
	teacher_id INT NOT NULL REFERENCES users(user_id),
	academic_id INT NOT NULL REFERENCES users(user_id),
	test_class_type INT NOT NULL,
	test_class_name VARCHAR(64) NOT NULL,
	test_class_description VARCHAR(64) NOT NULL,
	test_class_data TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_test_classes_modified_at BEFORE UPDATE
ON test_classes FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE test_classes_feedback (
	test_class_feedback_id SERIAL PRIMARY KEY,
	test_class_id INT NOT NULL REFERENCES test_classes(test_class_id),
	academic_id INT NOT NULL REFERENCES users(user_id),
	test_class_feedback_message TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_test_classes_feedback_modified_at BEFORE UPDATE
ON test_classes_feedback FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE subscriptions (
	subscription_id SERIAL PRIMARY KEY,
	subscription_name VARCHAR(64) UNIQUE NOT NULL,
	subscription_price DECIMAL(8,2) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_subscriptions_modified_at BEFORE UPDATE
ON subscriptions FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE purchases (
	purchase_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(user_id),
	purchase_total DECIMAL(8,2) NOT NULL,
	purchase_discount DECIMAL(8,2) NOT NULL,
	purchase_currency DECIMAL(8,2) NOT NULL,
	purchase_currency_value DECIMAL(8,2) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_purchases_modified_at BEFORE UPDATE
ON purchases FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE purchases_products (
	purchase_product_id SERIAL PRIMARY KEY,
	purchase_id INT NOT NULL REFERENCES purchases(purchase_id),
	purchase_product_name VARCHAR(64),
	purchase_product_price DECIMAL(8,2) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_purchases_products_modified_at BEFORE UPDATE
ON purchases_products FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_courses (
	user_course_id SERIAL PRIMARY KEY,
	purchase_id INT NOT NULL REFERENCES purchases(purchase_id),
	course_id INT NOT NULL REFERENCES courses(course_id),
	user_id INT NOT NULL REFERENCES users(user_id),
	user_course_approved BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_courses_modified_at BEFORE UPDATE
ON users_courses FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_quizzes (
	user_quiz_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(user_id),
	quiz_id INT NOT NULL REFERENCES quizzes(quiz_id),
	user_quiz_aproved BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_quizzes_modified_at BEFORE UPDATE
ON users_quizzes FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_answers (
	user_answer_id SERIAL PRIMARY KEY,
	user_quiz_id INT NOT NULL REFERENCES users_quizzes(user_quiz_id),
	user_id INT NOT NULL REFERENCES users(user_id),
	question_option_id INT NOT NULL REFERENCES questions_options(question_option_id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_answers_modified_at BEFORE UPDATE
ON users_answers FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_projects (
	user_project_id SERIAL PRIMARY KEY,
	course_project_id INT NOT NULL REFERENCES courses_projects(course_project_id),
	user_id INT NOT NULL REFERENCES users(user_id),
	user_project_type VARCHAR(64) NOT NULL,
	user_project_data TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_projects_modified_at BEFORE UPDATE
ON users_projects FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_subscriptions (
	users_subscriptions_id SERIAL PRIMARY KEY,
	purchase_id INT NOT NULL REFERENCES purchases(purchase_id),
	subscription_id INT NOT NULL REFERENCES subscriptions(subscription_id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);
CREATE TRIGGER update_users_subscriptions_modified_at BEFORE UPDATE
ON users_subscriptions FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_projects_history (
	user_project_history_id SERIAL PRIMARY KEY,
	user_project_id INT NOT NULL REFERENCES users_projects(user_project_id),
	user_project_type VARCHAR(64) NOT NULL,
	user_project_data TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);

CREATE TRIGGER update_users_projects_history_modified_at BEFORE UPDATE
ON users_projects_history FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();

CREATE TABLE users_academy (
	user_academy_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(user_id),
	degree_id INT NOT NULL REFERENCES degrees(degree_id),
	user_academy_name VARCHAR(64) NOT NULL,
	user_academy_institution VARCHAR(64) NOT NULL,
	user_academy_year SMALLINT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ
);

CREATE TRIGGER update_users_academy_modified_at BEFORE UPDATE
ON users_academy FOR EACH ROW EXECUTE PROCEDURE update_modified_at_column();