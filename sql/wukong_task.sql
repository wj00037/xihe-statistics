CREATE TABLE "public"."wukong_task" (
	 id bigserial PRIMARY KEY,
	 username VARCHAR(255) NOT NULL,
	 picture_style VARCHAR(255),
	 description VARCHAR(255) NOT NULL,
	 status VARCHAR(255) NOT NULL,
	 links TEXT,
	 created_at int8 NOT NULL DEFAULT extract(epoch from now())
);