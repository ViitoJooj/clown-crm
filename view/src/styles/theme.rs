// Advanced theme configuration with multiple color palettes

#[derive(Clone, Copy, Debug, PartialEq)]
pub enum ThemeVariant {
    Burgundy,
    DarkPurple,
    OceanBlue,
    ForestGreen,
    SunsetOrange,
    MidnightBlue,
}

pub struct Theme {
    pub primary: &'static str,
    pub primary_dark: &'static str,
    pub primary_light: &'static str,
    pub primary_gradient: &'static str,
    pub background: &'static str,
    pub background_gradient: &'static str,
    pub surface: &'static str,
    pub surface_light: &'static str,
    pub surface_elevated: &'static str,
    pub text_primary: &'static str,
    pub text_secondary: &'static str,
    pub text_muted: &'static str,
    pub border: &'static str,
    pub border_light: &'static str,
    pub success: &'static str,
    pub error: &'static str,
    pub warning: &'static str,
    pub info: &'static str,
    pub accent: &'static str,
    pub shadow: &'static str,
    pub glow: &'static str,
}

impl Theme {
    pub fn burgundy() -> Self {
        Theme {
            primary: "#8B0E0E",
            primary_dark: "#5C0A0A",
            primary_light: "#C91A1A",
            primary_gradient: "linear-gradient(135deg, #8B0E0E 0%, #C91A1A 100%)",
            background: "#0A0A0A",
            background_gradient: "linear-gradient(180deg, #0A0A0A 0%, #1A0505 100%)",
            surface: "#1A1A1A",
            surface_light: "#2A2A2A",
            surface_elevated: "#3A1A1A",
            text_primary: "#FFFFFF",
            text_secondary: "#B0B0B0",
            text_muted: "#707070",
            border: "#3A3A3A",
            border_light: "#4A4A4A",
            success: "#10B981",
            error: "#EF4444",
            warning: "#F59E0B",
            info: "#3B82F6",
            accent: "#EC4899",
            shadow: "rgba(139, 14, 14, 0.3)",
            glow: "rgba(201, 26, 26, 0.4)",
        }
    }

    pub fn dark_purple() -> Self {
        Theme {
            primary: "#7C3AED",
            primary_dark: "#5B21B6",
            primary_light: "#A78BFA",
            primary_gradient: "linear-gradient(135deg, #7C3AED 0%, #A78BFA 100%)",
            background: "#0F0A1E",
            background_gradient: "linear-gradient(180deg, #0F0A1E 0%, #1E1134 100%)",
            surface: "#1A1333",
            surface_light: "#2A1F47",
            surface_elevated: "#3A2B5C",
            text_primary: "#FFFFFF",
            text_secondary: "#C4B5FD",
            text_muted: "#8B7AC7",
            border: "#3A2B5C",
            border_light: "#4A3670",
            success: "#10B981",
            error: "#EF4444",
            warning: "#F59E0B",
            info: "#60A5FA",
            accent: "#EC4899",
            shadow: "rgba(124, 58, 237, 0.3)",
            glow: "rgba(167, 139, 250, 0.4)",
        }
    }

    pub fn ocean_blue() -> Self {
        Theme {
            primary: "#0EA5E9",
            primary_dark: "#0369A1",
            primary_light: "#38BDF8",
            primary_gradient: "linear-gradient(135deg, #0EA5E9 0%, #38BDF8 100%)",
            background: "#020617",
            background_gradient: "linear-gradient(180deg, #020617 0%, #0C1629 100%)",
            surface: "#0F172A",
            surface_light: "#1E293B",
            surface_elevated: "#334155",
            text_primary: "#F1F5F9",
            text_secondary: "#94A3B8",
            text_muted: "#64748B",
            border: "#334155",
            border_light: "#475569",
            success: "#10B981",
            error: "#EF4444",
            warning: "#F59E0B",
            info: "#3B82F6",
            accent: "#06B6D4",
            shadow: "rgba(14, 165, 233, 0.3)",
            glow: "rgba(56, 189, 248, 0.4)",
        }
    }

    pub fn forest_green() -> Self {
        Theme {
            primary: "#10B981",
            primary_dark: "#059669",
            primary_light: "#34D399",
            primary_gradient: "linear-gradient(135deg, #10B981 0%, #34D399 100%)",
            background: "#0A1810",
            background_gradient: "linear-gradient(180deg, #0A1810 0%, #112822 100%)",
            surface: "#132B23",
            surface_light: "#1E3D34",
            surface_elevated: "#2A4F45",
            text_primary: "#F0FDF4",
            text_secondary: "#A7F3D0",
            text_muted: "#6EE7B7",
            border: "#2A4F45",
            border_light: "#356156",
            success: "#10B981",
            error: "#EF4444",
            warning: "#F59E0B",
            info: "#3B82F6",
            accent: "#14B8A6",
            shadow: "rgba(16, 185, 129, 0.3)",
            glow: "rgba(52, 211, 153, 0.4)",
        }
    }

    pub fn sunset_orange() -> Self {
        Theme {
            primary: "#F97316",
            primary_dark: "#C2410C",
            primary_light: "#FB923C",
            primary_gradient: "linear-gradient(135deg, #F97316 0%, #FB923C 100%)",
            background: "#1C0E08",
            background_gradient: "linear-gradient(180deg, #1C0E08 0%, #2A1710 100%)",
            surface: "#2A1710",
            surface_light: "#3D2418",
            surface_elevated: "#503120",
            text_primary: "#FFF7ED",
            text_secondary: "#FDBA74",
            text_muted: "#FB923C",
            border: "#503120",
            border_light: "#63462F",
            success: "#10B981",
            error: "#EF4444",
            warning: "#F59E0B",
            info: "#3B82F6",
            accent: "#EC4899",
            shadow: "rgba(249, 115, 22, 0.3)",
            glow: "rgba(251, 146, 60, 0.4)",
        }
    }

    pub fn midnight_blue() -> Self {
        Theme {
            primary: "#3B82F6",
            primary_dark: "#1E40AF",
            primary_light: "#60A5FA",
            primary_gradient: "linear-gradient(135deg, #3B82F6 0%, #60A5FA 100%)",
            background: "#030712",
            background_gradient: "linear-gradient(180deg, #030712 0%, #0F1729 100%)",
            surface: "#111827",
            surface_light: "#1F2937",
            surface_elevated: "#374151",
            text_primary: "#F9FAFB",
            text_secondary: "#9CA3AF",
            text_muted: "#6B7280",
            border: "#374151",
            border_light: "#4B5563",
            success: "#10B981",
            error: "#EF4444",
            warning: "#F59E0B",
            info: "#3B82F6",
            accent: "#8B5CF6",
            shadow: "rgba(59, 130, 246, 0.3)",
            glow: "rgba(96, 165, 250, 0.4)",
        }
    }

    pub fn from_variant(variant: ThemeVariant) -> Self {
        match variant {
            ThemeVariant::Burgundy => Self::burgundy(),
            ThemeVariant::DarkPurple => Self::dark_purple(),
            ThemeVariant::OceanBlue => Self::ocean_blue(),
            ThemeVariant::ForestGreen => Self::forest_green(),
            ThemeVariant::SunsetOrange => Self::sunset_orange(),
            ThemeVariant::MidnightBlue => Self::midnight_blue(),
        }
    }

    pub fn default() -> Self {
        Self::burgundy()
    }
}

// Default theme constant
pub const THEME: Theme = Theme {
    primary: "#8B0E0E",
    primary_dark: "#5C0A0A",
    primary_light: "#C91A1A",
    primary_gradient: "linear-gradient(135deg, #8B0E0E 0%, #C91A1A 100%)",
    background: "#0A0A0A",
    background_gradient: "linear-gradient(180deg, #0A0A0A 0%, #1A0505 100%)",
    surface: "#1A1A1A",
    surface_light: "#2A2A2A",
    surface_elevated: "#3A1A1A",
    text_primary: "#FFFFFF",
    text_secondary: "#B0B0B0",
    text_muted: "#707070",
    border: "#3A3A3A",
    border_light: "#4A4A4A",
    success: "#10B981",
    error: "#EF4444",
    warning: "#F59E0B",
    info: "#3B82F6",
    accent: "#EC4899",
    shadow: "rgba(139, 14, 14, 0.3)",
    glow: "rgba(201, 26, 26, 0.4)",
};

pub fn global_styles() -> &'static str {
    r#"
    * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
    }

    html, body {
        width: 100%;
        height: 100%;
        font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
        background: linear-gradient(180deg, #0A0A0A 0%, #1A0505 100%);
        color: #FFFFFF;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
    }

    body {
        overflow: hidden;
    }

    a {
        color: #8B0E0E;
        text-decoration: none;
        transition: color 0.3s cubic-bezier(0.4, 0, 0.2, 1);
        position: relative;
    }

    a:hover {
        color: #C91A1A;
    }

    a::after {
        content: '';
        position: absolute;
        width: 0;
        height: 2px;
        bottom: -2px;
        left: 0;
        background: linear-gradient(90deg, #8B0E0E, #C91A1A);
        transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    }

    a:hover::after {
        width: 100%;
    }

    button {
        font-family: inherit;
        cursor: pointer;
        border: none;
        border-radius: 8px;
        transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
        position: relative;
        overflow: hidden;
    }

    button::before {
        content: '';
        position: absolute;
        top: 50%;
        left: 50%;
        width: 0;
        height: 0;
        border-radius: 50%;
        background: rgba(255, 255, 255, 0.1);
        transform: translate(-50%, -50%);
        transition: width 0.6s, height 0.6s;
    }

    button:hover::before {
        width: 300px;
        height: 300px;
    }

    button:active {
        transform: scale(0.98);
    }

    input, textarea, select {
        font-family: inherit;
        background-color: rgba(42, 42, 42, 0.7);
        backdrop-filter: blur(10px);
        -webkit-backdrop-filter: blur(10px);
        color: #FFFFFF;
        border: 1px solid rgba(58, 58, 58, 0.5);
        padding: 10px 14px;
        border-radius: 8px;
        transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    }

    input:hover, textarea:hover, select:hover {
        border-color: rgba(139, 14, 14, 0.5);
        background-color: rgba(42, 42, 42, 0.9);
    }

    input:focus, textarea:focus, select:focus {
        outline: none;
        border-color: #8B0E0E;
        background-color: rgba(42, 42, 42, 1);
        box-shadow: 0 0 0 4px rgba(139, 14, 14, 0.15),
                    0 4px 20px rgba(139, 14, 14, 0.2);
        transform: translateY(-1px);
    }

    input::placeholder, textarea::placeholder {
        color: #707070;
    }

    /* Scrollbar styling - Modern */
    ::-webkit-scrollbar {
        width: 10px;
        height: 10px;
    }

    ::-webkit-scrollbar-track {
        background: rgba(26, 26, 26, 0.3);
        border-radius: 10px;
    }

    ::-webkit-scrollbar-thumb {
        background: linear-gradient(180deg, #3A3A3A, #4A4A4A);
        border-radius: 10px;
        border: 2px solid rgba(26, 26, 26, 0.3);
    }

    ::-webkit-scrollbar-thumb:hover {
        background: linear-gradient(180deg, #4A4A4A, #5A5A5A);
    }

    /* Blur utility classes */
    .blur-light {
        backdrop-filter: blur(8px);
        -webkit-backdrop-filter: blur(8px);
    }

    .blur-medium {
        backdrop-filter: blur(16px);
        -webkit-backdrop-filter: blur(16px);
    }

    .blur-heavy {
        backdrop-filter: blur(24px);
        -webkit-backdrop-filter: blur(24px);
    }

    /* Glass effect */
    .glass {
        background: rgba(26, 26, 26, 0.6);
        backdrop-filter: blur(20px) saturate(180%);
        -webkit-backdrop-filter: blur(20px) saturate(180%);
        border: 1px solid rgba(255, 255, 255, 0.05);
    }

    /* Gradient animation */
    @keyframes gradient-shift {
        0% { background-position: 0% 50%; }
        50% { background-position: 100% 50%; }
        100% { background-position: 0% 50%; }
    }

    .gradient-animate {
        background-size: 200% 200%;
        animation: gradient-shift 8s ease infinite;
    }

    /* Smooth fade in */
    @keyframes fadeIn {
        from {
            opacity: 0;
            transform: translateY(10px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }

    .fade-in {
        animation: fadeIn 0.5s cubic-bezier(0.4, 0, 0.2, 1);
    }

    /* Glow effect */
    @keyframes glow {
        0%, 100% {
            box-shadow: 0 0 20px rgba(139, 14, 14, 0.3),
                        0 0 40px rgba(139, 14, 14, 0.2);
        }
        50% {
            box-shadow: 0 0 30px rgba(139, 14, 14, 0.5),
                        0 0 60px rgba(139, 14, 14, 0.3);
        }
    }

    .glow {
        animation: glow 3s ease-in-out infinite;
    }

    /* Spin animation */
    @keyframes spin {
        to { transform: rotate(360deg); }
    }

    /* Scale animation */
    @keyframes scaleIn {
        from {
            opacity: 0;
            transform: scale(0.9);
        }
        to {
            opacity: 1;
            transform: scale(1);
        }
    }

    .scale-in {
        animation: scaleIn 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    }

    /* Slide animations */
    @keyframes slideInRight {
        from {
            opacity: 0;
            transform: translateX(20px);
        }
        to {
            opacity: 1;
            transform: translateX(0);
        }
    }

    @keyframes slideInLeft {
        from {
            opacity: 0;
            transform: translateX(-20px);
        }
        to {
            opacity: 1;
            transform: translateX(0);
        }
    }

    @keyframes float {
        0%, 100% { 
            transform: translate(0, 0) scale(1); 
        }
        50% { 
            transform: translate(20px, 20px) scale(1.05); 
        }
    }

    .slide-in-right {
        animation: slideInRight 0.4s cubic-bezier(0.4, 0, 0.2, 1);
    }

    .slide-in-left {
        animation: slideInLeft 0.4s cubic-bezier(0.4, 0, 0.2, 1);
    }
    "#
}
