-- Создание таблиц

-- Таблица users
-- Описание: Таблица пользователей для хранения учетных данных.
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(100),
                       password VARCHAR(100)
);

INSERT INTO users (username, password)
VALUES
    ('Shodmon', 'password');

SELECT id, username, password FROM users WHERE id >= 1 AND id <=5;
SELECT * FROM users;
SELECT COUNT(*) AS all FROM users;

-- Таблица authors
--  Описание: Таблица авторов с именем и биографией. Каждый автор может иметь
--  множество книг.
CREATE TABLE authors (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(100),
                         biography TEXT
);

-- Таблица books
-- Описание: Таблица книг, связанная с таблицей авторов через author_id.
-- Отношение "один к многим" между авторами и книгами.
CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255),
                       author_id INT REFERENCES authors(id)
);

-- Таблица borrow
-- Описание: Таблица для отслеживания взятых и возвращенных книг.
-- Связь "много ко многим" между пользователями и книгами.
CREATE TABLE borrow (
                        id SERIAL PRIMARY KEY,
                        user_id INT REFERENCES users(id),
                        book_id INT REFERENCES books(id),
                        borrow_date DATE,
                        return_date DATE
);

-- Таблица profiles
-- Описание: Таблица профилей, связь "один к одному" с таблицей
-- пользователей. Хранит email и адрес пользователя.
CREATE TABLE profiles (
                          user_id INT PRIMARY KEY REFERENCES users(id),
                          email VARCHAR(255),
                          address VARCHAR(255)
);

-- Добавление начальных данных

-- Заполнение таблицы users
INSERT INTO users (username, password)
VALUES
        ('alice', 'pass123'),
        ('bob', 'pass456'),
        ('charlie', 'pass789');

-- Заполнение таблицы authors
INSERT INTO authors (name, biography)
VALUES
        ('J. K. Rowling', 'writer of Harry Potter'),
        ('George R. R. Martin', 'writer of Game of Thrones'),
        ('J. R. R. Tolkien', 'writer ofLord of the Rings');

-- Заполнение таблицы books
INSERT INTO books (title, author_id)
VALUES
        ('War and Peace', 1),
        ('Anna Karenina', 1),
        ('Crime and Punishment', 2),
        ('The Brothers Karamazov', 2);

-- Заполнение таблицы borrow
INSERT INTO borrow (user_id, book_id, borrow_date, return_date)
VALUES
        (1, 1, '2022-01-01', '2022-01-15'),

        (2, 2, '2022-02-01', NULL),

        (3, 3, '2022-03-01', '2022-03-15');

-- Заполнение таблицы profiles
INSERT INTO profiles (user_id, email, address)
VALUES
        (1, 'alice@example.com', '123 Maple St'),
        (2, 'bob@example.com', '456 Oak St'),
        (3, 'charlie@example.com', '789 Pine St');
