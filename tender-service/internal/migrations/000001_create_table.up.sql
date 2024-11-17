-- Create ENUM type for tender status
CREATE TYPE tender_status AS ENUM ('open', 'closed', 'awarded');

-- Create ENUM type for bid status
CREATE TYPE bid_status AS ENUM ('pending', 'accepted', 'rejected');

CREATE TABLE tenders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(50) DEFAULT 'open',
    client_id UUID NOT NULL ,
    budget NUMERIC(15, 2),
    deadline TIMESTAMP,
    file_url TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);


CREATE TABLE bids (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tender_id UUID NOT NULL REFERENCES tenders(id) ON DELETE CASCADE,
    contractor_id UUID NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    delivery_time INT NOT NULL,
    comments TEXT,
    status bid_status DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);


CREATE TABLE winners (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tender_id UUID NOT NULL REFERENCES tenders(id),
    bid_id UUID NOT NULL
);


CREATE TABLE notifications (
    id UUID PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    relation_id VARCHAR(255) NOT NULL, -- E.g., tender_id, bid_id
    type VARCHAR(50) NOT NULL,         -- Type: tender, bid, etc.
    created_at TIMESTAMP DEFAULT NOW() -- Timestamp for creation
);