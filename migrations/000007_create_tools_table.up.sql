CREATE TABLE IF NOT EXISTS tools (
      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
      code VARCHAR(255) NOT NULL UNIQUE,
      name VARCHAR(255) NOT NULL,
      description TEXT NULL,
      brand VARCHAR(255) NULL,
      category VARCHAR(100) NOT NULL DEFAULT 'primary'
        CHECK(
          category IN (
               'primary',
               'secondary',
               'additional',
               'special'
            )
          ),
      price NUMERIC(18,2) NOT NULL DEFAULT 0,
      usage_period INTEGER NOT NULL DEFAULT 1,
      usage_period_unit VARCHAR(4) DEFAULT 'Y'
                   CHECK ( usage_period_unit IN ('Y', 'M', 'W', 'D') ),
      photo_id BIGINT NULL REFERENCES upload_files(id),
      is_active BOOLEAN NOT NULL DEFAULT TRUE,

      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
      updated_at TIMESTAMPTZ NULL,
      deleted_at TIMESTAMPTZ NULL,
      created_by BIGINT NULL REFERENCES users(id),
      updated_by BIGINT NULL REFERENCES users(id),
      deleted_by BIGINT NULL REFERENCES users(id)
);

CREATE INDEX idx_tools_code
  ON tools(code);

CREATE INDEX idx_tools_name
  ON tools(name);

CREATE INDEX idx_tools_category
  ON tools(category);

CREATE INDEX idx_tools_photo_id
  ON tools(photo_id);

CREATE INDEX idx_tools_is_active
ON tools(is_active);
