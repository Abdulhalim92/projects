ALTER TABLE users
ADD UNIQUE (username);


CREATE TABLE reviews (
                         id SERIAL PRIMARY KEY,
                         user_id INT REFERENCES users(user_id),
                         book_id INT REFERENCES books(book_id),
                         review_text TEXT,
                         rating DECIMAL(2, 1) CHECK (rating >= 1.0 AND rating <= 5.0), -- Рейтинг с одним десятичным знаком
                         review_date DATE DEFAULT CURRENT_DATE
);

INSERT INTO reviews (user_id, book_id, review_text, rating)
VALUES
    (1, 1, 'Отличная книга, очень понравилась!', 4.0),
    (2, 1, 'Хорошая книга, но есть недочеты.', 3.5),
    (3, 2, 'Невероятно увлекательная книга!', 5.0),
    (1, 2, 'Не понравилось, ожидал большего.', 2.5),
    (2, 3, 'Интересная, но тяжеловата для чтения.', 4.5);

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
    SELECT book_id from borrow
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
SELECT * FROM borrow;
SELECT * FROM books
WHERE book_id IN (
    SELECT book_id FROM borrow
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
SET address = 'LA, California'
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