-- +goose Up
-- +goose StatementBegin
CREATE TABLE promo_stats (
    promo_id UInt64,
    date Date,
    views UInt32,
    likes UInt32,
    comments UInt32
) ENGINE = SummingMergeTree()
PARTITION BY toYYYYMM(date)
ORDER BY (promo_id, date);

CREATE TABLE post_stats (
    post_id UInt64,
    user_id UInt64,
    date Date,
    views UInt32,
    likes UInt32,
    comments UInt32
) ENGINE = SummingMergeTree()
PARTITION BY toYYYYMM(date)
ORDER BY (post_id, date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE post_stats;
DROP TABLE promo_stats;
-- +goose StatementEnd
