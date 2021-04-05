use std::fs;

pub fn create_collection(collection_name: &str) {
    fs::create_dir_all(format!("./.db/{}", collection_name));
}

pub fn get_document(collection_name: &str, doc_id: &str) -> String {
    return fs::read_to_string(format!("./.db/{}/{}.json", collection_name, doc_id))
        .expect("Unable to read file");
}

pub fn create_document(data: &str, collection_name: &str, doc_id: &str) {
    println!("./.db/{}/{}.json", collection_name, doc_id);
    fs::write(format!("./.db/{}/{}.json", collection_name, doc_id), data)
        .expect("Unable to write file");
}
