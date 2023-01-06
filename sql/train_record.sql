CREATE TABLE "public"."train_record" (
	 id bigserial PRIMARY KEY,
	 username VARCHAR(255) NOT NULL,
	 project_id VARCHAR(255) NOT NULL,
	 train_id VARCHAR(255) NOT NULL,
	 create_at int8 NOT NULL DEFAULT extract(epoch from now())
);