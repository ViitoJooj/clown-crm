use reqwest::Client;
use serde::{Deserialize, Serialize};
use crate::models::*;
use crate::utils::storage;

pub struct ApiClient {
    base_url: String,
    client: Client,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ApiError {
    pub status: u16,
    pub message: String,
}

impl ApiClient {
    pub fn new(base_url: String) -> Self {
        ApiClient {
            base_url: base_url.trim_end_matches('/').to_string(),
            client: Client::new(),
        }
    }

    fn get_auth_header(&self) -> Option<String> {
        storage::get_token().map(|token| format!("Bearer {}", token))
    }

    pub async fn login(&self, email: String, password: String) -> Result<LoginResponse, ApiError> {
        let url = format!("{}/api/v1/auth/login", self.base_url);
        let req_body = LoginRequest { email, password };

        let response = self
            .client
            .post(&url)
            .json(&req_body)
            .send()
            .await
            .map_err(|_| ApiError {
                status: 0,
                message: "Failed to connect to server".to_string(),
            })?;

        let status = response.status().as_u16();

        if status == 200 {
            response.json().await.map_err(|_| ApiError {
                status,
                message: "Invalid response format".to_string(),
            })
        } else {
            Err(ApiError {
                status,
                message: "Login failed".to_string(),
            })
        }
    }

    pub async fn register(
        &self,
        first_name: String,
        last_name: String,
        email: String,
        password: String,
    ) -> Result<RegisterResponse, ApiError> {
        let url = format!("{}/api/v1/auth/register", self.base_url);
        let req_body = RegisterRequest {
            first_name,
            last_name,
            email,
            password,
        };

        let response = self
            .client
            .post(&url)
            .json(&req_body)
            .send()
            .await
            .map_err(|_| ApiError {
                status: 0,
                message: "Failed to connect to server".to_string(),
            })?;

        let status = response.status().as_u16();

        if status == 201 || status == 200 {
            response.json().await.map_err(|_| ApiError {
                status,
                message: "Invalid response format".to_string(),
            })
        } else {
            Err(ApiError {
                status,
                message: "Registration failed".to_string(),
            })
        }
    }

    pub async fn get_current_user(&self) -> Result<User, ApiError> {
        let url = format!("{}/users/me", self.base_url);
        let auth_header = self.get_auth_header().ok_or(ApiError {
            status: 401,
            message: "Not authenticated".to_string(),
        })?;

        let response = self
            .client
            .get(&url)
            .header("Authorization", auth_header)
            .send()
            .await
            .map_err(|_| ApiError {
                status: 0,
                message: "Failed to connect to server".to_string(),
            })?;

        let status = response.status().as_u16();

        if status == 200 {
            response.json().await.map_err(|_| ApiError {
                status,
                message: "Invalid response format".to_string(),
            })
        } else {
            Err(ApiError {
                status,
                message: "Failed to fetch user".to_string(),
            })
        }
    }

    pub async fn get_all_users(&self) -> Result<Vec<User>, ApiError> {
        let url = format!("{}/api/v1/users", self.base_url);
        let auth_header = self.get_auth_header().ok_or(ApiError {
            status: 401,
            message: "Not authenticated".to_string(),
        })?;

        let response = self
            .client
            .get(&url)
            .header("Authorization", auth_header)
            .send()
            .await
            .map_err(|_| ApiError {
                status: 0,
                message: "Failed to connect to server".to_string(),
            })?;

        let status = response.status().as_u16();

        if status == 200 {
            response.json().await.map_err(|_| ApiError {
                status,
                message: "Invalid response format".to_string(),
            })
        } else {
            Err(ApiError {
                status,
                message: "Failed to fetch users".to_string(),
            })
        }
    }

    pub async fn send_message(
        &self,
        to_uuid: String,
        message: String,
    ) -> Result<Message, ApiError> {
        let url = format!("{}/api/v1/messages/send", self.base_url);
        let auth_header = self.get_auth_header().ok_or(ApiError {
            status: 401,
            message: "Not authenticated".to_string(),
        })?;

        let req_body = serde_json::json!({
            "to": to_uuid,
            "message": message
        });

        let response = self
            .client
            .post(&url)
            .header("Authorization", auth_header)
            .json(&req_body)
            .send()
            .await
            .map_err(|_| ApiError {
                status: 0,
                message: "Failed to connect to server".to_string(),
            })?;

        let status = response.status().as_u16();

        if status == 200 || status == 201 {
            response.json().await.map_err(|_| ApiError {
                status,
                message: "Invalid response format".to_string(),
            })
        } else {
            Err(ApiError {
                status,
                message: "Failed to send message".to_string(),
            })
        }
    }

    pub async fn get_chat_history(
        &self,
        user_uuid: String,
        page: usize,
    ) -> Result<ChatHistory, ApiError> {
        let url = format!("{}/api/v1/messages/history/{}", self.base_url, user_uuid);
        let auth_header = self.get_auth_header().ok_or(ApiError {
            status: 401,
            message: "Not authenticated".to_string(),
        })?;

        let response = self
            .client
            .get(&url)
            .query(&[("page", page.to_string())])
            .header("Authorization", auth_header)
            .send()
            .await
            .map_err(|_| ApiError {
                status: 0,
                message: "Failed to connect to server".to_string(),
            })?;

        let status = response.status().as_u16();

        if status == 200 {
            response.json().await.map_err(|_| ApiError {
                status,
                message: "Invalid response format".to_string(),
            })
        } else {
            Err(ApiError {
                status,
                message: "Failed to fetch chat history".to_string(),
            })
        }
    }

    pub async fn create_group(
        &self,
        name: String,
        description: Option<String>,
        member_uuids: Vec<String>,
    ) -> Result<Group, ApiError> {
        let url = format!("{}/api/v1/groups", self.base_url);
        let auth_header = self.get_auth_header().ok_or(ApiError {
            status: 401,
            message: "Not authenticated".to_string(),
        })?;

        let req_body = serde_json::json!({
            "name": name,
            "description": description,
            "members": member_uuids
        });

        let response = self
            .client
            .post(&url)
            .header("Authorization", auth_header)
            .json(&req_body)
            .send()
            .await
            .map_err(|_| ApiError {
                status: 0,
                message: "Failed to connect to server".to_string(),
            })?;

        let status = response.status().as_u16();

        if status == 200 || status == 201 {
            response.json().await.map_err(|_| ApiError {
                status,
                message: "Invalid response format".to_string(),
            })
        } else {
            Err(ApiError {
                status,
                message: "Failed to create group".to_string(),
            })
        }
    }

    pub async fn update_user(&self, first_name: Option<String>, last_name: Option<String>) -> Result<User, ApiError> {
        let url = format!("{}/api/v1/users/me", self.base_url);
        let auth_header = self.get_auth_header().ok_or(ApiError {
            status: 401,
            message: "Not authenticated".to_string(),
        })?;

        let mut req_body = serde_json::json!({});
        if let Some(first) = first_name {
            req_body["first_name"] = serde_json::json!(first);
        }
        if let Some(last) = last_name {
            req_body["last_name"] = serde_json::json!(last);
        }

        let response = self
            .client
            .patch(&url)
            .header("Authorization", auth_header)
            .json(&req_body)
            .send()
            .await
            .map_err(|_| ApiError {
                status: 0,
                message: "Failed to connect to server".to_string(),
            })?;

        let status = response.status().as_u16();

        if status == 200 {
            response.json().await.map_err(|_| ApiError {
                status,
                message: "Invalid response format".to_string(),
            })
        } else {
            Err(ApiError {
                status,
                message: "Failed to update user".to_string(),
            })
        }
    }

    pub async fn create_user(&self, email: String, password: String, first_name: String, last_name: String) -> Result<User, ApiError> {
        let url = format!("{}/api/v1/users", self.base_url);
        let auth_header = self.get_auth_header().ok_or(ApiError {
            status: 401,
            message: "Not authenticated".to_string(),
        })?;

        let req_body = serde_json::json!({
            "email": email,
            "password": password,
            "first_name": first_name,
            "last_name": last_name
        });

        let response = self
            .client
            .post(&url)
            .header("Authorization", auth_header)
            .json(&req_body)
            .send()
            .await
            .map_err(|_| ApiError {
                status: 0,
                message: "Failed to connect to server".to_string(),
            })?;

        let status = response.status().as_u16();

        if status == 200 || status == 201 {
            response.json().await.map_err(|_| ApiError {
                status,
                message: "Invalid response format".to_string(),
            })
        } else {
            Err(ApiError {
                status,
                message: "Failed to create user".to_string(),
            })
        }
    }

    pub fn get_base_url(&self) -> &str {
        &self.base_url
    }
}
