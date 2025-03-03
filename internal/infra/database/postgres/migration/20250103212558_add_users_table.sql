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

-- Comentários da tabela Users
COMMENT ON TABLE users IS 'Tabela que armazena os dados dos usuários';
COMMENT ON COLUMN users.id IS 'Identificador único do usuário';
COMMENT ON COLUMN users.name IS 'Nome do usuário';
COMMENT ON COLUMN users.email IS 'E-mail do usuário, deve ser único';
COMMENT ON COLUMN users.photo_path IS 'URL da foto do usuário';
COMMENT ON COLUMN users.created_at IS 'Data de criação do registro';
COMMENT ON COLUMN users.updated_at IS 'Data de atualização do registro, atualizado via função trigger';

-- Create Cryptocurrencies table
CREATE TABLE cryptocurrencies (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    symbol VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Comentários da tabela Cryptocurrencies
COMMENT ON TABLE cryptocurrencies IS 'Tabela que armazena as criptomoedas cadastradas pelos usuários';
COMMENT ON COLUMN cryptocurrencies.id IS 'Identificador único da criptomoeda';
COMMENT ON COLUMN cryptocurrencies.name IS 'Nome da criptomoeda';
COMMENT ON COLUMN cryptocurrencies.symbol IS 'Símbolo da criptomoeda, deve ser único';
COMMENT ON COLUMN cryptocurrencies.created_at IS 'Data de criação do registro';
COMMENT ON COLUMN cryptocurrencies.updated_at IS 'Data de atualização do registro, atualizado via função trigger';

-- Create Positions table
CREATE TABLE positions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
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

-- Comentários da tabela Positions
COMMENT ON TABLE positions IS 'Tabela que armazena as posições de investimento dos usuários';
COMMENT ON COLUMN positions.id IS 'Identificador único da posição';
COMMENT ON COLUMN positions.user_id IS 'ID do usuário';
COMMENT ON COLUMN positions.crypto_currency IS 'ID da criptomoeda associada';
COMMENT ON COLUMN positions.quantity IS 'Quantidade de criptomoeda na posição';
COMMENT ON COLUMN positions.purchase_price IS 'Preço de compra da criptomoeda';
COMMENT ON COLUMN positions.invested_amount IS 'Valor total investido';
COMMENT ON COLUMN positions.purchase_date IS 'Data de compra da criptomoeda';
COMMENT ON COLUMN positions.last_profit_price IS 'Preço de venda da criptomoeda';
COMMENT ON COLUMN positions.status IS 'Status da posição (ativo ou fechado)';
COMMENT ON COLUMN positions.created_at IS 'Data de criação do registro';
COMMENT ON COLUMN positions.updated_at IS 'Data de atualização do registro, atualizado via função trigger';

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

-- Comentários da tabela Profit_Takes
COMMENT ON TABLE profit_takes IS 'Tabela que armazena retiradas de lucro das posições dos usuários';
COMMENT ON COLUMN profit_takes.id IS 'Identificador único da retirada';
COMMENT ON COLUMN profit_takes.position IS 'ID da posição associada';
COMMENT ON COLUMN profit_takes.amount_withdrawn IS 'Quantidade de criptomoeda retirada';
COMMENT ON COLUMN profit_takes.price_at_withdraw IS 'Preço da criptomoeda no momento da retirada';
COMMENT ON COLUMN profit_takes.remaining_value IS 'Quantidade restante na posição após a retirada';
COMMENT ON COLUMN profit_takes.withdraw_date IS 'Data da retirada';
COMMENT ON COLUMN profit_takes.created_at IS 'Data de criação do registro';
COMMENT ON COLUMN profit_takes.updated_at IS 'Data de atualização do registro, atualizado via função trigger';

-- Indexes
CREATE INDEX idx_positions_user ON positions(user_id);
CREATE INDEX idx_positions_crypto ON positions(crypto_currency);
CREATE INDEX idx_profit_takes_position ON profit_takes(position);
CREATE INDEX idx_positions_active_status ON positions(status) WHERE status = 'active';

-- Trigger for updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

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
DROP TABLE IF EXISTS profit_takes;
DROP TABLE IF EXISTS positions;
DROP TABLE IF EXISTS cryptocurrencies;
DROP TABLE IF EXISTS users;
DROP FUNCTION IF EXISTS update_updated_at_column;
-- +goose StatementEnd
