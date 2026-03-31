use std::fs;
use std::path::PathBuf;

const JWT_TOKEN_KEY: &str = "clown_crm_jwt_token";
const BACKEND_URL_KEY: &str = "clown_crm_backend_url";
const CURRENT_USER_KEY: &str = "clown_crm_current_user";

fn get_data_dir() -> PathBuf {
    let data_dir = dirs::data_local_dir()
        .unwrap_or_else(|| PathBuf::from("."))
        .join("clown_crm");
    let _ = fs::create_dir_all(&data_dir);
    data_dir
}

pub fn save_token(token: &str) -> Result<(), Box<dyn std::error::Error>> {
    let path = get_data_dir().join(JWT_TOKEN_KEY);
    fs::write(path, token)?;
    Ok(())
}

pub fn get_token() -> Option<String> {
    let path = get_data_dir().join(JWT_TOKEN_KEY);
    fs::read_to_string(path).ok()
}

pub fn clear_token() {
    let path = get_data_dir().join(JWT_TOKEN_KEY);
    let _ = fs::remove_file(path);
}

pub fn save_backend_url(url: &str) -> Result<(), Box<dyn std::error::Error>> {
    let path = get_data_dir().join(BACKEND_URL_KEY);
    fs::write(path, url)?;
    Ok(())
}

pub fn get_backend_url() -> Option<String> {
    let path = get_data_dir().join(BACKEND_URL_KEY);
    fs::read_to_string(path).ok()
}

pub fn save_current_user(user_json: &str) -> Result<(), Box<dyn std::error::Error>> {
    let path = get_data_dir().join(CURRENT_USER_KEY);
    fs::write(path, user_json)?;
    Ok(())
}

pub fn get_current_user() -> Option<String> {
    let path = get_data_dir().join(CURRENT_USER_KEY);
    fs::read_to_string(path).ok()
}

pub fn clear_all() {
    let _ = fs::remove_dir_all(get_data_dir());
}
