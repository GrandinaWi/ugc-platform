CREATE TABLE IF NOT EXISTS comments (
                                     id BIGSERIAL PRIMARY KEY,
                                     post_id BIGINT NOT NULL,
                                     author_id BIGINT NOT NULL,
                                     parent_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
                             );
