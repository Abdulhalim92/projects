-- Легкие задачи (15)

-- Выбрать всех пользователей.
    SELECT * FROM users;
-- Выбрать все книги.
    SELECT * FROM books;
-- Найти всех авторов.
    SELECT * FROM authors;
-- Выбрать книги определенного автора по имени.
    SELECT * FROM books join authors ON books.author_id = authors.authors_id WHERE authors.name = 'George R. R. Martin';
-- Выбрать всех пользователей, у которых есть email.
    SELECT * FROM profiles WHERE profiles.email IS NOT NULL;
-- Обновить пароль пользователя 'alice'.
    UPDATE users SET password = 'pass888' WHERE username = 'alice';
-- Посчитать количество книг в базе данных.
    SELECT count(*) AS Количество_книг FROM books;
-- Вывести список книг, которые в данный момент находятся на руках у пользователей.
    SELECT * FROM books JOIN borrow ON books.books_id = borrow.book_id WHERE borrow.return_date IS NULL;
-- Показать профили всех пользователей.
    SELECT * FROM profiles JOIN users ON profiles.user_id = users.users_id;
-- Найти пользователя по email.
    SELECT * FROM users JOIN profiles ON users.users_id = profiles.user_id WHERE profiles.email = 'charlie@example.com';
-- Вывести список всех книг, которые были когда-либо взяты.
    SELeCT * FROM books JOIN borrow ON books.books_id = borrow.book_id;
-- Добавить нового пользователя с профилем.
        WITH new_user AS
         ( INSERT INTO users (username, password)
             VALUES
                 ('dave', 'pass101')
             RETURNING users_id
         )
        INSERT INTO profiles (user_id, email, address)
            SELECT users_id,
                'something@example.com',
                '4289428 something' FROM new_user;
-- Обновить адрес пользователя 'bob'.
    UPDATE profiles SET address = 'NewAddress' WHERE user_id = (SELECT users_id FROM users WHERE username = 'bob');
    SELECT * FROM users;
-- Показать автора книги 'War and Peace'.
    SELECT * FROM authors LEFT JOIN books ON authors.authors_id = books.author_id WHERE books.title = 'War and Peace';
-- Показать количество книг каждого автора.
    SELECT authors.name, count(*) FROM books JOIN authors ON books.author_id = authors.authors_id  GROUP BY authors_id;

-- Средние задачи (15)

-- Выбрать книги, которые возвращены позже установленной даты.
SELECT books.title FROM books JOIN borrow ON books.books_id = borrow.book_id WHERE borrow.return_date > '2022-01-30';
-- Вывести список книг, которые никогда не были возвращены.
    SELECT * FROM books JOIN borrow ON books.books_id = borrow.book_id WHERE borrow.return_date IS NULL;
-- Выбрать всех пользователей, у которых нет профиля.
    SELECT * FROM users WHERE users.users_id not in (SELECT user_id FROM profiles);
-- Вставить новую книгу для автора 'Fyodor Dostoevsky'.
    ---INSERT INTO books (author_id, title)VALUES((SELECT authors_id FROM authors WHERE title = 'Fyodor Dostoevsky'), 'NewBookOfFyodor');
-- Обновить биографию автора 'Leo Tolstoy'.
    UPDATE authors SET biography = 'Russian writer, best known for War and Peace and Anna Karenina.'
    WHERE name = 'Leo Tolstoy';
-- Показать всех пользователей, которые взяли книгу более одного раза.
    SELECT u.username, COUNT(br.book_id) FROM users u
    JOIN borrow br
    ON u.users_id = br.user_id
    GROUP BY u.username
    HAVING COUNT(br.book_id) > 1;
-- Вывести список книг, возвращенных в течение последнего месяца.
    SELECT b.title FROM books b
    JOIN borrow br
    ON b.books_id = br.book_id
    WHERE br.return_date BETWEEN NOW() - INTERVAL '1 month' AND NOW();
-- Выбрать книги, которые взяты, но не возвращены, и количество дней просрочки.
    SELECT b.title, CURRENT_DATE - br.borrow_date AS overdue_days FROM books b
    JOIN borrow br ON
    b.books_id = br.book_id
    WHERE br.return_date IS NULL;
-- Выбрать книги, которые были возвращены и имеют отзывы.
    SELECT DISTINCT b.title FROM books b
    JOIN borrow br
    ON b.books_id = br.book_id
    JOIN reviews r
    ON b.books_id = r.book_id
    WHERE br.return_date IS NOT NULL;
-- Добавить отзыв пользователя 'alice' на книгу 'War and Peace'.
    INSERT INTO reviews (book_id, user_id, rating, review_text)
    VALUES
        ((SELECT books.books_id FROM books WHERE title = 'War and Peace'),
        (SELECT users.users_id FROM users WHERE username = 'alice'),
        5,'Excellent read!');
-- Обновить рейтинг всех отзывов на книгу 'Anna Karenina' на 4.
    UPDATE reviews SET rating = 4
    WHERE book_id = (SELECT books.books_id FROM books WHERE title = 'Anna Karenina');
-- Выбрать авторов, которые имеют более двух книг.
    SELECT a.name FROM authors a
    JOIN books b
    ON a.authors_id = b.author_id
    GROUP BY a.name
    HAVING COUNT(b.books_id) > 2;
-- Показать книги, которые не были возвращены в течение 30 дней после взятия.
    SELECT b.title FROM books b
    JOIN borrow br
    ON b.books_id = br.book_id
    WHERE br.return_date IS NULL AND CURRENT_DATE - br.borrow_date > 30;
-- Выбрать книги, которые были взяты более пяти раз.
    SELECT b.title, COUNT(*) FROM books b
    JOIN borrow br
    ON b.books_id = br.book_id
    GROUP BY b.title
    HAVING COUNT(*) > 5;
-- Вывести список книг, которые были взяты в текущем году.
    SELECT b.title FROM books b
    JOIN borrow br
    ON b.books_id = br.book_id
    WHERE EXTRACT(YEAR FROM br.borrow_date) = EXTRACT(YEAR FROM CURRENT_DATE);

-- Сложные задачи (15)

-- Выбрать книги, которые взяли пользователи с более чем одним адресом.
    SELECT DISTINCT b.title FROM books b
    JOIN borrow br
    ON b.books_id = br.book_id
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
                RETURNING users_id
             )
    INSERT INTO profiles (user_id, email, address)
    SELECT users_id,
            'email@example.com',
            'Address' FROM ins;
-- Показать всех пользователей, которые взяли книги, но не оставили отзывы на них.
    SELECT DISTINCT u.username FROM users u
    JOIN borrow br
    ON u.users_id = br.user_id
    LEFT JOIN reviews r
    ON br.book_id = r.book_id AND br.user_id = r.user_id
    WHERE r.reviews_id IS NULL;
-- Обновить профили всех пользователей, которые вернули книги, установив address на 'Updated Address'.
    UPDATE profiles SET address = 'Updated Address'
    WHERE user_id IN (
        SELECT DISTINCT br.user_id FROM borrow br
        WHERE br.return_date IS NOT NULL
    );
-- Выбрать книги, которые взяты, но не возвращены, и количество дней просрочки.
    SELECT b.title, CURRENT_DATE - br.borrow_date AS overdue_days FROM books b
    JOIN borrow br
    ON b.books_id = br.book_id
    WHERE br.return_date IS NULL AND CURRENT_DATE > br.borrow_date;
-- Выбрать книги, которые были возвращены, но имеют разные оценки от разных пользователей.
    SELECT b.title FROM books b
    JOIN reviews r
    ON b.books_id = r.book_id
    GROUP BY b.title
    HAVING COUNT(DISTINCT r.rating) > 1;
-- Обновить пароль для всех пользователей, которые не взяли книги в этом году.
    UPDATE users SET password = 'new2023pass'
    WHERE users.users_id NOT IN
         (
              SELECT DISTINCT br.user_id FROM borrow br
              WHERE EXTRACT(YEAR FROM br.borrow_date) = EXTRACT(YEAR FROM CURRENT_DATE)
         );
-- Выбрать пользователей, которые взяли книги, исключая тех, кто взял книги определенного автора.
    SELECT DISTINCT u.username FROM users u
    JOIN borrow br
    ON u.users_id = br.user_id
    WHERE NOT EXISTS
            (
                SELECT 1 FROM borrow br2
                JOIN books b
                ON br2.book_id = b.books_id
                WHERE b.author_id = (SELECT authors.authors_id FROM authors WHERE name = 'Leo Tolstoy') AND br2.user_id = u.users_id
          );
-- Выбрать книги, которые были взяты и возвращены более трех раз.
    SELECT b.title, COUNT(*) FROM books b
    JOIN borrow br ON b.books_id = br.book_id
    WHERE br.return_date IS NOT NULL
    GROUP BY b.title
    HAVING COUNT(*) > 3;
-- Показать книги с наибольшим количеством отзывов.
    SELECT b.title, COUNT(*) AS review_count FROM books b
    JOIN reviews r ON b.books_id = r.book_id
    GROUP BY b.title
    ORDER BY review_count DESC
    LIMIT 1;
-- Показать пользователей, которые взяли книги более 5 раз, но не оставили ни одного отзыва.
    SELECT u.username FROM users u
    JOIN borrow br
    ON u.users_id = br.user_id
    LEFT JOIN reviews r
    ON br.user_id = r.user_id AND br.book_id = r.book_id
    GROUP BY u.username
    HAVING COUNT(br.borrow_id) > 5 AND COUNT(r.reviews_id) = 0;
-- Выбрать книги, которые были взяты в прошлом году, но не в этом.
    SELECT DISTINCT b.title FROM books b
    JOIN borrow br
    ON b.books_id = br.book_id
    WHERE EXTRACT(YEAR FROM br.borrow_date) = EXTRACT(YEAR FROM CURRENT_DATE) - 1
    AND b.books_id NOT IN
      (
          SELECT br2.book_id FROM borrow br2
          WHERE EXTRACT(YEAR FROM br2.borrow_date) = EXTRACT(YEAR FROM CURRENT_DATE)
      );
-- Показать книги, на которые есть только положительные отзывы.
    SELECT b.title FROM books b
    JOIN reviews r
    ON b.books_id = r.book_id
    GROUP BY b.title
    HAVING MIN(r.rating) > 3;
-- Обновить return_date в borrow для всех книг определенного автора.
    UPDATE borrow SET return_date = CURRENT_DATE
    WHERE book_id IN
      (
          SELECT books.books_id FROM books
          WHERE author_id = (SELECT authors.authors_id FROM authors WHERE name = 'Leo Tolstoy')
      );
-- Выбрать книги, которые взяты пользователями из определенного адреса.
    SELECT b.title FROM books b
    JOIN borrow br
    ON b.books_id = br.book_id
    JOIN profiles p
    ON br.user_id = p.user_id
    WHERE p.address = 'Specific Address';
