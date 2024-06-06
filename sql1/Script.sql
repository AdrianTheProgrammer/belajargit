create table users (
id serial primary key not null,
name varchar(100) not null,
email varchar(100) not null,
password varchar(100) not null,
address varchar(100) not null,
phone varchar(20) not null,
birth_date time not null);

create table loans (
id serial primary key not null,
user_id int not null,
book_id int not null,
deadline_date time not null,
date_loan time not null,
date_return time not null,
status_loan varchar(20) not null);

create table books (
id serial primary key not null,
genre_id int not null,
title varchar(100) not null,
author varchar(100) not null,
publisher varchar(100) not null,
publish_year int not null);

create table genres (
id serial primary key not null,
genre_name varchar(20) not null);

create table book_request (
id serial primary key not null,
user_id int not null,
approbed_admin_id int not null,
title varchar(100) not null,
author varchar(100) not null,
publisher varchar(100) not null,
publish_year int not null,
status_request varchar(20) not null);

create table admins (
id serial primary key not null,
name varchar(100) not null,
email varchar(100) not null,
password varchar(100) not null,
phone varchar(20) not null);

alter table loans add constraint fk_user_id foreign key (user_id) references users(id);
alter table loans add constraint fk_book_id foreign key (book_id) references books(id);
alter table books add constraint fk_genre_id foreign key (genre_id) references genres(id);
alter table book_request add constraint fk_user_id foreign key (user_id) references users(id);
alter table book_request rename column approbed_admin_id to approved_admin_id;
alter table book_request add constraint fk_approved_admin_id foreign key (approved_admin_id) references admins(id);
alter table users drop column birth_date;
alter table users add birth_date date;

insert into users (name, email, password, address, phone, birth_date) values ('admin1', 'admin1@admin.com', '12345', 'Bandung', '08123456789', '1999-01-01');
update users set name = 'John Doe', email = 'johndoe@gmail.com' where id = 1;
insert into genres (genre_name) values ('fantasy');
insert into books (genre_id, title, author, publisher, publish_year) values (1, 'A Game of Thrones', 'George R. R. Martin', 'Bantam Books', '1996');
alter table loans drop column deadline_date;
alter table loans drop column date_loan;
alter table loans drop column date_return;
alter table loans add deadline_date date;
alter table loans add date_loan date;
alter table loans add date_return date;
insert into loans (user_id, book_id, deadline_date, date_loan, date_return, status_loan) values (1, 1, '2024-06-30', '2024-06-06', '0001-01-01', 'Loaned');
insert into admins (name, email, password, phone) values ('admin1', 'admin1@gmail.com', '12345', '089987654321');
insert into book_request (user_id, approved_admin_id, title, author, publisher, publish_year, status_request) values (1, 1, 'Naruto', 'Masashi Kishimoto', 'Weekly Shonen Jump', '1999', 'Approved');