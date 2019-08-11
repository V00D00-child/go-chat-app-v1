-- Create database
create database chattie;
use chattie;

-- Create table
create table App_User
(
  User_ID           BIGINT not null,
  User_Name         VARCHAR(36) not null,
  User_Email VARCHAR(100) not null,
  User_Password VARCHAR(100) not null,
  Active           BOOLEAN NOT NULL,
  Online_Status  BOOLEAN NOT NULL 
) ;

--  Add primary key
alter table App_User
  add constraint App_User primary key (USER_ID);

-- Add sample users
insert into App_User (User_ID, User_Name,User_Email, User_Password, Active,Online_Status)
values (1, 'user1', 'user1@gmail.com','$2a$04$O9QVBQaIKxy9byajg1xfRuAf9PZ1vfliqZf6YiCtbkYDwKfwgRVhC', 1,0);
 
insert into App_User (User_ID, User_Name, User_Email,User_Password, Active,Online_Status)
values (2, 'user2', 'user2@gmail.com','$2a$04$rVe/Zbu1xpD8Hrt/0CaE8uto6tB3mI1ueLp7vPlRWz1s2ZUhoy.0e', 1,0);

-- Get all current users in database
select * from App_User;

