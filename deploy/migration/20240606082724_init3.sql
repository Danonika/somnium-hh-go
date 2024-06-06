-- +goose Up
-- +goose StatementBegin
CREATE TABLE job_applications (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    job_id INT NOT NULL,
    applied_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (job_id) REFERENCES job(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE job_applications;
-- +goose StatementEnd
