-- Deals table: Sales opportunities and pipeline management
CREATE TABLE IF NOT EXISTS deals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(250) NOT NULL,
    value DECIMAL(15, 2) DEFAULT 0,
    currency VARCHAR(3) DEFAULT 'USD',
    expected_close_date DATE,
    
    -- Stage and probability
    stage_id UUID NOT NULL REFERENCES pipeline_stages(id) ON DELETE RESTRICT,
    probability INTEGER DEFAULT 0 CHECK (probability >= 0 AND probability <= 100),
    
    -- Related entities
    contact_id UUID REFERENCES contacts(id) ON DELETE SET NULL,
    company_id UUID REFERENCES companies(id) ON DELETE SET NULL,
    
    -- Assignment and ownership
    assigned_to UUID REFERENCES users(uuid) ON DELETE SET NULL,
    owner_id UUID REFERENCES users(uuid) ON DELETE SET NULL,
    
    -- Source and metadata
    source VARCHAR(50), -- e.g., 'inbound', 'outbound', 'referral'
    notes TEXT,
    custom_fields JSONB DEFAULT '{}',
    tags JSONB DEFAULT '[]',
    
    -- Status tracking
    is_won BOOLEAN DEFAULT false,
    is_lost BOOLEAN DEFAULT false,
    lost_reason VARCHAR(250),
    
    -- Timestamps
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    closed_at TIMESTAMP,
    
    -- Constraints
    CONSTRAINT valid_close_state CHECK (
        (is_won = false AND is_lost = false) OR 
        (is_won = true AND is_lost = false) OR 
        (is_won = false AND is_lost = true)
    )
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_deals_stage ON deals(stage_id);
CREATE INDEX IF NOT EXISTS idx_deals_contact ON deals(contact_id);
CREATE INDEX IF NOT EXISTS idx_deals_company ON deals(company_id);
CREATE INDEX IF NOT EXISTS idx_deals_assigned ON deals(assigned_to);
CREATE INDEX IF NOT EXISTS idx_deals_owner ON deals(owner_id);
CREATE INDEX IF NOT EXISTS idx_deals_status ON deals(is_won, is_lost);
CREATE INDEX IF NOT EXISTS idx_deals_close_date ON deals(expected_close_date);
CREATE INDEX IF NOT EXISTS idx_deals_created ON deals(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_deals_tags ON deals USING gin(tags);
CREATE INDEX IF NOT EXISTS idx_deals_value ON deals(value DESC);

-- Trigger for updated_at
CREATE OR REPLACE FUNCTION update_deals_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_deals_updated_at
    BEFORE UPDATE ON deals
    FOR EACH ROW
    EXECUTE FUNCTION update_deals_updated_at();

-- Trigger to set closed_at when deal is won or lost
CREATE OR REPLACE FUNCTION set_deal_closed_at()
RETURNS TRIGGER AS $$
BEGIN
    IF (NEW.is_won = true OR NEW.is_lost = true) AND OLD.closed_at IS NULL THEN
        NEW.closed_at = NOW();
    ELSIF NEW.is_won = false AND NEW.is_lost = false THEN
        NEW.closed_at = NULL;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_deal_closed_at
    BEFORE UPDATE ON deals
    FOR EACH ROW
    EXECUTE FUNCTION set_deal_closed_at();
