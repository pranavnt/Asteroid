use rocket::data::FromDataSimple;
use rocket::data::Outcome;
use rocket::http::Status;
use rocket::Data;
use rocket::Outcome::Failure;
use rocket::Outcome::Success;
use std::io::Read;
use rocket::Request;

pub struct Store {
    pub contents: String,
}

 impl FromDataSimple for Store {
    type Error = String;

    fn from_data(req: &Request, data: Data) -> Outcome<Self, String> {
        let mut contents = String::new();

        if let Err(e) = data.open().take(256).read_to_string(&mut contents) {
            return Failure((Status::InternalServerError, format!("{:?}", e)));
        }

        Success(Store { contents })
    }
}