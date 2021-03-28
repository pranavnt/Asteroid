use serde_json::{Result, Value};

pub fn str_to_json(data: &str) -> Value {
    let v: Value = serde_json::from_str(data).unwrap();

    println!("{}", v["hi"]);

    v
}
