use dioxus::prelude::*;
use crate::models::User;
use crate::utils::storage;

#[derive(Clone, Debug)]
pub enum AuthState {
    Unauthenticated,
    Authenticating,
    Authenticated(User),
    Error(String),
}

pub fn use_auth() -> Signal<AuthState> {
    use_signal(|| {
        if let Some(user_json) = storage::get_current_user() {
            if let Ok(user) = serde_json::from_str::<User>(&user_json) {
                AuthState::Authenticated(user)
            } else {
                AuthState::Unauthenticated
            }
        } else {
            AuthState::Unauthenticated
        }
    })
}

pub fn set_authenticated_user(mut auth_signal: Signal<AuthState>, user: User) {
    if let Ok(user_json) = serde_json::to_string(&user) {
        let _ = storage::save_current_user(&user_json);
        auth_signal.set(AuthState::Authenticated(user));
    }
}
