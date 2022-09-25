ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "owner_currency_key";

ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

DROP TABLE IF EXISTS "users";

-- Way to find key names

-- SELECT nspname
-- FROM pg_catalog.pg_namespace;

-- SELECT con.*
--        FROM pg_catalog.pg_constraint con
--             INNER JOIN pg_catalog.pg_class rel
--                        ON rel.oid = con.conrelid
--             INNER JOIN pg_catalog.pg_namespace nsp
--                        ON nsp.oid = connamespace
--        WHERE nsp.nspname = 'public'
--              AND rel.relname = 'accounts';
