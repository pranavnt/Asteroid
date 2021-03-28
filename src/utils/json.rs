use serde_json::Value;

pub fn str_to_json(data: &str) -> Value {
    let v: Value = serde_json::from_str(data).unwrap();

    v
}
