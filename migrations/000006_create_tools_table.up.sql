CREATE TABLE tools (
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
      photo_url TEXT NULL,
      is_active BOOLEAN NOT NULL DEFAULT TRUE,
      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
      updated_at TIMESTAMPTZ NULL,
      deleted_at TIMESTAMPTZ NULL,
      created_by INTEGER NULL REFERENCES users(id),
      updated_by INTEGER NULL REFERENCES users(id)
);

CREATE INDEX idx_tools_code
  ON tools(code);

CREATE INDEX idx_tools_name
  ON tools(name);

CREATE INDEX idx_tools_category
  ON tools(category);
