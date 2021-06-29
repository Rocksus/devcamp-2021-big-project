CREATE TABLE product (
    id bigserial PRIMARY KEY,
    name varchar(50) NOT NULL,
    description text,
    price bigint NOT NULL DEFAULT 0,
    rating float NOT NULL DEFAULT 0,
    image_url text,
    additional_image_url  text[]
);