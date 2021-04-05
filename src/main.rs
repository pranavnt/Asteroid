#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use]
extern crate rocket;

use serde_json::Value;

mod utils;

#[post("/document/create", format = "application/json", data = "<input>")]
fn create_document(input: utils::store::Store) -> &'static str {
    let req_body: Value = utils::json::str_to_json(&String::from(input.contents));

    utils::file::create_document(&format!("{}",req_body), req_body["collection"], "hi");

    "{\"sucesss\": true}"
}

#[post("/document/get", format = "application/json", data = "<input>")]
fn create_document(input: utils::store::Store) -> &'static str {
    let req_body: Value = utils::json::str_to_json(&String::from(input.contents));

    // utils::file::create_document(&format!("{}",req_body), "test", "hi");

    println!("{}", req_body);

    "200"
}



fn main() {
    utils::file::create_collection("test");
    rocket::ignite().mount("/", routes![process_store]).launch();
}
