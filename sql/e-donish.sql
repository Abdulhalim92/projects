-- Таблица для студентов
CREATE TABLE "students" (
                            "student_id" serial PRIMARY KEY,
                            "name" varchar(255) NOT NULL,
                            "class" varchar(10) NOT NULL,
                            "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
                            "updated_at" timestamp
);

-- Таблица для учителей
CREATE TABLE "teachers" (
                            "teacher_id" serial PRIMARY KEY,
                            "name" varchar(255),
                            "user_id" integer NOT NULL,  -- связь с таблицей пользователей
                            "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
                            "updated_at" timestamp,
                            FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") -- Внешний ключ на пользователей
);

-- Таблица для записей посещаемости и оценок студентов
CREATE TABLE "student_daily_records" (
                                         "record_id" serial PRIMARY KEY,
                                         "student_id" integer NOT NULL,
                                         "subject_id" integer NOT NULL,
                                         "attendance" bool NOT NULL DEFAULT false,
                                         "grade" float,
                                         "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
                                         "updated_at" timestamp,
                                         FOREIGN KEY ("student_id") REFERENCES "students" ("student_id"),
                                         FOREIGN KEY ("subject_id") REFERENCES "subjects" ("subject_id")
);

-- Таблица для предметов
CREATE TABLE "subjects" (
                            "subject_id" serial PRIMARY KEY,
                            "name" varchar(255) NOT NULL,
                            "teacher_id" integer NOT NULL,
                            "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
                            "updated_at" timestamp,
                            FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("teacher_id")
);

-- Таблица для ролей пользователей
CREATE TABLE "roles" (
                         "role_id" serial PRIMARY KEY,
                         "name" varchar(255) NOT NULL,
                         "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
                         "updated_at" timestamp
);

-- Таблица для пользователей
CREATE TABLE "users" (
                         "user_id" serial PRIMARY KEY,
                         "username" varchar(255) NOT NULL,
                         "password" varchar(255) NOT NULL,
                         "role_id" integer NOT NULL,
                         "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
                         "updated_at" timestamp,
                         FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id") -- Внешний ключ на роли
);

-- Таблица для родителей
CREATE TABLE "parents" (
                           "parent_id" serial PRIMARY KEY,
                           "name" varchar(255),
                           "user_id" integer NOT NULL,   -- связь с таблицей пользователей
                           "student_id" integer NOT NULL, -- связь с таблицей студентов
                           "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
                           "updated_at" timestamp,
                           FOREIGN KEY ("user_id") REFERENCES "users" ("user_id"),   -- Внешний ключ на пользователей
                           FOREIGN KEY ("student_id") REFERENCES "students" ("student_id") -- Внешний ключ на студентов
);

-- Таблица для администрации (директора, завуча и т.д.)
CREATE TABLE "admins" (
                          "admin_id" serial PRIMARY KEY,
                          "name" varchar(255),
                          "user_id" integer NOT NULL,   -- связь с таблицей пользователей
                          "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
                          "updated_at" timestamp,
                          FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") -- Внешний ключ на пользователей
);
