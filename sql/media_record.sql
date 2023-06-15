CREATE TABLE "public"."media_record" (
	 id bigserial PRIMARY KEY,
	 name VARCHAR(255) NOT NULL,
	 create_at int8 NOT NULL DEFAULT extract(epoch from now())
);