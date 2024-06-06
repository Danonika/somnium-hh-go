-- +goose Up
-- +goose StatementBegin
CREATE TABLE status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO status (name) VALUES ('Active'), ('Inactive'), ('Deleted');

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO category (name) VALUES
('Software Development'),
('Network Administration'),
('Cybersecurity'),
('Database Management'),
('Cloud Computing'),
('IT Support and Helpdesk'),
('Web Development'),
('Mobile App Development'),
('System Administration'),
('Data Science and Analytics');

CREATE TABLE job (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    salary VARCHAR(255),
    categoryid int NOT NULL,
    description TEXT,
    location TEXT,
    contact_number VARCHAR(255),
    contact_email VARCHAR(255),
    skills TEXT[],
    date_posted TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expire_date TIMESTAMP,
    statusid int,
    applicants_count INT,
    FOREIGN KEY (categoryid) REFERENCES category(id),
    FOREIGN KEY (statusid) REFERENCES status(id)
);

CREATE INDEX job_categoryid_idx ON job (categoryid);
CREATE INDEX job_statusid_idx ON job (statusid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE job;

DROP TABLE category;

DROP TABLE status;
-- +goose StatementEnd
