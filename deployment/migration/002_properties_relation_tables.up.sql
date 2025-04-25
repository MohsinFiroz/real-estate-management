CREATE TABLE owners
(
    id                   VARCHAR(26) PRIMARY KEY,
    name                 VARCHAR(255) NOT NULL,
    mobile               VARCHAR(50),
    email                VARCHAR(255) NOT NULL UNIQUE,
    communication_medium VARCHAR(20) CHECK (communication_medium IN ('SMS', 'WeChat', 'WhatsApp')),
    insurance            TEXT,
    account_number       VARCHAR(100),
    bsb                  VARCHAR(20),
    identification       VARCHAR(255),
    address              TEXT,
    notes                TEXT,
    is_active            BOOLEAN   DEFAULT true,
    created_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_owners_name ON owners (name);
CREATE INDEX idx_owners_email ON owners (email);
CREATE INDEX idx_owners_mobile ON owners (mobile);

CREATE TABLE tenants
(
    id                   VARCHAR(26) PRIMARY KEY,
    name                 VARCHAR(255) NOT NULL,
    mobile               VARCHAR(50),
    email                VARCHAR(255) NOT NULL UNIQUE,
    communication_medium VARCHAR(20) CHECK (communication_medium IN ('SMS', 'WeChat', 'WhatsApp')),
    is_active            BOOLEAN   DEFAULT true,
    created_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

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
    is_active               BOOLEAN   DEFAULT true,
    management_fee          DECIMAL(5, 2),
    management_start_date   DATE         NOT NULL,
    management_end_date     DATE         NOT NULL,
    water_bill_account      VARCHAR(100),
    last_water_bill_reading DECIMAL(10, 2),
    notes                   TEXT,
    other                   TEXT,
    created_at              TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_properties_owner_id ON properties (owner_id);
CREATE INDEX idx_properties_suburb ON properties (suburb);
CREATE INDEX idx_properties_postcode ON properties (postcode);
CREATE INDEX idx_properties_address ON properties (address);

CREATE TABLE tenancies
(
    id                VARCHAR(26) PRIMARY KEY,
    property_id       VARCHAR(26) NOT NULL REFERENCES properties (id),
    primary_tenant_id VARCHAR(26) NOT NULL REFERENCES tenants (id),
    status            VARCHAR(20) CHECK DEFAULT 'Active' (status IN ('Active', 'Inactive', 'BondRefund', 'SACAT')),
    rent              INTEGER     NOT NULL,
    bond_amount       INTEGER     NOT NULL,
    bond_id           VARCHAR(100),
    rent_frequency    VARCHAR(20) NOT NULL DEFAULT 'Fortnightly' CHECK (rent_frequency IN ('Weekly', 'Fortnightly', 'Monthly')),
    start_date        DATE        NOT NULL,
    end_date          DATE        NOT NULL,
    notes             TEXT,
    notices           TEXT,
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_tenancies_property_id ON tenancies (property_id);
CREATE INDEX idx_tenancies_primary_tenant_id ON tenancies (primary_tenant_id);
CREATE INDEX idx_tenancies_start_date ON tenancies (start_date);
CREATE INDEX idx_tenancies_end_date ON tenancies (end_date);