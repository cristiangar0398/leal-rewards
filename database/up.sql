-- Eliminar tablas si ya existen
DROP TABLE IF EXISTS campaigns;
DROP TABLE IF EXISTS points;
DROP TABLE IF EXISTS cashback;
DROP TABLE IF EXISTS trades;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS transactions;


-- Crear tabla de usuarios
CREATE TABLE users (
  id VARCHAR(32) PRIMARY KEY,
  password VARCHAR(255) NOT NULL,
  document VARCHAR(32) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Crear tabla de comercios
CREATE TABLE trades (
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  user_id VARCHAR(32) NOT NULL,
  conversion_rate INT NOT NULL DEFAULT 1000, -- Factor de conversión por defecto
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Crear tabla de transacciones
CREATE TABLE transactions (
  id VARCHAR(32) PRIMARY KEY,
  user_id VARCHAR(32) NOT NULL,
  trade_id VARCHAR(32) NOT NULL,
  amount DECIMAL(10, 2) NOT NULL, -- Monto de la transacción
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

-- Crear tabla de cashback
CREATE TABLE cashback (
  id VARCHAR(32) PRIMARY KEY,
  user_id VARCHAR(32) NOT NULL,
  amount DECIMAL(10, 2) NOT NULL, -- Monto de cashback acumulado
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Crear tabla de campañas
CREATE TABLE campaigns (
  id VARCHAR(32) PRIMARY KEY,
  trade_id VARCHAR(32) NOT NULL,
  branch_name VARCHAR(255) NOT NULL, -- Nombre de la sucursal
  start_date TIMESTAMP NOT NULL,
  end_date TIMESTAMP NOT NULL,
  bonus_type VARCHAR(10) NOT NULL, -- Puede ser 'double' o 'percentage'
  bonus_value DECIMAL(5, 2) NOT NULL, -- Por ejemplo, 2.0 para el doble de puntos o 1.3 para un 30% adicional
  min_purchase_amount DECIMAL(10, 2), -- Monto mínimo de compra para aplicar la campaña
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY (trade_id) REFERENCES trades(id)
);
