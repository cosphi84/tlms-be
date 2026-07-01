CREATE TABLE users (
     id BIGSERIAL PRIMARY KEY,
     email VARCHAR(255) NOT NULL UNIQUE,
     password TEXT NOT NULL,
     name VARCHAR(255) NOT NULL,
     image TEXT NULL,
     office_id BIGINT NOT NULL REFERENCES offices(id) ON DELETE RESTRICT,
     is_active BOOLEAN NOT NULL DEFAULT TRUE,
     failed_login_attempts INTEGER NOT NULL DEFAULT 0,

     locked_until TIMESTAMPTZ NULL,
     last_login_at TIMESTAMPTZ NULL,
     last_login_from INET NULL,

     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
     updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,

     created_by BIGINT NULL REFERENCES users(id),
     updated_by BIGINT NULL REFERENCES users(id),
    deleted_by BIGINT NULL REFERENCES users(id)
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_office_id ON users(office_id);
CREATE INDEX idx_users_is_active ON users(is_active);
