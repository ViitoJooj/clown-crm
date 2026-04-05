-- Call Logs table: Track phone and video calls via Twilio
CREATE TABLE IF NOT EXISTS call_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Call details
    call_type VARCHAR(20) NOT NULL, -- 'inbound', 'outbound', 'video'
    contact_id UUID REFERENCES contacts(id) ON DELETE SET NULL,
    user_id UUID REFERENCES users(uuid) ON DELETE SET NULL,
    
    -- Call metrics
    duration_seconds INTEGER DEFAULT 0,
    status VARCHAR(50), -- 'completed', 'missed', 'no_answer', 'busy', 'failed', 'cancelled'
    
    -- Twilio integration
    twilio_call_sid VARCHAR(100) UNIQUE,
    from_number VARCHAR(50),
    to_number VARCHAR(50),
    recording_url TEXT,
    recording_duration INTEGER,
    
    -- Additional info
    notes TEXT,
    metadata JSONB DEFAULT '{}',
    
    -- Timestamps
    started_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ended_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_call_logs_contact ON call_logs(contact_id);
CREATE INDEX IF NOT EXISTS idx_call_logs_user ON call_logs(user_id);
CREATE INDEX IF NOT EXISTS idx_call_logs_type ON call_logs(call_type);
CREATE INDEX IF NOT EXISTS idx_call_logs_status ON call_logs(status);
CREATE INDEX IF NOT EXISTS idx_call_logs_twilio_sid ON call_logs(twilio_call_sid);
CREATE INDEX IF NOT EXISTS idx_call_logs_started ON call_logs(started_at DESC);
CREATE INDEX IF NOT EXISTS idx_call_logs_created ON call_logs(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_call_logs_from_number ON call_logs(from_number);
CREATE INDEX IF NOT EXISTS idx_call_logs_to_number ON call_logs(to_number);

-- Trigger to automatically create activity when call is logged
CREATE OR REPLACE FUNCTION create_activity_for_call()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO activities (
        activity_type,
        title,
        description,
        duration_minutes,
        contact_id,
        user_id,
        metadata
    ) VALUES (
        'call',
        CASE 
            WHEN NEW.call_type = 'video' THEN 'Video Call'
            WHEN NEW.call_type = 'inbound' THEN 'Inbound Call'
            ELSE 'Outbound Call'
        END,
        NEW.notes,
        CEIL(NEW.duration_seconds / 60.0),
        NEW.contact_id,
        NEW.user_id,
        jsonb_build_object(
            'call_log_id', NEW.id,
            'status', NEW.status,
            'duration_seconds', NEW.duration_seconds,
            'from_number', NEW.from_number,
            'to_number', NEW.to_number
        )
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_create_call_activity
    AFTER INSERT ON call_logs
    FOR EACH ROW
    WHEN (NEW.status = 'completed')
    EXECUTE FUNCTION create_activity_for_call();
