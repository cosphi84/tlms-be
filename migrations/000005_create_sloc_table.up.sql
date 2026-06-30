CREATE TABLE storage_locations (
    id BIGSERIAL PRIMARY KEY,
    code VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    office_id BIGINT NOT NULL REFERENCES offices(id) ON DELETE RESTRICT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    created_by BIGINT NULL REFERENCES users(id),
    updated_by BIGINT NULL REFERENCES users(id),
    deleted_at TIMESTAMPZ NULL,
    deleted_by BIGINT NULL REFERENCES users(id),
);

CREATE INDEX idx_storage_locs_code ON storage_locations(code);
CREATE INDEX idx_storage_locs_office_id ON storage_locations(office_id);
