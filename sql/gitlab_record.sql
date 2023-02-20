CREATE TABLE "public"."gitlab_record" (
	 id bigserial PRIMARY KEY,
	 counts bigint NOT NULL,
	 create_at int8 NOT NULL DEFAULT extract(epoch from now())
);