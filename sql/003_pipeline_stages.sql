-- Pipeline Stages: Define deal progression stages
CREATE TABLE IF NOT EXISTS pipeline_stages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    display_order INTEGER NOT NULL,
    probability INTEGER DEFAULT 0 CHECK (probability >= 0 AND probability <= 100),
    color VARCHAR(7) DEFAULT '#6B7280', -- hex color code
    is_active BOOLEAN DEFAULT true,
    is_closed_won BOOLEAN DEFAULT false,
    is_closed_lost BOOLEAN DEFAULT false,
    
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    
    CONSTRAINT unique_stage_order UNIQUE(display_order)
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_pipeline_stages_order ON pipeline_stages(display_order);
CREATE INDEX IF NOT EXISTS idx_pipeline_stages_active ON pipeline_stages(is_active);

-- Insert default pipeline stages
INSERT INTO pipeline_stages (name, display_order, probability, color, is_closed_won, is_closed_lost) VALUES
    ('Prospecting', 1, 10, '#6B7280', false, false),
    ('Qualification', 2, 25, '#3B82F6', false, false),
    ('Proposal', 3, 50, '#F59E0B', false, false),
    ('Negotiation', 4, 75, '#8B5CF6', false, false),
    ('Closed Won', 5, 100, '#10B981', true, false),
    ('Closed Lost', 6, 0, '#EF4444', false, true)
ON CONFLICT (display_order) DO NOTHING;

-- Trigger for updated_at
CREATE OR REPLACE FUNCTION update_pipeline_stages_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_pipeline_stages_updated_at
    BEFORE UPDATE ON pipeline_stages
    FOR EACH ROW
    EXECUTE FUNCTION update_pipeline_stages_updated_at();
