use dioxus::prelude::*;
use crate::components::common::{Icon, Avatar};

#[derive(Props, Clone, PartialEq)]
pub struct DealCardProps {
    pub title: String,
    pub value: f64,
    #[props(default = "USD".to_string())]
    pub currency: String,
    #[props(default = "qualification".to_string())]
    pub stage: String,
    #[props(default = 50)]
    pub progress: u8,
    #[props(default = "".to_string())]
    pub assigned_to: String,
    #[props(default = "".to_string())]
    pub assigned_avatar: String,
    #[props(default = "".to_string())]
    pub company: String,
    #[props(default = None)]
    pub on_click: Option<EventHandler<MouseEvent>>,
}

#[component]
pub fn DealCard(props: DealCardProps) -> Element {
    let stage_color = match props.stage.to_lowercase().as_str() {
        "prospecting" => "#6B7280",
        "qualification" => "#3B82F6",
        "proposal" => "#F59E0B",
        "negotiation" => "#8B5CF6",
        "won" => "#10B981",
        "lost" => "#EF4444",
        _ => "#6B7280",
    };

    let formatted_value = format!("{}{:.2}", 
        if props.currency == "USD" { "$" } else { &props.currency },
        props.value
    );

    rsx! {
        div {
            class: "deal-card",
            onclick: move |e| {
                if let Some(handler) = &props.on_click {
                    handler.call(e);
                }
            },
            style: r#"
                background: linear-gradient(135deg, rgba(26, 26, 26, 0.95) 0%, rgba(42, 42, 42, 0.85) 100%);
                backdrop-filter: blur(10px);
                -webkit-backdrop-filter: blur(10px);
                border: 1px solid rgba(58, 58, 58, 0.5);
                border-radius: 16px;
                padding: 20px;
                transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
                box-shadow: 0 4px 16px rgba(0, 0, 0, 0.4);
                cursor: pointer;
                position: relative;
                overflow: hidden;
            "#,

            div {
                style: "margin-bottom: 16px;",
                
                div {
                    style: "display: flex; justify-content: space-between; align-items: start; margin-bottom: 8px;",
                    
                    h3 {
                        style: "font-size: 16px; font-weight: 700; color: #FFFFFF; flex: 1; margin-right: 12px;",
                        "{props.title}"
                    }

                    span {
                        style: format!(r#"
                            display: inline-flex;
                            align-items: center;
                            padding: 4px 10px;
                            border-radius: 999px;
                            font-size: 11px;
                            font-weight: 600;
                            background: rgba({}, {}, {}, 0.2);
                            color: {};
                            border: 1px solid rgba({}, {}, {}, 0.4);
                            text-transform: capitalize;
                        "#, 
                            if props.stage.contains("prospecting") { "107, 114, 128" } else if props.stage.contains("qualification") { "59, 130, 246" } else if props.stage.contains("proposal") { "245, 158, 11" } else if props.stage.contains("negotiation") { "139, 92, 246" } else if props.stage.contains("won") { "16, 185, 129" } else { "239, 68, 68" },
                            if props.stage.contains("prospecting") { "107, 114, 128" } else if props.stage.contains("qualification") { "59, 130, 246" } else if props.stage.contains("proposal") { "245, 158, 11" } else if props.stage.contains("negotiation") { "139, 92, 246" } else if props.stage.contains("won") { "16, 185, 129" } else { "239, 68, 68" },
                            if props.stage.contains("prospecting") { "107, 114, 128" } else if props.stage.contains("qualification") { "59, 130, 246" } else if props.stage.contains("proposal") { "245, 158, 11" } else if props.stage.contains("negotiation") { "139, 92, 246" } else if props.stage.contains("won") { "16, 185, 129" } else { "239, 68, 68" },
                            stage_color,
                            if props.stage.contains("prospecting") { "107, 114, 128" } else if props.stage.contains("qualification") { "59, 130, 246" } else if props.stage.contains("proposal") { "245, 158, 11" } else if props.stage.contains("negotiation") { "139, 92, 246" } else if props.stage.contains("won") { "16, 185, 129" } else { "239, 68, 68" },
                            if props.stage.contains("prospecting") { "107, 114, 128" } else if props.stage.contains("qualification") { "59, 130, 246" } else if props.stage.contains("proposal") { "245, 158, 11" } else if props.stage.contains("negotiation") { "139, 92, 246" } else if props.stage.contains("won") { "16, 185, 129" } else { "239, 68, 68" },
                            if props.stage.contains("prospecting") { "107, 114, 128" } else if props.stage.contains("qualification") { "59, 130, 246" } else if props.stage.contains("proposal") { "245, 158, 11" } else if props.stage.contains("negotiation") { "139, 92, 246" } else if props.stage.contains("won") { "16, 185, 129" } else { "239, 68, 68" }
                        ),
                        "{props.stage}"
                    }
                }

                if !props.company.is_empty() {
                    p {
                        style: "font-size: 13px; color: #B0B0B0; display: flex; align-items: center; gap: 6px;",
                        Icon { name: "building".to_string(), size: "12".to_string(), color: "#B0B0B0".to_string(), class: "".to_string() }
                        "{props.company}"
                    }
                }
            }

            div {
                style: "margin-bottom: 16px;",
                
                div {
                    style: "display: flex; align-items: center; gap: 8px; margin-bottom: 6px;",
                    Icon { name: "dollar-sign".to_string(), size: "20".to_string(), color: "#10B981".to_string(), class: "".to_string() }
                    span {
                        style: "font-size: 24px; font-weight: 700; color: #10B981;",
                        "{formatted_value}"
                    }
                }

                div {
                    style: "display: flex; align-items: center; gap: 8px; margin-top: 12px;",
                    
                    div {
                        style: "flex: 1;",
                        
                        div {
                            style: "display: flex; justify-content: space-between; margin-bottom: 6px;",
                            span { style: "font-size: 12px; color: #B0B0B0;", "Progress" }
                            span { style: "font-size: 12px; font-weight: 600; color: #FFFFFF;", "{props.progress}%" }
                        }

                        div {
                            style: "width: 100%; height: 6px; background: rgba(58, 58, 58, 0.5); border-radius: 999px; overflow: hidden;",
                            
                            div {
                                style: format!(r#"
                                    width: {}%;
                                    height: 100%;
                                    background: linear-gradient(90deg, {} 0%, {} 100%);
                                    border-radius: 999px;
                                    transition: width 0.5s cubic-bezier(0.4, 0, 0.2, 1);
                                "#, props.progress, stage_color, stage_color),
                            }
                        }
                    }
                }
            }

            if !props.assigned_to.is_empty() {
                div {
                    style: "display: flex; align-items: center; gap: 10px; padding-top: 12px; border-top: 1px solid rgba(58, 58, 58, 0.5);",
                    
                    Avatar {
                        src: props.assigned_avatar.clone(),
                        alt: props.assigned_to.clone(),
                        size: "32px".to_string(),
                    }

                    div {
                        style: "flex: 1;",
                        p { style: "font-size: 11px; color: #707070; margin-bottom: 2px;", "Assigned to" }
                        p { style: "font-size: 13px; font-weight: 600; color: #FFFFFF;", "{props.assigned_to}" }
                    }

                    Icon { name: "chevron-right".to_string(), size: "20".to_string(), color: "#707070".to_string(), class: "".to_string() }
                }
            }

            style {
                r#"
                .deal-card:hover {
                    transform: translateY(-4px);
                    box-shadow: 0 12px 40px rgba(0, 0, 0, 0.5);
                    border-color: rgba(139, 14, 14, 0.3);
                }
                "#
            }
        }
    }
}
