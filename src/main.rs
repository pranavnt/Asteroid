#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use]
extern crate rocket;

use rocket::data::FromDataSimple;
use rocket::data::Outcome;
use rocket::Data;
use rocket::Outcome::Failure;
use rocket::Outcome::Success;
use rocket::Request;
use rocket::http::Status;
use std::io::Read;

struct Store {
    contents: String,
}

impl FromDataSimple for Store {
    type Error = String;

    fn from_data(req: &Request, data: Data) -> Outcome<Self, String> {
        let mut contents = String::new();

        if let Err(e) = data.open().take(256).read_to_string(&mut contents) {
            return Failure((Status::InternalServerError, format!("{:?}", e)));
        }

        println!("{}",contents);

        Success(Store { contents })
    }
}

#[post("/process", format = "application/json", data = "<input>")]
fn process_store(input: Store) -> &'static str {
    "200 Okey Dokey"
}

fn main() {
    rocket::ignite().mount("/", routes![process_store]).launch();
}
