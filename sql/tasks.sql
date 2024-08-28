-- Легкие задачи (15)

-- Выбрать всех пользователей.
    SELECT * FROM users WHERE users."UserId" = 5;
-- Выбрать все книги.
    SELECT * FROM books;
-- Найти всех авторов.
    SELECT * FROM authors;
-- Выбрать книги определенного автора по имени.
SELECT *FROM books join authors ON books.authorId = authors.authorId WHERE authors.name = 'George R. R. Martin';
-- Выбрать всех пользователей, у которых есть email.
    SELECT * FROM profiles WHERE profiles.email IS NOT NULL;
-- Обновить пароль пользователя 'alice'.
    UPDATE users SET password = 'pass888' WHERE username = 'alice';
-- Посчитать количество книг в базе данных.
    SELECT count(*) AS Количество_книг FROM books;
-- Вывести список книг, которые в данный момент находятся на руках у пользователей.
SELECT *FROM books JOIN borrow ON books.bookId = borrow.bookId WHERE borrow.returnDate IS NULL;
-- Показать профили всех пользователей.
    SELECT * FROM profiles JOIN users ON profiles.userId = users.userId;
-- Найти пользователя по email.
    SELECT * FROM users JOIN profiles ON users.userId = profiles.userId WHERE profiles.email = 'charlie@example.com';
-- Вывести список всех книг, которые были когда-либо взяты.
    SELeCT * FROM books JOIN borrow ON books.bookId = borrow.bookId;
-- Добавить нового пользователя с профилем.
        WITH new_user AS
         ( INSERT INTO users (username, password)
             VALUES
                 ('dave', 'pass101')
             RETURNING userId
         )
        INSERT INTO profiles (userId, email, address)
            SELECT userId,
                'something@example.com',
                '4289428 something' FROM new_user;
-- Обновить адрес пользователя 'bob'.
    UPDATE profiles SET address = 'NewAddress' WHERE userId = (SELECT userId FROM users WHERE username = 'bob');
    SELECT * FROM users;
-- Показать автора книги 'War and Peace'.
SELECT *FROM authors LEFT JOIN books ON authors.authorId = books.authorId WHERE books.title = 'War and Peace';
-- Показать количество книг каждого автора.
SELECT authors.name, count(*)FROM books JOIN authors ON books.authorId = authors.authorId GROUP BY authorId;

-- Средние задачи (15)

-- Выбрать книги, которые возвращены позже установленной даты.
SELECT books.title FROM books JOIN borrow ON books.bookId = borrow.bookId WHERE borrow.returnDate > '2022-01-30';
-- Вывести список книг, которые никогда не были возвращены.
SELECT *FROM books JOIN borrow ON books.bookId = borrow.bookId WHERE borrow.returnDate IS NULL;
-- Выбрать всех пользователей, у которых нет профиля.
    SELECT * FROM users WHERE users.userId not in (SELECT userId FROM profiles);
-- Вставить новую книгу для автора 'Fyodor Dostoevsky'.
    ---INSERT INTO books (author_id, title)VALUES((SELECT authors_id FROM authors WHERE title = 'Fyodor Dostoevsky'), 'NewBookOfFyodor');
-- Обновить биографию автора 'Leo Tolstoy'.
    UPDATE authors SET biography = 'Russian writer, best known for War and Peace and Anna Karenina.'
    WHERE name = 'Leo Tolstoy';
-- Показать всех пользователей, которые взяли книгу более одного раза.
SELECT u.username, COUNT(br.bookId)FROM users u
    JOIN borrow br
    ON u.userId = br.userId
GROUP BY u.username
HAVING COUNT(br.bookId) > 1;
-- Вывести список книг, возвращенных в течение последнего месяца.
    SELECT b.title FROM books b
    JOIN borrow br
    ON b.bookId = br.bookId
    WHERE br.returnDate BETWEEN NOW() - INTERVAL '1 month' AND NOW();
-- Выбрать книги, которые взяты, но не возвращены, и количество дней просрочки.
SELECT b.title, CURRENT_DATE - br.borrowDate AS overdue_days FROM books b
    JOIN borrow br ON
    b.bookId = br.bookId
WHERE br.returnDate IS NULL;
-- Выбрать книги, которые были возвращены и имеют отзывы.
    SELECT DISTINCT b.title FROM books b
    JOIN borrow br
    ON b.bookId = br.bookId
    JOIN reviews r
    ON b.bookId = r.bookId
    WHERE br.returnDate IS NOT NULL;
-- Добавить отзыв пользователя 'alice' на книгу 'War and Peace'.
    INSERT INTO reviews (bookId, userId, rating, reviewText)
    VALUES
        ((SELECT books.bookId FROM books WHERE title = 'War and Peace'),
        (SELECT users.userId FROM users WHERE username = 'alice'),
        5,'Excellent read!');
-- Обновить рейтинг всех отзывов на книгу 'Anna Karenina' на 4.
    UPDATE reviews SET rating = 4
    WHERE bookId = (SELECT books.bookId FROM books WHERE title = 'Anna Karenina');
-- Выбрать авторов, которые имеют более двух книг.
    SELECT a.name FROM authors a
    JOIN books b
    ON a.authorId = b.authorId
    GROUP BY a.name
    HAVING COUNT(b.bookId) > 2;
-- Показать книги, которые не были возвращены в течение 30 дней после взятия.
    SELECT b.title FROM books b
    JOIN borrow br
    ON b.bookId = br.bookId
    WHERE br.returnDate IS NULL AND CURRENT_DATE - br.borrowDate > 30;
-- Выбрать книги, которые были взяты более пяти раз.
    SELECT b.title, COUNT(*) FROM books b
    JOIN borrow br
    ON b.bookId = br.bookId
    GROUP BY b.title
    HAVING COUNT(*) > 5;
-- Вывести список книг, которые были взяты в текущем году.
    SELECT b.title FROM books b
    JOIN borrow br
    ON b.bookId = br.bookId
    WHERE EXTRACT(YEAR FROM br.borrowDate) = EXTRACT(YEAR FROM CURRENT_DATE);

-- Сложные задачи (15)

-- Выбрать книги, которые взяли пользователи с более чем одним адресом.
    SELECT DISTINCT b.title FROM books b
    JOIN borrow br
    ON b.bookId = br.bookId
    JOIN profiles p
    ON br.userId = p.userId
    WHERE p.address IS NOT NULL
    GROUP BY b.title, p.userId
    HAVING COUNT(DISTINCT p.address) > 1;
-- Вставить несколько пользователей и сразу же создать для них профили.
    WITH ins AS
             ( INSERT INTO users (username, password)
                 VALUES
                     ('newuser1', 'pass123'),
                    ('newuser2', 'pass123')
                RETURNING userId
             )
    INSERT INTO profiles (userId, email, address)
    SELECT userId,
            'email@example.com',
            'Address' FROM ins;
-- Показать всех пользователей, которые взяли книги, но не оставили отзывы на них.
    SELECT DISTINCT u.username FROM users u
    JOIN borrow br
    ON u.userId = br.userId
    LEFT JOIN reviews r
    ON br.bookId = r.bookId AND br.userId = r.userId
    WHERE r.reviewId IS NULL;
-- Обновить профили всех пользователей, которые вернули книги, установив address на 'Updated Address'.
    UPDATE profiles SET address = 'Updated Address'
    WHERE userId IN (
        SELECT DISTINCT br.userId FROM borrow br
        WHERE br.returnDate IS NOT NULL
    );
-- Выбрать книги, которые взяты, но не возвращены, и количество дней просрочки.
SELECT b.title, CURRENT_DATE - br.borrowDate AS overdue_days FROM books b
    JOIN borrow br
    ON b.bookId = br.bookId
WHERE br.returnDate IS NULL AND CURRENT_DATE > br.borrowDate;
-- Выбрать книги, которые были возвращены, но имеют разные оценки от разных пользователей.
    SELECT b.title FROM books b
    JOIN reviews r
    ON b.bookId = r.bookId
    GROUP BY b.title
    HAVING COUNT(DISTINCT r.rating) > 1;
-- Обновить пароль для всех пользователей, которые не взяли книги в этом году.
    UPDATE users SET password = 'new2023pass'
    WHERE users.userId NOT IN
         (
             SELECT DISTINCT br.userId FROM borrow br
             WHERE EXTRACT(YEAR FROM br.borrowDate) = EXTRACT(YEAR FROM CURRENT_DATE)
         );
-- Выбрать пользователей, которые взяли книги, исключая тех, кто взял книги определенного автора.
    SELECT DISTINCT u.username FROM users u
    JOIN borrow br
    ON u.userId = br.userId
    WHERE NOT EXISTS
            (
                SELECT 1 FROM borrow br2
                JOIN books b
                ON br2.bookId = b.bookId
                WHERE b.authorId = (SELECT authors.authorId FROM authors WHERE name = 'Leo Tolstoy')AND br2.userId = u.userId
          );
-- Выбрать книги, которые были взяты и возвращены более трех раз.
    SELECT b.title, COUNT(*) FROM books b
    JOIN borrow br ON b.bookId = br.bookId
    WHERE br.returnDate IS NOT NULL
    GROUP BY b.title
    HAVING COUNT(*) > 3;
-- Показать книги с наибольшим количеством отзывов.
    SELECT b.title, COUNT(*) AS review_count FROM books b
    JOIN reviews r ON b.bookId = r.bookId
    GROUP BY b.title
    ORDER BY review_count DESC
    LIMIT 1;
-- Показать пользователей, которые взяли книги более 5 раз, но не оставили ни одного отзыва.
    SELECT u.username FROM users u
    JOIN borrow br
    ON u.userId = br.userId
    LEFT JOIN reviews r
    ON br.userId = r.userId AND br.bookId = r.bookId
    GROUP BY u.username
    HAVING COUNT(br.borrowId) > 5 AND COUNT(r.reviewId) = 0;
-- Выбрать книги, которые были взяты в прошлом году, но не в этом.
    SELECT DISTINCT b.title FROM books b
    JOIN borrow br
    ON b.bookId = br.bookId
    WHERE EXTRACT(YEAR FROM br.borrowDate) = EXTRACT(YEAR FROM CURRENT_DATE) - 1
    AND b.bookId NOT IN
      (
          SELECT br2.bookId FROM borrow br2
          WHERE EXTRACT(YEAR FROM br2.borrowDate) = EXTRACT(YEAR FROM CURRENT_DATE)
      );
-- Показать книги, на которые есть только положительные отзывы.
    SELECT b.title FROM books b
    JOIN reviews r
    ON b.bookId = r.bookId
    GROUP BY b.title
    HAVING MIN(r.rating) > 3;
-- Обновить return_date в borrow для всех книг определенного автора.
    UPDATE borrow SET returnDate = CURRENT_DATE
    WHERE bookId IN
      (
          SELECT books.bookId FROM books
          WHERE authorId = (SELECT authors.authorId FROM authors WHERE name = 'Leo Tolstoy')
      );
-- Выбрать книги, которые взяты пользователями из определенного адреса.
    SELECT b.title FROM books b
    JOIN borrow br
    ON b.bookId = br.bookId
    JOIN profiles p
    ON br.userId = p.userId
    WHERE p.address = 'Specific Address';
