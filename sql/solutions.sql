-- Легкие задачи (15)

-- Выбрать всех пользователей.
SELECT * FROM users;

SELECT * FROM users order by user_id DESC LIMIT 1;

-- Выбрать все книги.
SELECT * FROM books;

-- Найти всех авторов.
SELECT * FROM authors;

-- Выбрать книги определенного автора по имени.
SELECT b.title
FROM books b
JOIN authors a
ON b.author_id = a.id
WHERE a.name = 'J. K. Rowling';

SELECT title FROM books WHERE author_id IN (SELECT id FROM authors WHERE name = 'J. K. Rowling');

-- Выбрать всех пользователей, у которых есть email.
SELECT u.username
FROM users u
JOIN profiles p
ON u.user_id = p.user_id
WHERE p.email IS NOT NULL OR p.email != '';

-- Обновить пароль пользователя 'alice'.
UPDATE users SET password = 'newpassword123' WHERE username = 'alice';

-- Посчитать количество книг в базе данных.
SELECT COUNT(*) FROM books;

-- Вывести список книг, которые в данный момент находятся на руках у
-- пользователей.
SELECT b.title FROM books b
JOIN borrows br
ON b.id = br.book_id
WHERE br.return_date IS NULL;

-- Показать профили всех пользователей.
SELECT * FROM profiles;

-- Найти пользователя по email.
SELECT u.username FROM users u
JOIN profiles p
ON u.user_id = p.user_id
WHERE p.email = 'alice@example.com';

-- Вывести список всех книг, которые были когда-либо взяты.
SELECT DISTINCT b.title FROM books b
JOIN borrows br
ON b.id = br.book_id;

-- Добавить нового пользователя с профилем.
WITH new_user AS
    ( INSERT INTO users (username, password)
        VALUES
            ('dave', 'pass101')
        RETURNING user_id
    )
INSERT INTO profiles (user_id, email, address)
    SELECT user_id,
           'dave@example.com',
           '1600 Amphitheatre' FROM new_user;

-- Обновить адрес пользователя 'bob'.
UPDATE profiles SET address = '789 New St'
WHERE user_id = (SELECT users.user_id FROM users WHERE username = 'bob');

-- Показать автора книги 'War and Peace'.
SELECT a.name FROM authors a
JOIN books b
ON a.id = b.author_id
WHERE b.title = 'War and Peace';

-- Показать количество книг каждого автора.
SELECT a.name, COUNT(b.id) AS book_count FROM authors a
JOIN books b
ON a.id = b.author_id
GROUP BY a.name;

-- Средние задачи (15)

-- Выбрать книги, которые возвращены позже установленной даты.
SELECT b.title FROM books b
JOIN borrows br USING (book_id)
WHERE br.return_date > '2022-01-30';

-- Вывести список книг, которые никогда не были возвращены.
SELECT b.title FROM books b
JOIN borrows br USING (book_id)
WHERE br.return_date IS NULL;

-- Выбрать всех пользователей, у которых нет профиля.
SELECT u.username FROM users u
LEFT JOIN profiles p
ON u.user_id = p.user_id
WHERE p.user_id IS NULL;

-- Вставить новую книгу для автора 'Fyodor Dostoevsky'.
INSERT INTO books (title, author_id)
VALUES
    ('Notes from Underground', (SELECT id FROM authors WHERE name = 'Fyodor Dostoevsky'));

-- Обновить биографию автора 'Leo Tolstoy'.
UPDATE authors SET biography = 'Russian writer, best known for War and Peace and Anna Karenina.'
WHERE name = 'Leo Tolstoy';

-- Показать всех пользователей, которые взяли книгу более одного раза.
SELECT u.username, COUNT(br.book_id) FROM users u
JOIN borrows br
ON u.user_id = br.user_id
GROUP BY u.username
HAVING COUNT(br.book_id) > 1;

-- Вывести список книг, возвращенных в течение последнего месяца.
SELECT b.title FROM books b
JOIN borrows br
ON b.id = br.book_id
WHERE br.return_date BETWEEN NOW() - INTERVAL '1 month' AND NOW();

-- Выбрать книги, которые взяты, но не возвращены, и количество дней просрочки.
SELECT b.title, CURRENT_DATE - br.borrow_date AS overdue_days FROM books b
JOIN borrows br ON
b.id = br.book_id
WHERE br.return_date IS NULL;

-- Выбрать книги, которые были возвращены и имеют отзывы.
SELECT DISTINCT b.title FROM books b
JOIN borrows br
ON b.id = br.book_id
JOIN reviews r
ON b.id = r.book_id
WHERE br.return_date IS NOT NULL;

-- Добавить отзыв пользователя 'alice' на книгу 'War and Peace'.
INSERT INTO reviews (book_id, user_id, rating, review_text)
VALUES
    ((SELECT id FROM books WHERE title = 'War and Peace'),
     (SELECT id FROM users WHERE username = 'alice'),
     5,
     'Excellent read!');

-- Обновить рейтинг всех отзывов на книгу 'Anna Karenina' на 4.
UPDATE reviews SET rating = 4
WHERE book_id = (SELECT id FROM books WHERE title = 'Anna Karenina');

-- Выбрать авторов, которые имеют более двух книг.
SELECT a.name FROM authors a
JOIN books b
ON a.id = b.author_id
GROUP BY a.name
HAVING COUNT(b.id) > 2;

-- Показать книги, которые не были возвращены в течение 30 дней после взятия.
SELECT b.title FROM books b
JOIN borrows br
ON b.book_id = br.book_id
WHERE br.return_date IS NULL AND CURRENT_DATE - br.borrow_date > 30;

-- Выбрать книги, которые были взяты более пяти раз.
SELECT b.title, COUNT(*) FROM books b
JOIN borrows br
ON b.id = br.book_id
GROUP BY b.title
HAVING COUNT(*) > 5;

-- Вывести список книг, которые были взяты в текущем году.
SELECT b.title FROM books b
JOIN borrows br
ON b.id = br.book_id
WHERE EXTRACT(YEAR FROM br.borrow_date) = EXTRACT(YEAR FROM CURRENT_DATE);

-- Сложные задачи (15)

-- Выбрать книги, которые взяли пользователи с более чем одним адресом.
SELECT DISTINCT b.title FROM books b
JOIN borrows br
ON b.id = br.book_id
JOIN profiles p
ON br.user_id = p.user_id
WHERE p.address IS NOT NULL
GROUP BY b.title, p.user_id
HAVING COUNT(DISTINCT p.address) > 1;

-- Вставить несколько пользователей и сразу же создать для них профили.
WITH ins AS
    ( INSERT INTO users (username, password)
        VALUES
            ('newuser1', 'pass123'),
            ('newuser2', 'pass123')
        RETURNING user_id
    )
INSERT INTO profiles (user_id, email, address)
       SELECT user_id,
              'email@example.com',
              'Address' FROM ins;

-- Показать всех пользователей, которые взяли книги, но не оставили отзывы
-- на них.
SELECT DISTINCT u.username FROM users u
JOIN borrows br
ON u.user_id = br.user_id
LEFT JOIN reviews r
ON br.book_id = r.book_id AND br.user_id = r.user_id
WHERE r.id IS NULL;

-- Обновить профили всех пользователей, которые вернули книги, установив
-- address на 'Updated Address'.
UPDATE profiles SET address = 'Updated Address'
WHERE user_id IN (
        SELECT DISTINCT br.user_id FROM borrows br
        WHERE br.return_date IS NOT NULL
        );

-- Выбрать книги, которые взяты, но не возвращены, и количество дней
-- просрочки.
SELECT b.title, CURRENT_DATE - br.borrow_date AS overdue_days FROM books b
JOIN borrows br
ON b.id = br.book_id
WHERE br.return_date IS NULL AND CURRENT_DATE > br.borrow_date;

-- Выбрать книги, которые были возвращены, но имеют разные оценки от разных
-- пользователей.
SELECT b.title FROM books b
JOIN reviews r
ON b.id = r.book_id
GROUP BY b.title
HAVING COUNT(DISTINCT r.rating) > 1;

-- Обновить пароль для всех пользователей, которые не взяли книги в этом году.
UPDATE users SET password = 'new2023pass'
WHERE users.user_id NOT IN
      (
        SELECT DISTINCT br.user_id FROM borrows br
        WHERE EXTRACT(YEAR FROM br.borrow_date) = EXTRACT(YEAR FROM CURRENT_DATE)
        );

-- Выбрать пользователей, которые взяли книги, исключая тех, кто взял книги
-- определенного автора.
SELECT DISTINCT u.username FROM users u
JOIN borrows br
ON u.user_id = br.user_id
WHERE NOT EXISTS
    (
        SELECT 1 FROM borrows br2
        JOIN books b
        ON br2.book_id = b.id
        WHERE b.author_id = (SELECT id FROM authors WHERE name = 'Leo Tolstoy') AND br2.user_id = u.user_id
    );

-- Выбрать книги, которые были взяты и возвращены более трех раз.
SELECT b.title, COUNT(*) FROM books b
JOIN borrows br ON b.id = br.book_id
WHERE br.return_date IS NOT NULL
GROUP BY b.title
HAVING COUNT(*) > 3;

-- Показать книги с наибольшим количеством отзывов.
SELECT b.title, COUNT(*) AS review_count FROM books b
JOIN reviews r ON b.id = r.book_id
GROUP BY b.title
ORDER BY review_count DESC
LIMIT 1;

-- Показать пользователей, которые взяли книги более 5 раз, но не оставили
-- ни одного отзыва.
SELECT u.username FROM users u
JOIN borrows br
ON u.user_id = br.user_id
LEFT JOIN reviews r
ON br.user_id = r.user_id AND br.book_id = r.book_id
GROUP BY u.username
HAVING COUNT(br.id) > 5 AND COUNT(r.id) = 0;

-- Выбрать книги, которые были взяты в прошлом году, но не в этом.
SELECT DISTINCT b.title FROM books b
JOIN borrows br
ON b.id = br.book_id
WHERE EXTRACT(YEAR FROM br.borrow_date) = EXTRACT(YEAR FROM CURRENT_DATE) - 1
  AND b.id NOT IN
      (
        SELECT br2.book_id FROM borrows br2
        WHERE EXTRACT(YEAR FROM br2.borrow_date) = EXTRACT(YEAR FROM CURRENT_DATE)
        );

-- Показать книги, на которые есть только положительные отзывы.
SELECT b.title FROM books b
JOIN reviews r
ON b.id = r.book_id
GROUP BY b.title
HAVING MIN(r.rating) > 3;

-- Обновить return_date в borrow для всех книг определенного автора.
UPDATE borrows SET return_date = CURRENT_DATE
WHERE book_id IN
      (
        SELECT id FROM books
        WHERE author_id = (SELECT id FROM authors WHERE name = 'Leo Tolstoy')
        );

-- Выбрать книги, которые взяты пользователями из определенного адреса.
SELECT b.title FROM books b
JOIN borrows br
ON b.id = br.book_id
JOIN profiles p
ON br.user_id = p.user_id
WHERE p.address = 'Specific Address';