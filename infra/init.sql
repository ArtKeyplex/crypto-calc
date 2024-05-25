CREATE TABLE IF NOT EXISTS public.currencies (
                                          id SERIAL PRIMARY KEY,
                                          name VARCHAR(100) NOT NULL,
                                          code VARCHAR(10) UNIQUE NOT NULL,
                                          available BOOLEAN NOT NULL DEFAULT TRUE,
                                          rate DECIMAL(20, 5) NOT NULL,
                                          type VARCHAR(10) NOT NULL
);

INSERT INTO public.currencies (name, code, available, rate, type)
VALUES
    ('Euro', 'EUR', true, 1.00000, 'fiat'),
    ('US Dollar', 'USD', true, 1.00000, 'fiat'),
    (
        'Chinese Yuan', 'CNY', true, 1.00000, 'fiat'
    ),
    ('Tether', 'USDT', true, 1.00000, 'crypto'),
    ('USD Coin', 'USDC', true, 1.00000, 'crypto'),
    ('Ethereum', 'ETH', true, 1.00000, 'crypto') ON CONFLICT (code) DO NOTHING;
