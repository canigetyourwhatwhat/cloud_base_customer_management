-- 記事データを格納するためのテーブル
create table if not exists customers (
                                        id integer unsigned auto_increment primary key,
                                        customer_id int(10) not null default 0,
                                        first_name varchar(255) not null default "",
                                        last_name varchar(255) not null default "",
                                        company_name varchar(255) not null default "",
                                        email varchar(255) not null default "",
                                        created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


#
# -- seed data
# insert into customers
# (customer_id, first_name, last_name, company_name, email) values
#     (1, 'Harry', 'Potter', 'Hogwarts', 'gryffindor@gmail.com'),
#     (2, 'Robert', 'Martin', 'Clean Architecture', 'unclebob@gmail.com');
#

