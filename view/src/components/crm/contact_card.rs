use dioxus::prelude::*;
use crate::components::common::{Icon, Avatar};

#[derive(Props, Clone, PartialEq)]
pub struct ContactCardProps {
    pub name: String,
    #[props(default = "".to_string())]
    pub company: String,
    #[props(default = "".to_string())]
    pub email: String,
    #[props(default = "".to_string())]
    pub phone: String,
    #[props(default = "".to_string())]
    pub avatar_url: String,
    #[props(default = "lead".to_string())]
    pub status: String,
    #[props(default = None)]
    pub on_call: Option<EventHandler<MouseEvent>>,
    #[props(default = None)]
    pub on_email: Option<EventHandler<MouseEvent>>,
    #[props(default = None)]
    pub on_view: Option<EventHandler<MouseEvent>>,
}

#[component]
pub fn ContactCard(props: ContactCardProps) -> Element {
    let status_color = match props.status.to_lowercase().as_str() {
        "lead" => "#3B82F6",
        "prospect" => "#F59E0B",
        "customer" => "#10B981",
        "inactive" => "#6B7280",
        _ => "#6B7280",
    };

    let status_rgb = match props.status.to_lowercase().as_str() {
        "lead" => "59, 130, 246",
        "prospect" => "245, 158, 11",
        "customer" => "16, 185, 129",
        _ => "107, 114, 128",
    };

    rsx! {
        div {
            class: "contact-card",
            style: r#"
                background: linear-gradient(135deg, rgba(26, 26, 26, 0.95) 0%, rgba(42, 42, 42, 0.85) 100%);
                backdrop-filter: blur(10px);
                -webkit-backdrop-filter: blur(10px);
                border: 1px solid rgba(58, 58, 58, 0.5);
                border-radius: 16px;
                padding: 20px;
                transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
                box-shadow: 0 4px 16px rgba(0, 0, 0, 0.4);
                position: relative;
                overflow: hidden;
            "#,

            div {
                style: format!("position: absolute; top: 0; left: 0; right: 0; height: 3px; background: {};", status_color),
            }

            div {
                style: "display: flex; gap: 16px; align-items: center; margin-bottom: 16px;",
                
                Avatar {
                    src: props.avatar_url.clone(),
                    alt: props.name.clone(),
                    size: "56px".to_string(),
                }

                div {
                    style: "flex: 1; min-width: 0;",
                    
                    h3 {
                        style: "font-size: 18px; font-weight: 700; color: #FFFFFF; margin-bottom: 4px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;",
                        "{props.name}"
                    }

                    if !props.company.is_empty() {
                        p {
                            style: "font-size: 14px; color: #B0B0B0; display: flex; align-items: center; gap: 6px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;",
                            Icon { name: "building".to_string(), size: "14".to_string(), color: "#B0B0B0".to_string(), class: "".to_string() }
                            span { "{props.company}" }
                        }
                    }
                }

                span {
                    style: format!(r#"
                        display: inline-flex;
                        align-items: center;
                        padding: 4px 12px;
                        border-radius: 999px;
                        font-size: 12px;
                        font-weight: 600;
                        background: rgba({}, 0.2);
                        color: {};
                        border: 1px solid rgba({}, 0.4);
                    "#, status_rgb, status_color, status_rgb),
                    "{props.status}"
                }
            }

            div {
                style: "display: flex; flex-direction: column; gap: 8px; margin-bottom: 16px;",
                
                if !props.email.is_empty() {
                    div {
                        style: "display: flex; align-items: center; gap: 8px; font-size: 14px; color: #B0B0B0;",
                        Icon { name: "mail".to_string(), size: "16".to_string(), color: "#8B0E0E".to_string(), class: "".to_string() }
                        span { style: "overflow: hidden; text-overflow: ellipsis; white-space: nowrap;", "{props.email}" }
                    }
                }

                if !props.phone.is_empty() {
                    div {
                        style: "display: flex; align-items: center; gap: 8px; font-size: 14px; color: #B0B0B0;",
                        Icon { name: "phone".to_string(), size: "16".to_string(), color: "#8B0E0E".to_string(), class: "".to_string() }
                        span { "{props.phone}" }
                    }
                }
            }

            div {
                style: "display: flex; gap: 8px; padding-top: 12px; border-top: 1px solid rgba(58, 58, 58, 0.5);",
                
                if let Some(on_call) = props.on_call {
                    button {
                        onclick: move |e| on_call.call(e),
                        class: "action-btn",
                        style: r#"
                            flex: 1;
                            padding: 10px;
                            border-radius: 10px;
                            background: rgba(42, 42, 42, 0.7);
                            border: 1px solid rgba(58, 58, 58, 0.5);
                            color: #FFFFFF;
                            display: flex;
                            align-items: center;
                            justify-content: center;
                            gap: 6px;
                            font-size: 13px;
                            font-weight: 600;
                            cursor: pointer;
                            transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
                        "#,
                        Icon { name: "phone".to_string(), size: "16".to_string(), color: "currentColor".to_string(), class: "".to_string() }
                        "Call"
                    }
                }

                if let Some(on_email) = props.on_email {
                    button {
                        onclick: move |e| on_email.call(e),
                        class: "action-btn",
                        style: r#"
                            flex: 1;
                            padding: 10px;
                            border-radius: 10px;
                            background: rgba(42, 42, 42, 0.7);
                            border: 1px solid rgba(58, 58, 58, 0.5);
                            color: #FFFFFF;
                            display: flex;
                            align-items: center;
                            justify-content: center;
                            gap: 6px;
                            font-size: 13px;
                            font-weight: 600;
                            cursor: pointer;
                            transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
                        "#,
                        Icon { name: "mail".to_string(), size: "16".to_string(), color: "currentColor".to_string(), class: "".to_string() }
                        "Email"
                    }
                }

                if let Some(on_view) = props.on_view {
                    button {
                        onclick: move |e| on_view.call(e),
                        class: "action-btn action-btn-primary",
                        style: r#"
                            flex: 1;
                            padding: 10px;
                            border-radius: 10px;
                            background: linear-gradient(135deg, #8B0E0E 0%, #C91A1A 100%);
                            border: 1px solid rgba(139, 14, 14, 0.5);
                            color: #FFFFFF;
                            display: flex;
                            align-items: center;
                            justify-content: center;
                            gap: 6px;
                            font-size: 13px;
                            font-weight: 600;
                            cursor: pointer;
                            transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
                        "#,
                        Icon { name: "user".to_string(), size: "16".to_string(), color: "currentColor".to_string(), class: "".to_string() }
                        "View"
                    }
                }
            }

            style {
                r#"
                .contact-card:hover {
                    transform: translateY(-4px);
                    box-shadow: 0 12px 40px rgba(0, 0, 0, 0.5);
                    border-color: rgba(139, 14, 14, 0.3);
                }
                
                .action-btn:hover {
                    transform: translateY(-2px);
                    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
                    border-color: rgba(139, 14, 14, 0.5);
                    background: rgba(139, 14, 14, 0.2);
                }
                
                .action-btn-primary:hover {
                    background: linear-gradient(135deg, #C91A1A 0%, #8B0E0E 100%);
                    box-shadow: 0 6px 20px rgba(139, 14, 14, 0.4);
                }
                "#
            }
        }
    }
}
