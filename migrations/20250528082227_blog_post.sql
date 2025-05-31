-- +goose Up
-- +goose StatementBegin
CREATE TABLE blog_post (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    author VARCHAR(100),
    published BOOLEAN DEFAULT FALSE,
    slug VARCHAR(255) UNIQUE NOT NULL,
    descriptions TEXT[] NOT NULL,
    meta_description TEXT
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON blog_post
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE INDEX idx_blog_post_descriptions ON blog_post USING GIN (descriptions);
CREATE INDEX idx_blog_post_published ON blog_post (published);
CREATE INDEX idx_blog_post_created_at ON blog_post (created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS trigger_update_updated_at ON blog_post;
DROP FUNCTION IF EXISTS update_updated_at_column;
DROP INDEX IF EXISTS idx_blog_post_descriptions;
DROP INDEX IF EXISTS idx_blog_post_published;
DROP INDEX IF EXISTS idx_blog_post_created_at;
DROP TABLE IF EXISTS blog_post;
-- +goose StatementEnd
