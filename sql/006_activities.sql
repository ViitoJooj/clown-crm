-- Activities table: Activity history and timeline
CREATE TABLE IF NOT EXISTS activities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Activity details
    activity_type VARCHAR(50) NOT NULL, -- 'call', 'email', 'meeting', 'note', 'status_change', 'deal_stage_change'
    title VARCHAR(250) NOT NULL,
    description TEXT,
    duration_minutes INTEGER,
    
    -- Related entities
    contact_id UUID REFERENCES contacts(id) ON DELETE CASCADE,
    deal_id UUID REFERENCES deals(id) ON DELETE CASCADE,
    company_id UUID REFERENCES companies(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(uuid) ON DELETE SET NULL,
    
    -- Additional metadata (flexible JSON storage)
    metadata JSONB DEFAULT '{}',
    
    -- Timestamp
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_activities_contact ON activities(contact_id);
CREATE INDEX IF NOT EXISTS idx_activities_deal ON activities(deal_id);
CREATE INDEX IF NOT EXISTS idx_activities_company ON activities(company_id);
CREATE INDEX IF NOT EXISTS idx_activities_user ON activities(user_id);
CREATE INDEX IF NOT EXISTS idx_activities_type ON activities(activity_type);
CREATE INDEX IF NOT EXISTS idx_activities_created ON activities(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_activities_metadata ON activities USING gin(metadata);

-- View for activity feed with related entity names
CREATE OR REPLACE VIEW activity_feed AS
SELECT 
    a.id,
    a.activity_type,
    a.title,
    a.description,
    a.duration_minutes,
    a.contact_id,
    CONCAT(c.first_name, ' ', c.last_name) as contact_name,
    a.deal_id,
    d.title as deal_title,
    a.company_id,
    comp.name as company_name,
    a.user_id,
    CONCAT(u.first_name, ' ', u.last_name) as user_name,
    a.metadata,
    a.created_at
FROM activities a
LEFT JOIN contacts c ON a.contact_id = c.id
LEFT JOIN deals d ON a.deal_id = d.id
LEFT JOIN companies comp ON a.company_id = comp.id
LEFT JOIN users u ON a.user_id = u.uuid
ORDER BY a.created_at DESC;
