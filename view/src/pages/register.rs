use dioxus::prelude::*;
use crate::components::common::*;
use crate::utils::api_client::ApiClient;
use crate::utils::storage;

#[component]
pub fn RegisterPage() -> Element {
    let mut backend_url = use_signal(|| storage::get_backend_url().unwrap_or_default());
    let mut first_name = use_signal(|| "".to_string());
    let mut last_name = use_signal(|| "".to_string());
    let mut email = use_signal(|| "".to_string());
    let mut password = use_signal(|| "".to_string());
    let mut password_confirm = use_signal(|| "".to_string());
    let mut loading = use_signal(|| false);
    let mut error = use_signal(|| "".to_string());
    let mut success = use_signal(|| false);

    let handle_register = move |_event: MouseEvent| {
        if backend_url.read().is_empty() {
            error.set("Please enter backend URL".to_string());
            return;
        }
        if first_name.read().is_empty() || last_name.read().is_empty() {
            error.set("Please enter your name".to_string());
            return;
        }
        if email.read().is_empty() || password.read().is_empty() {
            error.set("Please fill in all fields".to_string());
            return;
        }
        if password.read().to_string() != password_confirm.read().to_string() {
            error.set("Passwords do not match".to_string());
            return;
        }
        if password.read().len() < 6 {
            error.set("Password must be at least 6 characters".to_string());
            return;
        }

        loading.set(true);
        error.set("".to_string());

        let url = backend_url.read().to_string();
        let fname = first_name.read().to_string();
        let lname = last_name.read().to_string();
        let em = email.read().to_string();
        let pwd = password.read().to_string();

        spawn(async move {
            let client = ApiClient::new(url);
            
            match client.register(fname, lname, em, pwd).await {
                Ok(response) => {
                    if response.success {
                        success.set(true);
                        error.set("Account created! Please wait for admin approval.".to_string());
                    } else {
                        error.set(response.message);
                        loading.set(false);
                    }
                }
                Err(e) => {
                    error.set(format!("Registration failed: {}", e.message));
                    loading.set(false);
                }
            }
        });
    };

    rsx! {
        div {
            style: "display: flex; justify-content: center; align-items: center; min-height: 100vh; background-color: #0A0A0A; padding: 20px;",
            div {
                style: "width: 100%; max-width: 500px;",
                div {
                    style: "text-align: center; margin-bottom: 30px;",
                    h1 {
                        style: "color: #8B0E0E; margin-bottom: 10px;",
                        "Create Account"
                    }
                    p {
                        style: "color: #A0A0A0;",
                        "Join Clown CRM"
                    }
                }

                Card {
                    div {
                        style: "display: flex; flex-direction: column; gap: 16px;",
                        
                        if !error.read().is_empty() {
                            div {
                                style: if success.read().clone() {
                                    "background-color: rgba(16, 185, 129, 0.1); border: 1px solid #10B981; color: #6EE7B7; padding: 12px; border-radius: 6px;"
                                } else {
                                    "background-color: rgba(239, 68, 68, 0.1); border: 1px solid #EF4444; color: #FF6B6B; padding: 12px; border-radius: 6px;"
                                },
                                "{error}"
                            }
                        }

                        if !success.read().clone() {
                            div {
                                label {
                                    style: "display: block; color: #A0A0A0; font-size: 14px; margin-bottom: 6px;",
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
                                style: "display: grid; grid-template-columns: 1fr 1fr; gap: 12px;",
                                div {
                                    label {
                                        style: "display: block; color: #A0A0A0; font-size: 14px; margin-bottom: 6px;",
                                        "First Name"
                                    }
                                    Input {
                                        input_type: "text".to_string(),
                                        placeholder: "John".to_string(),
                                        value: first_name.read().to_string(),
                                        onchange: move |e: FormEvent| { first_name.set(e.value()); }
                                    }
                                }
                                div {
                                    label {
                                        style: "display: block; color: #A0A0A0; font-size: 14px; margin-bottom: 6px;",
                                        "Last Name"
                                    }
                                    Input {
                                        input_type: "text".to_string(),
                                        placeholder: "Doe".to_string(),
                                        value: last_name.read().to_string(),
                                        onchange: move |e: FormEvent| { last_name.set(e.value()); }
                                    }
                                }
                            }

                            div {
                                label {
                                    style: "display: block; color: #A0A0A0; font-size: 14px; margin-bottom: 6px;",
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
                                    style: "display: block; color: #A0A0A0; font-size: 14px; margin-bottom: 6px;",
                                    "Password"
                                }
                                Input {
                                    input_type: "password".to_string(),
                                    placeholder: "••••••••".to_string(),
                                    value: password.read().to_string(),
                                    onchange: move |e: FormEvent| { password.set(e.value()); }
                                }
                            }

                            div {
                                label {
                                    style: "display: block; color: #A0A0A0; font-size: 14px; margin-bottom: 6px;",
                                    "Confirm Password"
                                }
                                Input {
                                    input_type: "password".to_string(),
                                    placeholder: "••••••••".to_string(),
                                    value: password_confirm.read().to_string(),
                                    onchange: move |e: FormEvent| { password_confirm.set(e.value()); }
                                }
                            }

                            Button {
                                disabled: loading.read().clone(),
                                class: "primary".to_string(),
                                onclick: handle_register,
                                if loading.read().clone() {
                                    LoadingSpinner {}
                                    "Creating Account..."
                                } else {
                                    "Create Account"
                                }
                            }
                        }

                        div {
                            style: "text-align: center; color: #A0A0A0; font-size: 14px;",
                            if success.read().clone() {
                                a {
                                    href: "/login",
                                    style: "color: #8B0E0E; cursor: pointer; text-decoration: underline;",
                                    "Back to login"
                                }
                            } else {
                                p {
                                    "Already have an account? "
                                    a {
                                        href: "/login",
                                        style: "color: #8B0E0E; cursor: pointer; text-decoration: underline;",
                                        "Sign in"
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
