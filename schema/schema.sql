
CREATE TABLE Task (
    id SERIAL PRIMARY KEY START WITH 1,
    title VARCHAR(255) NOT NULL,  
    description TEXT NOT NULL,
    priority SMALLINT CHECK (Priority IN (1, 2, 3, 4, 5)) NOT NULL,
    createdat BIGINT NOT NULL,
    duedate BIGINT DEFAULT 0,
    status SMALLINT DEFAULT 1 CHECK (Status IN (1, 2, 3)), -- 1= pending, 2=done, 3=expired
    createdby VARCHAR(100) NOT NULL
);
