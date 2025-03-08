-- +goose Up
create table users (
  id serial primary key,
  email text unique not null,
  password_hash text not null,
  created_at timestamptz default current_timestamp
);

create table products (
  id serial primary key,
  name text not null,
  price decimal(10, 2) not null,
  quantity int not null,
  product_details jsonb,
  created_at timestamptz default current_timestamp,
  updated_at timestamptz default current_timestamp
);

create table purchases (
  id serial primary key,
  user_id int not null,
  product_id int not null,
  purchase_date timestamptz default current_timestamp,
  quantity int not null,
  created_at timestamptz default current_timestamp,
  foreign key (user_id) references users(id) on delete cascade,
  foreign key (product_id) references products(id) on delete cascade
);


-- +goose Down
drop table if exists purchases;
drop table if exists products;
drop table if exists users;
