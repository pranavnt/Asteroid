#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use]
extern crate rocket;

use serde_json::Value;

mod utils;

#[post("/document/create", format = "application/json", data = "<input>")]
fn document_create(input: utils::store::Store) -> &'static str {
    let req_body: Value = utils::json::str_to_json(&String::from(input.contents));

    utils::file::create_document(
        &format!("{}", req_body),
        &req_body["collection"].to_string(),
        &req_body["id"].to_string(),
    );

    println!("{}", req_body);

    "200"
}

#[post("/document/get", format = "application/json", data = "<input>")]
fn document_get(input: utils::store::Store) -> String {
    let req_body: Value = utils::json::str_to_json(&String::from(input.contents));

    println!("{}",req_body["collection"].to_string());
    println!("{}",req_body["id"].to_string());

    utils::file::get_document(
        &req_body["collection"].to_string(),
        &req_body["id"].to_string()
    );

    "200".to_string()
}

fn main() {
    rocket::ignite()
        .mount("/", routes![document_create, document_get])
        .launch();
}
