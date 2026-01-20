CREATE TABLE IF NOT EXISTS restaurants (
        id              UUID PRIMARY KEY,
        name            VARCHAR(255) NOT NULL,
        description     TEXT,
        address         VARCHAR(255) NOT NULL,
        phone           VARCHAR(20) NOT NULL,
        latitude        DECIMAL(10,8) NOT NULL,
        longitude       DECIMAL(11,8) NOT NULL,
        opening_time    TIME,
        closing_time    TIME,
        created_at      TIMESTAMP DEFAULT NOW(),
        updated_at      TIMESTAMP DEFAULT NOW()
    );
    
CREATE TABLE IF NOT EXISTS menu_items (
        id              UUID PRIMARY KEY,
        restaurant_id   UUID NOT NULL REFERENCES restaurants(id),
        name            VARCHAR(255) NOT NULL,
        description     TEXT,
        price           DECIMAL(10,2) NOT NULL,
        image_url       VARCHAR(500),
        is_available    BOOLEAN DEFAULT TRUE,
        category        VARCHAR(50),
        created_at      TIMESTAMP DEFAULT NOW(),
        updated_at      TIMESTAMP DEFAULT NOW()
    );
    