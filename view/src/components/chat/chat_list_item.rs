use dioxus::prelude::*;

#[component]
pub fn ChatListItem(
    #[props(default = "".to_string())] name: String,
    #[props(default = "".to_string())] last_message: String,
    #[props(default = 0)] unread_count: i32,
    #[props(default = false)] is_active: bool,
    onclick: EventHandler<MouseEvent>,
) -> Element {
    let item_style = if is_active {
        r#"
        display: flex;
        align-items: center;
        gap: 12px;
        padding: 12px 16px;
        background-color: #2A2A2A;
        cursor: pointer;
        border-left: 4px solid #8B0E0E;
        transition: all 0.2s ease;
        "#
    } else {
        r#"
        display: flex;
        align-items: center;
        gap: 12px;
        padding: 12px 16px;
        background-color: transparent;
        cursor: pointer;
        border-left: 4px solid transparent;
        transition: all 0.2s ease;
        "#
    };

    rsx! {
        div {
            style: item_style,
            onclick: move |e| onclick.call(e),
            
            // Avatar
            div {
                style: "width: 48px; height: 48px; border-radius: 50%; background: linear-gradient(135deg, #8B0E0E, #C41E3A); display: flex; align-items: center; justify-content: center; flex-shrink: 0;",
                span {
                    style: "color: white; font-weight: bold; font-size: 18px;",
                    "{name.chars().next().unwrap_or('?')}"
                }
            }
            
            // Chat info
            div {
                style: "flex: 1; min-width: 0;",
                div {
                    style: "display: flex; justify-content: space-between; align-items: center; margin-bottom: 4px;",
                    span {
                        style: "color: #FFFFFF; font-weight: 600; font-size: 15px;",
                        "{name}"
                    }
                    if unread_count > 0 {
                        span {
                            style: "background-color: #8B0E0E; color: white; border-radius: 12px; padding: 2px 8px; font-size: 12px; font-weight: bold;",
                            "{unread_count}"
                        }
                    }
                }
                p {
                    style: "margin: 0; color: #A0A0A0; font-size: 13px; white-space: nowrap; overflow: hidden; text-overflow: ellipsis;",
                    "{last_message}"
                }
            }
        }
    }
}
