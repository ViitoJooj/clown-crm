// Extended semantic colors for CRM-specific use cases

pub struct SemanticColors {
    // Status colors for contacts
    pub lead_status: &'static str,
    pub prospect_status: &'static str,
    pub customer_status: &'static str,
    pub inactive_status: &'static str,
    
    // Deal stage colors
    pub deal_prospecting: &'static str,
    pub deal_qualification: &'static str,
    pub deal_proposal: &'static str,
    pub deal_negotiation: &'static str,
    pub deal_won: &'static str,
    pub deal_lost: &'static str,
    
    // Priority colors
    pub priority_low: &'static str,
    pub priority_medium: &'static str,
    pub priority_high: &'static str,
    pub priority_urgent: &'static str,
    
    // Task status colors
    pub task_pending: &'static str,
    pub task_in_progress: &'static str,
    pub task_completed: &'static str,
    pub task_cancelled: &'static str,
    
    // Call status colors
    pub call_completed: &'static str,
    pub call_missed: &'static str,
    pub call_no_answer: &'static str,
    pub call_busy: &'static str,
    
    // Chart colors
    pub chart_primary: &'static str,
    pub chart_secondary: &'static str,
    pub chart_tertiary: &'static str,
    pub chart_quaternary: &'static str,
}

impl SemanticColors {
    pub fn default() -> Self {
        Self::burgundy()
    }
    
    pub fn burgundy() -> Self {
        SemanticColors {
            // Contact status
            lead_status: "#3B82F6",       // Blue for new leads
            prospect_status: "#F59E0B",   // Orange for prospects
            customer_status: "#10B981",   // Green for customers
            inactive_status: "#6B7280",   // Gray for inactive
            
            // Deal stages
            deal_prospecting: "#6B7280",  // Gray
            deal_qualification: "#3B82F6", // Blue
            deal_proposal: "#F59E0B",     // Orange
            deal_negotiation: "#8B5CF6",  // Purple
            deal_won: "#10B981",          // Green
            deal_lost: "#EF4444",         // Red
            
            // Priority
            priority_low: "#6B7280",      // Gray
            priority_medium: "#3B82F6",   // Blue
            priority_high: "#F59E0B",     // Orange
            priority_urgent: "#EF4444",   // Red
            
            // Task status
            task_pending: "#6B7280",      // Gray
            task_in_progress: "#3B82F6", // Blue
            task_completed: "#10B981",    // Green
            task_cancelled: "#EF4444",    // Red
            
            // Call status
            call_completed: "#10B981",    // Green
            call_missed: "#EF4444",       // Red
            call_no_answer: "#F59E0B",    // Orange
            call_busy: "#6B7280",         // Gray
            
            // Charts
            chart_primary: "#8B0E0E",     // Burgundy
            chart_secondary: "#3B82F6",   // Blue
            chart_tertiary: "#10B981",    // Green
            chart_quaternary: "#F59E0B",  // Orange
        }
    }
}

// Utility functions for color operations
pub fn lighten_color(hex: &str, amount: f32) -> String {
    // Simple color lightening (for hover states, etc.)
    // This is a simplified version - in production you might want a full color manipulation library
    format!("{}cc", hex) // Add some transparency for now
}

pub fn darken_color(hex: &str, amount: f32) -> String {
    // Simple color darkening
    format!("{}88", hex) // Add transparency
}

pub fn with_opacity(hex: &str, opacity: f32) -> String {
    let opacity_hex = format!("{:02x}", (opacity * 255.0) as u8);
    format!("{}{}", hex, opacity_hex)
}

// Gradient generators
pub fn gradient_vertical(from: &str, to: &str) -> String {
    format!("linear-gradient(180deg, {} 0%, {} 100%)", from, to)
}

pub fn gradient_horizontal(from: &str, to: &str) -> String {
    format!("linear-gradient(90deg, {} 0%, {} 100%)", from, to)
}

pub fn gradient_diagonal(from: &str, to: &str) -> String {
    format!("linear-gradient(135deg, {} 0%, {} 100%)", from, to)
}

pub fn gradient_radial(from: &str, to: &str) -> String {
    format!("radial-gradient(circle, {} 0%, {} 100%)", from, to)
}

// Shadow generators
pub fn box_shadow(color: &str, blur: i32, spread: i32) -> String {
    format!("0 {}px {}px {}", blur / 2, blur, with_opacity(color, 0.3))
}

pub fn box_shadow_elevated(color: &str) -> String {
    format!("0 8px 32px {}, 0 2px 8px {}", 
        with_opacity(color, 0.4), 
        with_opacity(color, 0.2))
}

pub fn text_shadow(color: &str) -> String {
    format!("0 2px 4px {}", with_opacity(color, 0.5))
}

// Status badge color helper
pub fn get_status_color(status: &str) -> &'static str {
    match status.to_lowercase().as_str() {
        "lead" | "pending" | "new" => "#3B82F6",
        "prospect" | "in_progress" | "ongoing" => "#F59E0B",
        "customer" | "completed" | "success" | "won" => "#10B981",
        "inactive" | "cancelled" | "lost" => "#EF4444",
        "qualified" | "proposal" => "#8B5CF6",
        _ => "#6B7280",
    }
}

// Priority color helper
pub fn get_priority_color(priority: &str) -> &'static str {
    match priority.to_lowercase().as_str() {
        "low" => "#6B7280",
        "medium" => "#3B82F6",
        "high" => "#F59E0B",
        "urgent" | "critical" => "#EF4444",
        _ => "#6B7280",
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_with_opacity() {
        assert_eq!(with_opacity("#8B0E0E", 0.5), "#8B0E0E7f");
    }
    
    #[test]
    fn test_gradients() {
        assert!(gradient_vertical("#000", "#fff").contains("180deg"));
        assert!(gradient_horizontal("#000", "#fff").contains("90deg"));
        assert!(gradient_diagonal("#000", "#fff").contains("135deg"));
    }
    
    #[test]
    fn test_status_colors() {
        assert_eq!(get_status_color("lead"), "#3B82F6");
        assert_eq!(get_status_color("won"), "#10B981");
        assert_eq!(get_status_color("lost"), "#EF4444");
    }
}
