CREATE TABLE "public"."cloud_record" (
	 id bigserial PRIMARY KEY,
	 username VARCHAR(255) NOT NULL,
	 cloud_id VARCHAR(255) NOT NULL,
	 create_at int8 NOT NULL DEFAULT extract(epoch from now())
);
CREATE UNIQUE INDEX idx_unq_cloud_record_index_username ON cloud_record USING btree (username);