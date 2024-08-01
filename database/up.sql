-- Eliminar tablas si ya existen
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS trades;

-- Crear tabla de usuarios
CREATE TABLE users (
  id VARCHAR(32) PRIMARY KEY,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Crear tabla de comercios
CREATE TABLE trades (
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  user_id VARCHAR(32) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id)
);