-- +goose Up
-- +goose StatementBegin
INSERT INTO users (email, password_hash, created_at, updated_at) VALUES
    ('alice@example.com', 'hash1', current_timestamp, current_timestamp),
    ('bob@example.com', 'hash2', current_timestamp, current_timestamp);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO products (name, price, quantity, product_details, created_at, updated_at) VALUES
    ('Toyota Camry', 24000.00, 10, '{"description": "Reliable and fuel-efficient", "category": "sedan", "year": 2023}', current_timestamp, current_timestamp),
    ('Ford Mustang', 35000.00, 5, '{"description": "Iconic sports car", "category": "sports", "year": 2022}', current_timestamp, current_timestamp),
    ('Tesla Model 3', 39999.99, 8, '{"description": "Electric sedan with autopilot", "category": "electric", "year": 2023}', current_timestamp, current_timestamp),
    ('Chevrolet', 45000.00, 4, '{"description": "Powerful pickup truck", "category": "truck", "year": 2023}', current_timestamp, current_timestamp);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO purchases (user_id, product_id, quantity, purchase_date, created_at, updated_at) VALUES
    (1, 1, 1, current_timestamp, current_timestamp, current_timestamp),
    (2, 3, 1, current_timestamp, current_timestamp, current_timestamp);
-- +goose StatementEnd

-- +goose Down
