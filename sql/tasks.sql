-- Легкие задачи (15)

-- Выбрать всех пользователей.
SELECT * FROM users;

-- Выбрать все книги.
SELECT * FROM books;

-- Найти всех авторов.
SELECT * FROM authors;

-- Выбрать книги определенного автора по имени.
SELECT title AS j_k_rowling FROM books
    JOIN authors ON books.author_id = authors.id
    WHERE authors.name = 'J. K. Rowling';

-- Выбрать всех пользователей, у которых есть email.
SELECT * FROM users
    JOIN profiles ON profiles.user_id = users.id
    WHERE profiles.email IS NOT NULL;

-- Обновить пароль пользователя 'alice'.
UPDATE users SET password = 'qwerty' WHERE username = 'alice';

-- Посчитать количество книг в базе данных.
SELECT COUNT(*) AS amount_of_books FROM books;

-- Вывести список книг, которые в данный момент находятся на руках у пользователей.
SELECT books.* FROM books
    JOIN borrow ON borrow.book_id = books.id
    WHERE borrow.return_date IS NULL;

-- Показать профили всех пользователей.
SELECT users.username, profiles.* FROM users
    JOIN profiles ON profiles.user_id = users.id;

-- Найти пользователя по email.
SELECT * FROM profiles WHERE email = 'alice@example.com';

-- Вывести список всех книг, которые были когда-либо взяты.
SELECT books.* FROM books
    JOIN borrow ON borrow.book_id = books.id
    WHERE borrow.borrow_date IS NOT NULL;

-- Добавить нового пользователя с профилем.
INSERT INTO users (username, password) VALUES
    ('said','saidis4');

INSERT INTO profiles (user_id, email, address) VALUES
    (4, 'said@said.com','M. Kholov');

-- Обновить адрес пользователя 'bob'.
UPDATE profiles SET address = 'somewhere St' WHERE user_id = 2;

-- Показать автора книги 'War and Peace'.
SELECT authors.* FROM authors
    JOIN books ON books.author_id = authors.id
    WHERE books.title = 'War and Peace';

-- Показать количество книг каждого автора.
SELECT authors.name AS authors, COUNT(books) AS books_amount
    FROM authors
    JOIN books ON books.author_id = authors.id
    GROUP BY authors
    ORDER BY COUNT(books) DESC;
-- Средние задачи (15)

-- Выбрать книги, которые возвращены позже установленной даты.
SELECT books.* FROM books
    JOIN borrow ON borrow.book_id = books.id
    WHERE borrow.return_date > '2022-02-01';

-- Вывести список книг, которые никогда не были возвращены.
SELECT books.* FROM books
    JOIN borrow ON borrow.book_id = books.id
    WHERE borrow.return_date IS NULL;

-- Выбрать всех пользователей, у которых нет профиля.
SELECT users.* FROM users
    LEFT JOIN profiles ON profiles.user_id = users.id
    WHERE profiles IS NULL;

-- Вставить новую книгу для автора 'Fyodor Dostoevsky'.
INSERT INTO books (title, author_id)
    VALUES ('White Nights', 2);

-- Обновить биографию автора 'Leo Tolstoy'.
UPDATE authors SET biography = 'writer of the "War and Peace"' WHERE name = 'Leo Tolstoy';

-- Показать всех пользователей, которые взяли книгу более одного раза.
SELECT users.username AS users, COUNT(borrow) AS times_borrowed FROM users
    JOIN borrow ON borrow.user_id = users.id
    -- WHERE times_borrowed > 1
    GROUP BY users
    ORDER BY times_borrowed;

-- Вывести список книг, возвращенных в течение последнего месяца.
SELECT books.*, borrow.return_date FROM books
    JOIN borrow ON borrow.book_id = books.id
    WHERE borrow.return_date >= NOW() - INTERVAL '1 MONTH' AND borrow.return_date <= NOW();

-- Выбрать книги, которые взяты, но не возвращены, и количество дней просрочки.
SELECT EXTRACT(DAY FROM books.*, NOW() - borrow.borrow_date) AS days_passed FROM books
    JOIN borrow ON borrow.book_id = books.id
    WHERE borrow.return_date IS NULL;

-- Выбрать книги, которые были возвращены и имеют отзывы.
SELECT books.*, reviews.review_text, borrow.return_date FROM books
    JOIN reviews ON reviews.book_id = books.id
    JOIN borrow ON borrow.return_date = books.id
    WHERE reviews.review_text IS NOT NULL AND borrow.return_date IS NOT NULL;

-- Добавить отзыв пользователя 'alice' на книгу 'War and Peace'.
INSERT INTO reviews (id, user_id, book_id, review_text, rating, review_date)
    VALUES (1, 1, 1, 'The best book I have ever read!', 5.0, NOW())
    ON CONFLICT (id)
        DO UPDATE SET
            review_text = EXCLUDED.review_text,
            rating = EXCLUDED.rating,
            review_date = EXCLUDED.review_date;

-- Обновить рейтинг всех отзывов на книгу 'Anna Karenina' на 4.
-- Выбрать авторов, которые имеют более двух книг.
-- Показать книги, которые не были возвращены в течение 30 дней после взятия.
-- Выбрать книги, которые были взяты более пяти раз.
-- Вывести список книг, которые были взяты в текущем году.

-- Сложные задачи (15)

-- Выбрать книги, которые взяли пользователи с более чем одним адресом.
-- Вставить несколько пользователей и сразу же создать для них профили.
-- Показать всех пользователей, которые взяли книги, но не оставили отзывы на них.
-- Обновить профили всех пользователей, которые вернули книги, установив address на 'Updated Address'.
-- Выбрать книги, которые взяты, но не возвращены, и количество дней просрочки.
-- Выбрать книги, которые были возвращены, но имеют разные оценки от разных пользователей.
-- Обновить пароль для всех пользователей, которые не взяли книги в этом году.
-- Выбрать пользователей, которые взяли книги, исключая тех, кто взял книги определенного автора.
-- Выбрать книги, которые были взяты и возвращены более трех раз.
-- Показать книги с наибольшим количеством отзывов.
-- Показать пользователей, которые взяли книги более 5 раз, но не оставили ни одного отзыва.
-- Выбрать книги, которые были взяты в прошлом году, но не в этом.
-- Показать книги, на которые есть только положительные отзывы.
-- Обновить return_date в borrow для всех книг определенного автора.
-- Выбрать книги, которые взяты пользователями из определенного адреса.
