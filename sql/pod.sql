CREATE TABLE "public"."pod" (
	id uuid DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
	cloud_id VARCHAR(255) NOT NULL,
	owner VARCHAR(255) NOT NULL,
	status VARCHAR(50) NOT NULL,
	expiry BIGINT NOT NULL,
	error VARCHAR(255),
	access_url VARCHAR(255),
	created_at int8 NOT NULL DEFAULT extract(epoch from now())
);