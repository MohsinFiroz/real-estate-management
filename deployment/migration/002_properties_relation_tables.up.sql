CREATE TABLE owners
(
    id                   VARCHAR(26) PRIMARY KEY,
    name                 VARCHAR(255) NOT NULL,
    mobile               VARCHAR(50),
    email                VARCHAR(255) NOT NULL UNIQUE,
    communication_medium VARCHAR(100),
    insurance            TEXT,
    account_number       VARCHAR(100),
    bsb                  VARCHAR(20),
    identification       VARCHAR(255),
    address              TEXT,
    notes                TEXT,
    created_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for owners table
CREATE INDEX idx_owners_name ON owners (name);
CREATE INDEX idx_owners_email ON owners (email);
CREATE INDEX idx_owners_mobile ON owners (mobile);

CREATE TABLE tenants
(
    id                   VARCHAR(26) PRIMARY KEY,
    name                 VARCHAR(255) NOT NULL,
    mobile               VARCHAR(50),
    email                VARCHAR(255) NOT NULL UNIQUE,
    communication_medium VARCHAR(100),
    notes                TEXT,
    created_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for tenants table
CREATE INDEX idx_tenants_name ON tenants (name);
CREATE INDEX idx_tenants_email ON tenants (email);
CREATE INDEX idx_tenants_mobile ON tenants (mobile);

CREATE TABLE properties
(
    id                      VARCHAR(26) PRIMARY KEY,
    owner_id                VARCHAR(26) REFERENCES owners (id),
    address                 VARCHAR(255) NOT NULL,
    suburb                  VARCHAR(100) NOT NULL,
    postcode                VARCHAR(20)  NOT NULL,
    key_no                  VARCHAR(100),
    management_fee          DECIMAL(5, 2),
    water_bill_account      VARCHAR(100),
    last_water_bill_reading DECIMAL(10, 2),
    notes                   TEXT,
    other                   TEXT,
    created_at              TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for properties table
CREATE INDEX idx_properties_owner_id ON properties (owner_id);
CREATE INDEX idx_properties_suburb ON properties (suburb);
CREATE INDEX idx_properties_postcode ON properties (postcode);
CREATE INDEX idx_properties_address ON properties (address);