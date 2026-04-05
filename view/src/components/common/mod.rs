use dioxus::prelude::*;

pub mod app_layout;
pub mod icon;

pub use app_layout::AppLayout;
pub use icon::Icon;

#[component]
pub fn Button(
    #[props(default = false)] disabled: bool,
    #[props(default = "".to_string())] class: String,
    onclick: EventHandler<MouseEvent>,
    children: Element,
) -> Element {
    let base_style = r#"
        padding: 12px 24px;
        border-radius: 10px;
        font-size: 14px;
        font-weight: 600;
        cursor: pointer;
        transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
        border: none;
        display: inline-flex;
        align-items: center;
        justify-content: center;
        gap: 8px;
        white-space: nowrap;
        position: relative;
        overflow: hidden;
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
    "#;

    let variant_style = if class.contains("primary") {
        "background: linear-gradient(135deg, #8B0E0E 0%, #C91A1A 100%); color: white; border: 1px solid rgba(139, 14, 14, 0.5); box-shadow: 0 4px 16px rgba(139, 14, 14, 0.3);"
    } else if class.contains("secondary") {
        "background: rgba(42, 42, 42, 0.7); backdrop-filter: blur(10px); color: #8B0E0E; border: 1px solid rgba(139, 14, 14, 0.5);"
    } else if class.contains("danger") {
        "background: linear-gradient(135deg, #EF4444 0%, #DC2626 100%); color: white; border: 1px solid rgba(239, 68, 68, 0.5); box-shadow: 0 4px 16px rgba(239, 68, 68, 0.3);"
    } else if class.contains("success") {
        "background: linear-gradient(135deg, #10B981 0%, #059669 100%); color: white; border: 1px solid rgba(16, 185, 129, 0.5); box-shadow: 0 4px 16px rgba(16, 185, 129, 0.3);"
    } else if class.contains("ghost") {
        "background: transparent; color: #B0B0B0; border: 1px solid rgba(58, 58, 58, 0.5);"
    } else {
        "background: rgba(42, 42, 42, 0.7); backdrop-filter: blur(10px); color: white; border: 1px solid rgba(58, 58, 58, 0.5);"
    };

    let disabled_style = if disabled {
        "opacity: 0.5; cursor: not-allowed; pointer-events: none;"
    } else {
        "&:hover { transform: translateY(-2px); box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3); }"
    };

    rsx! {
        button {
            disabled,
            onclick: move |e| onclick.call(e),
            style: "{base_style} {variant_style} {disabled_style}",
            {children}
        }
    }
}

#[component]
pub fn Input(
    #[props(default = "text".to_string())] input_type: String,
    #[props(default = "".to_string())] placeholder: String,
    #[props(default = "".to_string())] value: String,
    onchange: EventHandler<FormEvent>,
) -> Element {
    let base_style = r#"
        width: 100%;
        padding: 12px 16px;
        border-radius: 10px;
        font-size: 14px;
        background: rgba(42, 42, 42, 0.7);
        backdrop-filter: blur(10px);
        -webkit-backdrop-filter: blur(10px);
        color: #FFFFFF;
        border: 1px solid rgba(58, 58, 58, 0.5);
        transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    "#;

    rsx! {
        input {
            r#type: input_type,
            placeholder,
            value,
            onchange: move |e| onchange.call(e),
            style: base_style,
        }
    }
}

#[component]
pub fn LoadingSpinner(
    #[props(default = "24px".to_string())] size: String,
) -> Element {
    let spinner_style = format!(r#"
        display: inline-block;
        width: {};
        height: {};
        border: 3px solid rgba(58, 58, 58, 0.3);
        border-top-color: #8B0E0E;
        border-radius: 50%;
        animation: spin 0.8s linear infinite;
    "#, size, size);

    rsx! {
        div {
            style: "{spinner_style}",
        }
    }
}

#[component]
pub fn Card(
    #[props(default = "".to_string())] class: String,
    #[props(default = false)] glass: bool,
    children: Element,
) -> Element {
    let card_style = if glass {
        r#"
            background: rgba(26, 26, 26, 0.6);
            backdrop-filter: blur(20px) saturate(180%);
            -webkit-backdrop-filter: blur(20px) saturate(180%);
            border: 1px solid rgba(255, 255, 255, 0.05);
            border-radius: 16px;
            padding: 24px;
            transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
            box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
        "#
    } else {
        r#"
            background: linear-gradient(135deg, rgba(26, 26, 26, 0.95) 0%, rgba(42, 42, 42, 0.85) 100%);
            backdrop-filter: blur(10px);
            -webkit-backdrop-filter: blur(10px);
            border: 1px solid rgba(58, 58, 58, 0.5);
            border-radius: 16px;
            padding: 24px;
            transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
            box-shadow: 0 4px 16px rgba(0, 0, 0, 0.4);
        "#
    };

    let hover_style = r#"
        &:hover {
            transform: translateY(-2px);
            box-shadow: 0 12px 40px rgba(0, 0, 0, 0.5);
            border-color: rgba(139, 14, 14, 0.3);
        }
    "#;

    rsx! {
        div {
            style: "{card_style} {hover_style}",
            class,
            {children}
        }
    }
}

#[component]
pub fn Badge(
    #[props(default = "default".to_string())] variant: String,
    children: Element,
) -> Element {
    let badge_style = match variant.as_str() {
        "success" => "background: linear-gradient(135deg, rgba(16, 185, 129, 0.2), rgba(5, 150, 105, 0.3)); color: #34D399; border: 1px solid rgba(16, 185, 129, 0.4);",
        "error" => "background: linear-gradient(135deg, rgba(239, 68, 68, 0.2), rgba(220, 38, 38, 0.3)); color: #F87171; border: 1px solid rgba(239, 68, 68, 0.4);",
        "warning" => "background: linear-gradient(135deg, rgba(245, 158, 11, 0.2), rgba(217, 119, 6, 0.3)); color: #FBBF24; border: 1px solid rgba(245, 158, 11, 0.4);",
        "info" => "background: linear-gradient(135deg, rgba(59, 130, 246, 0.2), rgba(37, 99, 235, 0.3)); color: #60A5FA; border: 1px solid rgba(59, 130, 246, 0.4);",
        "primary" => "background: linear-gradient(135deg, rgba(139, 14, 14, 0.2), rgba(201, 26, 26, 0.3)); color: #C91A1A; border: 1px solid rgba(139, 14, 14, 0.4);",
        _ => "background: rgba(42, 42, 42, 0.7); color: #B0B0B0; border: 1px solid rgba(58, 58, 58, 0.5);",
    };

    rsx! {
        span {
            style: "display: inline-flex; align-items: center; padding: 4px 12px; border-radius: 999px; font-size: 12px; font-weight: 600; backdrop-filter: blur(10px); {badge_style}",
            {children}
        }
    }
}

#[component]
pub fn Modal(
    #[props(default = false)] show: bool,
    on_close: EventHandler<MouseEvent>,
    children: Element,
) -> Element {
    if !show {
        return rsx! { div {} };
    }

    rsx! {
        div {
            style: "position: fixed; inset: 0; z-index: 1000; display: flex; align-items: center; justify-content: center; animation: fadeIn 0.3s ease;",
            
            // Backdrop with blur
            div {
                style: "position: absolute; inset: 0; background: rgba(10, 10, 10, 0.8); backdrop-filter: blur(8px); -webkit-backdrop-filter: blur(8px);",
                onclick: move |e| on_close.call(e),
            }
            
            // Modal content
            div {
                style: "position: relative; z-index: 10; max-width: 90%; max-height: 90vh; overflow: auto; animation: scaleIn 0.3s cubic-bezier(0.4, 0, 0.2, 1);",
                Card {
                    glass: true,
                    {children}
                }
            }
        }
    }
}

#[component]
pub fn Divider() -> Element {
    rsx! {
        div {
            style: "height: 1px; background: linear-gradient(90deg, transparent, rgba(58, 58, 58, 0.5), transparent); margin: 16px 0;",
        }
    }
}

#[component]
pub fn Avatar(
    #[props(default = "".to_string())] src: String,
    #[props(default = "".to_string())] alt: String,
    #[props(default = "40px".to_string())] size: String,
) -> Element {
    let avatar_style = format!(
        "width: {}; height: {}; border-radius: 50%; object-fit: cover; border: 2px solid rgba(139, 14, 14, 0.5); box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);",
        size, size
    );

    if src.is_empty() {
        rsx! {
            div {
                style: "{avatar_style} display: flex; align-items: center; justify-content: center; background: linear-gradient(135deg, #8B0E0E, #C91A1A); color: white; font-weight: bold; font-size: 16px;",
                "{alt.chars().next().unwrap_or('?').to_uppercase()}"
            }
        }
    } else {
        rsx! {
            img {
                src,
                alt,
                style: avatar_style,
            }
        }
    }
}
