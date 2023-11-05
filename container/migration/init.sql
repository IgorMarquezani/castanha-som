CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  passwd TEXT NOT NULL
);

CREATE TABLE sessions (
  key_access UUID UNIQUE NOT NULL,
  expires_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  user_id UUID PRIMARY KEY,
  FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE products (
  category TEXT NOT NULL,
  brand TEXT NOT NULL,
  name TEXT NOT NULL,
  price MONEY,
  description TEXT,
  PRIMARY KEY(name)
);

CREATE TABLE addresses (
  receiver TEXT NOT NULL,
  street TEXT NOT NULL,
  hood TEXT NOT NULL,
  CEP TEXT NOT NULL,
  user_id UUID NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE carts (
  id UUID NOT NULL,
  user_id UUID NOT NULL,
  PRIMARY KEY(id),
  FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE cart_items (
  cart_id UUID NOT NULL,
  product_name INTEGER NOT NULL,
  quantity INTEGER NOT NULL,
  FOREIGN KEY(cart_id) REFERENCES carts(id),
  FOREIGN KEY(product_name) REFERENCES products(name),
  PRIMARY KEY(cart_id, product_name)
);


CREATE TABLE purchases (
  id UUID NOT NULL,
  user_id UUID NOT NULL,
  CEP TEXT NOT NULL,
  total MONEY,
  FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE purchased_item (
  purchase_id UUID NOT NULL,
  product TEXT NOT NULL,
  value MONEY,
  quantity INTEGER NOT NULL,
  FOREIGN KEY(purchase_id) REFERENCES purchases(id)
);
