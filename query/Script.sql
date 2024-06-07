-- No. 4a
insert into genres (genre_name) values ('Fantasy'), ('Sci-Fi'), ('Dystopian'), ('Horror'), ('Mystery');

-- No. 4b
insert into books (genre_id, title, author, publisher, publish_year) values
(2, 'The Lord of The Rings', 'J. R. R. Tolkien', 'Houghton Mifflin', '1954'),
(2, 'The Hobbit', 'J. R. R. Tolkien', 'Houghton Mifflin', '1937'),
(2, 'A Game of Thrones', 'George R. R. Martin', 'Bantam Books', '1996'),
(2, 'Harry Potter', 'J. K. Rowling', 'Bloomsbury', '1997'),
(2, 'The Chronicles of Narnia', 'C. S. Lewis', 'HarperCollins', '1950');

insert into books (genre_id, title, author, publisher, publish_year) values
(3, 'Dune', 'Frank Herbert', 'Chilton Books', '1965'),
(3, 'The Left Hand of Darkness', 'Ursula K. Le Guin', 'Ace Books', '1969'),
(3, 'Neuromancer', 'William Gibson', 'Ace Books', '1984'),
(3, 'Hyperion', 'Dan Simmons', 'Doubleday', '1989'),
(3, 'Nineteen Eighty-Four', 'George Orwell', 'Secker & Warburg', '1949');

insert into books (genre_id, title, author, publisher, publish_year) values
(4, 'The Hunger Games', 'Suzanne Collins', 'Scholastic Press', '2008'),
(4, 'Divergent', 'Veronica Roth', 'HarperCollins', '2011'),
(4, 'Fahrenheit 451', 'Ray Bradbury', 'Ballantine Books', '1953'),
(4, 'The Giver', 'Lois Lowry', 'Houghton Mifflin', '1993'),
(4, 'Brave New World', 'Aldous Huxley', 'Chatto & Windus', '1932');

insert into books (genre_id, title, author, publisher, publish_year) values
(5, 'Ring', 'Koji Suzuki', 'Kadokawa Shoten', '1991'),
(5, 'The Shining', 'Stephen King', 'Double Day', '1977'),
(5, 'It', 'Stephen King', 'Viking Press', '1986'),
(5, 'Bird Box', 'Josh Malerman', 'Harper Voyager', '2014'),
(5, 'World War Z', 'Max Brooks', 'Crown', '2006');

insert into books (genre_id, title, author, publisher, publish_year) values
(6, 'Shutter Island', 'Dennis Lehane', 'HarperCollins', '2003'),
(6, 'The Da Vinci Code', 'Dan Brown', 'Doubleday', '2003'),
(6, 'The Silence of the Lambs', 'Thomas Harris', 'St. Martins Press', '1988'),
(6, 'Gone Girl', 'Gillian Flynn', 'Crown Publishing Group', '2012'),
(6, 'Big Little Lies', 'Liane Moriarty', 'Penguin Publishing', '2014');

-- No. 4c
insert into users (name, email, password, address, phone, birth_date) values
('John Doe', 'johndoe@gmail.com', '123456', 'New York City', '081123456789', '2001-01-01'),
('Jane Doe', 'janedoe@gmail.com', '123456', 'Washington D. C.', '082234567891', '2002-02-02'),
('Sherlock Holmes', 'sherlockholmes@gmail.com', '123456', 'Baker Street', '083345678912', '2003-03-03'),
('John Watson', 'johnwatson@gmail.com', '123456', 'Baker Street', '084456789123', '2004-04-04'),
('James Moriarty', 'jamesmoriarty@gmail.com', '123456', 'London', '085567891234', '2005-05-05');

-- No. 4d
select * from books where title = 'A Game of Thrones';

-- No. 4e
select * from books where genre_id = 3;

-- No. 4f (Semua ID No. 1 di ganti jadi 2 dan seterusnya karena
-- setiap tabel sudah saya coba tambahkan 1 row data dan sudah saya hapus untuk tugas hari ini)
insert into loans (user_id, book_id, status_loan, deadline_date, date_loan, date_return) values
(2, 2, 'Loaned', '2024-06-30', '2024-06-07', '0001-01-01'),
(2, 3, 'Loaned', '2024-06-30', '2024-06-08', '0001-01-01'),
(2, 7, 'Loaned', '2024-06-30', '2024-06-09', '0001-01-01'),
(3, 12, 'Loaned', '2024-06-30', '2024-06-10', '0001-01-01'),
(4, 2, 'Loaned', '2024-06-30', '2024-06-11', '0001-01-01');

-- No. 4g
update loans set status_loan = 'Returned', date_return = '2024-06-08' where user_id = 2 and book_id = 2;

-- No. 4h
select u.name, b.title, l.status_loan, l.deadline_date, l.date_loan, l.date_return
from loans as l
join users as u on l.user_id = u.id
join books as b on l.book_id = b.id
where l.status_loan = 'Loaned';

-- No. 4i
select u.name, b.title, l.status_loan, l.deadline_date, l.date_loan, l.date_return
from loans as l
join users as u on l.user_id = u.id
join books as b on l.book_id = b.id
where l.status_loan = 'Returned';