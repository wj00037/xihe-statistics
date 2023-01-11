CREATE TABLE "public"."fileupload_record" (
	 id bigserial PRIMARY KEY,
	 username VARCHAR(255) NOT NULL,
	 upload_path VARCHAR(255) NOT NULL,
	 create_at int8 NOT NULL DEFAULT extract(epoch from now())
);
CREATE UNIQUE INDEX idx_unq_fileupload_record_index_username ON fileupload_record USING btree (username);