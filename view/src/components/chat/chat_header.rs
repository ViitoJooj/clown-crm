use crate::components::common::Icon;
use dioxus::prelude::*;

#[component]
pub fn ChatHeader(
    #[props(default = "".to_string())] chat_name: String,
    #[props(default = "".to_string())] status: String,
) -> Element {
    rsx! {
        div {
            style: "display: flex; justify-content: space-between; align-items: center; padding: 16px; background-color: #1A1A1A; border-bottom: 1px solid #2A2A2A;",

            div {
                style: "display: flex; align-items: center; gap: 12px;",

                div {
                    style: "width: 40px; height: 40px; border-radius: 50%; background: linear-gradient(135deg, #8B0E0E, #C41E3A); display: flex; align-items: center; justify-content: center;",
                    span {
                        style: "color: white; font-weight: bold;",
                        "{chat_name.chars().next().unwrap_or('?')}"
                    }
                }

                div {
                    h3 {
                        style: "margin: 0 0 2px 0; color: #FFFFFF; font-size: 16px; font-weight: 600;",
                        "{chat_name}"
                    }
                    p {
                        style: "margin: 0; color: #A0A0A0; font-size: 12px;",
                        "{status}"
                    }
                }
            }

            div {
                style: "display: flex; gap: 8px;",
                button {
                    style: "background: none; border: none; color: #8B0E0E; cursor: pointer; padding: 8px; display: flex; align-items: center; justify-content: center;",
                    Icon { name: "search".to_string(), size: "18".to_string(), color: "#8B0E0E".to_string() }
                }
                button {
                    style: "background: none; border: none; color: #8B0E0E; cursor: pointer; padding: 8px; display: flex; align-items: center; justify-content: center;",
                    Icon { name: "more-vertical".to_string(), size: "18".to_string(), color: "#8B0E0E".to_string() }
                }
            }
        }
    }
}
