--- This file is used to create the tables and insert the initial data

--- for storing the users
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    username text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    password text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

--- for storing e-wallets
CREATE TABLE IF NOT EXISTS ewallets (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    owner_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    balance NUMERIC(19, 4) NOT NULL DEFAULT 0.00, -- Assuming 4 decimal places for precision
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

--- for storing stock information
CREATE TABLE IF NOT EXISTS stocks (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    symbol text NOT NULL UNIQUE,
    name text NOT NULL,
    description text,
    current_price NUMERIC(19, 4) NOT NULL DEFAULT 0.00, -- Current market price
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

--- for storing user stock holdings
CREATE TABLE IF NOT EXISTS user_stock_holdings (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    stock_id uuid NOT NULL REFERENCES stocks(id) ON DELETE CASCADE,
    quantity NUMERIC(19, 8) NOT NULL DEFAULT 0.00000000, -- Allowing for fractional shares up to 8 decimal places
    average_purchase_price NUMERIC(19, 4) NOT NULL DEFAULT 0.00,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    UNIQUE (user_id, stock_id) -- Ensures a user has one holding record per stock
);

--- for storing e-wallet transactions
CREATE TABLE IF NOT EXISTS ewallet_transactions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    wallet_id uuid NOT NULL REFERENCES ewallets(id) ON DELETE CASCADE,
    transaction_type text NOT NULL, -- e.g., 'DEPOSIT', 'WITHDRAWAL', 'STOCK_PURCHASE', 'STOCK_SALE'
    amount NUMERIC(19, 4) NOT NULL,
    description text,
    related_stock_id uuid REFERENCES stocks(id) ON DELETE SET NULL, -- Nullable, for stock transactions
    price_at_transaction NUMERIC(19, 4), -- Price per share at the time of stock transaction
    quantity_transacted NUMERIC(19, 8), -- Quantity of stock bought/sold
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

-- Indexes for foreign keys to improve query performance
CREATE INDEX IF NOT EXISTS idx_ewallets_owner_id ON ewallets(owner_id);
CREATE INDEX IF NOT EXISTS idx_user_stock_holdings_user_id ON user_stock_holdings(user_id);
CREATE INDEX IF NOT EXISTS idx_user_stock_holdings_stock_id ON user_stock_holdings(stock_id);
CREATE INDEX IF NOT EXISTS idx_ewallet_transactions_wallet_id ON ewallet_transactions(wallet_id);
CREATE INDEX IF NOT EXISTS idx_ewallet_transactions_related_stock_id ON ewallet_transactions(related_stock_id);

-- Insert initial stock data
INSERT INTO stocks (symbol, name, description, current_price) VALUES
('TSLA', 'Tesla, Inc.', 'Electric vehicle and clean energy company', 170.00) ON CONFLICT (symbol) DO NOTHING;

INSERT INTO stocks (symbol, name, description, current_price) VALUES
('GOOG', 'Alphabet Inc. (Class C)', 'Multinational technology conglomerate', 175.00) ON CONFLICT (symbol) DO NOTHING;

