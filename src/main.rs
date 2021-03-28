#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use]
extern crate rocket;

mod utils;

#[post("/process", format = "application/json", data = "<input>")]
fn process_store(input: utils::store::Store) -> &'static str {
    println!("{}", input.contents);
    "200"
}

fn main() {
    rocket::ignite().mount("/", routes![process_store]).launch();
}
