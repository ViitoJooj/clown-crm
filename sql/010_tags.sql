-- Tags table: Reusable tags for categorization
CREATE TABLE IF NOT EXISTS tags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    color VARCHAR(7) DEFAULT '#6B7280', -- hex color code
    entity_type VARCHAR(50) NOT NULL, -- 'contact', 'deal', 'company', 'all'
    
    -- Usage tracking
    usage_count INTEGER DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    
    -- Constraints
    CONSTRAINT unique_tag_per_entity UNIQUE(name, entity_type)
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_tags_entity_type ON tags(entity_type);
CREATE INDEX IF NOT EXISTS idx_tags_name ON tags(name);
CREATE INDEX IF NOT EXISTS idx_tags_usage ON tags(usage_count DESC);

-- Common tags for quick start
INSERT INTO tags (name, color, entity_type) VALUES
    ('Hot Lead', '#EF4444', 'contact'),
    ('Cold Lead', '#3B82F6', 'contact'),
    ('VIP', '#F59E0B', 'contact'),
    ('High Priority', '#EF4444', 'deal'),
    ('Quick Win', '#10B981', 'deal'),
    ('Long Term', '#8B5CF6', 'deal'),
    ('Strategic', '#F59E0B', 'company'),
    ('SMB', '#3B82F6', 'company'),
    ('Enterprise', '#8B5CF6', 'company')
ON CONFLICT (name, entity_type) DO NOTHING;
