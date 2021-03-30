use std::fs;

fn create_collection(collection_name: &str) {
    fs::create_dir_all("./.db/" + collection_name);
}

fn read_file_to_string(collection_name: &str, doc_id: &str) -> str {
    return fs::read_to_string("/.db/" + collection_name + "/".doc_id + ".json")
        .expect("Unable to read file");
}

fn dump_to_file(data: &str) -> bool {}
