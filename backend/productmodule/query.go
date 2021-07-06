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
	WHERE product.name LIKE '%' || $1 || '%' OR product.description LIKE '%' || $1 || '%'
	LIMIT $2
	OFFSET $3
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
