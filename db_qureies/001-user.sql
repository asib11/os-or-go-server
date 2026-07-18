
-- database type: PostgreSQL
-- SMALLINT 16 bits
-- INT 32 bits
-- BIGINT 64 bits
-- SERIAL 32 bits auto-incrementing integer

-- REAL 32 bits
-- DOUBLE PRECISION 64 bits (best use for monetary values)
-- NUMERIC variable precision


-- CHAR(n) fixed length, blank padded
-- VARCHAR(n) variable length
-- TEXT variable unlimited length

-- boolean true/false

-- TIMESTAMP without time zone
-- TIMESTAMPTZ with time zone



CREATE TABLE users (
    id             SERIAL PRIMARY KEY,
    first_name     VARCHAR(100) NOT NULL,
    last_name      VARCHAR(100) NOT NULL,
    email          VARCHAR(255) NOT NULL UNIQUE,
    password       VARCHAR(255) NOT NULL,
    is_shop_owner  BOOLEAN NOT NULL DEFAULT FALSE,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

