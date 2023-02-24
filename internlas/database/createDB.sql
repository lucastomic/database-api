USE naturalYSalvaje;
CREATE TABLE user(
  name VARCHAR(255) NOT NULL,
  id INT AUTO_INCREMENT,
  nif VARCHAR(255),
  phone VARCHAR(255),
  business VARCHAR(255),
  ubication VARCHAR(255),

  PRIMARY KEY(id)
);

CREATE TABLE product(
  name VARCHAR(255),
  arrivalDate DATETIME,
  arrivalPlace VARCHAR(255),

  PRIMARY KEY(name)
);

CREATE TABLE caliber(
  name VARCHAR(255),
  amount INT NOT NULL,
  price DECIMAL NOT NULL,
  weight DECIMAL NOT NULL,
  product VARCHAR(255) NOT NULL,

  PRIMARY KEY(product,name),
  FOREIGN KEY (product) REFERENCES product(name)
);

CREATE TABLE sale(
  id INT AUTO_INCREMENT,
  date DATETIME NOT NULL,
  payedAtMoment boolean NOT NULL,
  amount INT NOT NULL,
  userID INT NOT NULL,
  product VARCHAR(255) NOT NULL,
  caliberName VARCHAR(255) NOT NULL,

  PRIMARY KEY(id),
  FOREIGN KEY (userID) REFERENCES user(id),
  FOREIGN KEY (product, caliberName) REFERENCES caliber(product,name)
  ON UPDATE NO ACTION
  ON DELETE NO ACTION
);





CREATE TABLE photo(
  link VARCHAR(255),
  product VARCHAR(255) NOT NULL,

  FOREIGN KEY (product) REFERENCES product(name)
);


CREATE TABLE video(
  link VARCHAR(255),
  product VARCHAR(255) NOT NULL,

  FOREIGN KEY (product) REFERENCES product(name)
);

CREATE TABLE fishingDay(
  date DATE,
  product VARCHAR(255) NOT NULL,

  FOREIGN KEY (product) REFERENCES product(name)
);

