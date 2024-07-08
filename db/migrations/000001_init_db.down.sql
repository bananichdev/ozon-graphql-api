DROP INDEX IF EXISTS idx_comments_id;
DROP TRIGGER IF EXISTS trigger_check_parent_id ON comments;
DROP FUNCTION IF EXISTS check_parent_id();
DROP TABLE IF EXISTS comments;

DROP INDEX IF EXISTS idx_posts_id;
DROP TABLE IF EXISTS posts;
