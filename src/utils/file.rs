use std::fs;

fn create_collection(collection_name: &str) {
    fs::create_dir_all(["./.db/", collection_name].to_owned());
}

fn get_document(collection_name: &str, doc_id: &str) -> String {
    return fs::read_to_string(["/.db/", collection_name, "/", doc_id, ".json"].concat())
        .expect("Unable to read file");
}

fn create_document(data: &str, collection_name: &str, doc_id: &str) {
    fs::write(
        concat!("/.db/", collection_name, "/", doc_id, ".json"),
        data,
    )
    .expect("Unable to write file");
}
