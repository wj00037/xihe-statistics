CREATE TABLE "public"."download_record" (
	 id bigserial PRIMARY KEY,
	 username VARCHAR(255) NOT NULL,
	 download_path VARCHAR(255) NOT NULL,
	 create_at int8 NOT NULL DEFAULT extract(epoch from now())
);