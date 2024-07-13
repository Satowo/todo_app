CREATE TABLE `boards` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	`title` varchar(20) NOT NULL,
	`deleted` boolean NOT NULL DEFAULT FALSE,
	PRIMARY KEY (`id`)
);

CREATE TABLE `categories` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	`board_id` BIGINT UNSIGNED NOT NULL,
	`title` varchar(20) NOT NULL,
	`deleted` boolean NOT NULL DEFAULT FALSE,
	PRIMARY KEY (`id`),
	CONSTRAINT `fk_boards`
	FOREIGN KEY (`board_id`) REFERENCES `boards` (`id`)
);

CREATE TABLE `items` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	`category_id` BIGINT UNSIGNED NOT NULL,
	`title` varchar(20) NOT NULL,
	`content` varchar(500) NOT NULL,
	`expired_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`deleted` boolean NOT NULL DEFAULT FALSE,
	PRIMARY KEY (`id`),
	CONSTRAINT `fk_categories`
	FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
);

INSERT INTO items
(
	id,
	category_id,
	title,
	content,
	expired_at,
	archived
)
VALUES
(1, 1, 'title_1', 'content_1', '2021-01-01 00:00:00', DEFAULT),
(2, 2, 'title_2', 'content_2', '2021-01-01 00:00:00', DEFAULT),
(3, 3, 'title_3', 'content_3', '2021-01-01 00:00:00', DEFAULT),
(4, 4, 'title_4', 'content_4', '2021-01-01 00:00:00', TRUE);

INSERT INTO categories
(
    id,
	board_id,
    title,
    deleted
)
VALUES
(1, 1, "category_1", DEFAULT),
(2, 1, "category_2", DEFAULT),
(3, 2, "category_3", DEFAULT),
(4, 3, "category_4", DEFAULT);

INSERT INTO boards
(
    id,
    title,
    deleted
)
VALUES
(1, 'board_1', DEFAULT),
(2, 'board_2', DEFAULT),
(3, 'board_3', DEFAULT);

