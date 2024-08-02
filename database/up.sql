-- Eliminar tablas si ya existen
DROP TABLE IF EXISTS campaigns;
DROP TABLE IF EXISTS points;
DROP TABLE IF EXISTS cashback;
DROP TABLE IF EXISTS trades;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS transactions;



CREATE TABLE users (
  id VARCHAR(32) PRIMARY KEY,
  password VARCHAR(255) NOT NULL,
  document VARCHAR(32) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE trades (
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  user_id VARCHAR(32) NOT NULL,
  conversion_rate INT NOT NULL DEFAULT 1000, 
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id)
);


CREATE TABLE transactions (
  id VARCHAR(32) PRIMARY KEY,
  user_id VARCHAR(32) NOT NULL,
  trade_id VARCHAR(32) NOT NULL,
  amount DECIMAL(10, 2) NOT NULL, 
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (trade_id) REFERENCES trades(id)
);

-- Crear tabla de puntos
CREATE TABLE points (
  id VARCHAR(32) PRIMARY KEY,
  user_id VARCHAR(32) NOT NULL,
  trade_id VARCHAR(32) NOT NULL,
  points INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (trade_id) REFERENCES trades(id)
);


CREATE TABLE cashback (
  id VARCHAR(32) PRIMARY KEY,
  user_id VARCHAR(32) NOT NULL,
  amount DECIMAL(10, 2) NOT NULL, 
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Crear tabla de campañas
CREATE TABLE campaigns (
  id VARCHAR(32) PRIMARY KEY,
  trade_id VARCHAR(32) NOT NULL,
  branch_name VARCHAR(255) NOT NULL, 
  start_date TIMESTAMP NOT NULL,
  end_date TIMESTAMP NOT NULL,
  bonus_type VARCHAR(10) NOT NULL,
  bonus_value DECIMAL(5, 2) NOT NULL, 
  min_purchase_amount DECIMAL(10, 2), 
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY (trade_id) REFERENCES trades(id)
);
