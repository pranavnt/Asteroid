#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use]
extern crate rocket;

use serde_json::{Result, Value};

mod utils;

#[post("/process", format = "application/json", data = "<input>")]
fn process_store(input: utils::store::Store) -> &'static str {
    println!("{}", input.contents);
    let req_body:Result<()> = utils::json::str_to_json(&String::from(input.contents));
    "200"
}

fn main() {
    rocket::ignite().mount("/", routes![process_store]).launch();
}
