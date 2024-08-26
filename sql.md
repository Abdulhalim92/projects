# Вход в PostgreSQL
```bash
sudo -u postgres psql
```

# Создание пользователя humo
```sql
CREATE USER humo WITH PASSWORD 'humo';
```

# Назначение привилегий пользователю humo
```sql
ALTER USER humo CREATEDB;
```

# Создание базы данных humo_db
```sql
CREATE DATABASE humo_db OWNER humo;
```

# Подключение к базе данных humo_db
```sql
\c humo_db
```

# Создание таблицы users
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);
```

# Создание таблицы books
```sql
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL
);
```

# Добавление данных в таблицу users
```sql
INSERT INTO users (username, password) VALUES 
('user1', 'pass1'),
('user2', 'pass2'),
('user3', 'pass3'),
('user4', 'pass4'),
('user5', 'pass5'),
('user6', 'pass6'),
('user7', 'pass7'),
('user8', 'pass8'),
('user9', 'pass9'),
('user10', 'pass10');
```

# Добавление записей в таблицу books
```sql
INSERT INTO books (title, author) VALUES 
('Book1', 'Author1'),
('Book2', 'Author2'),
('Book3', 'Author3'),
('Book4', 'Author4'),
('Book5', 'Author5'),
('Book6', 'Author6'),
('Book7', 'Author7'),
('Book8', 'Author8'),
('Book9', 'Author9'),
('Book10', 'Author10');
```

# Просмотр содержимого таблицы users
```sql
SELECT * FROM users;
```

# Просмотр содержимого таблицы books
```sql
SELECT * FROM books;
```

# Обновление записи в таблице users
```sql
UPDATE users SET password = 'newpass' WHERE id = 1;
```

# Получение обновлённых данных из таблицыusers
```sql
SELECT * FROM users WHERE id = 1;
```

# Выход из psql
```bash
\q
```