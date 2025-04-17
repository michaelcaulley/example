-- Create new schemas
CREATE SCHEMA "app";
CREATE SCHEMA "todo";

-- Move tables to their new schemas
ALTER TABLE "public"."users" SET SCHEMA "app";
ALTER TABLE "public"."moderators" SET SCHEMA "app";
ALTER TABLE "public"."todos" SET SCHEMA "todo";
ALTER TABLE "public"."reminders" SET SCHEMA "todo";
ALTER TABLE "public"."todo_reminders" SET SCHEMA "todo";

-- Create/update indexes if needed
CREATE INDEX IF NOT EXISTS "user_name" ON "app"."users" USING gin ("name" gin_trgm_ops);
CREATE INDEX IF NOT EXISTS "todo_owner_id" ON "todo"."todos" ("owner_id");
CREATE INDEX IF NOT EXISTS "moderator_user_id" ON "app"."moderators" ("user_id");
CREATE INDEX IF NOT EXISTS "moderator_moderator_user_id" ON "app"."moderators" ("moderator_user_id");

-- Add comments
COMMENT ON TABLE "todo"."reminders" IS 'Reminder for a user to take action.';
COMMENT ON TABLE "todo"."todo_reminders" IS 'A join table holding the relationships of todos to reminders';