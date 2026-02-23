CREATE TABLE api_keys (
    key_hash      VARCHAR(64) PRIMARY KEY,  
    service_name  VARCHAR(100) NOT NULL,    
    is_active     BOOLEAN DEFAULT TRUE,     
    created_at    TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX idx_api_keys_active ON api_keys (is_active) WHERE is_active = TRUE;

CREATE TABLE notifications (
    id              BIGSERIAL,
    user_id         BIGINT NOT NULL,
    title           VARCHAR(255),
    body            TEXT,
    status          SMALLINT DEFAULT 0,
    source_service  VARCHAR(100),
    author_id       VARCHAR(100),
    payload         JSONB,               
    created_at      TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (id, created_at)
) PARTITION BY RANGE (created_at);
CREATE TABLE notifications_202602 PARTITION OF notifications
    FOR VALUES FROM ('2026-02-01') TO ('2026-03-01');
CREATE INDEX idx_notifications_user_date ON notifications (user_id, created_at DESC);
CREATE INDEX idx_notifications_status ON notifications (status) WHERE status = 0;