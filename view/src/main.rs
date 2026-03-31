mod components;
mod hooks;
mod models;
mod pages;
mod styles;
mod utils;

use dioxus::prelude::*;
use hooks::{use_auth, AuthState};
use pages::LoginPage;
use styles::theme::global_styles;
use utils::storage;
use components::common::AppLayout;

fn main() {
    dioxus::launch(App);
}

#[component]
fn App() -> Element {
    rsx! {
        head {
            title { "Clown CRM" }
            style { "{global_styles()}" }
        }
        
        body {
            style: "margin: 0; padding: 0;",
            AppContent {}
        }
    }
}

#[component]
fn AppContent() -> Element {
    let auth_signal = use_auth();
    let auth_state = auth_signal.read().clone();

    match auth_state {
        AuthState::Authenticated(user) => {
            rsx! {
                AppLayout {
                    user,
                    auth_signal
                }
            }
        },
        AuthState::Unauthenticated | AuthState::Error(_) => {
            rsx! {
                LoginPage {
                    initial_backend_url: storage::get_backend_url().unwrap_or_default(),
                    auth_signal
                }
            }
        },
        AuthState::Authenticating => {
            rsx! {
                div {
                    style: "display: flex; justify-content: center; align-items: center; height: 100vh;",
                    "Authenticating..."
                }
            }
        },
    }
}