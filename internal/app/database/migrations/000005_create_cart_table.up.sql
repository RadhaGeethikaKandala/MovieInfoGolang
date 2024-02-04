

CREATE TABLE IF NOT EXISTS cart (
  customer_id INTEGER NOT NULL,
  movie_id INTEGER NOT NULL,
  constraint pk_cart PRIMARY KEY (customer_id, movie_id),
  constraint fk_cart_customer foreign key (customer_id) REFERENCES customers (id),
  constraint fk_cart_movies foreign key (movie_id) REFERENCES movies (id)
);
