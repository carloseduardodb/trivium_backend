-- +goose Up

INSERT INTO public.cryptocurrencies ("name", symbol) VALUES
('Bitcoin', 'BTC'),
('Ethereum', 'ETH'),
('Cardano', 'ADA'),
('Solana', 'SOL'),
('Polkadot', 'DOT'),
('Ripple', 'XRP'),
('Litecoin', 'LTC'),
('Chainlink', 'LINK'),
('Stellar', 'XLM'),
('Dogecoin', 'DOGE');

INSERT INTO public.users ("name", email, photo_path) VALUES
('Alice Johnson', 'alice@example.com', NULL),
('Bob Smith', 'bob@example.com', NULL),
('Charlie Brown', 'charlie@example.com', NULL),
('David Lee', 'david@example.com', NULL),
('Emma Wilson', 'emma@example.com', NULL);

INSERT INTO public.positions (user_id, crypto_currency, quantity, purchase_price, invested_amount, purchase_date, last_profit_price, status) VALUES
(1, 1, 0.5, 45000.00, 22500.00, '2024-03-01 10:00:00', 48000.00, 'active'),
(2, 2, 10.0, 3000.00, 30000.00, '2024-03-02 11:00:00', 3500.00, 'active'),
(3, 3, 1000.0, 1.20, 1200.00, '2024-03-03 12:00:00', 1.50, 'active'),
(4, 4, 200.0, 150.00, 30000.00, '2024-03-04 13:00:00', 180.00, 'active'),
(5, 5, 500.0, 40.00, 20000.00, '2024-03-05 14:00:00', 45.00, 'active'),
(1, 6, 1000.0, 0.80, 800.00, '2024-03-06 15:00:00', 1.00, 'active'),
(2, 7, 50.0, 200.00, 10000.00, '2024-03-07 16:00:00', 220.00, 'active'),
(3, 8, 150.0, 25.00, 3750.00, '2024-03-08 17:00:00', 28.00, 'active'),
(4, 9, 2000.0, 0.10, 200.00, '2024-03-09 18:00:00', 0.12, 'active'),
(5, 10, 5000.0, 0.05, 250.00, '2024-03-10 19:00:00', 0.07, 'active');

INSERT INTO public.profit_takes ("position", amount_withdrawn, price_at_withdraw, remaining_value, withdraw_date) VALUES
(1, 5000.00, 50000.00, 17500.00, '2024-03-05 09:00:00'),
(2, 10000.00, 3500.00, 20000.00, '2024-03-06 10:00:00'),
(3, 200.00, 1.50, 1000.00, '2024-03-07 11:00:00'),
(4, 3000.00, 180.00, 27000.00, '2024-03-08 12:00:00'),
(5, 2500.00, 45.00, 17500.00, '2024-03-09 13:00:00');

-- +goose Down

DELETE FROM public.profit_takes;
DELETE FROM public.positions;
DELETE FROM public.users;
DELETE FROM public.cryptocurrencies;
