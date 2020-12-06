SET database = defaultdb;
CREATE TABLE users (
    id
        CHAR(26) NOT NULL,
    first_name
        VARCHAR(200) NOT NULL,
    last_name
        VARCHAR(200) NOT NULL,
    data
        JSONB NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    INVERTED INDEX users__data__inv (data)
);
INSERT
INTO
    users
        (
            id,
            first_name,
            last_name,
            data
        )
VALUES
    (
        '01CT3MF9WZN99FA1909BAHXHTN',
        'John',
        'Doe',
        '{"foo": "bar"}'
    ),
    (
        '01CT3MJ458P6YDP5E2Q4Z2ZVSP',
        'Jane',
        'Doe',
        NULL
    );

