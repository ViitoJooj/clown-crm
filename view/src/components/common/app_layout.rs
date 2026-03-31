use dioxus::prelude::*;
use crate::models::{User, Message};
use crate::utils::storage;
use crate::utils::api_client::ApiClient;
use crate::hooks::{AuthState, set_authenticated_user};

#[component]
pub fn AppLayout(user: User, auth_signal: Signal<AuthState>) -> Element {
    // Clone UUID once to avoid borrow issues
    let user_uuid = user.uuid.clone();
    let mut selected_chat = use_signal(|| Option::<(String, String)>::None);
    let mut show_menu = use_signal(|| false);
    let mut show_profile_edit = use_signal(|| false);
    let mut show_create_account = use_signal(|| false);
    let mut profile_pic_url = use_signal(|| user.profile_picture.clone().unwrap_or_default());
    let mut users_list = use_signal(|| Vec::<User>::new());
    let mut messages = use_signal(|| Vec::<Message>::new());
    let mut message_input = use_signal(|| String::new());
    let mut loading_users = use_signal(|| true);
    let mut current_user_state = use_signal(|| user.clone());
    let mut first_name_edit = use_signal(|| user.first_name.clone());
    let mut last_name_edit = use_signal(|| user.last_name.clone());
    let mut new_account_email = use_signal(|| String::new());
    let mut new_account_password = use_signal(|| String::new());
    let mut new_account_first = use_signal(|| String::new());
    let mut new_account_last = use_signal(|| String::new());

    // Load users on mount
    use_effect(move || {
        let base_url = storage::get_backend_url().unwrap_or_default();
        let current_user_uuid = user.uuid.clone();
        spawn(async move {
            let client = ApiClient::new(base_url);
            match client.get_all_users().await {
                Ok(fetched_users) => {
                    users_list.set(fetched_users.into_iter().filter(|u| u.uuid != current_user_uuid).collect());
                    loading_users.set(false);
                }
                Err(_) => {
                    loading_users.set(false);
                }
            }
        });
    });

    rsx! {
        div {
            style: "display: flex; height: 100vh; background-color: #0A0A0A; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;",
            
            // SIDEBAR
            div {
                style: "width: 360px; background-color: #0A0A0A; display: flex; flex-direction: column; height: 100vh; overflow: hidden; border-right: 1px solid #2A2A2A;",
                
                // Header with Menu
                div {
                    style: "padding: 12px 16px; background-color: #1A1A1A; border-bottom: 1px solid #2A2A2A; display: flex; justify-content: space-between; align-items: center; position: relative;",
                    h1 {
                        style: "color: #8B0E0E; margin: 0; font-size: 24px; font-weight: bold;",
                        "💬 Clown CRM"
                    }
                    div {
                        style: "position: relative;",
                        button {
                            style: "background: none; border: none; color: #8B0E0E; cursor: pointer; font-size: 20px; padding: 8px; border-radius: 50%; hover:background-color: #2A2A2A;",
                            onclick: move |_| {
                                show_menu.toggle();
                            },
                            "⋮"
                        }
                        
                        if show_menu() {
                            div {
                                style: "position: absolute; top: 100%; right: 0; background-color: #2A2A2A; border: 1px solid #3A3A3A; border-radius: 8px; min-width: 150px; z-index: 100; box-shadow: 0 4px 12px rgba(0,0,0,0.5);",
                                
                                button {
                                    style: "display: block; width: 100%; text-align: left; padding: 12px 16px; background: none; border: none; color: #FFFFFF; cursor: pointer; hover:background-color: #3A3A3A; border-bottom: 1px solid #3A3A3A; font-size: 14px;",
                                    onclick: move |_| {
                                        show_profile_edit.toggle();
                                        show_menu.set(false);
                                    },
                                    "✏️ Edit Profile"
                                }
                                
                                if current_user_state.read().is_admin() {
                                    button {
                                        style: "display: block; width: 100%; text-align: left; padding: 12px 16px; background: none; border: none; color: #FFFFFF; cursor: pointer; hover:background-color: #3A3A3A; border-bottom: 1px solid #3A3A3A; font-size: 14px;",
                                        onclick: move |_| {
                                            show_create_account.set(true);
                                            show_menu.set(false);
                                        },
                                        "➕ Create Account"
                                    }
                                }
                                
                                button {
                                    style: "display: block; width: 100%; text-align: left; padding: 12px 16px; background: none; border: none; color: #FFFFFF; cursor: pointer; hover:background-color: #3A3A3A; font-size: 14px;",
                                    onclick: move |_| {
                                        storage::clear_token();
                                        storage::clear_all();
                                        auth_signal.set(AuthState::Unauthenticated);
                                    },
                                    "🚪 Logout"
                                }
                            }
                        }
                    }
                }
                
                // User Profile
                div {
                    style: "padding: 12px 16px; background-color: #1A1A1A; display: flex; gap: 12px; align-items: center; border-bottom: 1px solid #2A2A2A; cursor: pointer; hover:background-color: #2A2A2A;",
                    onclick: move |_| {
                        show_profile_edit.toggle();
                    },
                    
                    div {
                        style: if !profile_pic_url().is_empty() {
                            format!("width: 40px; height: 40px; border-radius: 50%; background-image: url('{}'); background-size: cover; flex-shrink: 0;", profile_pic_url())
                        } else {
                            "width: 40px; height: 40px; border-radius: 50%; background: linear-gradient(135deg, #8B0E0E, #C41E3A); display: flex; align-items: center; justify-content: center; color: white; font-weight: bold; flex-shrink: 0;".to_string()
                        },
                        {
                            if profile_pic_url().is_empty() {
                                let initial = user.first_name.chars().next().unwrap_or('U');
                                rsx! { "{initial}" }
                            } else {
                                rsx! {}
                            }
                        }
                    }
                    
                    div {
                        style: "flex: 1; min-width: 0;",
                        p {
                            style: "margin: 0; color: #FFFFFF; font-weight: 600; font-size: 14px; white-space: nowrap; overflow: hidden; text-overflow: ellipsis;",
                            "{user.first_name}"
                        }
                        p {
                            style: "margin: 4px 0 0 0; color: #A0A0A0; font-size: 11px;",
                            "Click to edit"
                        }
                    }
                }
                
                // Search Bar
                div {
                    style: "padding: 12px 16px; background-color: #0A0A0A; display: flex; gap: 8px;",
                    input {
                        r#type: "text",
                        placeholder: "Search chats...",
                        style: "flex: 1; padding: 8px 12px; border-radius: 20px; border: none; background-color: #2A2A2A; color: #FFFFFF; font-size: 14px;",
                    }
                    button {
                        style: "background: none; border: none; color: #8B0E0E; cursor: pointer; font-size: 18px; padding: 4px 8px;",
                        "➕"
                    }
                }
                
                // Chat List - Real Users
                div {
                    style: "flex: 1; overflow-y: auto; background-color: #0A0A0A;",
                    if loading_users() {
                        div {
                            style: "padding: 20px; text-align: center; color: #A0A0A0;",
                            "Loading users..."
                        }
                    } else {
                        for other_user in users_list.read().iter() {
                            div {
                                key: "{other_user.uuid}",
                                style: if let Some((selected_id, _)) = selected_chat.read().as_ref() {
                                    if selected_id == &other_user.uuid {
                                        "padding: 12px 8px; margin: 0 8px; display: flex; gap: 12px; align-items: center; border-radius: 8px; cursor: pointer; background-color: #1A1A1A;"
                                    } else {
                                        "padding: 12px 8px; margin: 0 8px; display: flex; gap: 12px; align-items: center; border-radius: 8px; cursor: pointer; hover:background-color: #1A1A1A;"
                                    }
                                } else {
                                    "padding: 12px 8px; margin: 0 8px; display: flex; gap: 12px; align-items: center; border-radius: 8px; cursor: pointer; hover:background-color: #1A1A1A;"
                                },
                                onclick: {
                                    let user_uuid = other_user.uuid.clone();
                                    let user_name = other_user.first_name.clone();
                                    move |_| {
                                        selected_chat.set(Some((user_uuid.clone(), user_name.clone())));
                                        messages.clear();
                                    }
                                },
                                
                                div {
                                    style: "width: 48px; height: 48px; border-radius: 50%; background: linear-gradient(135deg, #8B0E0E, #C41E3A); display: flex; align-items: center; justify-content: center; color: white; font-weight: bold; flex-shrink: 0;",
                                    "{other_user.first_name.chars().next().unwrap_or('U')}"
                                }
                                
                                div {
                                    style: "flex: 1; min-width: 0;",
                                    p {
                                        style: "margin: 0; color: #FFFFFF; font-weight: 500; font-size: 14px; white-space: nowrap; overflow: hidden; text-overflow: ellipsis;",
                                        "{other_user.first_name}"
                                    }
                                    p {
                                        style: "margin: 4px 0 0 0; color: #A0A0A0; font-size: 12px; white-space: nowrap; overflow: hidden; text-overflow: ellipsis;",
                                        "Online"
                                    }
                                }
                            }
                        }
                    }
                }
            }
            
            // CHAT AREA
            div {
                style: "flex: 1; display: flex; flex-direction: column; overflow: hidden; background-color: #0A0A0A;",
                
                div {
                    style: "display: flex; justify-content: space-between; align-items: center; padding: 16px; background-color: #1A1A1A; border-bottom: 1px solid #2A2A2A;",
                    {
                        if let Some((_, chat_user_name)) = selected_chat.read().as_ref() {
                            rsx! {
                                div {
                                    h3 {
                                        style: "margin: 0; color: #FFFFFF; font-size: 16px; font-weight: 600;",
                                        "{chat_user_name}"
                                    }
                                    p {
                                        style: "margin: 4px 0 0 0; color: #A0A0A0; font-size: 12px;",
                                        "Online"
                                    }
                                }
                                div {
                                    style: "display: flex; gap: 8px;",
                                    button {
                                        style: "background: none; border: none; color: #8B0E0E; cursor: pointer; font-size: 14px; padding: 8px 12px; border-radius: 4px;",
                                        "📞 Call"
                                    }
                                    button {
                                        style: "background: none; border: none; color: #8B0E0E; cursor: pointer; font-size: 14px; padding: 8px 12px; border-radius: 4px;",
                                        "🖥️ Share"
                                    }
                                }
                            }
                        } else {
                            rsx! {
                                h3 {
                                    style: "margin: 0; color: #A0A0A0; font-size: 16px;",
                                    "Select a user to start"
                                }
                            }
                        }
                    }
                }
                
                div {
                    style: "flex: 1; overflow-y: auto; display: flex; flex-direction: column; padding: 20px; gap: 16px;",
                    {
                        if selected_chat.read().is_none() {
                            rsx! {
                                div {
                                    style: "display: flex; flex-direction: column; justify-content: center; align-items: center; height: 100%; text-align: center;",
                                    h2 {
                                        style: "color: #8B0E0E; margin-bottom: 10px; font-size: 32px;",
                                        "👋"
                                    }
                                    p {
                                        style: "color: #A0A0A0; margin: 10px 0; font-size: 14px;",
                                        "Select a user from the left"
                                    }
                                }
                            }
                        } else {
                            rsx! {
                                for (idx, msg) in messages.read().iter().enumerate() {
                                    div {
                                        key: "{idx}",
                                        style: "display: flex; justify-content: flex-start; gap: 8px;",
                                        
                                        div {
                                            style: "width: 32px; height: 32px; border-radius: 50%; background: linear-gradient(135deg, #8B0E0E, #C41E3A); flex-shrink: 0;"
                                        }
                                        
                                        div {
                                            style: "background-color: #1A1A1A; color: white; padding: 12px 16px; border-radius: 12px 12px 12px 4px; max-width: 60%; word-wrap: break-word;",
                                            p {
                                                style: "margin: 0; font-size: 14px;",
                                                "{msg.message}"
                                            }
                                            span {
                                                style: "margin-top: 4px; display: block; font-size: 11px; opacity: 0.7;",
                                                "{msg.time}"
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
                
                div {
                    style: "display: flex; align-items: flex-end; gap: 12px; padding: 12px 16px; background-color: #1A1A1A; border-top: 1px solid #2A2A2A;",
                    button {
                        style: "background: none; border: none; color: #8B0E0E; cursor: pointer; font-size: 18px; padding: 8px; border-radius: 50%;",
                        "📎"
                    }
                    input {
                        r#type: "text",
                        placeholder: "Message...",
                        value: message_input.read().to_string(),
                        onchange: move |e: FormEvent| message_input.set(e.value()),
                        style: "flex: 1; padding: 10px 16px; border-radius: 24px; border: none; background-color: #2A2A2A; color: white; font-size: 14px;",
                    }
                    button {
                        style: "background: none; border: none; color: #8B0E0E; cursor: pointer; font-size: 18px; padding: 8px; border-radius: 50%;",
                        "😊"
                    }
                    button {
                        style: "background: none; border: none; color: #8B0E0E; cursor: pointer; font-size: 18px; padding: 8px; border-radius: 50%;",
                        "📤"
                    }
                }
            }
            
            if show_profile_edit() {
                div {
                    style: "position: fixed; top: 0; left: 0; right: 0; bottom: 0; background-color: rgba(0,0,0,0.7); display: flex; align-items: center; justify-content: center; z-index: 1000;",
                    onclick: move |_| {
                        show_profile_edit.set(false);
                    },
                    
                    div {
                        style: "background-color: #1A1A1A; border-radius: 12px; padding: 24px; width: 90%; max-width: 400px;",
                        onclick: move |e: MouseEvent| {
                            e.stop_propagation();
                        },
                        
                        h2 {
                            style: "margin: 0 0 20px 0; color: white; font-size: 18px;",
                            "Edit Profile"
                        }
                        
                        div {
                            style: "display: flex; flex-direction: column; gap: 16px;",
                            
                            div {
                                style: "display: flex; flex-direction: column; align-items: center; gap: 12px;",
                                
                                div {
                                    style: if !profile_pic_url().is_empty() {
                                        format!("width: 80px; height: 80px; border-radius: 50%; background-image: url('{}'); background-size: cover;", profile_pic_url())
                                    } else {
                                        "width: 80px; height: 80px; border-radius: 50%; background: linear-gradient(135deg, #8B0E0E, #C41E3A); display: flex; align-items: center; justify-content: center; color: white; font-weight: bold; font-size: 32px;".to_string()
                                    },
                                    {
                                        if profile_pic_url().is_empty() {
                                            let initial = current_user_state.read().first_name.chars().next().unwrap_or('U');
                                            rsx! { "{initial}" }
                                        } else {
                                            rsx! {}
                                        }
                                    }
                                }
                                
                                input {
                                    r#type: "file",
                                    accept: "image/*",
                                    style: "display: none;",
                                    id: "profile-pic-input",
                                }
                                
                                label {
                                    r#for: "profile-pic-input",
                                    style: "background-color: #8B0E0E; color: white; padding: 8px 16px; border-radius: 6px; cursor: pointer; font-size: 13px;",
                                    "📷 Upload Photo"
                                }
                            }
                            
                            div {
                                style: "display: flex; flex-direction: column; gap: 12px;",
                                
                                div {
                                    style: "display: flex; flex-direction: column; gap: 4px;",
                                    label {
                                        style: "color: #A0A0A0; font-size: 12px; font-weight: 500;",
                                        "First Name"
                                    }
                                    input {
                                        r#type: "text",
                                        value: first_name_edit.read().to_string(),
                                        onchange: move |e: FormEvent| first_name_edit.set(e.value()),
                                        placeholder: "First name",
                                        style: "padding: 10px 12px; border-radius: 6px; border: 1px solid #3A3A3A; background-color: #2A2A2A; color: white; font-size: 13px;",
                                    }
                                }
                                
                                div {
                                    style: "display: flex; flex-direction: column; gap: 4px;",
                                    label {
                                        style: "color: #A0A0A0; font-size: 12px; font-weight: 500;",
                                        "Last Name"
                                    }
                                    input {
                                        r#type: "text",
                                        value: last_name_edit.read().to_string(),
                                        onchange: move |e: FormEvent| last_name_edit.set(e.value()),
                                        placeholder: "Last name",
                                        style: "padding: 10px 12px; border-radius: 6px; border: 1px solid #3A3A3A; background-color: #2A2A2A; color: white; font-size: 13px;",
                                    }
                                }
                            }
                            
                            button {
                                style: "width: 100%; padding: 12px; background-color: #8B0E0E; color: white; border: none; border-radius: 6px; cursor: pointer; font-weight: 500;",
                                onclick: move |_| {
                                    let base_url = storage::get_backend_url().unwrap_or_default();
                                    let first_name = first_name_edit.read().to_string();
                                    let last_name = last_name_edit.read().to_string();
                                    
                                    spawn(async move {
                                        let client = ApiClient::new(base_url);
                                        match client.update_user(Some(first_name.clone()), Some(last_name.clone())).await {
                                            Ok(updated_user) => {
                                                current_user_state.set(updated_user.clone());
                                                if let Ok(user_json) = serde_json::to_string(&updated_user) {
                                                    let _ = storage::save_current_user(&user_json);
                                                }
                                                set_authenticated_user(auth_signal, updated_user);
                                                show_profile_edit.set(false);
                                            }
                                            Err(_) => {
                                                // Show error
                                            }
                                        }
                                    });
                                },
                                "Save Changes"
                            }
                            
                            button {
                                style: "width: 100%; padding: 10px; background-color: transparent; color: #A0A0A0; border: 1px solid #3A3A3A; border-radius: 6px; cursor: pointer; font-size: 13px;",
                                onclick: move |_| {
                                    show_profile_edit.set(false);
                                },
                                "Cancel"
                            }
                        }
                    }
                }
            }
            
            if show_create_account() {
                div {
                    style: "position: fixed; top: 0; left: 0; right: 0; bottom: 0; background-color: rgba(0,0,0,0.7); display: flex; align-items: center; justify-content: center; z-index: 1000;",
                    onclick: move |_| {
                        show_create_account.set(false);
                    },
                    
                    div {
                        style: "background-color: #1A1A1A; border-radius: 12px; padding: 24px; width: 90%; max-width: 400px;",
                        onclick: move |e: MouseEvent| {
                            e.stop_propagation();
                        },
                        
                        h2 {
                            style: "margin: 0 0 20px 0; color: white; font-size: 18px;",
                            "Create New Account"
                        }
                        
                        div {
                            style: "display: flex; flex-direction: column; gap: 12px;",
                            
                            div {
                                style: "display: flex; flex-direction: column; gap: 4px;",
                                label {
                                    style: "color: #A0A0A0; font-size: 12px; font-weight: 500;",
                                    "Email"
                                }
                                input {
                                    r#type: "email",
                                    value: new_account_email.read().to_string(),
                                    onchange: move |e: FormEvent| new_account_email.set(e.value()),
                                    placeholder: "user@example.com",
                                    style: "padding: 10px 12px; border-radius: 6px; border: 1px solid #3A3A3A; background-color: #2A2A2A; color: white; font-size: 13px;",
                                }
                            }
                            
                            div {
                                style: "display: flex; flex-direction: column; gap: 4px;",
                                label {
                                    style: "color: #A0A0A0; font-size: 12px; font-weight: 500;",
                                    "Password"
                                }
                                input {
                                    r#type: "password",
                                    value: new_account_password.read().to_string(),
                                    onchange: move |e: FormEvent| new_account_password.set(e.value()),
                                    placeholder: "Enter password",
                                    style: "padding: 10px 12px; border-radius: 6px; border: 1px solid #3A3A3A; background-color: #2A2A2A; color: white; font-size: 13px;",
                                }
                            }
                            
                            div {
                                style: "display: flex; flex-direction: column; gap: 4px;",
                                label {
                                    style: "color: #A0A0A0; font-size: 12px; font-weight: 500;",
                                    "First Name"
                                }
                                input {
                                    r#type: "text",
                                    value: new_account_first.read().to_string(),
                                    onchange: move |e: FormEvent| new_account_first.set(e.value()),
                                    placeholder: "First name",
                                    style: "padding: 10px 12px; border-radius: 6px; border: 1px solid #3A3A3A; background-color: #2A2A2A; color: white; font-size: 13px;",
                                }
                            }
                            
                            div {
                                style: "display: flex; flex-direction: column; gap: 4px;",
                                label {
                                    style: "color: #A0A0A0; font-size: 12px; font-weight: 500;",
                                    "Last Name"
                                }
                                input {
                                    r#type: "text",
                                    value: new_account_last.read().to_string(),
                                    onchange: move |e: FormEvent| new_account_last.set(e.value()),
                                    placeholder: "Last name",
                                    style: "padding: 10px 12px; border-radius: 6px; border: 1px solid #3A3A3A; background-color: #2A2A2A; color: white; font-size: 13px;",
                                }
                            }
                            
                            button {
                                style: "width: 100%; padding: 12px; background-color: #8B0E0E; color: white; border: none; border-radius: 6px; cursor: pointer; font-weight: 500;",
                                onclick: move |_| {
                                    let base_url = storage::get_backend_url().unwrap_or_default();
                                    let email = new_account_email.read().to_string();
                                    let password = new_account_password.read().to_string();
                                    let first = new_account_first.read().to_string();
                                    let last = new_account_last.read().to_string();
                                    
                                    if !email.is_empty() && !password.is_empty() && !first.is_empty() && !last.is_empty() {
                                        spawn(async move {
                                            let client = ApiClient::new(base_url);
                                            match client.create_user(email, password, first, last).await {
                                                Ok(_) => {
                                                    new_account_email.set(String::new());
                                                    new_account_password.set(String::new());
                                                    new_account_first.set(String::new());
                                                    new_account_last.set(String::new());
                                                    show_create_account.set(false);
                                                }
                                                Err(_) => {
                                                    // Show error
                                                }
                                            }
                                        });
                                    }
                                },
                                "Create Account"
                            }
                            
                            button {
                                style: "width: 100%; padding: 10px; background-color: transparent; color: #A0A0A0; border: 1px solid #3A3A3A; border-radius: 6px; cursor: pointer; font-size: 13px;",
                                onclick: move |_| {
                                    show_create_account.set(false);
                                },
                                "Cancel"
                            }
                        }
                    }
                }
            }
        }
    }
}
