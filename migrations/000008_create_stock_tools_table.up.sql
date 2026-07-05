CREATE TABLE IF NOT EXISTS stock_tools (
          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

          tool_id UUID NOT NULL REFERENCES tools(id),
          storage_loc_id BIGINT NOT NULL REFERENCES storage_locations(id),

          qty INTEGER NOT NULL DEFAULT 0 CHECK ( qty >= 0 ),
          sbin VARCHAR(255) NOT NULL DEFAULT '',

          reference_type VARCHAR(50) NOT NULL
            CHECK (
              reference_type IN (
                 'initial_stock',
                 'replacement',
                 'procurement',
                 'adjustment'
                )
              ) DEFAULT 'initial_stock',

          stock_counter INT NOT NULL DEFAULT 1,
          reference_number VARCHAR(255) NULL,

          expired_at TIMESTAMPTZ NULL DEFAULT NULL,
          created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
          created_by INTEGER NULL REFERENCES users(id),
          updated_at TIMESTAMPTZ NULL,
          updated_by INTEGER NULL REFERENCES users(id),
          deleted_at TIMESTAMPTZ NULL,
          deleted_by INTEGER NULL REFERENCES users(id),

          UNIQUE(tool_id, storage_loc_id)
);

CREATE INDEX idx_stock_tools_id
  ON stock_tools(tool_id);

CREATE INDEX idx_stock_s_loc_id
  ON stock_tools(storage_loc_id);
