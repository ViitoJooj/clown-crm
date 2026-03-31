use dioxus::prelude::*;
use crate::components::common::*;
use crate::utils::api_client::ApiClient;
use crate::utils::storage;
use crate::models::User;
use crate::hooks::{use_auth, set_authenticated_user, AuthState, Signal};

#[component]
pub fn LoginPage(
    #[props(default = "".to_string())] initial_backend_url: String,
    auth_signal: Signal<AuthState>,
) -> Element {
    let mut backend_url = use_signal(|| initial_backend_url);
    let mut email = use_signal(|| "".to_string());
    let mut password = use_signal(|| "".to_string());
    let mut loading = use_signal(|| false);
    let mut error = use_signal(|| "".to_string());

    let handle_login = move |_event: MouseEvent| {
        if backend_url.read().is_empty() {
            error.set("Please enter backend URL".to_string());
            return;
        }
        if email.read().is_empty() || password.read().is_empty() {
            error.set("Please fill in all fields".to_string());
            return;
        }

        loading.set(true);
        error.set("".to_string());

        let url = backend_url.read().to_string();
        let em = email.read().to_string();
        let pwd = password.read().to_string();

        spawn(async move {
            let client = ApiClient::new(url);
            
            match client.login(em, pwd).await {
                Ok(response) => {
                    if response.success {
                        if let Some(token) = response.token {
                            let _ = storage::save_token(&token);
                            let _ = storage::save_backend_url(&client.get_base_url());
                            
                            if let Some(user_dto) = response.user {
                                let user = User {
                                    uuid: user_dto.uuid,
                                    first_name: user_dto.first_name,
                                    last_name: user_dto.last_name,
                                    email: user_dto.email,
                                    role: user_dto.role,
                                    profile_picture: user_dto.profile_picture,
                                    updated_at: user_dto.updated_at,
                                    created_at: user_dto.created_at,
                                };
                                set_authenticated_user(auth_signal, user.clone());
                            }
                        }
                    } else {
                        error.set(response.message);
                        loading.set(false);
                    }
                }
                Err(e) => {
                    error.set(format!("Login failed: {}", e.message));
                    loading.set(false);
                }
            }
        });
    };

    rsx! {
        div {
            style: "display: flex; justify-content: center; align-items: center; height: 100vh; background: linear-gradient(180deg, #0A0A0A 0%, #1A0505 100%); position: relative; overflow: hidden;",
            
            // Animated background elements
            div {
                style: "position: absolute; inset: 0; overflow: hidden; pointer-events: none;",
                div {
                    style: "position: absolute; top: 10%; left: 10%; width: 300px; height: 300px; background: radial-gradient(circle, rgba(139, 14, 14, 0.15) 0%, transparent 70%); border-radius: 50%; filter: blur(60px); animation: float 20s ease-in-out infinite;",
                }
                div {
                    style: "position: absolute; bottom: 10%; right: 10%; width: 400px; height: 400px; background: radial-gradient(circle, rgba(201, 26, 26, 0.1) 0%, transparent 70%); border-radius: 50%; filter: blur(80px); animation: float 25s ease-in-out infinite reverse;",
                }
            }
            
            div {
                style: "width: 100%; max-width: 440px; padding: 20px; position: relative; z-index: 1; animation: fadeIn 0.8s ease;",
                
                div {
                    style: "text-align: center; margin-bottom: 40px;",
                    div {
                        style: "display: inline-block; background: linear-gradient(135deg, #8B0E0E, #C91A1A); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; margin-bottom: 12px;",
                        h1 {
                            style: "font-size: 42px; font-weight: 900; margin: 0; letter-spacing: -1px;",
                            "Clown CRM"
                        }
                    }
                    p {
                        style: "color: #707070; font-size: 16px; margin: 0;",
                        "Team Communication Platform"
                    }
                }

                Card {
                    glass: true,
                    div {
                        style: "display: flex; flex-direction: column; gap: 20px;",
                        
                        if !error.read().is_empty() {
                            div {
                                style: "background: linear-gradient(135deg, rgba(239, 68, 68, 0.15), rgba(220, 38, 38, 0.1)); backdrop-filter: blur(10px); border: 1px solid rgba(239, 68, 68, 0.3); color: #FF6B6B; padding: 14px 16px; border-radius: 12px; font-size: 14px; display: flex; align-items: center; gap: 10px; animation: slideInRight 0.3s ease;",
                                span {
                                    style: "font-size: 18px;",
                                    "⚠️"
                                }
                                span {
                                    "{error}"
                                }
                            }
                        }

                        div {
                            label {
                                style: "display: block; color: #B0B0B0; font-size: 13px; font-weight: 600; margin-bottom: 8px; text-transform: uppercase; letter-spacing: 0.5px;",
                                "Backend URL"
                            }
                            Input {
                                input_type: "text".to_string(),
                                placeholder: "http://localhost:8080".to_string(),
                                value: backend_url.read().to_string(),
                                onchange: move |e: FormEvent| { backend_url.set(e.value()); }
                            }
                        }

                        div {
                            label {
                                style: "display: block; color: #B0B0B0; font-size: 13px; font-weight: 600; margin-bottom: 8px; text-transform: uppercase; letter-spacing: 0.5px;",
                                "Email"
                            }
                            Input {
                                input_type: "email".to_string(),
                                placeholder: "your@email.com".to_string(),
                                value: email.read().to_string(),
                                onchange: move |e: FormEvent| { email.set(e.value()); }
                            }
                        }

                        div {
                            label {
                                style: "display: block; color: #B0B0B0; font-size: 13px; font-weight: 600; margin-bottom: 8px; text-transform: uppercase; letter-spacing: 0.5px;",
                                "Password"
                            }
                            Input {
                                input_type: "password".to_string(),
                                placeholder: "••••••••".to_string(),
                                value: password.read().to_string(),
                                onchange: move |e: FormEvent| { password.set(e.value()); }
                            }
                        }

                        Button {
                            disabled: loading.read().clone(),
                            class: "primary".to_string(),
                            onclick: handle_login,
                            if loading.read().clone() {
                                div {
                                    style: "display: flex; align-items: center; gap: 10px;",
                                    LoadingSpinner { size: "18px".to_string() }
                                    span { "Logging in..." }
                                }
                            } else {
                                div {
                                    style: "display: flex; align-items: center; gap: 10px;",
                                    span { "🔐" }
                                    span { "Login" }
                                }
                            }
                        }

                        Divider {}

                        div {
                            style: "text-align: center; color: #707070; font-size: 14px;",
                            "Don't have an account? "
                            a {
                                href: "/register",
                                style: "color: #C91A1A; font-weight: 600; text-decoration: none; transition: color 0.3s ease;",
                                "Sign up"
                            }
                        }
                    }
                }
            }
        }
    }
}
