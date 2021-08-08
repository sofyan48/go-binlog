-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(40) DEFAULT NULL,
  status enum('active','deleted') DEFAULT 'active',
  created timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=latin1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd
