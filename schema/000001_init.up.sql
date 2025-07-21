CREATE TABLE subs
(
    id serial primary key not null unique,
    user_id varchar(255) not null,
    service_name varchar(255) not null,
    price int,
    start_date DATE,
    end_date DATE
);