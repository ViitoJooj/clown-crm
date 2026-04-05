-- Notifications table: User notifications and alerts
CREATE TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Notification details
    user_id UUID NOT NULL REFERENCES users(uuid) ON DELETE CASCADE,
    notification_type VARCHAR(50) NOT NULL, -- 'task_due', 'deal_stage', 'mention', 'assignment', 'call_missed'
    title VARCHAR(250) NOT NULL,
    message TEXT,
    
    -- Related entity (polymorphic)
    related_to_type VARCHAR(50), -- 'contact', 'deal', 'task', 'call'
    related_to_id UUID,
    action_url TEXT,
    
    -- Read status
    is_read BOOLEAN DEFAULT false,
    read_at TIMESTAMP,
    
    -- Timestamps
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_notifications_user ON notifications(user_id);
CREATE INDEX IF NOT EXISTS idx_notifications_type ON notifications(notification_type);
CREATE INDEX IF NOT EXISTS idx_notifications_read ON notifications(is_read);
CREATE INDEX IF NOT EXISTS idx_notifications_related ON notifications(related_to_type, related_to_id);
CREATE INDEX IF NOT EXISTS idx_notifications_created ON notifications(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_notifications_unread_user ON notifications(user_id, is_read, created_at DESC);

-- Trigger to set read_at timestamp
CREATE OR REPLACE FUNCTION set_notification_read_at()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.is_read = true AND OLD.is_read = false THEN
        NEW.read_at = NOW();
    ELSIF NEW.is_read = false THEN
        NEW.read_at = NULL;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_notification_read_at
    BEFORE UPDATE ON notifications
    FOR EACH ROW
    EXECUTE FUNCTION set_notification_read_at();
