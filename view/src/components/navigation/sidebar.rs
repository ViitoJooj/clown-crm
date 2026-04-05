use dioxus::prelude::*;
use crate::models::User;
use crate::components::chat::ChatListItem;
use crate::components::common::Icon;

#[component]
pub fn Sidebar(user: User) -> Element {
    let mut search_query = use_signal(|| "".to_string());
    let mut show_logout_menu = use_signal(|| false);

    let sidebar_style = r#"
        width: 360px;
        background-color: #0A0A0A;
        display: flex;
        flex-direction: column;
        height: 100vh;
        overflow-y: auto;
    "#;

    let header_style = r#"
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 16px;
        background-color: #1A1A1A;
        border-bottom: 1px solid #2A2A2A;
    "#;

    let search_container_style = r#"
        padding: 16px;
        border-bottom: 1px solid #2A2A2A;
    "#;

    let search_input_style = r#"
        width: 100%;
        padding: 10px 16px;
        border-radius: 24px;
        border: none;
        background-color: #2A2A2A;
        color: #FFFFFF;
        font-size: 14px;
    "#;

    let user_profile_style = r#"
        display: flex;
        align-items: center;
        gap: 12px;
        padding: 12px 16px;
        background-color: #1A1A1A;
        border-bottom: 1px solid #2A2A2A;
    "#;

    let chats_container_style = r#"
        flex: 1;
        overflow-y: auto;
    "#;

    rsx! {
        div {
            style: sidebar_style,
            
            // Header
            div {
                style: header_style,
                h1 {
                    style: "color: #8B0E0E; margin: 0; font-size: 28px; font-weight: bold; display: flex; align-items: center;",
                    Icon { name: "message-square".to_string(), size: "28".to_string(), color: "#8B0E0E".to_string() }
                }
                div {
                    style: "display: flex; gap: 8px;",
                    button {
                        style: "background: none; border: none; color: #8B0E0E; cursor: pointer; padding: 8px; border-radius: 50%; transition: all 0.2s ease; display: flex; align-items: center;",
                        onclick: move |_| show_logout_menu.toggle(),
                        Icon { name: "more-vertical".to_string(), size: "20".to_string(), color: "#8B0E0E".to_string() }
                    }
                }
            }

            // User Profile
            div {
                style: user_profile_style,
                div {
                    style: "width: 48px; height: 48px; border-radius: 50%; background: linear-gradient(135deg, #8B0E0E, #C41E3A); display: flex; align-items: center; justify-content: center; flex-shrink: 0;",
                    span {
                        style: "color: white; font-weight: bold; font-size: 18px;",
                        "{user.first_name.chars().next().unwrap_or('U')}"
                    }
                }
                div {
                    style: "flex: 1;",
                    p {
                        style: "margin: 0 0 4px 0; color: #FFFFFF; font-weight: 600; font-size: 14px;",
                        "{user.full_name()}"
                    }
                    p {
                        style: "margin: 0; color: #A0A0A0; font-size: 12px;",
                        "{user.email}"
                    }
                }
                if user.is_admin() {
                    span {
                        style: "background-color: #8B0E0E; color: white; padding: 4px 8px; border-radius: 4px; font-size: 11px; font-weight: bold; white-space: nowrap;",
                        "Admin"
                    }
                }
            }

            // Logout Menu
            if show_logout_menu.read().clone() {
                div {
                    style: "position: fixed; top: 60px; right: 10px; background-color: #2A2A2A; border: 1px solid #3A3A3A; border-radius: 8px; z-index: 1000; box-shadow: 0 4px 12px rgba(0,0,0,0.5);",
                    button {
                        style: "width: 200px; padding: 12px 16px; background: none; border: none; color: #FFFFFF; cursor: pointer; text-align: left; font-size: 14px; transition: all 0.2s ease; display: flex; align-items: center; gap: 8px;",
                        onclick: move |_| {
                            // TODO: Implement logout
                            println!("Logout clicked");
                        },
                        Icon { name: "log-out".to_string(), size: "16".to_string() }
                        "Logout"
                    }
                }
            }

            // Search
            div {
                style: search_container_style,
                div {
                    style: "position: relative;",
                    div {
                        style: "position: absolute; left: 12px; top: 50%; transform: translateY(-50%); pointer-events: none;",
                        Icon { name: "search".to_string(), size: "16".to_string(), color: "#666".to_string() }
                    }
                    input {
                        r#type: "text",
                        placeholder: "Search chats...",
                        value: search_query.read().to_string(),
                        onchange: move |e: FormEvent| search_query.set(e.value()),
                        style: "{search_input_style} padding-left: 40px;",
                    }
                }
            }

            // Chat List
            div {
                style: chats_container_style,
                
                ChatListItem {
                    name: "John Doe".to_string(),
                    last_message: "Hey, how are you?".to_string(),
                    unread_count: 2,
                    is_active: true,
                    onclick: move |_| {}
                }
                
                ChatListItem {
                    name: "Team Dev".to_string(),
                    last_message: "Project meeting at 3pm".to_string(),
                    unread_count: 5,
                    is_active: false,
                    onclick: move |_| {}
                }
                
                ChatListItem {
                    name: "Sarah Smith".to_string(),
                    last_message: "Thanks for the update!".to_string(),
                    unread_count: 0,
                    is_active: false,
                    onclick: move |_| {}
                }

                ChatListItem {
                    name: "Support Team".to_string(),
                    last_message: "Issue has been resolved".to_string(),
                    unread_count: 0,
                    is_active: false,
                    onclick: move |_| {}
                }
            }
        }
    }
}
