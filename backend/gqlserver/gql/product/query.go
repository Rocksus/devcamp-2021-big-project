package product

const (
	addProductQuery = `
	INSERT INTO product (
		name,
		description,
		price,
		rating,
		image_url,
		preview_image_url,
		slug
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	) returning id
`
	getProductQuery = `
	SELECT
		name,
		description,
		price,
		rating,
		image_url,
		preview_image_url,
		slug
	FROM
		product
	WHERE
		id=$1
`

	getProductBatchQuery = `
	SELECT
		*
	FROM
		product
	WHERE
		id > $1
	LIMIT $2
`

	editProductQuery = `
	UPDATE
		product
	SET
		%s
	WHERE
		id=%d
`
)
