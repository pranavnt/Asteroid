use serde_json::{Result, Value};

pub fn str_to_json(data: &str) ->Result<()>{
    let v: Value = serde_json::from_str(data)?;

    println!("{}", v["hi"]);

    Ok(())
}