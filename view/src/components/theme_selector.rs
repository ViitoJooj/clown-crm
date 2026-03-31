use dioxus::prelude::*;
use crate::styles::theme::ThemeVariant;

#[component]
pub fn ThemeSelector(
    current_theme: Signal<ThemeVariant>,
) -> Element {
    let mut show_selector = use_signal(|| false);

    let themes = vec![
        (ThemeVariant::Burgundy, "🍷", "Burgundy", "#8B0E0E"),
        (ThemeVariant::DarkPurple, "💜", "Purple", "#7C3AED"),
        (ThemeVariant::OceanBlue, "🌊", "Ocean", "#0EA5E9"),
        (ThemeVariant::ForestGreen, "🌲", "Forest", "#10B981"),
        (ThemeVariant::SunsetOrange, "🌅", "Sunset", "#F97316"),
        (ThemeVariant::MidnightBlue, "🌙", "Midnight", "#3B82F6"),
    ];

    rsx! {
        div {
            style: "position: relative;",
            
            // Theme button
            button {
                style: r#"
                    width: 44px;
                    height: 44px;
                    border-radius: 50%;
                    background: rgba(26, 26, 26, 0.8);
                    backdrop-filter: blur(10px);
                    border: 2px solid rgba(139, 14, 14, 0.5);
                    color: white;
                    font-size: 20px;
                    cursor: pointer;
                    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
                "#,
                onclick: move |_| {
                    show_selector.toggle();
                },
                "🎨"
            }

            // Theme selector dropdown
            if show_selector() {
                div {
                    style: r#"
                        position: absolute;
                        top: calc(100% + 12px);
                        right: 0;
                        background: rgba(26, 26, 26, 0.95);
                        backdrop-filter: blur(20px) saturate(180%);
                        -webkit-backdrop-filter: blur(20px) saturate(180%);
                        border: 1px solid rgba(255, 255, 255, 0.1);
                        border-radius: 16px;
                        padding: 16px;
                        min-width: 280px;
                        z-index: 1000;
                        box-shadow: 0 12px 40px rgba(0, 0, 0, 0.5);
                        animation: scaleIn 0.2s cubic-bezier(0.4, 0, 0.2, 1);
                    "#,
                    
                    div {
                        style: "margin-bottom: 12px; font-size: 14px; font-weight: 600; color: #B0B0B0;",
                        "Choose Theme"
                    }

                    div {
                        style: "display: grid; grid-template-columns: repeat(2, 1fr); gap: 8px;",
                        
                        for (variant, emoji, name, color) in themes {
                            {
                                let is_current = current_theme() == variant;
                                let border_style = if is_current {
                                    format!("border: 2px solid {};", color)
                                } else {
                                    "border: 1px solid rgba(58, 58, 58, 0.5);".to_string()
                                };
                                
                                rsx! {
                                    button {
                                        key: "{name}",
                                        style: r#"
                                            padding: 12px;
                                            border-radius: 12px;
                                            background: rgba(42, 42, 42, 0.6);
                                            backdrop-filter: blur(10px);
                                            {border_style}
                                            color: white;
                                            cursor: pointer;
                                            transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
                                            display: flex;
                                            flex-direction: column;
                                            align-items: center;
                                            gap: 6px;
                                            position: relative;
                                        "#,
                                        onclick: move |_| {
                                            current_theme.set(variant);
                                            show_selector.set(false);
                                        },
                                        
                                        div {
                                            style: "font-size: 24px;",
                                            "{emoji}"
                                        }
                                        div {
                                            style: "font-size: 12px; font-weight: 600;",
                                            "{name}"
                                        }
                                        div {
                                            style: format!("width: 32px; height: 4px; border-radius: 2px; background: {};", color),
                                        }
                                        
                                        if is_current {
                                            div {
                                                style: "position: absolute; top: 4px; right: 4px; font-size: 12px;",
                                                "✓"
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }

                    div {
                        style: "margin-top: 12px; padding-top: 12px; border-top: 1px solid rgba(58, 58, 58, 0.5); font-size: 11px; color: #707070; text-align: center;",
                        "Theme preference saved automatically"
                    }
                }
            }
        }
    }
}
