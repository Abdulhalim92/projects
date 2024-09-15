

--
-- CREATE TABLE reviews (
--                          id SERIAL PRIMARY KEY,
--                          user_id INT REFERENCES users(user_id),
--                          book_id INT REFERENCES books(book_id),
--                          review_text TEXT,
--                          rating DECIMAL(2, 1) CHECK (rating >= 1.0 AND rating <= 5.0), -- Рейтинг с одним десятичным знаком
--                          review_date DATE DEFAULT CURRENT_DATE
-- );


SELECT * FROM reviews;



SELECT * FROM users;
SELECT * FROM profiles;
SELECT * FROM books;
SELECT * FROM authors;

-- 4
SELECT * FROM books
WHERE author_id = (
    SELECT authors.author_id
    FROM authors
    WHERE name = 'J. K. Rowling'
);

-- 5
SELECT username, password FROM users
RIGHT JOIN profiles p
ON users.user_id = p.user_id
WHERE p.email IS NOT NULL OR p.email != '';

-- 6
UPDATE users
SET password = 'newpassword'
WHERE username = 'alice';
SELECT * FROM users;

-- 7
SELECT count(*)
FROM books;

--8
SELECT book_id, title FROM books
WHERE book_id IN (
    SELECT book_id from borrows
    WHERE return_date IS NULL
    );

-- 9
SELECT * FROM profiles
RIGHT JOIN users u on u.user_id = profiles.user_id;

-- 10
SELECT user_id, username, password from users u
RIGHT JOIN profiles p
on u.user_id = p.user_id
WHERE p.email = 'afanofmartialarts@gmail.com';

-- 11
SELECT * FROM borrows;
SELECT * FROM books
WHERE book_id IN (
    SELECT book_id FROM borrows
    );

--12

-- 13
INSERT INTO users (username, password)
VALUES
    ('bob', 'verysecret');
SELECT * FROM users;
INSERT INTO profiles
VALUES
    (8, 'mail@mail.ru', 'some_address');

UPDATE profiles
SET phone = '+10 30230324'
WHERE user_id in (
    SELECT user_id FROM users
    WHERE username = 'bob'
    );
SELECT * FROM profiles;

-- Показать автора книги 'War and Peace'
SELECT name FROM authors
WHERE author_id = (
    SELECT author_id FROM books
    WHERE title = 'War and Peace'
    );

-- Показать количество книг каждого автора

SELECT count(*) FROM books
RIGHT JOIN authors a on a.author_id = books.author_id
WHERE author_id IN (
    SELECT author_id FROM authors
    );

DELETE FROM users WHERE username = 'lonelyrabbit';

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS books CASCADE;
DROP TABLE IF EXISTS borrows CASCADE;
DROP TABLE IF EXISTS authors CASCADE;
DROP TABLE IF EXISTS reviews CASCADE;
DROP TABLE IF EXISTS profiles CASCADE;

-- Users table

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(100),
    password VARCHAR(100)
);

ALTER TABLE users ADD UNIQUE (username);
ALTER TABLE users ADD COLUMN created_at DATE NOT NULL DEFAULT now();
ALTER TABLE users ADD COLUMN updated_at DATE;
ALTER TABLE users ALTER COLUMN password TYPE VARCHAR(255);

CREATE TABLE authors (
    author_id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    biography TEXT
);

ALTER TABLE authors ADD COLUMN date_of_birth DATE;
ALTER TABLE authors ADD COLUMN created_at DATE DEFAULT now();
ALTER TABLE authors ADD COLUMN updated_at DATE;

CREATE TABLE books(
    book_id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    author_id INT,
    CONSTRAINT UQ_books_title UNIQUE (title),
    CONSTRAINT FK_books_author_id FOREIGN KEY (author_id) REFERENCES authors(author_id)
);

ALTER TABLE books ADD COLUMN created_at DATE DEFAULT now();
ALTER TABLE books ADD COLUMN updated_at DATE;

CREATE TABLE borrows (
    borrow_id SERIAL PRIMARY KEY,
    user_id INT,
    CONSTRAINT FK_borrow_user_id FOREIGN KEY (user_id) REFERENCES users(user_id),
    book_id INT,
    CONSTRAINT FK_borrow_book_id FOREIGN KEY (book_id) REFERENCES books(book_id),
    borrow_date DATE DEFAULT now(),
    return_date DATE
);

ALTER TABLE borrows
ALTER COLUMN borrow_id TYPE SERIAL;
-- ALTER TABLE borrows
-- ALTER COLUMN borrow_date TYPE NULL



CREATE TABLE profiles(
    user_id INT PRIMARY KEY,
    email VARCHAR(255),
    phone VARCHAR(50),
    CONSTRAINT FK_profile_user_id FOREIGN KEY (user_id) REFERENCES users(user_id),
    created_at DATE NOT NULL DEFAULT now(),
    updated_at DATE
);

CREATE TABLE reviews(
    review_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL ,
    CONSTRAINT FK_review_user_id FOREIGN KEY (user_id) REFERENCES users(user_id),
    book_id INT NOT NULL ,
    CONSTRAINT FK_review_book_id FOREIGN KEY (book_id) REFERENCES books(book_id),
    review_text TEXT,
    rating DECIMAL(2, 1) CHECK (rating >= 1.0 AND rating <= 5.0),
    created_at DATE DEFAULT now(),
    updated_at DATE
);

ALTER TABLE reviews ALTER COLUMN updated_at SET DEFAULT NULL;

INSERT INTO users (username, password)
VALUES
    ('harry_black', 'asd1234'),
    ('lora_tonks', 'password'),
    ('barty_crouch', 'securepassword'),
    ('matin_adams', 'verysecure');

INSERT INTO profiles (user_id, email, phone)
VALUES
    (1, 'mail@mail.ru', '+7 27324442'),
    (2, 'lora@mail.ru', '+ 32 3243234'),
    (3, 'barty@mail.ru', '+10 93209489'),
    (4, 'martin@mail.ru', '+2 348866372');

INSERT INTO authors (name, biography)
VALUES
    ('J. K. Rowling', 'writer of Harry Potter'),
    ('J. R. R. Tolkien', 'writer of the Lord of Rings'),
    ('J. Austen', 'writer of Pride and Prejudice'),
    ('E. Brontë', 'writer of Wuthering Heights'),
    ('A. Dumas', 'The Three Musketeers'),
    ('J. London', 'writer of Martin Eden');

INSERT INTO books (title, author_id)
VALUES
    ('Harry Potter and the Goblet of Fire', 1),
    ('Harry Potter and the Chamber of Secrets', 1),
    ('Harry Potter and the Prisoner of Azkaban', 1),
    ('Harry Potter and the Philosopher''s Stone', 1),
    ('Harry Potter and the Half-Blood Prince', 1),
    ('The Lord of the Rings', 2),
    ('Pride and Prejudice', 3),
    ('Wuthering Heights', 4),
    ('The Three Musketeers', 5),
    ('Martin Eden', 6);

INSERT INTO borrows (user_id, book_id)
VALUES
    (1, 1),
    (2, 1),
    (3, 2),
    (1, 2),
    (2, 3);


UPDATE borrows
SET return_date = now()
WHERE user_id = 1 AND book_id = 1;

INSERT INTO borrows (user_id, book_id)
VALUES
    (4, 5),
    (3, 8),
    (5, 2),
    (1, 9),
    (2, 6);

INSERT INTO borrows (user_id, book_id)
VALUES
    (4, 7);

DELETE FROM borrows
WHERE borrows.borrow_id = 7;

INSERT INTO reviews (user_id, book_id, review_text, rating)
VALUES
    (1, 1, 'Отличная книга, очень понравилась!', 4.0),
    (2, 1, 'Хорошая книга, но есть недочеты.', 3.5),
    (3, 2, 'Невероятно увлекательная книга!', 5.0),
    (1, 2, 'Не понравилось, ожидал большего.', 2.5),
    (2, 3, 'Интересная, но тяжеловата для чтения.', 4.5);


SELECT * FROM reviews
WHERE book_id = 3;

