-- Remove foreign key constraints
ALTER TABLE "follows" DROP CONSTRAINT IF EXISTS "follows_follower_id_fkey";
ALTER TABLE "follows" DROP CONSTRAINT IF EXISTS "follows_follows_id_fkey";
ALTER TABLE "liked_playlist" DROP CONSTRAINT IF EXISTS "liked_playlist_playlist_id_fkey";
ALTER TABLE "liked_playlist" DROP CONSTRAINT IF EXISTS "liked_playlist_user_id_fkey";
ALTER TABLE "playlists" DROP CONSTRAINT IF EXISTS "playlists_user_id_fkey";

-- Drop tables in reverse order of dependencies
DROP TABLE IF EXISTS "follows";
DROP TABLE IF EXISTS "liked_playlist";
DROP TABLE IF EXISTS "playlists";
DROP TABLE IF EXISTS "users";
