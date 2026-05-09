-- +goose Up
-- +goose StatementBegin
CREATE TABLE price_alerts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    crypto_currency BIGINT NOT NULL REFERENCES cryptocurrencies(id),
    symbol VARCHAR(50) NOT NULL,
    target_price DECIMAL(20,2) NOT NULL,
    direction VARCHAR(10) NOT NULL CHECK (direction IN ('above', 'below')),
    active BOOLEAN NOT NULL DEFAULT true,
    triggered_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE price_alerts IS 'Tabela que armazena alertas de preço dos usuários';
COMMENT ON COLUMN price_alerts.direction IS 'Direção do alerta: above (acima) ou below (abaixo)';
COMMENT ON COLUMN price_alerts.active IS 'Se o alerta está ativo ou já foi disparado';
COMMENT ON COLUMN price_alerts.triggered_at IS 'Data em que o alerta foi disparado';

CREATE INDEX idx_price_alerts_user ON price_alerts(user_id);
CREATE INDEX idx_price_alerts_active ON price_alerts(active) WHERE active = true;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS price_alerts;
-- +goose StatementEnd
