CREATE TABLE casbin_rule
(
    id BIGSERIAL PRIMARY KEY,

    ptype VARCHAR(16) NOT NULL,

    v0 VARCHAR(256),
    v1 VARCHAR(256),
    v2 VARCHAR(256),
    v3 VARCHAR(256),
    v4 VARCHAR(256),
    v5 VARCHAR(256),

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_casbin_ptype
ON casbin_rule(ptype);

CREATE INDEX idx_casbin_lookup
ON casbin_rule(ptype,v0,v1,v2);
