#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use]
extern crate rocket;

use serde_json::Value;

mod utils;

#[post("/process", format = "application/json", data = "<input>")]
fn process_store(input: utils::store::Store) -> &'static str {
    let req_body: Value = utils::json::str_to_json(&String::from(input.contents));

    println!("{}", req_body);

    "200"
}

fn main() {
    utils::file::create_collection("test");
    rocket::ignite().mount("/", routes![process_store]).launch();
}
