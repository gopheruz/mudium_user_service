package querys

var UserCreateQuery string = `
	INSERT INTO users(
			first_name,
			last_name,
			phone_number,
			email,
			gender,
			password,
			username,
			profile_image_url,
			type
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at
`
