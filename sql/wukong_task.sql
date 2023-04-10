CREATE TABLE "public"."wukong_task" (
	 id bigserial PRIMARY KEY,
	 user VARCHAR(255) NOT NULL,
	 style VARCHAR(255) NOT NULL,
	 desc VARCHAR(255) NOT NULL,
	 created_at int8 NOT NULL DEFAULT extract(epoch from now())
);