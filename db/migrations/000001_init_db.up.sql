CREATE TABLE posts(
    id BIGSERIAL PRIMARY KEY,
    author VARCHAR(100) NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    comments_disabled BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_posts_id ON posts (id);

CREATE TABLE comments(
    id BIGSERIAL PRIMARY KEY,
    author VARCHAR(100) NOT NULL,
    post_id BIGINT NOT NULL,
    parent_id BIGINT NULL,
    content VARCHAR(2000) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (post_id) REFERENCES posts (id)
);

CREATE INDEX idx_comments_id ON comments (id);

CREATE OR REPLACE FUNCTION check_parent_id()
RETURNS TRIGGER AS $$
BEGIN
    IF COALESCE(NEW.parent_id, 0) > (SELECT COALESCE(MAX(id), 0) FROM comments) THEN
        RAISE EXCEPTION 'parent_id cannot be greater than the maximum id in the comments table';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_check_parent_id
BEFORE INSERT OR UPDATE ON comments
FOR EACH ROW EXECUTE FUNCTION check_parent_id();
