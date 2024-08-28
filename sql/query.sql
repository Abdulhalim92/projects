-- Создание таблиц

-- Таблица users
-- Описание: Таблица пользователей для хранения учетных данных.
CREATE TABLE users (
                       userId SERIAL PRIMARY KEY,
                       username VARCHAR(100),
                       password VARCHAR(100)
);

-- Таблица authors
--  Описание: Таблица авторов с именем и биографией. Каждый автор может иметь
--  множество книг.
CREATE TABLE authors (
                         authorId SERIAL PRIMARY KEY,
                         name VARCHAR(100),
                         biography TEXT
);

-- Таблица books
-- Описание: Таблица книг, связанная с таблицей авторов через author_id.
-- Отношение "один к многим" между авторами и книгами.
CREATE TABLE books (
                       bookId SERIAL PRIMARY KEY,
                       title VARCHAR(255) UNIQUE,
                       authorId INT REFERENCES authors(authorId)
);

-- Таблица borrow
-- Описание: Таблица для отслеживания взятых и возвращенных книг.
-- Связь "много ко многим" между пользователями и книгами.
CREATE TABLE borrow (
                        borrowId SERIAL PRIMARY KEY,
                        userId INT REFERENCES users(userId),
                        bookId INT REFERENCES books(bookId),
                        borrowDate DATE,
                        returnDate DATE
);

ALTER TABLE borrow
    ADD CONSTRAINT FK_borrow_user_id FOREIGN KEY (userId) REFERENCES users(userId);

ALTER TABLE borrow
    ADD CONSTRAINT FK_borrow_book_id FOREIGN KEY (bookId) REFERENCES books(bookId);

-- Таблица profiles
-- Описание: Таблица профилей, связь "один к одному" с таблицей
-- пользователей. Хранит email и адрес пользователя.
CREATE TABLE profiles (
                          userId INT PRIMARY KEY REFERENCES users(userId),
                          email VARCHAR(255),
                          address VARCHAR(255)
);

-- Таблица reviews
-- Описание: Таблица для хранения отзывов на книги. Связана с таблицами users и books.
CREATE TABLE reviews (
                         reviewId SERIAL PRIMARY KEY,
                         userId INT REFERENCES users(userId),
                         bookId INT REFERENCES books(bookId),
                         reviewText TEXT,
                         rating DECIMAL(2, 1) CHECK (rating >= 1.0 AND rating <= 5.0), -- Рейтинг с одним десятичным знаком
                         reviewDate DATE DEFAULT CURRENT_DATE
);


-- Добавление начальных данных

TRUNCATE TABLE books, borrow, reviews, users, profiles, authors RESTART IDENTITY;
TRUNCATE TABLE borrow RESTART IDENTITY;
TRUNCATE TABLE users RESTART IDENTITY CASCADE;
TRUNCATE TABLE profiles RESTART IDENTITY;
TRUNCATE TABLE reviews RESTART IDENTITY;
TRUNCATE TABLE authors RESTART IDENTITY CASCADE;


INSERT INTO authors (name, biography)
VALUES
    ('J. K. Rowling', 'writer of Harry Potter'),
    ('George R. R. Martin', 'writer of Game of Thrones'),
    ('J. R. R. Tolkien', 'writer ofLord of the Rings'),
    ('Leo Tolstoy', 'writer of Leo Tolstoy'),
    ('Charles Darwin', 'writer of Darwin'),
    ('Fyodor Dostoevsky', 'writer of Dostoevsky'),
    ('Victor Hugo', 'writer of Hugo'),
    ('William Shakespeare', 'writer of Shakespeare'),
    ('Mark Twain', 'writer of Twain'),
    ('William Wordsworth', 'writer of Wordsworth'),
    ('William Butler', 'writer of Butler'),
    ('Charles Dickens', 'writer of Fiction'),
    ('George Eliot', 'writer of Eliot'),
    ('Oscar Wilde', 'writer of Wilde'),
    ('Homer', 'writer of Homer');

INSERT INTO books (title, authorId)
VALUES
    ('Harry Potter and the Philosophers Stone', 1),
    ('Harry Potter and the Chamber of Secrets', 1),
    ('A Game of Thrones', 2),
    ('A Clash of Kings', 2),
    ('The Fellowship of the Ring', 3),
    ('The Two Towers', 3),
    ('War and Peace', 4),
    ('Anna Karenina', 4),
    ('On the Origin of Species', 5),
    ('The Descent of Man', 5),
    ('Crime and Punishment', 6),
    ('The Brothers Karamazov', 6),
    ('Les Misérables', 7),
    ('The Hunchback of Notre-Dame', 7),
    ('Romeo and Juliet', 8),
    ('Hamlet', 8),
    ('Adventures of Huckleberry Finn', 9),
    ('The Adventures of Tom Sawyer', 9),
    ('Lyrical Ballads', 10),
    ('The Prelude', 10),
    ('The Second Coming', 11),
    ('The Wild Swans at Coole', 11),
    ('A Tale of Two Cities', 12),
    ('Great Expectations', 12),
    ('Middlemarch', 13),
    ('The Mill on the Floss', 13),
    ('The Picture of Dorian Gray', 14),
    ('The Importance of Being Earnest', 14),
    ('The Iliad', 15),
    ('The Odyssey', 15);

INSERT INTO users (username, password)
VALUES
    ('johndoe', 'password123'),
    ('janedoe', 'securepassword'),
    ('alice', 'alicepass'),
    ('bob', 'bobpassword'),
    ('charlie', 'charlie123'),
    ('david', 'davidpass'),
    ('eva', 'evapass'),
    ('frank', 'frankpass'),
    ('grace', 'gracepass'),
    ('henry', 'henrypass'),
    ('isabel', 'isabelpass'),
    ('jack', 'jackpass'),
    ('karen', 'karenpass'),
    ('leo', 'leopass'),
    ('mia', 'miapass');

INSERT INTO profiles (userId, email, address)
VALUES
    (1, 'johndoe@example.com', '123 Main St'),
    (2, 'janedoe@example.com', NULL),          -- Пустое поле адреса
    (3, NULL, '789 Elm St'),                   -- Пустое поле email
    (4, 'bob@example.com', '321 Pine St'),
    (5, NULL, NULL),                           -- Оба поля пустые
    (6, 'david@example.com', '987 Cedar St'),
    (7, 'eva@example.com', NULL),
    (8, NULL, '654 Oak Ave'),
    (9, 'grace@example.com', '321 Elm St'),
    (10, 'henry@example.com', NULL),
    (11, NULL, NULL),
    (12, 'jack@example.com', '987 Oak St'),
    (13, 'karen@example.com', '456 Maple Ave'),
    (14, NULL, '369 Pine St'),
    (15, 'mia@example.com', '159 Cedar St');

INSERT INTO borrow (userId, bookId, borrowDate, returnDate)
VALUES
    (1, 1, '2024-08-01', '2024-08-15'),
    (2, 2, '2024-08-05', '2024-08-20'),
    (3, 3, '2024-08-10', NULL),                -- Пустое поле return_date
    (4, 4, '2024-08-15', '2024-08-30'),
    (5, 5, '2024-08-20', NULL),
    (6, 6, '2024-08-25', '2024-09-05'),
    (7, 7, '2024-08-30', NULL),
    (8, 8, '2024-09-01', '2024-09-15'),
    (9, 9, '2024-09-05', '2024-09-20'),
    (10, 10, '2024-09-10', NULL),
    (11, 11, '2024-09-15', '2024-09-30'),
    (12, 12, '2024-09-20', NULL),
    (13, 13, '2024-09-25', '2024-10-10'),
    (14, 14, '2024-09-30', NULL),
    (15, 15, '2024-10-05', '2024-10-20');

INSERT INTO reviews (userId, bookId, reviewText, rating)
VALUES
    (1, 1, 'An amazing start to a magical series.', 5.0),
    (2, 2, 'A gripping tale of power and betrayal.', 4.5),
    (3, 3, NULL, 4.8),                             -- Пустое поле review_text
    (4, 4, 'A deep and thought-provoking narrative.', 4.7),
    (5, 5, NULL, 4.9),                             -- Пустое поле review_text
    (6, 6, 'A fantastic sequel.', 4.6),
    (7, 7, 'A bit long but worth the read.', 4.4),
    (8, 8, NULL, 4.3),                             -- Пустое поле review_text
    (9, 9, 'A classic that everyone should read.', 4.9),
    (10, 10, 'Very interesting and insightful.', 4.7),
    (11, 11, NULL, 4.2),                           -- Пустое поле review_text
    (12, 12, 'A masterpiece of world literature.', 5.0),
    (13, 13, 'A complex and engaging story.', 4.8),
    (14, 14, NULL, 4.4),                           -- Пустое поле review_text
    (15, 15, 'An epic tale from antiquity.', 4.9);
DELETE FROM books WHERE books.bookId = 0;
DELETE FROM books;
DELETE FROM users;
DELETE FROM borrow;
DELETE FROM profiles;
DELETE FroM reviews;
DELETE FroM authors;

INSERT INTO users(username, password) VALUES ('Said', 'something');
SELECT * FROM users;