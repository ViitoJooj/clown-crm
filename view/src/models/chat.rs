use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, Serialize, Deserialize, PartialEq)]
pub struct Message {
    pub from: String,
    pub to: String,
    pub message: String,
    pub time: String,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct Chat {
    pub id: String,
    pub name: String,
    pub is_group: bool,
    pub members: Vec<String>,
    pub last_message: Option<Message>,
    pub created_at: String,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct ChatHistory {
    pub messages: Vec<Message>,
    pub total_count: usize,
    pub page: usize,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct Group {
    pub id: String,
    pub name: String,
    pub description: Option<String>,
    pub members: Vec<String>,
    pub admin: String,
    pub created_at: String,
}
