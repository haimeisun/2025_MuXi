/*1、建表*/
/*书籍*/
CREATE TABLE Book (
    id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    author VARCHAR(50)
);
/*人物*/
CREATE TABLE Person (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL
);
/*库存*/
CREATE TABLE Storage (
    id INT PRIMARY KEY AUTO_INCREMENT,
    book_id VARCHAR(36),
    stock INT,
    FOREIGN KEY (book_id) REFERENCES Book(id)
);
/*书籍偏好*/
CREATE TABLE ReadingPreference (
    person_id INT,
    book_id VARCHAR(36),
    PRIMARY KEY (person_id, book_id),
    FOREIGN KEY (person_id) REFERENCES Person(id),
    FOREIGN KEY (book_id) REFERENCES Book(id)
);

/*2、插入数据*/
INSERT INTO Book (id, title, author) VALUES
('go-away', 'the way to go', 'Ivo'),
('go-lang', 'Go语言圣经', 'Alan, Brian'),
('go-web', 'Go Web编程', 'Anonymous'),
('con-cur', 'Concurrency in Go', 'Katherine');

INSERT INTO Person (name) VALUES
('小明'),
('张三'),
('翟曙');

INSERT INTO Storage (book_id, stock) VALUES
('go-away', 20),
('go-lang', 17),
('go-web', 4),
('con-cur', 9);

INSERT INTO ReadingPreference (person_id, book_id) VALUES
(1, 'go-away'), -- 小明喜欢#1
(1, 'go-web'),   -- 小明喜欢#3
(1, 'con-cur'),  -- 小明喜欢#4
(2, 'go-away'),  -- 张三喜欢#1
(3, 'go-web'),   -- 翟曙喜欢#3
(3, 'con-cur');  -- 翟曙喜欢#4

/*SQL查询*/
--1.查询喜欢阅读#3的⼈--
SELECT p.name
FROM Person p
INNER JOIN ReadingPreference rp ON p.id = rp.person_id
INNER JOIN Book b ON rp.book_id = b.id
WHERE b.id = 'go-web';

--2.查询没有被喜欢阅读的书的信息(id,作者,标题,库存)--
SELECT b.id, b.author, b.title, s.stock
FROM Book b
LEFT JOIN ReadingPreference rp ON b.id = rp.book_id
LEFT JOIN Storage s ON b.id = s.book_id
WHERE rp.person_id IS NULL;

--3.查询哪些⼈喜欢哪本书,列出⼈名和书名--
SELECT p.name, b.title
FROM Person p
INNER JOIN ReadingPreference rp ON p.id = rp.person_id
INNER JOIN Book b ON rp.book_id = b.id;