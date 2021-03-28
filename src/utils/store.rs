use rocket::data::FromDataSimple;
use rocket::data::Outcome;
use rocket::http::Status;
use rocket::Data;
use rocket::Outcome::Failure;
use rocket::Outcome::Success;
use rocket::Request;
use std::io::Read;

pub struct Store {
    pub contents: String,
}

impl FromDataSimple for Store {
    type Error = String;

    fn from_data(_req: &Request, data: Data) -> Outcome<Self, String> {
        let mut contents = String::new();

        if let Err(e) = data.open().take(256).read_to_string(&mut contents) {
            return Failure((Status::InternalServerError, format!("{:?}", e)));
        }

        Success(Store { contents })
    }
}
