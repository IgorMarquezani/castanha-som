CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  passwd TEXT NOT NULL
);

CREATE TABLE user_features (
  admin_access BOOLEAN NOT NULL DEFAULT false,
  user_id UUID NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id)
  PRIMARY KEY(user_id)
);

CREATE TABLE sessions (
  key_access UUID UNIQUE NOT NULL,
  expires_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  user_id UUID PRIMARY KEY,
  FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE products (
  name TEXT NOT NULL,
  type TEXT NOT NULL,
  in_cash_value DECIMAL(1000, 2) NOT NULL,
  installment_value DECIMAL(1000, 2) NOT NULL,
  image_name TEXT NOT NULL,
  PRIMARY KEY(name)
);

CREATE TABLE product_descriptions (
  product_name TEXT NOT NULL,
  value TEXT NOT NULL,
  FOREIGN KEY(product_name) REFERENCES products(name),
  PRIMARY KEY(product_name, value)
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
  product_name TEXT NOT NULL,
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
