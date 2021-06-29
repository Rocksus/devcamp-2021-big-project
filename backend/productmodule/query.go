package productmodule

const (
	addProductQuery = `
	INSERT INTO product (
		name,
		description,
		price,
		rating,
		image_url,
		additional_image_url
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	) returning id
`
	getProductQuery = `
	SELECT
		name,
		description,
		price,
		rating,
		image_url,
		additional_image_url
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
	LIMIT $1
	OFFSET $2
`

	updateProductQuery = `
	UPDATE
		product
	SET
		%s
	WHERE
		id=%d
`
)
