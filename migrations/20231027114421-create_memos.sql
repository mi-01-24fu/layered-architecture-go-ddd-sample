-- +migrate Up
create table if not exists memos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    date DATE
);

-- +migrate Down
drop table if exists memos;