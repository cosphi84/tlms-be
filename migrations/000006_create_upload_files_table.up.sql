CREATE TABLE IF NOT EXISTS upload_files (
              id BIGSERIAL PRIMARY KEY,
              uuid UUID NOT NULL DEFAULT gen_random_uuid(),

              disk_name VARCHAR(255) NOT NULL,
              original_name VARCHAR(255) NOT NULL,
              mime_type VARCHAR(100) NOT NULL,
              extension VARCHAR(20) NOT NULL,
              size BIGINT NOT NULL,
              checksum VARCHAR(64) NOT NULL,

              path VARCHAR(500) NOT NULL,
              storage VARCHAR(50) NOT NULL DEFAULT 'local',

              is_archived BOOLEAN NOT NULL DEFAULT FALSE,
              archived_path VARCHAR(500) NULL,

              created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
              updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
              created_by BIGINT NULL REFERENCES users(id),
              updated_by BIGINT NULL REFERENCES users(id),
              deleted_at TIMESTAMPTZ NULL,
              deleted_by BIGINT NULL REFERENCES users(id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_upload_files_uuid ON upload_files(uuid);
CREATE INDEX IF NOT EXISTS idx_upload_files_checksum ON upload_files(checksum);
CREATE INDEX IF NOT EXISTS idx_upload_files_is_archived ON upload_files(is_archived);
CREATE INDEX IF NOT EXISTS idx_upload_files_deleted_at ON upload_files(deleted_at);
