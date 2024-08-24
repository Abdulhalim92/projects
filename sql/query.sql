-- Создание таблиц

-- Таблица users
-- Описание: Таблица пользователей для хранения учетных данных.
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(100),
                       password VARCHAR(100)
);

-- Таблица authors
--  Описание: Таблица авторов с именем и биографией. Каждый автор может иметь
--  множество книг.
CREATE TABLE authors (
                         author_id SERIAL PRIMARY KEY,
                         name VARCHAR(100),
                         biography TEXT
);

-- Таблица books
-- Описание: Таблица книг, связанная с таблицей авторов через author_id.
-- Отношение "один к многим" между авторами и книгами.
CREATE TABLE books (
                       book_id SERIAL PRIMARY KEY,
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

-- Таблица reviews
-- Описание: Таблица для хранения отзывов на книги. Связана с таблицами users и books.
CREATE TABLE reviews (
                         id SERIAL PRIMARY KEY,
                         user_id INT REFERENCES users(user_id),
                         book_id INT REFERENCES books(id),
                         review_text TEXT,
                         rating DECIMAL(2, 1) CHECK (rating >= 1.0 AND rating <= 5.0), -- Рейтинг с одним десятичным знаком
                         review_date DATE DEFAULT CURRENT_DATE
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
        (3, 'charlie@example.com', '789 Pine St'),
        (4, '', 'Dushanbe');

INSERT INTO profiles (user_id, address)
VALUES
    (5, 'Dushanbe');

-- Заполнение таблицы reviews
INSERT INTO reviews (user_id, book_id, review_text, rating)
VALUES
    (1, 1, 'Отличная книга, очень понравилась!', 4.0),
    (2, 1, 'Хорошая книга, но есть недочеты.', 3.5),
    (3, 2, 'Невероятно увлекательная книга!', 5.0),
    (1, 2, 'Не понравилось, ожидал большего.', 2.5),
    (2, 3, 'Интересная, но тяжеловата для чтения.', 4.5);

