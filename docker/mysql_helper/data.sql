ALTER TABLE user AUTO_INCREMENT = 1;
ALTER TABLE task AUTO_INCREMENT = 1;

insert into user (email, name, provider_id, avatar_url) values ('user1@gmail.com', 'user1', '1', '');
insert into user (email, name, provider_id, avatar_url) values ('user2@gmail.com', 'user2', '2', '');
insert into user (email, name, provider_id, avatar_url) values ('user3@gmail.com', 'user3', '3', '');
insert into user (email, name, provider_id, avatar_url) values ('user4@gmail.com', 'user4', '4', '');
insert into user (email, name, provider_id, avatar_url) values ('user5@gmail.com', 'user5', '5', '');
insert into user (email, name, provider_id, avatar_url) values ('user6@gmail.com', 'user6', '6', '');
insert into user (email, name, provider_id, avatar_url) values ('user7@gmail.com', 'user7', '7', '');
insert into user (email, name, provider_id, avatar_url) values ('user8@gmail.com', 'user8', '8', '');
insert into user (email, name, provider_id, avatar_url) values ('user9@gmail.com', 'user9', '9', '');

insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 10, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 10, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 20, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 30, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 40, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 50, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 60, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 70, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 80, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 90, 1);



