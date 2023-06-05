CREATE TABLE "public"."async_task" (
	 id bigserial PRIMARY KEY,
	 username VARCHAR(255) NOT NULL,
     task_type VARCHAR(255) NOT NULL,
	 status VARCHAR(255) NOT NULL,
	 created_at int8 NOT NULL DEFAULT extract(epoch from now()),
     metadata JSON
);