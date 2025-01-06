-- +goose Up
-- +goose StatementBegin
-- Create Users table
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    photo_path VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Cryptocurrencies table
CREATE TABLE cryptocurrencies (
    id BIGSERIAL PRIMARY KEY,
    user BIGINT NOT NULL REFERENCES users(id),
    name VARCHAR(255) NOT NULL,
    symbol VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(symbol)
);

-- Create Positions table
CREATE TABLE positions (
    id BIGSERIAL PRIMARY KEY,
    user BIGINT NOT NULL REFERENCES users(id),
    crypto_currency BIGINT NOT NULL REFERENCES cryptocurrencies(id),
    quantity DECIMAL(20,8) NOT NULL,
    purchase_price DECIMAL(20,2) NOT NULL,
    invested_amount DECIMAL(20,2) NOT NULL,
    purchase_date TIMESTAMP WITH TIME ZONE NOT NULL,
    last_profit_price DECIMAL(20,2),
    status VARCHAR(20) NOT NULL CHECK (status IN ('active', 'closed')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Profit_Takes table
CREATE TABLE profit_takes (
    id BIGSERIAL PRIMARY KEY,
    position BIGINT NOT NULL REFERENCES positions(id),
    amount_withdrawn DECIMAL(20,2) NOT NULL,
    price_at_withdraw DECIMAL(20,2) NOT NULL,
    remaining_value DECIMAL(20,2) NOT NULL,
    withdraw_date TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_cryptocurrencies_user ON cryptocurrencies(user);
CREATE INDEX idx_positions_user ON positions(user);
CREATE INDEX idx_positions_crypto ON positions(crypto_currency);
CREATE INDEX idx_profit_takes_position ON profit_takes(position);
CREATE INDEX idx_positions_status ON positions(status);

-- Trigger for updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Add triggers to all tables
CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_cryptocurrencies_updated_at
    BEFORE UPDATE ON cryptocurrencies
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_positions_updated_at
    BEFORE UPDATE ON positions
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_profit_takes_updated_at
    BEFORE UPDATE ON profit_takes
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
