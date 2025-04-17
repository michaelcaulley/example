-- Create index "user_name" to table: "users"
CREATE INDEX "user_name" ON "public"."users" USING gin ("name" gin_trgm_ops);
