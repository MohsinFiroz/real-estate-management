DROP INDEX IF EXISTS idx_properties_address;
DROP INDEX IF EXISTS idx_properties_postcode;
DROP INDEX IF EXISTS idx_properties_suburb;
DROP INDEX IF EXISTS idx_properties_owner_id;
DROP TABLE IF EXISTS properties;

DROP INDEX IF EXISTS idx_tenants_mobile;
DROP INDEX IF EXISTS idx_tenants_email;
DROP INDEX IF EXISTS idx_tenants_name;
DROP TABLE IF EXISTS tenants;

DROP INDEX IF EXISTS idx_owners_mobile;
DROP INDEX IF EXISTS idx_owners_email;
DROP INDEX IF EXISTS idx_owners_name;
DROP TABLE IF EXISTS owners;