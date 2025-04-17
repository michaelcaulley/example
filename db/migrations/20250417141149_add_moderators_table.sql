-- Create "moderators" table
CREATE TABLE "moderators" ("user_id" bigint NOT NULL, "moderator_user_id" bigint NOT NULL, PRIMARY KEY ("user_id", "moderator_user_id"), CONSTRAINT "moderators_users_moderator" FOREIGN KEY ("moderator_user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "moderators_users_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "moderator_moderator_user_id" to table: "moderators"
CREATE INDEX "moderator_moderator_user_id" ON "moderators" ("moderator_user_id");
-- Create index "moderator_user_id" to table: "moderators"
CREATE INDEX "moderator_user_id" ON "moderators" ("user_id");
